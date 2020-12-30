package queue

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type RedisQueueJob interface {
	handle() error
}

type RedisQueue struct {
	WorkerNumber  int
	RedisAddr     string
	RedisDb       int
	RedisPassword string
	pool          *redis.Pool
	handler       func(Msg) error
}

type Msg string

func (q *RedisQueue) Consume(handle func(Msg) error) {
	q.handler = handle
	q.pool = q.redisPoolInit()
	defer func() {
		_ = q.pool.Close()
	}()

	stopFetcherCtx, cancelFetcher := context.WithCancel(context.Background())
	defer func() {
		cancelFetcher()
	}()

	messagesChan := q.fetch(stopFetcherCtx)
	confirms := make(chan Msg)
	defer close(confirms)

	consume := func (in <-chan Msg) {
		for n := range in {
			if err := q.handler(n); err != nil {
				log.Printf("handle error: %s\n", err.Error())
				continue
			}
			confirms<-n
		}
	}

	for i := 0; i < q.WorkerNumber; i++ {
		go consume(messagesChan)
	}

	go q.acknowledge(confirms)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("end")
}

func (q *RedisQueue) redisPoolInit() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     2,
		IdleTimeout: 300 * time.Second,
		MaxActive:   3,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp",
				q.RedisAddr,
				redis.DialDatabase(q.RedisDb))
			if err != nil {
				return nil, err
			}

			if q.RedisPassword != "" {
				if _, err := c.Do("AUTH", q.RedisPassword); err != nil {
					_ = c.Close()
					return nil, err
				}
			}

			return c, err
		},
	}
}

func (q *RedisQueue) fetch(ctx context.Context) <-chan Msg {
	out := make(chan Msg)
	go func() {
		defer close(out)

		olds := q.fetchOld()
		for _, o := range olds {
			out <- Msg(o)
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
				out <- q.doFetch()
			}
		}
	}()
	return out
}

func (q *RedisQueue) fetchOld() []string {
	log.Println("fetch old jobs")
	conn := q.pool.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	olds, err := redis.Strings(conn.Do("lrange", q.getInprogressName(), 0, -1))
	if err != nil {
		log.Print(err)
		return nil
	}

	return olds
}

func (q *RedisQueue) doFetch() Msg {
	conn := q.pool.Get()
	defer conn.Close()

	message, err := redis.String(conn.Do("brpoplpush", q.getQueueName(), q.getInprogressName(), 1))
	if err != nil {
		log.Println("ERR: ", err)
		return ""
	}
	log.Printf("Get: %q \n", message)

	return Msg(message)
}

func (q *RedisQueue) acknowledge(confirm <-chan Msg) {
	conn := q.pool.Get()
	defer conn.Close()

	for c := range confirm {
		_, _ = conn.Do("lrem", q.getInprogressName(), -1, string(c))
	}
}

func (q *RedisQueue) getQueueName() string {
	return fmt.Sprintf(
		"%s:%s:%s",
		viper.GetString("app_name"),
		viper.GetString("app_env"),
		"email_send",
	)
}

func (q *RedisQueue) getInprogressName() string {
	return q.getQueueName() + ":inprogress"
}

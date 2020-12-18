package sending_email

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func GenerateCommand() *cobra.Command {
	return &cobra.Command{
		Use: "work:email",
		Short: "worker for sending email",
		Run: func(cmd *cobra.Command, args []string) {
			pool := redisPoolInit()
			defer func() {
				_ = pool.Close()
			}()

			stopFetcherCtx, cancelFetcher := context.WithCancel(context.Background())
			defer func() {
				cancelFetcher()
			}()

			messagesChan := fetch(stopFetcherCtx, pool)
			concurrent := viper.GetInt("queue.mail_queue.worker_count")
			if concurrent < 1 {
				concurrent = 1
			}

			var wg sync.WaitGroup
			confirms := make(chan *Msg)
			collect := func(in <-chan *Msg) {
				defer wg.Done()
				for n := range in {
					confirms <- n
				}
			}

			for i := 0 ; i < concurrent; i++ {
				wg.Add(1)
				go collect(work(messagesChan))
			}

			go func() {
				wg.Wait()
				close(confirms)
			}()

			go acknowledge(confirms, pool)

			log.Println("looping")

			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM)
			<-signalChan

			log.Println("end")
		},
	}
}

func fetch(ctx context.Context, redis *redis.Pool) <-chan *Msg {
	out := make(chan *Msg)
	go func() {
		close(out)
		olds := fetchOldInprogress(redis)
		for _, o := range olds {
			out <- o
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg := doFetch(redis)
				if msg != nil {
					out <- msg
				}
			}
		}
	}()
	return out
}

func work(msgCh <-chan *Msg) <-chan *Msg {
	out := make(chan *Msg)

	go func() {
		defer close(out)
		for m := range msgCh {
			err := handle(m)
			if err != nil {
				out <- m
			}
		}
	}()

	return out
}

func fetchOldInprogress(pool *redis.Pool) []*Msg {
	log.Println("fetchOldInprogress")
	conn := pool.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	olds, err := redis.Strings(conn.Do("lrange", getInprogressName(), 0, -1))
	log.Println(olds)
	if err != nil {
		log.Print(err)
		return nil
	}

	msgs := make([]*Msg, 0)
	for _, o := range olds {
		m, err := NewMsg(o)
		if err == nil {
			msgs = append(msgs, m)
		}
	}

	return msgs
}

func doFetch (pool *redis.Pool) *Msg {
	conn := pool.Get()
	defer conn.Close()

	message, err := redis.String(conn.Do("brpoplpush", getQueueName(), getInprogressName(), 1))
	if err != nil {
		// If redis returns null, the queue is empty. Just ignore the error.
		if err.Error() != "redigo: nil returned" {
			log.Println("ERR: ", err)
			return nil
		}
	}
	log.Printf("Get %q \n", message)

	msg, err := NewMsg(message)
	if err != nil {
		return nil
	}
	return msg
}

func acknowledge(confirm <-chan *Msg, pool *redis.Pool) {
	for c := range confirm {
		ack(pool, c.original)
	}
}

func ack(pool *redis.Pool, m string) {
	conn := pool.Get()
	defer conn.Close()
	conn.Do("lrem", getInprogressName(), -1, m)
}

func handle(m *Msg) error {
	log.Printf("Handling %q\n", m.content.Email)
	return nil
}

func getQueueName() string {
	return fmt.Sprintf(
		"%s:%s:%s",
		viper.GetString("app_name"),
		viper.GetString("app_env"),
		"email_send",
		)
}

func getInprogressName() string {
	return getQueueName() + ":inprogress"
}

func redisPoolInit() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 2,
		IdleTimeout: 300 * time.Second,
		MaxActive: 3,
		Dial: func () (redis.Conn, error) {
			c, err := redis.Dial("tcp", viper.GetString("redis_addr"))
			if err != nil {
				return nil, err
			}

			password := viper.GetString("redis_password")
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					_ = c.Close()
					return nil, err
				}
			}

			return c, err
		},
	}
}

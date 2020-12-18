package sending_email

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func GenerateCommand() *cobra.Command {
	return &cobra.Command{
		Use: "work:email",
		Short: "worker for sending email",
		Run: func(cmd *cobra.Command, args []string) {
			opt, err := redis.ParseURL("redis://localhost:6379/<db>")
			if err != nil {
				panic(err)
			}

			rdb := redis.NewClient(opt)

			stopFetcherCtx, cancelFetcher := context.WithCancel(context.Background())
			defer cancelFetcher()

			messagesChan := fetch(stopFetcherCtx, rdb)
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
				go collect(work(messagesChan))

			}

			go func() {
				wg.Wait()
				close(confirms)
			}()

			go acknowledge(confirms, rdb)

			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM)
			<-signalChan

			// todo end app
			cancelFetcher()
		},
	}
}

func fetch(ctx context.Context, redis *redis.Client) <-chan *Msg {
	out := make(chan *Msg)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(out)
				return
			default:
				out <- doFetch(redis)
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

func doFetch (rdb *redis.Client) *Msg {
	// get old in-progress job

	return &Msg{}
}

func acknowledge(confirm <-chan *Msg, rdb *redis.Client) {
}

func handle(m *Msg) error {
	return nil
}

type Msg struct {
	original string
}

func (m *Msg)Get(i string) interface{} {
	return gjson.Get(m.original, i)
}


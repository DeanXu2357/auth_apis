package sending_email

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"os"
	"os/signal"
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

			messagesChan := fetch(ctx, rdb)
			concurrent := viper.GetInt("queue.mail_queue.worker_count")
			if concurrent < 1 {
				concurrent = 1
			}
			out := make(chan *Msg)
			for i := 0 ; i < concurrent; i++ {
				work(messagesChan, out)
			}

			go acknowledge(out, rdb)

			//stopWorkerCtx, cancelWorker := context.WithCancel(context.Background())
			//
			//msgChan := make(chan *Msg)
			//confirmMsgChan := make(chan *Msg)
			//var f = &Fetcher{"send_email", stopFetcherCtx, msgChan}
			//Process(stopWorkerCtx, msgChan, viper.GetInt("queue.mail_queue.worker_count"))
			//f.DoFetch()
			//f.DoAcknowledge()

			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM)
			<-signalChan
			go func() {
				for sig := range signalChan {
					switch sig {
					case syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM:
						// do quit()
					}
				}
			}()
		},
	}
}

func fetch(ctx context.Context, redis *redis.Client) <-chan *Msg {
	out := make(chan *Msg)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				close(out)
				return
			default:
				msg := dofetch(redis)
				out <- msg
			}
		}
	}()
	return out
}

func work(msgCh <-chan *Msg, out chan *Msg) {
	go func() {
		defer close(out)
		for m := range msgCh {
			err := handle(m)
			if err != nil {
				out <- m
			}
		}
	}()
}

func acknowledge(confirm <-chan *Msg, rdb *redis.Client) {
	var 
}

func Process(stopCtx context.Context, input chan *Msg, concurrent int) chan *Msg {
	output := make(chan *Msg)
	for i := 0 ; i < concurrent ; i ++ {
		go work(stopCtx, input, output)
	}

	return output
}

//func work(ctx context.Context, input chan *Msg, output chan *Msg) {
//	for {
//		select {
//		case msg := <-input:
//			err := handle(msg)
//			if err != nil {
//				output <- msg
//			}
//		case <-ctx.Done():
//			return
//		}
//	}
//}

func handle(m *Msg) error {
	return nil
}

type Fetcher struct {
	queueName string
	stopCtx context.Context
	messages chan *Msg
	confirms chan *Msg
}

func (f *Fetcher) DoFetch() {}

type Msg struct {
	original string
}

func (m *Msg)Get(i string) interface{} {
	return gjson.Get(m.original, i)
}


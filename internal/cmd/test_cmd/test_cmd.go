package test_cmd

import (
	"auth/internal/events"
	"auth/internal/listeners"
	"auth/lib/email"
	"auth/lib/event_listener"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"time"
)

func GenerateTestCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "for testing",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test success")

			dispatcher := event_listener.NewDispatcher()
			dispatcher.AttachListener(events.Test, listeners.PrintMsgListener{})
			defer dispatcher.Close()

			log.Print("do something")
			var e events.TestEvent
			dispatcher.Dispatch(e)
			log.Print("do rest of works")

			time.Sleep(5 * time.Second)
			//testSendEmail()
		},
	}
}

func testSendEmail() {
	info := email.NewInfo()
	err := email.NewEmail(info).SendMail(
		[]string{"jasugun0000+receiver@gmail.com"},
		"test mail subject",
		"this is a test mail")
	if err != nil {
		fmt.Println(err)
	}
}

package sending_email

import (
	"fmt"
	"github.com/spf13/cobra"
)

func GenerateCommand() *cobra.Command {
	return &cobra.Command{
		Use: "work:email",
		Short: "worker for sending email",
		Run: func(cmd *cobra.Command, args []string) {
			for true {
				// get job
				// do job
				// if failed
				fmt.Println("test")
			}
		},
	}
}

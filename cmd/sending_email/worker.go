package sending_email

import 	"github.com/spf13/cobra"

func GenerateCommand() *cobra.Command {
	return &cobra.Command{
		Use: "work:email",
		Short: "worker for sending email",
	}
}

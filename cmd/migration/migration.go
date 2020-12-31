package migration

import "github.com/spf13/cobra"

func GenerateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "auto prepare db connection info to run go-migrate",
		Run: func(cmd *cobra.Command, args []string) {
			// command refresh db

			// command up

			// command down

			// command rollback

			// command create table
			// ./cmd/migrate.linux-amd64 create -dir ./db/migrations -ext sql  test
		},
	}
}

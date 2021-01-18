package main

import (
	"auth/internal/cmd/migration"
	"auth/internal/cmd/sending_email"
	"auth/internal/cmd/server"
	"auth/internal/cmd/test_cmd"
	"auth/internal/config"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	rootCmd := &cobra.Command{
		Short: "Console commands for this project",
	}

	config.InitialConfigurations()

	rootCmd.AddCommand(sending_email.GenerateCommand())
	rootCmd.AddCommand(test_cmd.GenerateTestCmd())
	rootCmd.AddCommand(server.GenerateServerCmd())
	rootCmd.AddCommand(migration.GenerateCommand())

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

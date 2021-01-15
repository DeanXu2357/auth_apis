package migration

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os/exec"
)

func GenerateCommand() *cobra.Command {
	upCmd := &cobra.Command{
		Use:   "up",
		Short: "migrate database to latest",
		Run: func(cmd *cobra.Command, args []string) {
			Up(prepareCommandString())
		},
	}

	downCmd := &cobra.Command{
		Use:   "down",
		Short: "rollback database",
		Run: func(cmd *cobra.Command, args []string) {
			Down(prepareCommandString())
		},
	}

	refreshCmd := &cobra.Command{
		Use:   "refresh",
		Short: "refresh database",
		Run: func(cmd *cobra.Command, args []string) {
			RefreshDatabase()
		},
	}

	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "auto prepare db connection info to run go-migrate",
	}

	cmd.AddCommand(upCmd, refreshCmd, downCmd)
	return cmd
}

// prepareCommandString
// migration command cheat sheet:
// ./cmd/migrate.linux-amd64 -database "postgres://postgres:fortestpwd@localhost:45487/auth?sslmode=disable" -verbose -path db/migrations up
func prepareCommandString() (string, []string) {
	connection := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("db_user"),
		viper.GetString("db_password"),
		viper.GetString("db_host"),
		viper.GetString("db_port"),
		viper.GetString("db_name"))

	return "./db/migrate.linux-amd64", []string{"-database", connection, "-verbose", "-path", "./db/migrations/"}
}

// RefreshDatabase
// TODO: modify to pipe , prevent concurrent issue
func RefreshDatabase() {
	cmd, args := prepareCommandString()

	Down(cmd, args)

	Up(cmd, args)
}

func Down(cmd string, args []string) {
	cmdDown := exec.Command(cmd, append(args, "down", "-all")...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmdDown.Stdout = &out
	cmdDown.Stderr = &stderr
	cmdDown.Dir = "/go/src/app"
	if err := cmdDown.Run(); err != nil {
		log.Printf("\nout:%s\nerr:%s\n", out.String(), stderr.String())
		log.Print(cmdDown.Args)
		log.Fatal(err)
	}
}

func Up(cmd string, args []string) {
	cmdUp := exec.Command(cmd, append(args, "up")...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmdUp.Stdout = &out
	cmdUp.Stderr = &stderr
	// todo: find way to set flexible abs paths
	cmdUp.Dir = "/go/src/app"
	if err := cmdUp.Run(); err != nil {
		log.Printf("out:%s\nerr:%s\n", out.String(), stderr.String())
		log.Panic(err)
	}
}

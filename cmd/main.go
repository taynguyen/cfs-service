package main

import (
	"log"
	"os"

	"github.com/k0kubun/pp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	cfsservice "gitlab.com/cfs-service"
	"gitlab.com/cfs-service/server"
	"gitlab.com/cfs-service/store"
)

func init() {

}

func main() {
	// 	var cmdEcho = &cobra.Command{
	// 		Use:   "echo [string to echo]",
	// 		Short: "Echo anything to the screen",
	// 		Long: `echo is for echoing anything back.
	// Echo works a lot like print, except it has a child command.`,
	// 		Args: cobra.MinimumNArgs(1),
	// 		Run: func(cmd *cobra.Command, args []string) {
	// 			fmt.Println("Echo: " + strings.Join(args, " "))
	// 		},
	// 	}

	config := &cfsservice.RuntimeConfig{}

	var rootCmd = &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			pp.Println("config:", config) // TODO: Remove this

			// init store
			var dbStore store.IStore
			var err error
			if dbStore, err = store.NewMySQLStore(config.DBConnectionString); err != nil {
				logrus.Fatal("DB connection failed", err)
			}
			if err := dbStore.Migrate(); err != nil {
				logrus.Fatal("DB migration failed", err)
			}

			if err := server.Start(config.Port, dbStore); err != nil {
				log.Fatal(err)
			}
		},
	}

	rootCmd.Flags().Uint64Var(&config.Port, "port", 8080, "Port")
	rootCmd.Flags().StringVar(&config.DBConnectionString, "db-conn", os.Getenv("DB_CONNECTION_STRING"), "DB-Connection string")

	// rootCmd.AddCommand(cmdEcho)
	rootCmd.Execute()
}

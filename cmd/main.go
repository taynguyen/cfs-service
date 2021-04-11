package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	cfsservice "gitlab.com/cfs-service"
	"gitlab.com/cfs-service/server"
	"gitlab.com/cfs-service/service"
	"gitlab.com/cfs-service/store"
)

func init() {

}

func main() {
	config := &cfsservice.RuntimeConfig{}

	var rootCmd = &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			// Init key service
			if err := service.InitializeKeyService(config.JWTPublicKeyPath); err != nil {
				logrus.Fatal("Missing public key file", err)
			}

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

	// TODO: Add sub command as cli tool.
	//   + Migrate down (migrate up is automatically when start app)

	rootCmd.Flags().Uint64Var(&config.Port, "port", 8080, "Port")
	rootCmd.Flags().StringVar(&config.DBConnectionString, "db-conn", os.Getenv("DB_CONNECTION_STRING"), "DB-Connection string")
	rootCmd.Flags().StringVar(&config.JWTPublicKeyPath, "public-key-path", os.Getenv("PUBLIC_KEY_PATH"), "Public key file path, used to validate JWT token")

	// rootCmd.AddCommand(cmdEcho)
	rootCmd.Execute()
}

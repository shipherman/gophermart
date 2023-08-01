/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/shipherman/gophermart/internal/accrual"
	"github.com/shipherman/gophermart/internal/db"
	"github.com/shipherman/gophermart/internal/handlers"
	"github.com/shipherman/gophermart/internal/transport/middleware"
	"github.com/shipherman/gophermart/internal/transport/routes"

	"github.com/caarlos0/env/v8"
	"github.com/spf13/cobra"
)

type Options struct {
	DSN     string `env:"DATABASE_URI"`
	Accrual string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	Address string `env:"RUN_ADDRESS"`
}

var cfg Options

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gophermart",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	dbclient := db.NewClient(cfg.DSN)
	err = dbclient.Start()
	if err != nil {
		os.Exit(1)
	}

	// место для Graceful shutdown
	defer dbclient.Stop()

	// Set accruall address
	accrual.SetAccrualAddress(cfg.Accrual)

	// Run server
	handler := handlers.NewHandler(dbclient)
	authenticator := middleware.NewAuthenticator(dbclient)
	router := routes.NewRouter(handler, &authenticator)

	log.Fatal(http.ListenAndServe(cfg.Address, router))
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Read Environment variables
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gophermart.yaml)")
	if cfg.DSN == "" {
		rootCmd.PersistentFlags().StringVarP(&cfg.DSN,
			"dsn",
			"d",
			"host=localhost port=5432 dbname=postgres user=postgres password=pass sslmode=disable",
			"DataBase connection string")
	}
	if cfg.Accrual == "" {
		rootCmd.PersistentFlags().StringVarP(&cfg.Accrual, "Accrual", "r", "http://localhost:8080", "Accrual service address")
	}
	if cfg.Address == "" {
		rootCmd.PersistentFlags().StringVarP(&cfg.Address, "address", "a", "localhost:9090", "Gophermart address string")
	}
}

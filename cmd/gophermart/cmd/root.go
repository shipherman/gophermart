/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/shipherman/gophermart/internal/clients"
	"github.com/shipherman/gophermart/internal/db"
	"github.com/shipherman/gophermart/internal/handlers"
	gmiddw "github.com/shipherman/gophermart/internal/transport/middleware"
	"github.com/shipherman/gophermart/internal/transport/routes"
	"github.com/shipherman/gophermart/internal/transport/worker"

	"github.com/caarlos0/env/v8"
	"github.com/spf13/cobra"
)

type Options struct {
	DSN     string `env:"DATABASE_URI"`
	Accrual string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	Address string `env:"RUN_ADDRESS"`
}

var cfg Options

// Init new logger;
// To do
// Make it general to gophermart app
// var logEntry = middleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags)}
var aWorker worker.Worker
var server http.Server
var dbclient *db.DBClient
var wg = sync.WaitGroup{}
var idleConnectionsClosed = make(chan struct{})

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

	dbclient = db.NewClient(cfg.DSN)
	err = dbclient.Start()
	if err != nil {
		os.Exit(1)
	}

	// Set accruall address
	clients.ConfigureAccrual(cfg.Accrual, time.Second*10)

	// Init accrual worker here
	aWorker = *worker.New(dbclient)

	// Worker pluc through oders in DB those are does not have final status
	wg.Add(1)
	go aWorker.Run(&wg)

	// Run server
	handler := handlers.NewHandler(dbclient)
	authenticator := gmiddw.NewAuthenticator(dbclient)
	router := routes.NewRouter(handler, &authenticator)

	server = http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	// место для Graceful shutdown

	go gracefullShutdown()
	log.Fatal(server.ListenAndServe())

	<-idleConnectionsClosed

}

func gracefullShutdown() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sigint
	log.Println("Shutting down server")

	aWorker.CloseCh <- true
	close(aWorker.CloseCh)
	dbclient.Stop()

	wg.Wait()

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP Server Shutdown Error: %v", err)
	}
	close(idleConnectionsClosed)
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

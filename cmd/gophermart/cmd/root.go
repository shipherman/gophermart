/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"net/http"
	"os"

	"github.com/shipherman/gophermart/lib/db"
	"github.com/shipherman/gophermart/lib/transport/routes"

	"github.com/spf13/cobra"
)

type Options struct {
	DSN string
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

	client := db.NewClient()
	defer client.Close()

	db.SetClient(client)

	router := routes.NewRouter()
	http.ListenAndServe(":9090", router)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gophermart.yaml)")
	rootCmd.PersistentFlags().StringVarP(&cfg.DSN, "dsn", "d", "host=localhost, port=port", "Connection string")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

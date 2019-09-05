package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var providerURL string

var (
	rootCmd = &cobra.Command{
		Use:   "doggy-cli",
		Short: "Simple command line app to CRUD dogs",
		Long:  "This project is a sample to test pact as a consumer point of view.",
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&providerURL, "url", "http://localhost:8080", "URL of the provider to fetch the dogs")
}

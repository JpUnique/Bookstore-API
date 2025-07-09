//nolint:gochecknoinits,gochecknoglobals //lint issue ignored
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Verbatim",
	Short: "Bookstore Api",
	Long:  `An API used to Delete, Put, Get, Post books into the database.`,
}

// Execute adds all subcommands and starts the CLI.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

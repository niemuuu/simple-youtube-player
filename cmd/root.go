package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yapla",
	Short: "Simple Youtube Audio Player",
	Run:   func(cmd *cobra.Command, args []string) { RunProcess(cmd, args) },
}

// Execute root command
func Execute() {
	rootCmd.PersistentFlags().StringP("query", "q", "", "specify search query in youtube")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

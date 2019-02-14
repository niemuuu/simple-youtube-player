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

func init() {
	rootCmd.PersistentFlags().StringP("query", "q", "", "specify search query in youtube")
	rootCmd.PersistentFlags().BoolP("random", "r", false, "whether yapla plays random")
}

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

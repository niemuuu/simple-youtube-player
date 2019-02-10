package cmd

import (
	"fmt"
	"os"

	"github.com/niemuuu/yapla/pkg/play"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yapla",
	Short: "Simple Youtube Audio Player",
	Run: func(cmd *cobra.Command, args []string) {
		play.MpvPlay()
	},
}

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

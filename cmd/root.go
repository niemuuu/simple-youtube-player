package cmd

import (
	"fmt"
	"os"

	"github.com/niemuuu/yapla/pkg/mpv"
	"github.com/niemuuu/yapla/pkg/youtube"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yapla",
	Short: "Simple Youtube Audio Player",
	Run: func(cmd *cobra.Command, args []string) {
		youtubeSvc, err := youtube.NewService()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		list, err := youtubeSvc.SearchWithQuery("future bass")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		urls, err := youtube.BuildURLsFromItems(list.Items)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = mpv.Play(urls...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	},
}

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

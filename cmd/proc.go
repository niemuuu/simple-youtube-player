package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/niemuuu/yapla/pkg/mpv"
	"github.com/niemuuu/yapla/pkg/youtube"
	"github.com/spf13/cobra"
)

// RunProcess does the main process
func RunProcess(cmd *cobra.Command, args []string) {
	query, err := cmd.PersistentFlags().GetString("query")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if strings.Trim(query, " ") == "" {
		fmt.Fprintln(os.Stderr, `Please specify the search query with '-q' option, like: [yapla -q "Hip Hop"]`)
		return
	}

	rand, err := cmd.PersistentFlags().GetBool("random")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get flag which is random option")
	}

	youtubeSvc, err := youtube.NewService()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	list, err := youtubeSvc.SearchWithQuery(query)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if rand {
		Shuffle(list.Items)
	}

	urls, err := youtube.BuildURLsFromItems(list.Items)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = mpv.Play(urls...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

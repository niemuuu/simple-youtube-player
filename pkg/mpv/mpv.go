package mpv

import (
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

const (
	novid = "--no-video"
	fs    = "--fs"
)

// Play audio
func Play(urls ...string) error {
	opts := []string{novid, fs}
	opts = append(opts, urls...)

	cmd := exec.Command("mpv", opts...)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return errors.Wrap(err, "error occured in mpv-player")
	}

	return nil
}

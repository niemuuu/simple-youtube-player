package mpv

import (
	"log"
	"os"
	"os/exec"
)

// Play audio
func Play() {
	var (
		novid = "--no-video"
		fs    = "--fs"
		url   = "https://www.youtube.com/watch?v=HOlF5sW26Yw"
	)
	opts := []string{novid, fs, url}

	cmd := exec.Command("mpv", opts...)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

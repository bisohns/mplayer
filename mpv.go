package mplayer

import (
	"os"
	"os/exec"
	"runtime"

	log "github.com/sirupsen/logrus"
)

// MPVPlayer : Plays a movie in a HTML web page
type MPVPlayer struct {
	Props
}

// Play : Opens browser and play media
func (p *MPVPlayer) Play() {
	mpv := getMPVExecutable()
	cmd := exec.Command(mpv, p.URL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// Opens URL in browser
func getMPVExecutable() string {
	var executable string
	switch runtime.GOOS {
	case "linux", "darwin":
		executable = "mpv"
	case "windows":
		executable = "mpv.exe"
	default:
		executable = "mpv"
	}
	return executable
}

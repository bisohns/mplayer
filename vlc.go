package mplayer

import (
	"os"
	"os/exec"
	"runtime"

	log "github.com/sirupsen/logrus"
)

// VLCPlayer : Plays a movie in a HTML web page
type VLCPlayer struct {
	Props
}

// Play : Opens browser and play media
func (p *VLCPlayer) Play() {
	vlc := getVLCExecutable()
	cmd := exec.Command(vlc, p.URL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// Opens URL in browser
func getVLCExecutable() string {
	var executable string
	switch runtime.GOOS {
	case "linux", "darwin":
		executable = "vlc"
	case "windows":
		executable = "vlc.exe"
	default:
		executable = "vlc"
	}
	return executable
}

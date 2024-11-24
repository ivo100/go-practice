package main

import (
	"os/exec"
	"time"
)

func main() {
	// Path to a system sound file (you can replace it with another sound file)
	soundPath := "/System/Library/Sounds"

	sounds := []string{
		"Basso.aiff",
		"Frog.aiff",
		"Pop.aiff",
		"Submarine.aiff",
		"Blow.aiff",
		"Funk.aiff",
		"Morse.aiff",
		"Purr.aiff",
		"Tink.aiff",
		"Bottle.aiff",
		"Glass.aiff",
		"Ping.aiff",
		"Sosumi.aiff",
	}
	// Execute the afplay command
	for _, sound := range sounds {
		cmd := exec.Command("afplay", soundPath+"/"+sound)
		_ = cmd.Run()
		time.Sleep(1 * time.Second)
	}
}

package main

import (
	"os"
	"os/exec"
)

func validateBinaries() {
	if _, err := exec.LookPath(ffmpegPath()); err == nil {
		ffmpegFound = true
	}

	if _, err := exec.LookPath(video2xPath()); err == nil {
		video2xFound = true
	} else if _, err := os.Stat(video2xPath()); err == nil {
		video2xFound = true
	}
}

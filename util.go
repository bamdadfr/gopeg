package main

import (
	"os"
	"os/exec"

	"gopeg/binary"
)

// TODO: relocate me to binary package?
func validateBinaries() {
	ffmpeg := binary.Ffmpeg()
	video2x := binary.Video2x()

	if _, err := exec.LookPath(ffmpeg.Command); err == nil {
		ffmpegFound = true
	}

	if _, err := exec.LookPath(video2x.Command); err == nil {
		video2xFound = true
	} else if _, err := os.Stat(video2x.Path.Windows); err == nil {
		video2xFound = true
	}
}

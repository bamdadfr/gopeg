package main

import (
	"log"
	"os"
	"os/exec"
)

func run(args []string) {
	if !isRunning.CompareAndSwap(false, true) {
		return
	}

	defer isRunning.Store(false)

	updateStatus("Computing...")

	var binaryPath string

	if args[0] == ffmpegExec() {
		binaryPath = ffmpegPath()
	}

	if args[0] == video2xExec() {
		binaryPath = video2xPath()
	}

	cmd := exec.Command(binaryPath, args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("Run:", cmd.String())

	if err := cmd.Run(); err != nil {
		log.Println("Error:", err)
		updateStatus("Error:" + err.Error())
		return
	}

	updateStatus("Done!")
}

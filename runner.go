package main

import (
	"log"
	"os"
	"os/exec"

	"gopeg/preset"
)

func run(p *preset.Preset, args []string) {
	if !isRunning.CompareAndSwap(false, true) {
		return
	}

	defer isRunning.Store(false)

	updateStatus("Computing...")

	var cmd *exec.Cmd
	if isWindows {
		cmd = exec.Command(p.Binary.Path.Windows, args...)
	} else {
		cmd = exec.Command(p.Binary.Path.Linux, args...)
	}

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

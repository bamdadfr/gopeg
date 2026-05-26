package main

import (
	"gopeg/binary"
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

	cmd := exec.Command(binary.ResolvedPath(p.Binary.Name), args...)
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

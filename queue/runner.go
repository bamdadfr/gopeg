package queue

import (
	"gopeg/binary"
	"gopeg/env"
	"log"
	"os"
	"os/exec"

	"gopeg/preset"
)

func run(p *preset.Preset, args []string) {
	env.UpdateStatus("Computing...")

	cmd := exec.Command(binary.ResolvedPath(p.Binary.Name), args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("Run:", cmd.String())

	if err := cmd.Run(); err != nil {
		log.Println("Error:", err)
		env.UpdateStatus("Error:" + err.Error())
		return
	}

	env.UpdateStatus("Done!")
}

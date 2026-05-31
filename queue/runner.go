package queue

import (
	"context"
	"fmt"
	"gopeg/binary"
	"gopeg/env"
	"log"
	"os"
	"os/exec"

	"gopeg/preset"
)

func run(ctx context.Context, p *preset.Preset, args []string) error {
	env.UpdateStatus("Computing...")

	cmd := exec.CommandContext(ctx, binary.ResolvedPath(p.Binary.Name), args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("Run:", cmd.String())

	if err := cmd.Run(); err != nil {
		log.Println("Error:", err)
		env.UpdateStatus("Error: " + err.Error())
		return fmt.Errorf("%s: %w", p.Binary.Name, err)
	}

	env.UpdateStatus("Done!")
	return nil
}

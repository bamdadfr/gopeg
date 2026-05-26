package main

import (
	"context"
	"gopeg/binary"
	"gopeg/env"
	"gopeg/queue"
	"gopeg/ui"

	"gopeg/preset"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env.UpdateStatus("Loading...")

	binary.Init()
	presets := preset.List()

	q := queue.NewQueue(ctx)
	w := ui.Init(presets, q)

	w.SetCloseIntercept(func() {
		cancel()
		w.Close()
	})

	env.UpdateStatus("Ready")
	w.ShowAndRun()
}

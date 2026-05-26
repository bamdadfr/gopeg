package main

import (
	"gopeg/binary"
	"gopeg/env"
	"gopeg/queue"
	"gopeg/ui"

	"gopeg/preset"
)

func main() {
	env.UpdateStatus("Loading...")

	binary.Init()
	presets := preset.List()

	q := queue.NewQueue()
	w := ui.Init(presets, q)

	env.UpdateStatus("Ready")
	w.ShowAndRun()
}

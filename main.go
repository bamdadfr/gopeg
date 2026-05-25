package main

import (
	"sort"

	"fyne.io/fyne/v2/container"

	"gopeg/preset"
)

func main() {
	updateStatus("Loading...")
	isRunning.Store(false)

	presets := []preset.Preset{
		preset.Remux(),
		preset.Archive(),
		preset.DecodeAndLoop(),
		preset.Upscale4x(),
		preset.Interpolate4x(),
	}

	sort.Slice(presets, func(i int, j int) bool {
		a := presets[i].Name
		b := presets[j].Name
		return a < b
	})

	validateBinaries()

	w := createWindow()
	list := createList(presets)
	statusBar := createStatusBar()
	root := container.NewBorder(nil, statusBar, nil, nil, list)
	w.SetContent(root)

	handleDropEvent(w, list, presets)
	updateStatus("Ready")

	w.ShowAndRun()
}

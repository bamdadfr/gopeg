package main

import (
	"slices"
	"strings"

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

	slices.SortFunc(presets, func(a, b preset.Preset) int {
		return strings.Compare(a.Name, b.Name)
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

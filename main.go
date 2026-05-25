package main

import (
	"fyne.io/fyne/v2/container"

	"gopeg/preset"
)

func main() {
	updateStatus("Loading...")
	isRunning.Store(false)
	presets := preset.List()

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

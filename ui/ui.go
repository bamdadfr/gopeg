package ui

import (
	"gopeg/preset"
	"gopeg/queue"
	"gopeg/ui/events"
)
import "fyne.io/fyne/v2"
import "fyne.io/fyne/v2/container"

func Init(presets []preset.Preset, q *queue.Queue) fyne.Window {
	w := CreateWindow()
	list := CreatePresetList(presets)
	statusBar := CreateStatusBar()
	root := container.NewBorder(nil, statusBar, nil, nil, list)
	w.SetContent(root)
	events.HandleSelect(list, presets)
	events.HandleDrop(w, q)
	return w
}

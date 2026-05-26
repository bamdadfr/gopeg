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
	statusBar := CreateStatusBar()
	presetList := CreatePresetList(presets)
	queueUi := CreateQueue(q)

	tabs := container.NewAppTabs(
		container.NewTabItem("Presets", presetList),
		container.NewTabItem("Queue", queueUi),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	root := container.NewBorder(nil, statusBar, nil, nil, tabs)

	w.SetContent(root)
	events.HandleSelect(presetList, presets)
	events.HandleDrop(w, q)
	return w
}

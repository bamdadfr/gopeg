package ui

import (
	"gopeg/pipeline"
	"gopeg/preset"
	"gopeg/queue"
	"gopeg/ui/events"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Init(
	q *queue.Queue,
	presets []preset.Preset,
	pipelines []pipeline.Pipeline,
) fyne.Window {
	w := CreateWindow()
	statusBar := CreateStatusBar()
	presetList := CreatePresetList(presets)
	pipelineList := CreatePipelineList(pipelines)
	queueUi := CreateQueue(q)

	tabs := container.NewAppTabs(
		container.NewTabItem("Presets", presetList),
		container.NewTabItem("Pipelines", pipelineList),
		container.NewTabItem("Queue", queueUi),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	root := container.NewBorder(nil, statusBar, nil, nil, tabs)

	w.SetContent(root)
	events.HandleSelect(presetList, presets, pipelineList)
	events.HandlePipelineSelect(pipelineList, pipelines, presetList)
	events.HandleDrop(w, q)
	return w
}

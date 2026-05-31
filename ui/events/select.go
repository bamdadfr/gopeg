package events

import (
	"gopeg/pipeline"
	"gopeg/preset"

	"fyne.io/fyne/v2/widget"
)

func HandleSelect(list *widget.List, presets []preset.Preset, pipelineList *widget.List) {
	list.OnSelected = func(id widget.ListItemID) {
		selectedPreset = &presets[id]
		selectedPipeline = nil
		selectionKind = SelectionPreset
		pipelineList.UnselectAll()
	}
}

func HandlePipelineSelect(list *widget.List, pipelines []pipeline.Pipeline, presetList *widget.List) {
	list.OnSelected = func(id widget.ListItemID) {
		selectedPipeline = &pipelines[id]
		selectedPreset = nil
		selectionKind = SelectionPipeline
		presetList.UnselectAll()
	}
}

package events

import (
	"gopeg/pipeline"
	"gopeg/preset"
)

var selectedPreset *preset.Preset
var selectedPipeline *pipeline.Pipeline

type SelectionKind int

const (
	SelectionNone SelectionKind = iota
	SelectionPreset
	SelectionPipeline
)

var selectionKind = SelectionNone

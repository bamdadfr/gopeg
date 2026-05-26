package events

import (
	"gopeg/preset"

	"fyne.io/fyne/v2/widget"
)

func HandleSelect(list *widget.List, presets []preset.Preset) {
	list.OnSelected = func(id widget.ListItemID) {
		selectedPreset = &presets[id]
	}
}

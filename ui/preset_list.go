package ui

import (
	"gopeg/preset"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreatePresetList(presets []preset.Preset) *widget.List {
	list := widget.NewList(
		func() int {
			return len(presets)
		},
		func() fyne.CanvasObject {
			return container.NewVBox(
				widget.NewLabel(""),
				widget.NewLabel(""),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			c := o.(*fyne.Container)

			name := c.Objects[0].(*widget.Label)
			name.SetText(presets[i].Name)
			name.TextStyle.Bold = true

			description := c.Objects[1].(*widget.Label)
			description.SetText(presets[i].Description)
		},
	)

	return list
}

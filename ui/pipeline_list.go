package ui

import (
	"gopeg/pipeline"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreatePipelineList(pipelines []pipeline.Pipeline) *widget.List {
	list := widget.NewList(
		func() int {
			return len(pipelines)
		},
		func() fyne.CanvasObject {
			return container.NewVBox(
				widget.NewLabel(""),
				widget.NewLabel(""),
				widget.NewLabel(""),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			c := o.(*fyne.Container)

			name := c.Objects[0].(*widget.Label)
			name.SetText(pipelines[i].Name)
			name.TextStyle.Bold = true

			desc := c.Objects[1].(*widget.Label)
			desc.SetText(pipelines[i].Description)

			steps := c.Objects[2].(*widget.Label)
			names := make([]string, len(pipelines[i].Steps))
			for j, s := range pipelines[i].Steps {
				names[j] = s.Name
			}
			steps.SetText("Steps: " + strings.Join(names, " → "))
		},
	)

	return list
}

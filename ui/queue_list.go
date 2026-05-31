package ui

import (
	"gopeg/env"
	"gopeg/queue"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateQueue(q *queue.Queue) *fyne.Container {
	playButton := widget.NewButton("Start", func() {
		if q.IsRunning() {
			return
		}

		q.Execute()
	})

	purgeButton := widget.NewButton("Purge", func() {
		if q.IsRunning() {
			return
		}

		q.Purge()
	})

	buttons := container.New(layout.NewHBoxLayout(), playButton, purgeButton)

	list := widget.NewList(
		func() int {
			return q.Length()
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
			job := q.Job(i)

			t1 := c.Objects[0].(*widget.Label)
			t2 := c.Objects[1].(*widget.Label)
			t3 := c.Objects[2].(*widget.Label)

			var icon string
			if job.IsDone {
				icon = env.Icons["success"]
			} else if job.IsRunning {
				icon = env.Icons["play"]
			} else {
				icon = env.Icons["wait"]
			}

			t1.SetText(job.Preset.Name + " " + icon)
			t1.TextStyle.Bold = true

			t2.SetText("In: " + job.InputPath)
			t3.SetText("Out: " + job.OutputPath)
		},
	)

	q.OnNotify(func() {
		fyne.Do(func() {
			list.Refresh()
		})
	})

	root := container.NewBorder(buttons, nil, nil, nil, list)
	return root
}

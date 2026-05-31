package events

import (
	"gopeg/binary"
	"gopeg/env"
	"gopeg/queue"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func HandleDrop(w fyne.Window, q *queue.Queue) {
	w.SetOnDropped(func(pos fyne.Position, uris []fyne.URI) {
		if len(uris) == 0 {
			env.UpdateStatus("No file detected!")
			return
		}

		if selectionKind == SelectionNone {
			env.UpdateStatus("No preset or pipeline selected!")
			return
		}

		if q.IsRunning() {
			env.UpdateStatus("Already running!")
			time.AfterFunc(2*time.Second, func() {
				env.UpdateStatus("Computing...")
			})
			return
		}

		for _, uri := range uris {
			inputPath := uri.Path()

			switch selectionKind {
			case SelectionPreset:
				if !binary.IsAvailable(selectedPreset.Binary.Name) {
					env.UpdateStatus(selectedPreset.Binary.Name + " not found!")
					return
				}

				if !selectedPreset.IsValidInput(inputPath) {
					dialog.ShowInformation(
						"Invalid file extension",
						"Expected: "+strings.Join(selectedPreset.Accept, ", "),
						w,
					)
					return
				}

				q.AddJob(inputPath, selectedPreset)

			case SelectionPipeline:
				for _, step := range selectedPipeline.Steps {
					if !binary.IsAvailable(step.Binary.Name) {
						env.UpdateStatus(step.Binary.Name + " not found!")
						return
					}
				}

				// Validate input against first step only
				first := selectedPipeline.Steps[0]
				if !first.IsValidInput(inputPath) {
					dialog.ShowInformation(
						"Invalid file extension",
						"Expected: "+strings.Join(first.Accept, ", "),
						w,
					)
					return
				}

				q.AddPipeline(inputPath, selectedPipeline)
			}
		}
	})
}

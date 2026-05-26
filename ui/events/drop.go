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

		if selectedPreset == nil {
			env.UpdateStatus("No preset selected!")
			return
		}

		if !binary.IsAvailable(selectedPreset.Binary.Name) {
			env.UpdateStatus(selectedPreset.Binary.Name + " not found!")
			return
		}

		if q.IsLocked {
			env.UpdateStatus("Already running!")

			time.AfterFunc(2*time.Second, func() {
				env.UpdateStatus("Computing...")
			})

			return
		}

		for _, uri := range uris {
			inputPath := uri.Path()

			if !selectedPreset.IsValidInput(inputPath) {
				dialog.ShowInformation(
					"Invalid file extension",
					"Expected: "+strings.Join(selectedPreset.Accept, ", "),
					w,
				)
				return
			}

			q.AddJob(inputPath, selectedPreset)
		}
	})
}

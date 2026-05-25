package main

import (
	"log"
	"os"
	"strings"
	"time"

	"gopeg/preset"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createWindow() fyne.Window {
	a := app.New()
	a.Settings().SetTheme(theme.DefaultTheme())

	if err := os.Setenv("FYNE_SCALE", "1.5"); err != nil {
		log.Fatal(err)
	}

	w := a.NewWindow("gopeg")
	w.Resize(fyne.NewSize(600, 400))

	return w
}

func createStatusBar() *fyne.Container {
	var ffmpegWidget fyne.Widget
	var video2xWidget fyne.Widget

	if ffmpegFound {
		ffmpegWidget = widget.NewLabel("ffmpeg " + icons["success"])
	} else {
		ffmpegWidget = widget.NewLabel("ffmpeg " + icons["error"])
	}

	if video2xFound {
		video2xWidget = widget.NewLabel("video2x " + icons["success"])
	} else {
		video2xWidget = widget.NewLabel("video2x " + icons["error"])
	}

	statusBar := container.NewHBox(
		widget.NewLabelWithData(statusText),
		layout.NewSpacer(),
		ffmpegWidget,
		video2xWidget,
	)

	return statusBar
}

func createList(presets []preset.Preset) *widget.List {
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
			c.Objects[0].(*widget.Label).SetText(presets[i].Name)
			c.Objects[1].(*widget.Label).SetText(presets[i].Description)
		},
	)

	return list
}

func updateStatus(text string) {
	err := statusText.Set(text)
	if err != nil {
		log.Println("Status error:", err)
	}
}

func handleDropEvent(w fyne.Window, list *widget.List, presets []preset.Preset) {
	var selectedPreset *preset.Preset // default value is nil

	list.OnSelected = func(id widget.ListItemID) {
		selectedPreset = &presets[id]
	}

	w.SetOnDropped(func(pos fyne.Position, uris []fyne.URI) {
		if len(uris) == 0 {
			updateStatus("No file detected!")
			return
		}

		if selectedPreset == nil {
			updateStatus("No preset selected!")
			return
		}

		binaryName := selectedPreset.Binary()

		if binaryName == "ffmpeg" && !ffmpegFound {
			updateStatus("ffmpeg not found!")
			return
		}

		if binaryName == "video2x" && !video2xFound {
			updateStatus("video2x not found!")
			return
		}

		if isRunning.Load() {
			updateStatus("Already running!")

			time.AfterFunc(2*time.Second, func() {
				updateStatus("Computing...")
			})

			return
		}

		inputPath := uris[0].Path()
		outputPath := selectedPreset.OutputPath(inputPath)

		if !selectedPreset.IsValidInput(inputPath) {
			dialog.ShowInformation(
				"Invalid file extension",
				"Expected: "+strings.Join(selectedPreset.Accept, ", "),
				w,
			)
			return
		}

		args := selectedPreset.Args(inputPath, outputPath)
		log.Println("command:", args)

		if selectedPreset.IsExistPath(outputPath) {
			dialog.ShowConfirm(
				"Output file exists!",
				"The output path already exists. Overwrite?",
				func(overwrite bool) {
					if !overwrite {
						return
					}

					if args[0] == "ffmpeg" {
						args = append([]string{args[0], "-y"}, args[1:]...)
					}

					go func() {
						run(args)
						w.Canvas().Refresh(w.Content())
					}()
				},
				w,
			)
			return
		}

		go func() {
			run(args)
			w.Canvas().Refresh(w.Content())
		}()
	})
}

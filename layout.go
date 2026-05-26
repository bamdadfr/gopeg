package main

import (
	"gopeg/binary"
	"log"
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
	// a.Settings().SetTheme(theme.DefaultTheme())
	a.Settings().SetTheme(&CustomTheme{Theme: theme.DefaultTheme()})

	w := a.NewWindow("gopeg")
	w.Resize(fyne.NewSize(600, 400))

	return w
}

func createStatusBar() *fyne.Container {
	var ffmpegText *widget.Label
	var video2xText *widget.Label

	ffmpegIsAvailable := binary.IsAvailable(binary.Ffmpeg().Name)
	video2xIsAvailable := binary.IsAvailable(binary.Video2x().Name)

	if ffmpegIsAvailable {
		ffmpegText = widget.NewLabel("ffmpeg " + icons["success"])
	} else {
		ffmpegText = widget.NewLabel("ffmpeg " + icons["error"])
	}

	if video2xIsAvailable {
		video2xText = widget.NewLabel("video2x " + icons["success"])
	} else {
		video2xText = widget.NewLabel("video2x " + icons["error"])
	}

	statusBar := container.NewHBox(
		widget.NewLabelWithData(statusText),
		layout.NewSpacer(),
		ffmpegText,
		video2xText,
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

			name := c.Objects[0].(*widget.Label)
			name.SetText(presets[i].Name)
			name.TextStyle.Bold = true

			description := c.Objects[1].(*widget.Label)
			description.SetText(presets[i].Description)
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

		if selectedPreset.Binary.Name == binary.Ffmpeg().Name && !binary.IsAvailable(binary.Ffmpeg().Name) {
			updateStatus("ffmpeg not found!")
			return
		}

		if selectedPreset.Binary.Name == binary.Video2x().Name && !binary.IsAvailable(binary.Video2x().Name) {
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

					if selectedPreset.Binary.OverwriteFlag != "" {
						args = append(
							[]string{
								selectedPreset.Binary.OverwriteFlag,
							},
							args...,
						)
					}

					go func() {
						run(selectedPreset, args)
					}()
				},
				w,
			)

			return
		}

		go func() {
			run(selectedPreset, args)
		}()
	})
}

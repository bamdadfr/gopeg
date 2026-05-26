package ui

import (
	"gopeg/binary"
	"gopeg/env"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateStatusBar() *fyne.Container {
	var ffmpegText *widget.Label
	var video2xText *widget.Label

	ffmpegIsAvailable := binary.IsAvailable(binary.Ffmpeg().Name)
	video2xIsAvailable := binary.IsAvailable(binary.Video2x().Name)

	if ffmpegIsAvailable {
		ffmpegText = widget.NewLabel("ffmpeg " + env.Icons["success"])
	} else {
		ffmpegText = widget.NewLabel("ffmpeg " + env.Icons["error"])
	}

	if video2xIsAvailable {
		video2xText = widget.NewLabel("video2x " + env.Icons["success"])
	} else {
		video2xText = widget.NewLabel("video2x " + env.Icons["error"])
	}

	statusBar := container.NewHBox(
		widget.NewLabelWithData(env.StatusText),
		layout.NewSpacer(),
		ffmpegText,
		video2xText,
	)

	return statusBar
}

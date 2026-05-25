package main

import (
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"gopeg/preset"
)

var (
	ffmpegFound  = false
	video2xFound = false
	statusText   = binding.NewString()
	icons        = map[string]string{
		"success": "✅",
		"error":   "❌",
	}
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

func handleDropEvent(w fyne.Window, list *widget.List, presets []preset.Preset) {
	var selectedPreset *preset.Preset // default value is nil

	list.OnSelected = func(id widget.ListItemID) {
		selectedPreset = &presets[id]
	}

	w.SetOnDropped(func(pos fyne.Position, uris []fyne.URI) {
		if len(uris) == 0 {
			statusText.Set("No file detected!")
			return
		}

		if selectedPreset == nil {
			statusText.Set("No preset selected!")
			return
		}

		dumbArgs := selectedPreset.Args("", "")

		if dumbArgs[0] == "ffmpeg" && !ffmpegFound {
			statusText.Set("ffmpeg not found!")
			return
		}

		if dumbArgs[0] == "video2x" && !video2xFound {
			statusText.Set("video2x not found!")
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

					args = append([]string{args[0], "-y"}, args[1:]...)
					statusText.Set("Computing...")
					run(args)
					statusText.Set("Done!")
				},
				w,
			)
			return
		}

		statusText.Set("Computing...")
		run(args)
		statusText.Set("Done!")
	})
}

func run(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("Run:", cmd.String())

	if err := cmd.Run(); err != nil {
		log.Println("ffmpeg error:", err)
		return
	}

	log.Println("Done.")
}

func validateBinaries() {
	if _, err := exec.LookPath("ffmpeg"); err == nil {
		ffmpegFound = true
	}

	if _, err := exec.LookPath("video2x"); err == nil {
		video2xFound = true
	}
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

func main() {
	statusText.Set("Loading...")

	presets := []preset.Preset{
		preset.Remux(),
		preset.Archive(),
		preset.DecodeAndLoop(),
		preset.Upscale4x(),
	}

	sort.Slice(presets, func(i int, j int) bool {
		a := presets[i].Name
		b := presets[j].Name
		return a < b
	})

	validateBinaries()

	w := createWindow()
	list := createList(presets)
	statusBar := createStatusBar()
	root := container.NewBorder(nil, statusBar, nil, nil, list)
	w.SetContent(root)

	handleDropEvent(w, list, presets)
	statusText.Set("Ready")

	w.ShowAndRun()
}

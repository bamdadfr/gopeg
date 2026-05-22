package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"gopeg/preset"
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

func createList(w fyne.Window, presets []preset.Preset) *widget.List {
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

	w.SetContent(list)

	return list
}

func handleDropEvent(w fyne.Window, list *widget.List, presets []preset.Preset) {
	var selectedPreset *preset.Preset // default value is nil

	list.OnSelected = func(id widget.ListItemID) {
		selectedPreset = &presets[id]
	}

	w.SetOnDropped(func(pos fyne.Position, uris []fyne.URI) {
		if len(uris) == 0 || selectedPreset == nil {
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

		args := selectedPreset.BuildArgs(inputPath, outputPath)
		log.Println("ffmpeg args:", args)

		if selectedPreset.IsExistPath(outputPath) {
			log.Println("lol")
			dialog.ShowConfirm("Output file exists!", "The output path already exists. Overwrite?", func(overwrite bool) {
				if !overwrite {
					return
				}

				args = append([]string{"-y"}, args...)
				run(args)
			},
				w)
			return
		}

		run(args)
	})
}

func run(args []string) {
	cmd := exec.Command("ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("Run:", cmd.String())

	if err := cmd.Run(); err != nil {
		log.Println("ffmpeg error:", err)
		return
	}

	log.Println("Done.")
}

func main() {
	presets := []preset.Preset{
		preset.Remux(),
	}

	if _, err := exec.LookPath("ffmpeg"); err != nil {
		log.Fatal(err)
	}

	w := createWindow()
	list := createList(w, presets)
	handleDropEvent(w, list, presets)

	w.ShowAndRun()
}

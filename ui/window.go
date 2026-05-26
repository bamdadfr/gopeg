package ui

import (
	"gopeg/env"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func CreateWindow() fyne.Window {
	a := app.New()
	// a.Settings().SetTheme(theme.DefaultTheme())
	a.Settings().SetTheme(&CustomTheme{Theme: theme.DefaultTheme()})

	w := a.NewWindow(env.AppName)
	w.Resize(fyne.NewSize(600, 400))

	return w
}

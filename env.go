package main

import (
	"runtime"

	"fyne.io/fyne/v2/data/binding"
)

const (
	isWindows = runtime.GOOS == "windows"
)

var (
	ffmpegFound  = false
	video2xFound = false
	isRunning    = false
	statusText   = binding.NewString()
	icons        = map[string]string{
		"success": "✅",
		"error":   "❌",
	}
)

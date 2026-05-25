package main

import (
	"runtime"
	"sync/atomic"

	"fyne.io/fyne/v2/data/binding"
)

const (
	isWindows = runtime.GOOS == "windows"
)

var isRunning atomic.Bool

var (
	ffmpegFound  = false
	video2xFound = false
	statusText   = binding.NewString()
	icons        = map[string]string{
		"success": "✅",
		"error":   "❌",
	}
)

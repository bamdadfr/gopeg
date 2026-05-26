package main

import (
	"sync/atomic"

	"fyne.io/fyne/v2/data/binding"
)

var isRunning atomic.Bool

var (
	statusText = binding.NewString()
	icons      = map[string]string{
		"success": "✅",
		"error":   "❌",
	}
)

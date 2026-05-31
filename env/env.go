package env

import (
	"log"
	"runtime"

	"fyne.io/fyne/v2/data/binding"
)

var (
	AppName    = "gopeg"
	StatusText = binding.NewString()
	Icons      = map[string]string{
		"success": "✅",
		"error":   "❌",
		"wait":    "⌛",
		"play":    "▶️",
	}
)

func UpdateStatus(text string) {
	go func() {
		err := StatusText.Set(text)
		if err != nil {
			log.Println("Status error:", err)
		}
	}()
}

func NullDevice() string {
	if runtime.GOOS == "windows" {
		return "NUL"
	}
	return "/dev/null"
}

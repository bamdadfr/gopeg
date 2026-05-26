package env

import (
	"log"

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
	err := StatusText.Set(text)
	if err != nil {
		log.Println("Status error:", err)
	}
}

// Package preset defines FFmpeg preset types.
package preset

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Preset struct {
	Name        string
	Description string
	Accept      []string
	Target      string
	BuildArgs   func(inputPath string, outputPath string) []string
}

func (p Preset) IsValidInput(inputPath string) bool {
	if len(p.Accept) == 0 {
		return true
	}

	ext := filepath.Ext(inputPath)
	return slices.Contains(p.Accept, ext)
}

func (p Preset) OutputPath(inputPath string) string {
	ext := filepath.Ext(inputPath)
	base := strings.TrimSuffix(inputPath, ext)
	return base + p.Target
}

func (p Preset) IsExistPath(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}

// Package preset defines binary usage with specific arguments
package preset

import (
	"os"
	"path/filepath"
	"slices"
	"strings"

	"gopeg/binary"
)

type Preset struct {
	Name        string
	Binary      binary.Binary
	Description string
	Accept      []string // pass empty for wildcard
	Ext         string   // extension
	Suffix      string
	Args        func(inputPath string, outputPath string) []string
}

func List() []Preset {
	presets := []Preset{
		ArchiveCrf4(),
		ArchiveCrf8(),
		ArchiveCrf10(),
		ArchiveCrf12(),
		ArchiveCrf16(),
		Remux(),
		Decode(),
		DecodeAndLoop(),
		InstagramVerticalV1(),
		InstagramVerticalV2(),
		Interpolate4x(),
		Upscale4x(),
	}

	slices.SortFunc(presets, func(a, b Preset) int {
		return strings.Compare(a.Name, b.Name)
	})

	return presets
}

func (p Preset) IsValidInput(inputPath string) bool {
	if len(p.Accept) == 0 {
		return true
	}

	ext := filepath.Ext(inputPath)
	return slices.Contains(p.Accept, ext)
}

func (p Preset) IsExistPath(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}

func (p Preset) OutputPath(inputPath string) string {
	base := pathBaseName(inputPath)
	return base + "_" + p.Suffix + "." + p.Ext
}

func pathBaseName(inputPath string) string {
	ext := filepath.Ext(inputPath)
	base := strings.TrimSuffix(inputPath, ext)
	return base
}

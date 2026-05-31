// Package preset defines binary usage with specific arguments
package preset

import (
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
	Args        func(inputPath string, outputPath string) [][]string
}

func List() []Preset {
	presets := []Preset{
		// ffv1
		Ffv1(),
		Ffv1Loop(3),
		Ffv1Loop(4),
		// x264 CBR
		X264(1),
		X264(4),
		X264(8),
		X264(10),
		X264(12),
		X264(16),
		// x264 CBR 10 bits
		X264TenBits(1),
		X264TenBits(4),
		X264TenBits(8),
		X264TenBits(12),
		X264TenBits(16),
		// x264 VBR 2 pass
		X264TwoPass(8),
		X264TwoPass(10),
		X264TwoPass(12),
		X264TwoPass(20),
		// repacks
		RepackMkvMp4(),
		// video2x interpolation
		Video2xRife(2),
		Video2xRife(3),
		Video2xRife(4),
		Video2xRife(5),
		Video2xRife(6),
		// video2x upscale
		Video2xUpscale4x(),
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

func (p Preset) OutputPath(inputPath string) string {
	base := pathBaseName(inputPath)
	return base + "_" + p.Suffix + "." + p.Ext
}

func pathBaseName(inputPath string) string {
	ext := filepath.Ext(inputPath)
	base := strings.TrimSuffix(inputPath, ext)
	return base
}

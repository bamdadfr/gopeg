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
	Args        func(inputPath string, outputPath string) [][]string
}

func List() []Preset {
	presets := []Preset{
		// ffv1
		ffv1(),
		ffv1Loop(3),
		ffv1Loop(4),
		// x264 CBR
		x264(1),
		x264(4),
		x264(8),
		x264(10),
		x264(12),
		x264(16),
		// x264 CBR 10 bits
		x264TenBits(1),
		x264TenBits(4),
		x264TenBits(8),
		x264TenBits(12),
		x264TenBits(16),
		// x264 VBR 2 pass
		x264TwoPass(8),
		x264TwoPass(10),
		x264TwoPass(12),
		x264TwoPass(20),
		// repacks
		repackMkvMp4(),
		// instagram
		instagramV1Vertical(),
		instagramV2Vertical(),
		instagramV2Square(),
		instagramV3Vbr1PremiereVertical(),
		instagramV3Vbr1PremiereSquare(),
		// video2x interpolation
		video2xRife(2),
		video2xRife(3),
		video2xRife(4),
		video2xRife(5),
		video2xRife(6),
		// video2x upscale
		video2xUpscale4x(),
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

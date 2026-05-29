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
		// archives lossless
		ArchiveFfv1(),
		ArchiveFfv1Loop3(),
		ArchiveFfv1Loop4(),
		// archives lossy
		ArchiveX264Crf4(),
		ArchiveX264Crf8(),
		ArchiveX264Crf10(),
		ArchiveX264Crf12(),
		ArchiveX264Crf16(),
		// repacks
		RepackMkvMp4(),
		// instagram
		InstagramV1Vertical(),
		InstagramV2Vertical(),
		InstagramV2Square(),
		InstagramV3Vbr1PremiereVertical(),
		InstagramV3Vbr1PremiereSquare(),
		// video2x interpolation
		Video2xRife4x(),
		Video2xRife5x(),
		Video2xRife6x(),
		// video2x upscale
		Video2xUpscale4x(),
		Video2xUpscale5x(),
		Video2xUpscale6x(),
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

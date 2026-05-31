// Package pipeline defines sequential chains of presets
package pipeline

import (
	"gopeg/preset"
	"path/filepath"
	"strings"
)

type Pipeline struct {
	Name        string
	Description string
	Steps       []preset.Preset
}

func List() []Pipeline {
	return []Pipeline{
		InstagramSquare8M(),
		Upscale(),
	}
}

// IntermediatePath returns the output path for a non-final step.
// It lives next to the original input so cleanup is straightforward.
func IntermediatePath(inputPath string, step preset.Preset) string {
	dir := filepath.Dir(inputPath)
	base := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
	return filepath.Join(dir, base+"_tmp_"+step.Suffix+"."+step.Ext)
}

// FinalPath returns the output path for the last step, using a
// pipeline-level suffix so it doesn't look like a single-preset output.
func (p Pipeline) FinalPath(inputPath string) string {
	last := p.Steps[len(p.Steps)-1]
	dir := filepath.Dir(inputPath)
	base := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
	return filepath.Join(dir, base+"_"+p.Suffix()+"."+last.Ext)
}

// Suffix joins step suffixes with underscores for the final output name.
func (p Pipeline) Suffix() string {
	parts := make([]string, len(p.Steps))
	for i, s := range p.Steps {
		parts[i] = s.Suffix
	}
	return "pipe_" + strings.Join(parts, "_")
}

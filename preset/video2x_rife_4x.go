package preset

import (
	"gopeg/binary"
)

func Video2xRife4x() Preset {
	return Preset{
		Name:        "Video2X RIFE 4x",
		Binary:      binary.Video2x(),
		Description: "Frame interpolation. ProRes 4444 container with 4:2:2 12bits stream.",
		Accept:      []string{},
		Ext:         "mov",
		Suffix:      "rife_4x",
		Args: func(inputPath string, outputPath string) [][]string {
			return [][]string{
				video2xRifeArgs(inputPath, outputPath, 4),
			}
		},
	}
}

package preset

import "gopeg/binary"

func Video2xRife6x() Preset {
	return Preset{
		Name:        "Video2X RIFE 6x",
		Binary:      binary.Video2x(),
		Description: "Frame interpolation. ProRes 4444 container with 4:2:2 12bits stream.",
		Accept:      []string{},
		Ext:         "mov",
		Suffix:      "rife_6x",
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-i", inputPath,
				"-o", outputPath,
				"-p", "rife",
				"-m", "6",
				"-c", "prores_ks",
				"-e", "profile=4",
			}
		},
	}
}

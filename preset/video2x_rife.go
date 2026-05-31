package preset

import (
	"gopeg/binary"
	"strconv"
)

func video2xRife(factor int) Preset {
	factorString := strconv.Itoa(factor) + "x"

	return Preset{
		Name:        "Video2X RIFE " + factorString,
		Binary:      binary.Video2x(),
		Description: "Frame interpolation. ProRes 4444 container with 4:2:2 12bits stream.",
		Accept:      []string{},
		Ext:         "mov",
		Suffix:      "rife_" + factorString,
		Args: func(inputPath string, outputPath string) [][]string {
			return [][]string{
				{
					"-i", inputPath,
					"-o", outputPath,
					"-p", "rife",
					"-m", strconv.Itoa(factor),
					"-c", "prores_ks",
					"-e", "profile=4",
				},
			}
		},
	}
}

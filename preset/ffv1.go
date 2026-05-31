package preset

import "gopeg/binary"

func Ffv1() Preset {
	return Preset{
		Name:        "FFV1",
		Binary:      binary.Ffmpeg(),
		Description: "Lossless.",
		Accept:      []string{},
		Ext:         "mkv",
		Suffix:      "ffv1",
		Args: func(inputPath string, outputPath string) [][]string {
			return [][]string{{
				"-i", inputPath,
				"-c:v", "ffv1",
				outputPath,
			}}
		},
	}
}

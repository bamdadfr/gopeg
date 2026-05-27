package preset

import "gopeg/binary"

func Interpolate4x() Preset {
	return Preset{
		Name:        "RIFE interpolate 4x",
		Binary:      binary.Video2x(),
		Description: "Frame interpolation via video2x RIFE.",
		Accept:      []string{},
		Ext:         "mov",
		Suffix:      "rife_4x",
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-i", inputPath,
				"-o", outputPath,
				"-p", "rife",
				"-m", "4",
				"-c", "prores_ks",
				"-e", "profile=4",
			}
		},
	}
}

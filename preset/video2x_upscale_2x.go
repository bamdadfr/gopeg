package preset

import "gopeg/binary"

func Video2xUpscale2x() Preset {
	return Preset{
		Name:        "Video2X upscale 2x",
		Binary:      binary.Video2x(),
		Description: "Real ESRGAN. ProRes 4444 container with 4:2:2 12bits stream.",
		Accept:      []string{},
		Ext:         "mov",
		Suffix:      "esrgan_2x",
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-i", inputPath,
				"-o", outputPath,
				"-p", "realesrgan",
				"--realesrgan-model", "realesrgan-plus",
				"-s", "2",
				"-c", "prores_ks",
				"-e", "profile=4",
			}
		},
	}
}

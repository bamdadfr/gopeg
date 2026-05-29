package preset

import "gopeg/binary"

func Video2xUpscale4x() Preset {
	return Preset{
		Name:        "Video2X upscale 4x",
		Binary:      binary.Video2x(),
		Description: "Real ESRGAN. ProRes 4444 container with 4:2:2 12bits stream.",
		Accept:      []string{},
		Ext:         "mov",
		Suffix:      "esrgan_4x",
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-i", inputPath,
				"-o", outputPath,
				"-p", "realesrgan",
				"--realesrgan-model", "realesrgan-plus",
				"-s", "4",
				"-c", "prores_ks",
				"-e", "profile=4",
			}
		},
	}
}

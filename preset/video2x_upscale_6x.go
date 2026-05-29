package preset

import "gopeg/binary"

func Video2xUpscale6x() Preset {
	return Preset{
		Name:        "Video2X upscale 6x",
		Binary:      binary.Video2x(),
		Description: "Real ESRGAN. ProRes 4444 container with 4:2:2 12bits stream.",
		Accept:      []string{},
		Ext:         "mov",
		Suffix:      "esrgan_6x",
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-i", inputPath,
				"-o", outputPath,
				"-p", "realesrgan",
				"--realesrgan-model", "realesrgan-plus",
				"-s", "6",
				"-c", "prores_ks",
				"-e", "profile=4",
			}
		},
	}
}

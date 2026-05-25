package preset

func Upscale4x() Preset {
	return Preset{
		Name:        "Upscale 4x",
		Description: "video2x real ESRGAN uspcaler",
		Accept:      []string{},
		OutputPath: func(inputPath string) string {
			base := pathBaseName(inputPath)
			return base + "_upscale_4x.mov"
		},
		Args: func(inputPath string, outputPath string) []string {
			// video2x -i output_lossless.mkv -o output.mov -p realesrgan --realesrgan-model realesrgan-plus -s 4 -c prores_ks -e profile=4
			return []string{"video2x", "-i", inputPath, "-o", outputPath, "-p", "realesrgan", "--realesrgan-model", "realesrgan-plus", "-s", "4", "-c", "prores_ks", "-e", "profile=4"}
		},
	}
}

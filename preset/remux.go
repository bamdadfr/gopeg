package preset

func Remux() Preset {
	return Preset{
		Name:        "Remux MKV → MP4",
		Description: "Convert Matroska to MP4. Useful for video2x outputs. Allows easy consumption in Adobe After Effects.",
		Accept:      []string{".mkv"},
		OutputPath: func(inputPath string) string {
			base := pathBaseName(inputPath)
			return base + "_remux.mp4"
		},
		Args: func(inputPath string, outputPath string) []string {
			return []string{"ffmpeg", "-i", inputPath, "-c", "copy", outputPath}
		},
	}
}

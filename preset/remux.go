package preset

func Remux() Preset {
	return Preset{
		Name:        "Remux MKV → MP4",
		Description: "Convert Matroska to MP4. Useful for video2x outputs. Allows easy consumption in Adobe After Effects.",
		Accept:      []string{".mkv"},
		Target:      ".mp4",
		BuildArgs: func(inputPath string, outputPath string) []string {
			return []string{"-i", inputPath, "-c", "copy", outputPath}
		},
	}
}

package preset

func Interpolate4x() Preset {
	return Preset{
		Name:        "RIFE interpolate 4x",
		Description: "Frame interpolation via video2x RIFE.",
		Accept:      []string{},
		OutputPath: func(inputPath string) string {
			base := pathBaseName(inputPath)
			return base + "_rife_4x.mov"
		},
		Args: func(inputPath string, outputPath string) []string {
			return []string{"video2x", "-i", inputPath, "-o", outputPath, "-p", "rife", "-m", "4", "-c", "prores_ks", "-e", "profile=4"}
		},
	}
}

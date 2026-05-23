package preset

func Archive() Preset {
	return Preset{
		Name:        "Archive",
		Description: "Near losless quality. H.264 codec. CRF 4. Maximum quality. Slow encoding.",
		Accept:      []string{},
		Target:      ".mp4",
		BuildArgs: func(inputPath string, outputPath string) []string {
			return []string{"-i", inputPath, "-c:v", "libx264", "-crf", "4", "-maxrate", "240M", "-bufsize", "250M", "-pix_fmt", "yuv420p", "-preset", "placebo", "-g", "30", "-profile:v", "high", "-level", "4.1", "-movflags", "+faststart", "-c:a", "aac", "-ar", "48000", "-b:a", "384k", outputPath}
		},
	}
}

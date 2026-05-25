package preset

func DecodeAndLoop() Preset {
	return Preset{
		Name:        "Decode & loop",
		Description: "Decode to independent frames. Loop the whole video 3 times.",
		Accept:      []string{},
		OutputPath: func(inputPath string) string {
			base := pathBaseName(inputPath)
			return base + "_decode_loop.mkv"
		},
		Args: func(inputPath string, outputPath string) []string {
			// example
			// ffmpeg -stream_loop 3 -i input.mp4 -c:v ffv1 output_lossless.mkv
			return []string{"ffmpeg", "-stream_loop", "3", "-i", inputPath, "-c:v", "ffv1", outputPath}
		},
	}
}

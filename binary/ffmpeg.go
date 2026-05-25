package binary

func Ffmpeg() Binary {
	return Binary{
		Name:    "ffmpeg",
		Command: "ffmpeg",
		Path: BinaryPath{
			Linux: "ffmpeg",
			// we assume ffmpeg is in PATH, installed through choco
			// TODO: this could be configured through dedicated yaml file, later
			Windows: "ffmpeg",
		},
		OverwriteFlag: "-y",
	}
}

package binary

func Ffmpeg() Binary {
	return Binary{
		Name:          "ffmpeg",
		OverwriteFlag: "-y",
		Candidates: []string{
			"ffmpeg",
		},
	}
}

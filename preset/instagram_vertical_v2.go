package preset

import "gopeg/binary"

func InstagramVerticalV2() Preset {
	return Preset{
		Name:        "Instagram Vertical v2",
		Binary:      binary.Ffmpeg(),
		Description: "Compressing for instagram 9:16 (1080p lanczos 20M/25M)",
		Accept:      []string{},
		OutputPath: func(inputPath string) string {
			base := pathBaseName(inputPath)
			return base + "_instagram_vertical_v1.mp4"
		},
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-i", inputPath,
				"-vf", "scale=1080:1920:flags=lanczos",
				"-c:v", "libx264",
				"-crf", "4",
				"-maxrate", "20M",
				"-bufsize", "25M",
				"-pix_fmt", "yuv420p",
				"-preset", "placebo",
				"-g", "30",
				"-movflags",
				"+faststart",
				"-c:a", "aac",
				"-b:a", "320k",
				outputPath,
			}
		},
	}
}

package preset

import "gopeg/binary"

func InstagramVerticalV2() Preset {
	return Preset{
		Name:        "Instagram Vertical v2",
		Binary:      binary.Ffmpeg(),
		Description: "320kbps AAC",
		Accept:      []string{},
		Ext:         "mp4",
		Suffix:      "instagram_vertical_v2",
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

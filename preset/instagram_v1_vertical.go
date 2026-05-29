package preset

import "gopeg/binary"

func InstagramV1Vertical() Preset {
	return Preset{
		Name:        "Instagram v1 vertical",
		Binary:      binary.Ffmpeg(),
		Description: "1080p lanczos 20M/25M CBR 4",
		Accept:      []string{},
		Ext:         "mp4",
		Suffix:      "instagram_v1_vertical",
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
				"-b:a", "128k",
				outputPath,
			}
		},
	}
}

package preset

import "gopeg/binary"

func instagramV1Vertical() Preset {
	return Preset{
		Name:        "Instagram v1 CBR Vertical",
		Binary:      binary.Ffmpeg(),
		Description: "1080p, lanczos, CBR 20M/25M, CRF 4, AAC 128k",
		Accept:      []string{},
		Ext:         "mp4",
		Suffix:      "instagram_v1_vertical",
		Args: func(inputPath string, outputPath string) [][]string {
			return [][]string{{
				"-i", inputPath,
				"-vf", "scale=1080:1920:flags=lanczos",
				"-c:v", "libx264",
				"-crf", "4",
				"-maxrate", "20M",
				"-bufsize", "25M",
				"-pix_fmt", "yuv420p",
				"-preset", "veryslow",
				"-g", "30",
				"-movflags",
				"+faststart",
				"-c:a", "aac",
				"-b:a", "128k",
				outputPath,
			}}
		},
	}
}

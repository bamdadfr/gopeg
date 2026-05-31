package preset

import "gopeg/binary"

func InstagramV2Square() Preset {
	return Preset{
		Name:        "Instagram v2 CBR Square",
		Binary:      binary.Ffmpeg(),
		Description: "v1 + 320kbps AAC",
		Accept:      []string{},
		Ext:         "mp4",
		Suffix:      "instagram_v2_square",
		Args: func(inputPath string, outputPath string) [][]string {
			return [][]string{{
				"-i", inputPath,
				"-vf", "scale=1080:1080:flags=lanczos",
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
				"-b:a", "320k",
				outputPath,
			}}
		},
	}
}

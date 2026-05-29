package preset

import "gopeg/binary"

func InstagramV3Vbr1PremiereSquare() Preset {
	return Preset{
		Name:        "Instagram v3 VBR1 Premiere Square",
		Binary:      binary.Ffmpeg(),
		Description: "Match Source - Adaptive High Bitrate.",
		Accept:      []string{},
		Ext:         "mp4",
		Suffix:      "instagram_v3_vbr1_premiere_square",
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-i", inputPath,
				"-vf", "scale=1080:1080:flags=lanczos",
				"-c:v", "libx264",
				"-profile:v", "main",
				"-level:v", "4.1",
				"-pix_fmt", "yuv420p",
				"-color_primaries", "bt709",
				"-color_trc", "bt709",
				"-colorspace", "bt709",
				"-r", "30", // lock 30 fps
				"-g", "60", // every 2 seconds
				"-b:v", "20M",
				"-maxrate", "24M",
				"-bufsize", "48M", // convention
				"-preset", "veryslow",
				"-movflags", "+faststart",
				"-c:a", "aac",
				"-b:a", "320k",
				"-ar", "48000", // force 38kHz
				"-ac", "2",
				outputPath,
			}
		},
	}
}

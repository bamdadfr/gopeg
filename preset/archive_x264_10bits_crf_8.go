package preset

import "gopeg/binary"

func ArchiveX26410bitsCrf8() Preset {
	return Preset{
		Name:        "Archive x264 10 bits CRF 8",
		Binary:      binary.Ffmpeg(),
		Description: "Preserves gradients.",
		Accept:      []string{},
		Ext:         "mp4",
		Suffix:      "archive_x264_10bits_crf_8",
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-i", inputPath,
				"-c:v", "libx264",
				"-crf", "8",
				"-maxrate", "240M",
				"-bufsize", "250M",
				"-pix_fmt", "yuv420p10le",
				"-preset", "veryslow",
				"-g", "30",
				"-profile:v", "high10",
				"-level", "5.1",
				"-movflags",
				"+faststart",
				"-c:a", "aac",
				"-ar", "48000",
				"-b:a", "384k",
				outputPath,
			}
		},
	}
}

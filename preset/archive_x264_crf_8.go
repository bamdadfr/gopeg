package preset

import "gopeg/binary"

func ArchiveX264Crf8() Preset {
	return Preset{
		Name:        "Archive x264 CRF 8",
		Binary:      binary.Ffmpeg(),
		Description: "Visually seamless. ~11x from ProRes 4444. Use this for intermediary generative files.",
		Accept:      []string{},
		Ext:         "mp4",
		Suffix:      "archive_x264_crf_8",
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-i", inputPath,
				"-c:v", "libx264",
				"-crf", "8",
				"-maxrate", "240M",
				"-bufsize", "250M",
				"-pix_fmt", "yuv420p",
				"-preset", "veryslow",
				"-g", "30",
				"-profile:v", "high",
				"-level", "4.1",
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

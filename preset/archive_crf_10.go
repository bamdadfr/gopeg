package preset

import "gopeg/binary"

func ArchiveCrf10() Preset {
	return Preset{
		Name:        "Archive CRF 10",
		Binary:      binary.Ffmpeg(),
		Description: "h264. Very light softening. ~15x from ProRes 4444.",
		Accept:      []string{},
		Ext:         "mp4",
		Suffix:      "archive_crf_10",
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-i", inputPath,
				"-c:v", "libx264",
				"-crf", "10",
				"-maxrate", "240M",
				"-bufsize", "250M",
				"-pix_fmt", "yuv420p",
				"-preset", "placebo",
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

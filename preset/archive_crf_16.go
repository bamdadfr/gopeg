package preset

import "gopeg/binary"

func ArchiveCrf16() Preset {
	return Preset{
		Name:        "Archive CRF 16",
		Binary:      binary.Ffmpeg(),
		Description: "h264. Light detail smearing. ~32x from ProRes 4444. Use this for terminal files.",
		Accept:      []string{},
		Ext:         "mp4",
		Suffix:      "archive_crf_16",
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-i", inputPath,
				"-c:v", "libx264",
				"-crf", "16",
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

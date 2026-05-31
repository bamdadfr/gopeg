package preset

import "gopeg/binary"

func ArchiveX264Crf12() Preset {
	return Preset{
		Name:        "Archive x264 CRF 12",
		Binary:      binary.Ffmpeg(),
		Description: "Very light softening and tiny detail smearing. ~20x from ProRes 4444.",
		Accept:      []string{},
		Ext:         "mp4",
		Suffix:      "archive_x264_crf_12",
		Args: func(inputPath string, outputPath string) [][]string {
			return [][]string{{
				"-i", inputPath,
				"-c:v", "libx264",
				"-crf", "12",
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
			}}
		},
	}
}

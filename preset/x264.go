package preset

import (
	"gopeg/binary"
	"strconv"
)

func X264(crf int) Preset {
	crfSuffix := strconv.Itoa(crf)

	return Preset{
		Name:        "x264 CRF " + crfSuffix,
		Binary:      binary.Ffmpeg(),
		Description: "",
		Accept:      []string{},
		Ext:         "mp4",
		Suffix:      "x264_crf_" + crfSuffix,
		Args: func(inputPath string, outputPath string) [][]string {
			return [][]string{
				{
					"-i", inputPath,
					"-c:v", "libx264",
					"-crf", strconv.Itoa(crf),
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
				},
			}
		},
	}
}

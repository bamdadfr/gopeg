package preset

import (
	"gopeg/binary"
	"strconv"
)

func X264TenBits(crf int) Preset {
	crfString := strconv.Itoa(crf)
	desc := "10 bits. Constant bitrate."

	if crf == 8 {
		desc = desc + " Preferred preset for archives and intermediaries"
	}

	return Preset{
		Name:        "x264 10 bits CRF " + crfString,
		Binary:      binary.Ffmpeg(),
		Description: desc,
		Accept:      []string{},
		Ext:         "mp4",
		Suffix:      "x264_10bits_crf_" + crfString,
		Args: func(inputPath string, outputPath string) [][]string {
			return [][]string{
				{
					"-i", inputPath,
					"-c:v", "libx264",
					"-crf", strconv.Itoa(crf),
					"-maxrate", "240M",
					"-bufsize", "250M",
					"-pix_fmt", "yuv420p10le",
					"-preset", "veryslow",
					"-g", "30",
					"-profile:v", "high10",
					"-level", "5.1",
					"-movflags",
					"+faststart",
					"-c:a", "alac",
					outputPath,
				},
			}
		},
	}
}

package preset

import (
	"gopeg/binary"
	"gopeg/env"
	"strconv"
)

func x264TwoPass(rate int) Preset {
	rateSuffix := strconv.Itoa(rate) + "M"

	return Preset{
		Name:        "x264 Two-Pass " + rateSuffix,
		Binary:      binary.Ffmpeg(),
		Description: "30 fps lock",
		Ext:         "mp4",
		Suffix:      "x264_2pass_" + rateSuffix,
		Args: func(inputPath string, outputPath string) [][]string {
			passLog := pathBaseName(inputPath)
			maxRate := rate + 4
			bufferSize := maxRate * 2
			fps := 30

			common := []string{
				"-c:v", "libx264",
				"-profile:v", "main",
				"-level:v", "4.1",
				"-pix_fmt", "yuv420p",
				"-color_primaries", "bt709",
				"-color_trc", "bt709",
				"-colorspace", "bt709",
				"-r", strconv.Itoa(fps),
				"-g", strconv.Itoa(fps * 2),
				"-b:v", strconv.Itoa(rate) + "M",
				"-maxrate", strconv.Itoa(maxRate) + "M",
				"-bufsize", strconv.Itoa(bufferSize) + "M",
				"-preset", "veryslow",
			}

			pass1 := append([]string{"-i", inputPath}, common...)

			pass1 = append(pass1,
				"-pass", "1",
				"-passlogfile", passLog,
				"-an",
				"-f", "null",
				env.NullDevice(),
			)

			pass2 := append([]string{"-i", inputPath}, common...)

			pass2 = append(pass2,
				"-pass", "2",
				"-passlogfile", passLog,
				"-movflags", "+faststart",
				"-c:a", "aac",
				"-b:a", "320k",
				"-ar", "48000",
				"-ac", "2",
				outputPath,
			)

			return [][]string{pass1, pass2}
		},
	}
}

package preset

import (
	"strconv"

	"gopeg/binary"
)

func Ffv1Rescale(width int, height int, fps int) Preset {
	w := strconv.Itoa(width)
	h := strconv.Itoa(height)
	rescaleString := w + "x" + h
	filterString := "scale=" + w + ":" + h + ":flags=lanczos"

	return Preset{
		Name:        "FFV1 " + rescaleString,
		Binary:      binary.Ffmpeg(),
		Description: "Lossless. Rescale.",
		Accept:      []string{},
		Ext:         "mkv",
		Suffix:      "ffv1_rescale_" + rescaleString,
		Args: func(inputPath string, outputPath string) [][]string {
			return [][]string{{
				"-i", inputPath,
				"-filter:v", filterString,
				"-c:v", "ffv1",
				"-level", "3",
				"-slicecrc", "1",
				"-pix_fmt", "yuv420p",
				"-c:a", "pcm_s24le",
				"-r", strconv.Itoa(fps),
				"-shortest",
				outputPath,
			}}
		},
	}
}

package preset

import (
	"gopeg/binary"
	"strconv"
)

func ffv1Loop(loop int) Preset {
	loopString := strconv.Itoa(loop)

	return Preset{
		Name:        "FFV1 & Loop " + loopString,
		Binary:      binary.Ffmpeg(),
		Description: "Lossless. " + loopString + " total loops.",
		Accept:      []string{},
		Ext:         "mkv",
		Suffix:      "ffv1_loop_" + loopString,
		Args: func(inputPath string, outputPath string) [][]string {
			renderingLoops := loop - 1 // because first iteration is implied as first loop by ffmpeg

			return [][]string{
				{
					"-stream_loop", strconv.Itoa(renderingLoops),
					"-i", inputPath,
					"-c:v", "ffv1",
					outputPath,
				},
			}
		},
	}
}

package preset

import "gopeg/binary"

func DecodeAndLoop() Preset {
	return Preset{
		Name:        "Decode & loop",
		Binary:      binary.Ffmpeg(),
		Description: "Decode to independent frames. Loop the whole video 3 times.",
		Accept:      []string{},
		Ext:         "mkv",
		Suffix:      "decode_loop",
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-stream_loop", "3",
				"-i", inputPath,
				"-c:v", "ffv1",
				outputPath,
			}
		},
	}
}

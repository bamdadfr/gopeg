package preset

import "gopeg/binary"

func Decode() Preset {
	return Preset{
		Name:        "Decode",
		Binary:      binary.Ffmpeg(),
		Description: "Decode to independent frames. FFV1 codec.",
		Accept:      []string{},
		OutputPath: func(inputPath string) string {
			base := pathBaseName(inputPath)
			return base + "_decode.mkv"
		},
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-i", inputPath,
				"-c:v", "ffv1",
				outputPath,
			}
		},
	}
}

package preset

import "gopeg/binary"

func ArchiveFfv1Loop3() Preset {
	return Preset{
		Name:        "Archive FFV1 & Loop 3",
		Binary:      binary.Ffmpeg(),
		Description: "Lossless. 3 total loops.",
		Accept:      []string{},
		Ext:         "mkv",
		Suffix:      "archive_ffv1_loop_3",
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-stream_loop", "2", // adds to input (total - 1)
				"-i", inputPath,
				"-c:v", "ffv1",
				outputPath,
			}
		},
	}
}

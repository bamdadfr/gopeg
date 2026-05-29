package preset

import "gopeg/binary"

func ArchiveFfv1Loop4() Preset {
	return Preset{
		Name:        "Archive FFV1 & Loop 4",
		Binary:      binary.Ffmpeg(),
		Description: "Lossless. 4 total loops.",
		Accept:      []string{},
		Ext:         "mkv",
		Suffix:      "archive_ffv1_loop_4",
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-stream_loop", "3", // adds to input
				"-i", inputPath,
				"-c:v", "ffv1",
				outputPath,
			}
		},
	}
}

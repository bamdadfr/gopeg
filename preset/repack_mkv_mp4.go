package preset

import "gopeg/binary"

func RepackMkvMp4() Preset {
	return Preset{
		Name:        "Repack MKV to MP4",
		Binary:      binary.Ffmpeg(),
		Description: "Smoother compatibility for Adobe After Effects",
		Accept:      []string{".mkv"},
		Ext:         "mp4",
		Suffix:      "repack_from_mkv",
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-i", inputPath,
				"-c", "copy",
				outputPath,
			}
		},
	}
}

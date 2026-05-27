package preset

import "gopeg/binary"

func Remux() Preset {
	return Preset{
		Name:        "Remux MKV → MP4",
		Binary:      binary.Ffmpeg(),
		Description: "Convert Matroska to MP4. Useful for video2x outputs. Allows easy consumption in Adobe After Effects.",
		Accept:      []string{".mkv"},
		Ext:         "mp4",
		Suffix:      "remux",
		Args: func(inputPath string, outputPath string) []string {
			return []string{
				"-i", inputPath,
				"-c", "copy",
				outputPath,
			}
		},
	}
}

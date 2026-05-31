package preset

import (
	"gopeg/binary"
	"gopeg/env"
)

func ArchiveX264TwoPass() Preset {
	return Preset{
		Name:        "Archive x264 Two-Pass 10M",
		Binary:      binary.Ffmpeg(),
		Description: "Two-pass ABR targeting 10 Mbps.",
		Ext:         "mp4",
		Suffix:      "x264_2pass_10m",
		Args: func(inputPath string, outputPath string) [][]string {
			return [][]string{
				{
					"-i", inputPath,
					"-c:v", "libx264",
					"-b:v", "10M",
					"-pass", "1",
					"-passlogfile", pathBaseName(inputPath),
					"-an",
					"-f", "null",
					env.NullDevice(),
				},
				{
					"-i", inputPath,
					"-c:v", "libx264",
					"-b:v", "10M",
					"-pass", "2",
					"-passlogfile", pathBaseName(inputPath),
					"-c:a", "aac",
					"-b:a", "384k",
					"-movflags", "+faststart",
					outputPath,
				},
			}
		},
	}
}

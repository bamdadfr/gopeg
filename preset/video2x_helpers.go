package preset

import "strconv"

func video2xRifeArgs(inputPath string, outputPath string, factor int) []string {
	return []string{
		"-i", inputPath,
		"-o", outputPath,
		"-p", "rife",
		"-m", strconv.Itoa(factor),
		"-c", "prores_ks",
		"-e", "profile=4",
	}
}

package main

func video2xExec() string {
	return "video2x"
}

func video2xPath() string {
	if isWindows {
		return "C:\\Program Files\\Video2X Qt6\\video2x.exe"
	}

	return video2xExec()
}

func ffmpegExec() string {
	return "ffmpeg"
}

func ffmpegPath() string {
	return ffmpegExec()
}

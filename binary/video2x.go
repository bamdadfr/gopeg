package binary

func Video2x() Binary {
	return Binary{
		Name:    "video2x",
		Command: "video2x",
		Path: BinaryPath{
			Linux:   "video2x",
			Windows: "C:\\Program Files\\Video2X Qt6\\video2x.exe",
		},
	}
}

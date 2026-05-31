package pipeline

import "gopeg/preset"

func InstagramSquare8M() Pipeline {
	return Pipeline{
		Name:        "Instagram Square 8M",
		Description: "FFV1 1080x1080 30 fps → x264 VBR2 8M",
		Steps: []preset.Preset{
			preset.Ffv1Rescale(1080, 1080, 30),
			preset.X264TwoPass(8),
		},
	}
}

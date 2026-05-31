package pipeline

import "gopeg/preset"

func Upscale() Pipeline {
	return Pipeline{
		Name:        "Upscale",
		Description: "From AI generations",
		Steps: []preset.Preset{
			preset.Ffv1Loop(3),
			preset.Video2xUpscale4x(),
			preset.Video2xRife(4),
			preset.X264TenBits(8),
		},
	}
}

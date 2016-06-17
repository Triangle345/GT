package Opengl

import (
	"github.com/go-gl/gl/v3.2-core/gl"
)

type CardSpecs struct {
	MaxTextureSize int32
}

// Probe probes the video card and return a struct of specifications for it
func Probe() CardSpecs {
	cs := CardSpecs{}
	gl.GetIntegerv(gl.MAX_TEXTURE_SIZE, &cs.MaxTextureSize)

	return cs
}

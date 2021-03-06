package Opengl

import "github.com/go-gl/gl/v3.3-core/gl"

type CardSpecs struct {
	MaxTextureSize       int32
	MaxTextureImageUnits int32
	ShaderVersion        string
}

// Probe probes the video card and return a struct of specifications for it
func Probe() CardSpecs {
	cs := CardSpecs{}
	gl.GetIntegerv(gl.MAX_TEXTURE_SIZE, &cs.MaxTextureSize)
	gl.GetIntegerv(gl.MAX_TEXTURE_IMAGE_UNITS, &cs.MaxTextureImageUnits)
	cs.ShaderVersion = gl.GoStr((gl.GetString(gl.SHADING_LANGUAGE_VERSION)))

	return cs
}

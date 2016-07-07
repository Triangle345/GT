package main

import (
	"GT"
	"GT/Graphics/Opengl"
	"GT/Logging"
)

func main() {
	// create the foundation: a new window, basic scene, and game component (node)
	GT.EngineStart()
	probe := Opengl.Probe()
	Logging.Info("Max Texture Size: ", probe.MaxTextureSize)
	Logging.Info("Max Texture Image Units: ", probe.MaxTextureImageUnits)
	Logging.Info("Shader Version: ", probe.ShaderVersion)
}

package main

import (
	"GT/Graphics"
	"GT/Graphics/Opengl"
	"GT/Logging"
	"GT/Window"
)

func main() {
	// create the foundation: a new window, basic scene, and game component (node)
	simpleWindow := Window.NewWindowedWindow("SimpleWindowTitle", 600, 800)
	Graphics.NewBasicScene(&simpleWindow)
	probe := Opengl.Probe()
	Logging.Info("Max Texture Size: ", probe.MaxTextureSize)
}

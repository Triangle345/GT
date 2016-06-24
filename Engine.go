package GT

import (
	"GT/Graphics/Font"
	"GT/Graphics/Image"
	"GT/Graphics/Opengl"
	"GT/Logging"
	"GT/Window"
)

// EngineStart initializes everything in order within the engine. Should be called first
func EngineStart() {
	Logging.Info("Engine initializing...")
	Window.Start()
	Opengl.Start()
	Font.LoadFonts("../Assets/Fonts")
	Image.Start()
	Opengl.CreateBuffers()
	Logging.Info("Engine initialization finished.")
}

func main() {

}

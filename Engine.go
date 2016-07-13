package GT

import (
	"GT/Graphics/Font"
	"GT/Graphics/Image"
	"GT/Graphics/Opengl"
	"GT/Logging"
	"GT/Window"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	Assets       = flag.String("Assets", "./Assets", "Location of Assets folder")
	OGL_VERSION  = flag.String("OGL_VERSION", "130", "OpenGL version used. The actual version in shaders")
	AssetsFonts  string
	AssetsImages string
	AssetsModels string
)

func getAbsPath(folder string) string {
	path, _ := filepath.Abs(*Assets + folder)

	if _, err := os.Stat(path); err != nil {
		fmt.Println("Please specify the correct Assets path with the -Assets flag")
		panic(err)
	}

	return path
}

func readFlags() {
	flag.Parse()

	// eventually add log flags (filtering, naming, etc.)
	logPath, _ := filepath.Abs(*Assets + "/Logs")
	logPath += string(os.PathSeparator)

	Logging.Init(logPath, log.Ldate|log.Ltime|log.Lshortfile)

	imgsPath := getAbsPath("/Images")

	fontsPath := getAbsPath("/Fonts")

	modelsPath := getAbsPath("/Models")

	AssetsFonts = fontsPath
	AssetsImages = imgsPath + string(os.PathSeparator)
	AssetsModels = modelsPath + string(os.PathSeparator)

	Logging.Info("Assets Path: ", *Assets)
	Logging.Info("Images Path: ", AssetsImages)
	Logging.Info("Fonts Path: ", AssetsFonts)
	Logging.Info("Models Path: ", AssetsModels)
	Logging.Info("OGL Version Used: ", *OGL_VERSION)
}

// EngineStart initializes everything in order within the engine. Should be called first
func EngineStart() {

	readFlags()

	Window.Start()
	Opengl.OGL_VERSION = *OGL_VERSION
	Opengl.Start()

	Font.LoadFonts(AssetsFonts)
	Image.LoadImages(AssetsImages)
	Opengl.CreateBuffers()
	Logging.Info("Engine initialization finished.")
	// Image.AggrImg.Print("aggr.png")
}

func main() {

}

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
	AssetsFonts  string
	AssetsImages string
)

func readFlags() {
	flag.Parse()

	// eventually add log flags (filtering, naming, etc.)
	logPath, _ := filepath.Abs(*Assets + "/Logs")
	logPath += string(os.PathSeparator)
	// logPath += inputFileNameFlag
	Logging.Init(logPath, log.Ldate|log.Ltime|log.Lshortfile)

	imgsPath, _ := filepath.Abs(*Assets + "/Images")

	if _, err := os.Stat(imgsPath); err != nil {
		fmt.Println("Please specify the correct Assets path with the -Assets flag")
		panic(err)
	}

	fontsPath, _ := filepath.Abs(*Assets + "/Fonts")
	if _, err := os.Stat(fontsPath); err != nil {

		fmt.Println("Please specify the correct Assets path with the -Assets flag")
		panic(err)
	}

	AssetsFonts = fontsPath
	AssetsImages = imgsPath + string(os.PathSeparator)

	Logging.Info("Assets Path: ", *Assets)
	Logging.Info("Images Path: ", AssetsImages)
	Logging.Info("Fonts Path: ", AssetsFonts)
}

// EngineStart initializes everything in order within the engine. Should be called first
func EngineStart() {

	readFlags()

	Window.Start()
	Opengl.Start()

	Font.LoadFonts(AssetsFonts)
	Image.LoadImages(AssetsImages)
	Opengl.CreateBuffers()
	Logging.Info("Engine initialization finished.")
}

func main() {

}

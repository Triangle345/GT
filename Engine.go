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
	// logPath += inputFileNameFlag
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
	// Image.AggrImg.Print("aggr.png")
}

func main() {

}

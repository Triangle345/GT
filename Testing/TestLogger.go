package main

import (
	"GT/Logging"
	//"fmt"
)

func main() {

	// TODO: test filtering and simultaneous outputs once implemented and configured
	Logging.Debug("debug")
	Logging.Info("info")

	// Note: this has been tested to work on multiple code call levels
	// so inserting debug logs in SimpleWindow.go as well as Scene.go
	// will place all info into GT/Logging/default.log
}

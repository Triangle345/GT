package main

import (
	"GT"
	"GT/Logging"
)

func main() {

	// The logger now initializes in the Engine's start. That is where flags get parsed etc.
	GT.EngineStart()

	// TODO: test filtering and simultaneous outputs once implemented and configured
	Logging.Debug("debug")
	Logging.Info("info")

	// Note: this has been tested to work on multiple code call levels
	// so inserting debug logs in SimpleWindow.go as well as Scene.go
	// will place all info into GT/Logging/default.log
}

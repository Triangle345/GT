// Logger.go contains simple logic to wrap around golang's native logger in order to provide multi-logging and (eventually?) filtered logging

package Logging

import (
	"log"
	"os"
	"path/filepath"
)

// make the loggers global to the class for easy use (no need to pass stuff around)
var logfile *os.File
var flogger *log.Logger
var stdlogger *log.Logger

// Note: this can be moved elsewhere but risks being coupled (i.e. Engine.go)
func init() {
	// TODO maybe add log flags (filtering, naming, etc.)
	// for flagging may want / need to abstract Engine.go's flags to a flag package or Assets or something
	path, _ := filepath.Abs("../Assets")
	logPath, _ := filepath.Abs(path + "/Logs")
	logPath += string(os.PathSeparator)
	logPath += ""
	Init(logPath, log.Ldate|log.Ltime|log.Lshortfile)
}

// Init creates a logger based on input from Engine.go (which parses flags and starts the engine)
// this now means that the logger is coupled to the engine, but that can be reverted later if deemed necessary
func Init(logFilePath string, loggerFlags int) {

	// only allow this if we did not already create a logger...
	if logfile == nil {

		// can specify a desired file name here
		fileLocation, err := newLogger(logFilePath, loggerFlags)
		if err != nil {
			panic(err)
		}
		Info("log file location: " + fileLocation)
	}
}

// newLogger creates a file logger and stdout logger from an inputted file name and logger flags
func newLogger(inFileName string, loggerFlags int) (string, error) {

	// set up our logger file's properties
	fileDir := filepath.Dir(inFileName) + string(os.PathSeparator)
	fileName := filepath.Base(inFileName)
	fileExtension := filepath.Ext(fileName)

	// verify or create our directory
	_, err := os.Stat(fileDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(fileDir, os.ModeDir)
		if err != nil {
			panic(err)
		}
	}

	// verify or default our file name
	if len(inFileName) <= 0 || fileName == "/" || fileName == "Logs" {
		fileName = "default"
	}

	// verify or default our extension
	// if we have an extension, then it will be included with the file name
	if len(fileExtension) > 0 && fileExtension != "." {
		fileExtension = ""
	} else {
		fileExtension = ".log"
	}

	// create the log file in the designated location
	fileName = fileDir + fileName + fileExtension
	logfile, err := os.Create(fileName)
	if err != nil {
		return fileName, err
	}

	// create our loggers based on the file, and some flags
	flogger = log.New(logfile, "", loggerFlags)
	stdlogger = log.New(os.Stdout, "", loggerFlags)

	// cannot defer here since the scope of the logger is developer designed...
	// (any file that imports this file will use it...)
	// defer logfile.Close()

	return fileName, nil
}

// Debug sets the prefix and prints to stdout as well as a file
func Debug(msg ...interface{}) {
	// TODO: only send these if filtered logging allows
	flogger.SetPrefix("DEBUG:\t")
	flogger.Println(msg)

	stdlogger.SetPrefix("DEBUG:\t")
	stdlogger.Println(msg)
}

// Info sets the prefix and prints to stdout as well as a file
func Info(msg ...interface{}) {
	// TODO: only send these if filtered logging allows
	flogger.SetPrefix("INFO:\t")
	flogger.Println(msg)

	stdlogger.SetPrefix("INFO:\t")
	stdlogger.Println(msg)
}

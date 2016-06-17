// Logger.go contains simple logic to wrap around golang's native logger in order to provide multi-logging and (eventually?) filtered logging

package Logging

import (
	"log"
	"os"
	"path/filepath"
)

// make the loggers global for easy use (no need to pass stuff around)
var logfile *os.File
var flogger *log.Logger
var stdlogger *log.Logger

func init() {
	// can specify a desired file name here
	fileLocation, err := NewGTLogger("", log.Ldate|log.Ltime|log.Lshortfile)
	if err != nil {
		panic(err)
	}
	Info("log file location: " + fileLocation)
}

// NewGTLogger creates a file logger and stdout logger from an inputted file name and logger flags
func NewGTLogger(inFileName string, loggerFlags int) (string, error) {

	// set up our logger file's name
	fileName := inFileName
	fileExtension := ".log"
	if len(inFileName) > 0 {
		fileName = inFileName
	} else {
		fileName = "default"
	}

	// create the log file in the logging directory
	path, err := filepath.Abs("../../GT/Logging/")
	if err != nil {
		return fileName, err
	}
	fileName = path + fileName + fileExtension
	logfile, err = os.Create(fileName)
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

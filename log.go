package logger

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	// DBG is the logger for messages at debug level.
	DBG *log.Logger
	// INF is the logger for messages at informational level.
	INF *log.Logger
	// WRN is the logger for messages at warning level.
	WRN *log.Logger
	// ERR is the logger for messages at error level.
	ERR *log.Logger
)

// Initialise sets up the logging facility.
func Initialise(level string) {
	//
	// prepare the logging subsystem as per the command line
	//
	switch strings.ToUpper(level)[:1] {
	case "D":
		DBG = log.New(os.Stdout, "[D] ", log.Ldate|log.Ltime|log.Lshortfile)
		INF = log.New(os.Stdout, "[I] ", log.Ldate|log.Ltime|log.Lshortfile)
		WRN = log.New(os.Stdout, "[W] ", log.Ldate|log.Ltime|log.Lshortfile)
		ERR = log.New(os.Stdout, "[E] ", log.Ldate|log.Ltime|log.Lshortfile)
	case "I":
		DBG = log.New(ioutil.Discard, "[D] ", log.Ldate|log.Ltime|log.Lshortfile)
		INF = log.New(os.Stdout, "[I] ", log.Ldate|log.Ltime|log.Lshortfile)
		WRN = log.New(os.Stdout, "[W] ", log.Ldate|log.Ltime|log.Lshortfile)
		ERR = log.New(os.Stdout, "[E] ", log.Ldate|log.Ltime|log.Lshortfile)
	case "W":
		DBG = log.New(ioutil.Discard, "[D] ", log.Ldate|log.Ltime|log.Lshortfile)
		INF = log.New(ioutil.Discard, "[I] ", log.Ldate|log.Ltime|log.Lshortfile)
		WRN = log.New(os.Stdout, "[W] ", log.Ldate|log.Ltime|log.Lshortfile)
		ERR = log.New(os.Stdout, "[E] ", log.Ldate|log.Ltime|log.Lshortfile)
	case "E":
		DBG = log.New(ioutil.Discard, "[D] ", log.Ldate|log.Ltime|log.Lshortfile)
		INF = log.New(ioutil.Discard, "[I] ", log.Ldate|log.Ltime|log.Lshortfile)
		WRN = log.New(ioutil.Discard, "[W] ", log.Ldate|log.Ltime|log.Lshortfile)
		ERR = log.New(os.Stdout, "[E] ", log.Ldate|log.Ltime|log.Lshortfile)
	}
}

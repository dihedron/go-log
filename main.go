package log

import (
	"io/ioutil"
	golog "log"
	"os"
	"strings"
)

var (
	// DBG is the logger for messages at debug level.
	DBG *golog.Logger
	// INF is the logger for messages at informational level.
	INF *golog.Logger
	// WRN is the logger for messages at warning level.
	WRN *golog.Logger
	// ERR is the logger for messages at error level.
	ERR *golog.Logger
)

// SetLevel sets up the logger's level to the given value.
func SetLevel(level string) {
	//
	// prepare the logging subsystem as per the command line
	//
	switch strings.ToUpper(level)[:1] {
	case "D":
		DBG = golog.New(os.Stdout, "[D] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		INF = golog.New(os.Stdout, "[I] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		WRN = golog.New(os.Stdout, "[W] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		ERR = golog.New(os.Stdout, "[E] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
	case "I":
		DBG = golog.New(ioutil.Discard, "[D] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		INF = golog.New(os.Stdout, "[I] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		WRN = golog.New(os.Stdout, "[W] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		ERR = golog.New(os.Stdout, "[E] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
	case "W":
		DBG = golog.New(ioutil.Discard, "[D] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		INF = golog.New(ioutil.Discard, "[I] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		WRN = golog.New(os.Stdout, "[W] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		ERR = golog.New(os.Stdout, "[E] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
	case "E":
		DBG = golog.New(ioutil.Discard, "[D] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		INF = golog.New(ioutil.Discard, "[I] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		WRN = golog.New(ioutil.Discard, "[W] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		ERR = golog.New(os.Stdout, "[E] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
	}
}

func init() {
	level := "Info"
	SetLevel(level)
}

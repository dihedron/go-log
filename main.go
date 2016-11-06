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

var level = "Information"

// GetLevel returns the current logging level.
func GetLevel() string {
	return level
}

// SetLevel sets up the logger's level to the given value.
func SetLevel(value string) {

	// initialise the logger streams as per the input parameter
	switch strings.ToUpper(value)[:1] {
	case "D":
		level = "Debug"
		DBG = golog.New(os.Stdout, "[D] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		INF = golog.New(os.Stdout, "[I] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		WRN = golog.New(os.Stdout, "[W] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		ERR = golog.New(os.Stdout, "[E] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
	case "I":
		level = "Information"
		DBG = golog.New(ioutil.Discard, "[D] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		INF = golog.New(os.Stdout, "[I] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		WRN = golog.New(os.Stdout, "[W] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		ERR = golog.New(os.Stdout, "[E] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
	case "W":
		level = "Warning"
		DBG = golog.New(ioutil.Discard, "[D] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		INF = golog.New(ioutil.Discard, "[I] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		WRN = golog.New(os.Stdout, "[W] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		ERR = golog.New(os.Stdout, "[E] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
	case "E":
		level = "Error"
		DBG = golog.New(ioutil.Discard, "[D] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		INF = golog.New(ioutil.Discard, "[I] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		WRN = golog.New(ioutil.Discard, "[W] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
		ERR = golog.New(os.Stdout, "[E] ", golog.Ldate|golog.Ltime|golog.Lshortfile)
	}
}

// init simply sets the logging level to the default value.
func init() {
	SetLevel(level)
}

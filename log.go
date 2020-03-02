// Copyright 2017-present Andrea Funtò. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/mattn/go-colorable"

	"github.com/fatih/color"
)

// LogLevel represents the log level.
type LogLevel int

const (
	//TraceLevel is the LogLevel for trace messages
	TraceLevel LogLevel = iota
	// DebugLevel is the LogLevel for debug messages.
	DebugLevel
	// InfoLevel is the LogLevel for informational messages.
	InfoLevel
	// WarnLevel is the LogLevel for warning messages.
	WarnLevel
	// ErrorLevel is the LogLevel for error messages.
	ErrorLevel
	// FatalLevel is the LogLevel for fatal error messages.
	FatalLevel
	// PanicLevel is the LogLevel for fatal error messages that cause a panic.
	PanicLevel
	// NoneLevel is the LogLevel corresponding to no log output.
	NoneLevel
)

// Flag is used to influence some aspects of the logger's behaviour such as
// automatically including runtime information (source file, caller function).
type Flag int8

const (
	// FlagSourceInfo specifies whether the log message should automatically
	// include the source location (note: this feature can be computationally
	// expensive since it uses reflection at runtime).
	FlagSourceInfo = 1 << iota
	// FlagFunctionInfo specifies whether the log message should automatically
	// include the name of the containing function (note: this feature can be
	// computationally expensive since it uses reflection at runtime).
	FlagFunctionInfo
)

// FunctionWidth represents the maximum width of the function name width in
// logging messages.
//const FunctionWidth int = 32

// LevelFromString returns a log Level value by parsing the user-provided string
// in a lenient way; if the parsing fails, returns and error.
func LevelFromString(s string) (LogLevel, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "t", "trc", "trace":
		return TraceLevel, nil
	case "d", "dbg", "debug":
		return DebugLevel, nil
	case "i", "inf", "info", "informational":
		return InfoLevel, nil
	case "w", "wrn", "warn", "warning":
		return WarnLevel, nil
	case "e", "err", "error":
		return ErrorLevel, nil
	case "f", "ftl", "fatal":
		return FatalLevel, nil
	case "p", "pan", "panic":
		return PanicLevel, nil
	case "nil", "null", "none":
		return NoneLevel, nil
	default:
		return ErrorLevel, fmt.Errorf("unparseable log level: %q", s)
	}
}

// String returns a string representation of the log level for use in traces.
func (l LogLevel) String() string {
	switch l {
	case TraceLevel:
		return "[T]"
	case DebugLevel:
		return "[D]"
	case InfoLevel:
		return "[I]"
	case WarnLevel:
		return "[W]"
	case ErrorLevel:
		return "[E]"
	case FatalLevel:
		return "[F]"
	case PanicLevel:
		return "[P]"
	}
	return ""
}

// logln is the prototype of log functions writing a line to a stream.
type logln func(writer io.Writer, args ...interface{}) (int, error)

// logf is the prototype of log functions writing a formatted output to a stream.
type logf func(writer io.Writer, format string, args ...interface{}) (int, error)

const (
	// SourceInfoNone is the constant that specifies that no source file information
	// (file and line) should be printed out.
	SourceInfoNone int8 = iota
	// SourceInfoShort is the constants that specifies that the source file
	// information should be printed in short form (file name only).
	SourceInfoShort
	// SourceInfoLong is the constants that specifies that the source file
	// information should be printed in log form (complete file path).
	SourceInfoLong
)

var (
	logLevel               LogLevel
	logLevelLock           sync.RWMutex
	logStream              io.Writer
	logStreamLock          sync.RWMutex
	logTimeFormat          string
	logTimeFormatLock      sync.RWMutex
	logPrintSourceInfo     int8
	logPrintSourceInfoLock sync.RWMutex
	logPrintCallerInfo     bool
	logPrintCallerInfoLock sync.RWMutex
	logTracef              logf
	logDebugf              logf
	logInfof               logf
	logWarnf               logf
	logErrorf              logf
	logFatalf              logf
	logPanicf              logf
	logTraceln             logln
	logDebugln             logln
	logInfoln              logln
	logWarnln              logln
	logErrorln             logln
	logFatalln             logln
	logPanicln             logln
)

func init() {
	SetLevel(DebugLevel)
	SetStream(os.Stderr, true)
	SetTimeFormat("2006-01-02@15:04:05.000")
	SetPrintCallerInfo(true)
	SetPrintSourceInfo(SourceInfoShort)
}

// SetLevel sets the log level for the application.
func SetLevel(level LogLevel) {
	logLevelLock.Lock()
	defer logLevelLock.Unlock()
	logLevel = level
}

// GetLevel retur s the current log level.
func GetLevel() LogLevel {
	logLevelLock.RLock()
	defer logLevelLock.RUnlock()
	return logLevel
}

// SetStream sets the stream to write messages to; if the colorise flag is set,
// the logger will wrap the stream so it always produces properly coloured output
// messages; this might be less appropriate when writing to a file.
func SetStream(stream io.Writer, colorise bool) {
	logStreamLock.Lock()
	defer logStreamLock.Unlock()
	if stream, ok := stream.(*os.File); colorise && ok {
		logStream = colorable.NewColorable(stream)
		logTracef = color.New(color.FgWhite).Fprintf
		logDebugf = color.New(color.FgWhite).Fprintf
		logInfof = color.New(color.FgGreen).Fprintf
		logWarnf = color.New(color.FgYellow).Fprintf
		logErrorf = color.New(color.FgRed).Fprintf
		logFatalf = color.New(color.FgBlue).Fprintf
		logPanicf = color.New(color.FgMagenta).Fprintf
		logTraceln = color.New(color.FgWhite).Fprintln
		logDebugln = color.New(color.FgWhite).Fprintln
		logInfoln = color.New(color.FgGreen).Fprintln
		logWarnln = color.New(color.FgYellow).Fprintln
		logErrorln = color.New(color.FgRed).Fprintln
		logFatalln = color.New(color.FgBlue).Fprintln
		logPanicln = color.New(color.FgMagenta).Fprintln
	} else {
		logStream = stream
		logTracef = fmt.Fprintf
		logDebugf = fmt.Fprintf
		logInfof = fmt.Fprintf
		logWarnf = fmt.Fprintf
		logErrorf = fmt.Fprintf
		logFatalf = fmt.Fprintf
		logPanicf = fmt.Fprintf
		logTraceln = fmt.Fprintln
		logDebugln = fmt.Fprintln
		logInfoln = fmt.Fprintln
		logWarnln = fmt.Fprintln
		logErrorln = fmt.Fprintln
		logFatalln = fmt.Fprintln
		logPanicln = fmt.Fprintln
	}
}

// GetStream returns the current log stream.
func GetStream() io.Writer {
	logStreamLock.RLock()
	defer logStreamLock.RUnlock()
	return logStream
}

// SetTimeFormat sets the format for log messages time.
func SetTimeFormat(format string) {
	logTimeFormatLock.Lock()
	defer logTimeFormatLock.Unlock()
	logTimeFormat = format
}

// GetTimeFormat returns the current format of log messages time.
func GetTimeFormat() string {
	logTimeFormatLock.RLock()
	defer logTimeFormatLock.RUnlock()
	return logTimeFormat
}

// SetPrintCallerInfo enables or disables the automatic addition of the calling
// function (with package) to the log messages. NOTE: enabling this feature can
// have severe impacts on performances since it uses reflection at runtime.
func SetPrintCallerInfo(enabled bool) {
	logPrintCallerInfoLock.Lock()
	defer logPrintCallerInfoLock.Unlock()
	logPrintCallerInfo = enabled
}

// GetPrintCallerInfo returns whether the automatic addition of the calling
// function (with package) to the log messages is enabled.
func GetPrintCallerInfo() bool {
	logPrintCallerInfoLock.RLock()
	defer logPrintCallerInfoLock.RUnlock()
	return logPrintCallerInfo
}

// SetPrintSourceInfo enables or disables the automatic addition of the source
// and line number info to the log messages; use one among SourceFileNone,
// SourceFileShort and SourceFileLong here. NOTE: enabling this feature can
// have severe impacts on performances since it uses reflection at runtime.
func SetPrintSourceInfo(value int8) {
	logPrintSourceInfoLock.Lock()
	defer logPrintSourceInfoLock.Unlock()
	logPrintSourceInfo = value
}

// GetPrintSourceInfo returns whether the automatic addition of the source and
// line number info to the log messages is enabled, and whether the file name
// will be printed in short or long form.
func GetPrintSourceInfo() int8 {
	logPrintSourceInfoLock.RLock()
	defer logPrintSourceInfoLock.RUnlock()
	return logPrintSourceInfo
}

// IsTrace returns whether the trace (TraceLevel) log elevel is enabled.
func IsTrace() bool {
	return GetLevel() <= TraceLevel
}

// IsDebug returns whether the debug (DebugLevel) log elevel is enabled.
func IsDebug() bool {
	return GetLevel() <= DebugLevel
}

// IsInfo returns whether the informational (InfoLevel) log elevel is enabled.
func IsInfo() bool {
	return GetLevel() <= InfoLevel
}

// IsWarning returns whether the warning (WarnLevel) log elevel is enabled.
func IsWarning() bool {
	return GetLevel() <= WarnLevel
}

// IsError returns whether the error (ErrorLevel) log elevel is enabled.
func IsError() bool {
	return GetLevel() <= ErrorLevel
}

// IsFatal returns whether the fatal (FatalLevel) log elevel is enabled.
func IsFatal() bool {
	return GetLevel() <= FatalLevel
}

// IsPanic returns whether the panic (PanicLevel) log elevel is enabled.
func IsPanic() bool {
	return GetLevel() <= PanicLevel
}

// IsDisabled returns whether the log is disabled.
func IsDisabled() bool {
	return GetLevel() <= NoneLevel
}

// Traceln writes a trace message to the current output stream, appending a new
// line.
func Traceln(args ...interface{}) (int, error) {
	if IsTrace() {
		args = prepareArgs(DebugLevel, args...)
		return logTraceln(GetStream(), args...)
	}
	return 0, nil
}

// Debugln writes a debug message to the current output stream, appending a new
// line.
func Debugln(args ...interface{}) (int, error) {
	if IsDebug() {
		args = prepareArgs(DebugLevel, args...)
		return logDebugln(GetStream(), args...)
	}
	return 0, nil
}

// Infoln writes an informational message to the current output stream,
// appending a new line.
func Infoln(args ...interface{}) (int, error) {
	if IsInfo() {
		args = prepareArgs(InfoLevel, args...)
		return logInfoln(GetStream(), args...)
	}
	return 0, nil
}

// Warnln writes a warning message to the current output stream, appending a new
// line.
func Warnln(args ...interface{}) (int, error) {
	if IsWarning() {
		args = prepareArgs(WarnLevel, args...)
		return logWarnln(GetStream(), args...)
	}
	return 0, nil
}

// Errorln writes an error message to the current output stream, appending a new
// line.
func Errorln(args ...interface{}) (int, error) {
	if IsError() {
		args = prepareArgs(ErrorLevel, args...)
		return logErrorln(GetStream(), args...)
	}
	return 0, nil
}

// Fatalln writes an error message to the current output stream, appending a new
// line.
func Fatalln(args ...interface{}) (int, error) {
	if IsFatal() {
		args = prepareArgs(FatalLevel, args...)
		logFatalln(GetStream(), args...)
	}
	return 0, nil
}

// Panicln writes an error message to the current output stream, appending a new
// line; then it panics.
func Panicln(args ...interface{}) (int, error) {
	if IsPanic() {
		args = prepareArgs(PanicLevel, args...)
		logPanicln(GetStream(), args...)
	}
	panic("unrecoverable error")
}

// Tracef writes a trace message to the current output stream, appending a new line.
func Tracef(format string, args ...interface{}) (int, error) {
	if IsTrace() {
		format, args = prepareFormatAndArgs(TraceLevel, format, args...)
		if !strings.HasSuffix(format, "\n") && !strings.HasSuffix(format, "\r") {
			format = format + "\n"
		}
		return logTracef(GetStream(), format, args...)
	}
	return 0, nil
}

// Debugf writes a debug message to the current output stream, appending a new line.
func Debugf(format string, args ...interface{}) (int, error) {
	if IsDebug() {
		format, args = prepareFormatAndArgs(DebugLevel, format, args...)
		if !strings.HasSuffix(format, "\n") && !strings.HasSuffix(format, "\r") {
			format = format + "\n"
		}
		return logDebugf(GetStream(), format, args...)
	}
	return 0, nil
}

// Infof writes an informational message to the current output stream,
// appending a new line.
func Infof(format string, args ...interface{}) (int, error) {
	if IsInfo() {
		format, args = prepareFormatAndArgs(InfoLevel, format, args...)
		if !strings.HasSuffix(format, "\n") && !strings.HasSuffix(format, "\r") {
			format = format + "\n"
		}
		return logInfof(GetStream(), format, args...)
	}
	return 0, nil
}

// Warnf writes a warning message to the current output stream, appending a new line.
func Warnf(format string, args ...interface{}) (int, error) {
	if IsWarning() {
		format, args = prepareFormatAndArgs(WarnLevel, format, args...)
		if !strings.HasSuffix(format, "\n") && !strings.HasSuffix(format, "\r") {
			format = format + "\n"
		}
		return logWarnf(GetStream(), format, args...)
	}
	return 0, nil
}

// Errorf writes an error message to the current output stream, appending a new
// line.
func Errorf(format string, args ...interface{}) (int, error) {
	if IsError() {
		format, args = prepareFormatAndArgs(ErrorLevel, format, args...)
		if !strings.HasSuffix(format, "\n") && !strings.HasSuffix(format, "\r") {
			format = format + "\n"
		}
		return logErrorf(GetStream(), format, args...)
	}
	return 0, nil
}

// Fatalf writes an error message to the current output stream, appending a new
// line.
func Fatalf(format string, args ...interface{}) (int, error) {
	if IsFatal() {
		format, args = prepareFormatAndArgs(FatalLevel, format, args...)
		if !strings.HasSuffix(format, "\n") && !strings.HasSuffix(format, "\r") {
			format = format + "\n"
		}
		logFatalf(GetStream(), format, args...)
	}
	return 0, nil
}

// Panicf writes an error message to the current output stream, appending a new
// line; then it panics.
func Panicf(format string, args ...interface{}) (int, error) {
	if IsPanic() {
		format, args = prepareFormatAndArgs(PanicLevel, format, args...)
		if !strings.HasSuffix(format, "\n") && !strings.HasSuffix(format, "\r") {
			format = format + "\n"
		}
		logPanicf(GetStream(), format, args...)
	}
	panic("unrecoverable error")
}

// Println is a raw version of the debug functions; it tries to interpret the
// message by checking if it starts with anthing like "[D]" or "[W]"; if so, it
// delegates to the corresponding logging function, otherwise it just prints to
// the log stream as is, with no additional formatting.
func Println(args ...interface{}) (int, error) {
	if len(args) > 0 {
		if value, ok := args[0].(string); ok {
			switch {
			case strings.HasPrefix(value, "[T]"):
				return Traceln(args[1:]...)
			case strings.HasPrefix(value, "[D]"):
				return Debugln(args[1:]...)
			case strings.HasPrefix(value, "[I]"):
				return Infoln(args[1:]...)
			case strings.HasPrefix(value, "[W]"):
				return Warnln(args[1:]...)
			case strings.HasPrefix(value, "[E]"):
				return Errorln(args[1:]...)
			case strings.HasPrefix(value, "[F]"):
				return Fatalln(args[1:]...)
			case strings.HasPrefix(value, "[P]"):
				return Panicln(args[1:]...)
			}
		}
	}
	return fmt.Fprintln(GetStream(), args...)
}

// Printf is a raw version of the debug functions; it tries to interpret the
// message by checking if it starts with anything like "[D]" or "[W]"; if so, it
// delegates to the corresponding logging function, otherwise it just prints to
// the log stream as is, with no additional formatting.
func Printf(format string, args ...interface{}) (int, error) {
	re := regexp.MustCompile(`^\[(T|D|I|W|E|F|P)\]\s`)
	switch {
	case strings.HasPrefix(format, "[T]"):
		return Tracef(re.ReplaceAllString(format, ""), args...)
	case strings.HasPrefix(format, "[D]"):
		return Debugf(re.ReplaceAllString(format, ""), args...)
	case strings.HasPrefix(format, "[I]"):
		return Infof(re.ReplaceAllString(format, ""), args...)
	case strings.HasPrefix(format, "[W]"):
		return Warnf(re.ReplaceAllString(format, ""), args...)
	case strings.HasPrefix(format, "[E]"):
		return Errorf(re.ReplaceAllString(format, ""), args...)
	case strings.HasPrefix(format, "[F]"):
		return Fatalf(re.ReplaceAllString(format, ""), args...)
	case strings.HasPrefix(format, "[P]"):
		return Panicf(re.ReplaceAllString(format, ""), args...)
	}
	return fmt.Fprintf(GetStream(), format, args...)
}

// prepareFormatAndArgs prepares the format and args array for logf, depending
// on the active runtime logging options (e.g. caller function, source file and
// line number).
func prepareFormatAndArgs(level LogLevel, format string, args ...interface{}) (string, []interface{}) {

	leadFormat := "%s %s - "
	tailFormat := ""
	leadArgs := []interface{}{level.String(), time.Now().Format(GetTimeFormat())}
	tailArgs := []interface{}{}

	if GetPrintCallerInfo() || GetPrintSourceInfo() > 0 {
		var fun, file string
		var line int
		pc, file, line, ok := runtime.Caller(2)
		if !ok {
			fun = "<unknown>"
			file = "???"
			line = -1
		} else {
			if GetPrintCallerInfo() {
				f := runtime.FuncForPC(pc)
				if f == nil {
					fun = "<unknown>"
				} else {
					fun = f.Name()
				}
				fun = fun[strings.LastIndex(fun, "/")+1:]
				leadFormat = leadFormat + "%s: "
				leadArgs = append(leadArgs, fun)
			}
			switch GetPrintSourceInfo() {
			case SourceInfoShort:
				file = file[strings.LastIndex(file, "/")+1:]
				fallthrough
				// tailFormat = " (%s:%d)"
				// tailArgs = append(tailArgs, []interface{}{file, line}...)
			case SourceInfoLong:
				tailFormat = " (%s:%d)"
				tailArgs = append(tailArgs, []interface{}{file, line}...)
				format = strings.TrimSuffix(format, "\n")
			default:
			}
		}
	}
	format = leadFormat + format + tailFormat
	args = append(leadArgs, append(args, tailArgs...)...)
	return format, args
}

// prepareArgs prepares the aray of args for logln , depending on the active
// runtime logging options (e.g. caller function, source file and line number);
// it is similar to prepareFormatAndArgs but logln does not require a format.
func prepareArgs(level LogLevel, args ...interface{}) []interface{} {

	list := []interface{}{fmt.Sprintf("%s %s -", level.String(), time.Now().Format(GetTimeFormat()))}
	if GetPrintCallerInfo() || GetPrintSourceInfo() > 0 {
		var fun, file string
		var line int
		pc, file, line, ok := runtime.Caller(2)
		if !ok {
			fun = "<unknown>"
			file = "???"
			line = -1
		} else {
			if GetPrintCallerInfo() {
				f := runtime.FuncForPC(pc)
				if f == nil {
					fun = "<unknown>"
				} else {
					fun = f.Name()
				}
				fun = fun[strings.LastIndex(fun, "/")+1:]
				list = append(list, fmt.Sprintf("%s:", fun))
			}
			switch GetPrintSourceInfo() {
			case SourceInfoShort:
				file = file[strings.LastIndex(file, "/")+1:]
				fallthrough
			case SourceInfoLong:
				last := strings.TrimSuffix(fmt.Sprintf("%v", args[len(args)-1]), "\n")
				args = append(args[:len(args)-1], last)
				args = append(args, fmt.Sprintf("(%s:%d)", file, line))
			default:
			}
		}
	}
	args = append(list, args...)
	return args
}

// ToJSON converts an object into pretty-printed JSON format.
func ToJSON(object interface{}) string {
	if bytes, err := json.MarshalIndent(object, "", "  "); err == nil {
		return string(bytes)
	}
	return ""
}

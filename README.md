# go-log - Yet another logger wrapper for golang
A simple logger for Go; it provides log messages colorising, automatic addition
or source file, line number and/or calling function to log messages. All top
level messages are synchronised, so it is safe to reconfigure the logger from
different goroutines.

// Copyright 2017-present Andrea Funtò. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/dihedron/go-log/log"
)

func main() {

	log.SetLevel(log.DBG)
	log.SetStream(os.Stdout, true)
	log.SetTimeFormat("15:04:05.000")
	log.SetPrintCallerInfo(true)
	log.SetPrintSourceInfo(log.SourceInfoShort)

	log.Debugf("debug message")
	log.Infof("info message")
	log.Warnf("warn message")
	log.Errorf("error message")

	log.Debugf("debug message with newline\n")
	log.Infof("info message with newline\n")
	log.Warnf("warn message with newline\n")
	log.Errorf("error message with newline\n")

	log.Debugln("debug message", "a", "b", "c")
	log.Infoln("info message")
	log.Warnln("warn message")
	log.Errorln("error message")

	log.Debugln("debug message", "a", "b", "c", "with newline\n")
	log.Infoln("info message with newline\n")
	log.Warnln("warn message with newline\n")
	log.Errorln("error message with newline\n")

	log.SetStream(os.Stdout, false)
	log.Debugln("debug message", "a", "b", "c", "no colour")
	log.Infoln("info message with newline", "no colour")
	log.Warnln("warn message with newline", "no colour")
	log.Errorln("error message with newline", "no colour")

}

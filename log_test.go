// Copyright 2017-present Andrea Funtò. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"os"
	"testing"
)

func TestLog(t *testing.T) {

	SetLevel(DBG)
	SetStream(os.Stdout, true)
	SetTimeFormat("15:04:05.000")
	SetPrintCallerInfo(true)
	SetPrintSourceInfo(SourceInfoShort)

	Debugf("debug message")
	Infof("info message")
	Warnf("warn message")
	Errorf("error message")

	Debugf("debug message with newline\n")
	Infof("info message with newline\n")
	Warnf("warn message with newline\n")
	Errorf("error message with newline\n")

	Debugln("debug message", "a", "b", "c")
	Infoln("info message")
	Warnln("warn message")
	Errorln("error message")

	Debugln("debug message", "a", "b", "c", "with newline\n")
	Infoln("info message with newline\n")
	Warnln("warn message with newline\n")
	Errorln("error message with newline\n")

	SetStream(os.Stdout, false)
	Debugln("debug message", "a", "b", "c", "no colour")
	Infoln("info message with newline", "no colour")
	Warnln("warn message with newline", "no colour")
	Errorln("error message with newline", "no colour")

}

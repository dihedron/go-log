// Copyright 2017-present Andrea Funtò. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"os"
	"testing"
)

func TestLog(t *testing.T) {

	SetLevel(TraceLevel)
	SetStream(os.Stdout, true)
	SetTimeFormat("15:04:05.000")
	SetPrintCallerInfo(true)
	SetPrintSourceInfo(SourceInfoShort)

	defer func() {
		if r := recover(); r != nil {
			t.Log("Recovered", r)
		}
	}()

	Tracef("trace message")
	Debugf("debug message")
	Infof("info message")
	Warnf("warn message")
	Errorf("error message")
	Fatalf("fatal message")
	//Panicf("panic message")

	Tracef("trace message with newline\n")
	Debugf("debug message with newline\n")
	Infof("info message with newline\n")
	Warnf("warn message with newline\n")
	Errorf("error message with newline\n")
	Fatalf("fatal message with newline\n")

	Traceln("trace message", "a", "b", "c")
	Debugln("debug message", "a", "b", "c")
	Infoln("info message", "a", "b", "c")
	Warnln("warn message", "a", "b", "c")
	Errorln("error message", "a", "b", "c")
	Fatalln("fatal message", "a", "b", "c")
	//Panicln("panic message", "a", "b", "c")

	Traceln("trace message", "a", "b", "c", "with newline\n")
	Debugln("debug message", "a", "b", "c", "with newline\n")
	Infoln("info message", "a", "b", "c", "with newline\n")
	Warnln("warn message", "a", "b", "c", "with newline\n")
	Errorln("error message", "a", "b", "c", "with newline\n")
	Fatalln("fatal message", "a", "b", "c", "with newline\n")

	SetStream(os.Stdout, false)
	Traceln("trace message", "a", "b", "c", "no colour")
	Debugln("debug message", "a", "b", "c", "no colour")
	Infoln("info message with newline", "no colour")
	Warnln("warn message with newline", "no colour")
	Errorln("error message with newline", "no colour")
	Fatalln("fatal message with newline", "no colour")

}

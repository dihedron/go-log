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
	log.SetStream(os.Stdout)
	log.SetTimeFormat("15:04:05.000")
	log.SetColorise(true)
	log.SetPrintCallerInfo(true)
	log.SetPrintSourceInfo(true)
	log.Debugln("debug message")
	log.Infoln("info message")
	log.Warnln("warn message")
	log.Warnln("error message")
}

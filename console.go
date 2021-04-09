// Copyright 2015 Daniel Theophanes.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

package service

import (
	"log"
	"os"

	"github.com/fatih/color"
)

// ConsoleLogger logs to the std err.
var ConsoleLogger = consoleLogger{}

type consoleLogger struct {
	info, warn, err *log.Logger
}

func init() {
	ConsoleLogger.info = log.New(os.Stderr, "[INF] ", (log.Ldate | log.Ltime | log.Lmsgprefix))
	ConsoleLogger.warn = log.New(os.Stderr, "[WAR] ", (log.Ldate | log.Ltime | log.Lmsgprefix))
	ConsoleLogger.err = log.New(os.Stderr, "[ERR] ", (log.Ldate | log.Ltime | log.Lmsgprefix))
}

func (c consoleLogger) Error(v ...interface{}) error {
	color.Set(color.FgRed)
	c.err.Print(v...)
	color.Unset()
	return nil
}
func (c consoleLogger) Warning(v ...interface{}) error {
	color.Set(color.FgHiWhite)
	c.warn.Print(v...)
	color.Unset()
	return nil
}
func (c consoleLogger) Info(v ...interface{}) error {
	c.info.Print(v...)
	return nil
}
func (c consoleLogger) Errorf(format string, a ...interface{}) error {
	color.Set(color.FgRed)
	c.err.Printf(format, a...)
	color.Unset()
	return nil
}
func (c consoleLogger) Warningf(format string, a ...interface{}) error {
	color.Set(color.FgHiWhite)
	c.warn.Printf(color.HiWhiteString(format), a...)
	color.Unset()
	return nil
}
func (c consoleLogger) Infof(format string, a ...interface{}) error {
	c.info.Printf(format, a...)
	return nil
}

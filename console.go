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
	trace, debug, info, warn, err *log.Logger
}

func init() {
	ConsoleLogger.trace = log.New(os.Stderr, "[TRC] ", (log.Ltime | log.Lmsgprefix))
	ConsoleLogger.debug = log.New(os.Stderr, "[DBG] ", (log.Ltime | log.Lmsgprefix))
	ConsoleLogger.info = log.New(os.Stderr, "[INF] ", (log.Ltime | log.Lmsgprefix))
	ConsoleLogger.warn = log.New(os.Stderr, "[WRN] ", (log.Ltime | log.Lmsgprefix))
	ConsoleLogger.err = log.New(os.Stderr, "[ERR] ", (log.Ltime | log.Lmsgprefix))
}

func (c consoleLogger) Error(v ...interface{}) error {
	color.Set(color.FgHiRed)
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
	color.Set(color.FgWhite)
	c.info.Print(v...)
	color.Unset()
	return nil
}
func (c consoleLogger) Debug(v ...interface{}) error {
	color.Set(color.FgHiCyan)
	c.debug.Print(v...)
	color.Unset()
	return nil
}
func (c consoleLogger) Trace(v ...interface{}) error {
	color.Set(color.FgCyan)
	c.trace.Print(v...)
	color.Unset()
	return nil
}
func (c consoleLogger) Errorf(format string, a ...interface{}) error {
	color.Set(color.FgHiRed)
	c.err.Printf(format, a...)
	color.Unset()
	return nil
}
func (c consoleLogger) Warningf(format string, a ...interface{}) error {
	color.Set(color.FgHiWhite)
	c.warn.Printf(format, a...)
	color.Unset()
	return nil
}
func (c consoleLogger) Infof(format string, a ...interface{}) error {
	color.Set(color.FgWhite)
	c.info.Printf(format, a...)
	color.Unset()
	return nil
}
func (c consoleLogger) Debugf(format string, a ...interface{}) error {
	color.Set(color.FgHiCyan)
	c.debug.Printf(format, a...)
	color.Unset()
	return nil
}
func (c consoleLogger) Tracef(format string, a ...interface{}) error {
	color.Set(color.FgCyan)
	c.trace.Printf(format, a...)
	color.Unset()
	return nil
}

// +build !windows

package service

import "github.com/fatih/color"

func (c consoleLogger) Trace(v ...interface{}) error {
	color.Set(color.Faint, color.Italic)
	c.trace.Print(v...)
	color.Unset()
	return nil
}
func (c consoleLogger) Tracef(format string, a ...interface{}) error {
	color.Set(color.Faint, color.Italic)
	c.trace.Printf(format, a...)
	color.Unset()
	return nil
}

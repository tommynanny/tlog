package logger

import (
	"bytes"
	"fmt"
	"log"
	"runtime"

	"github.com/gookit/goutil/dump"
)

type TLogger struct {
	// Logger is the underlying log.Logger instance used to output log messages.
	Logger  log.Logger
	Options *LoggerOptions
}

// wrap formats the given values using the dump package, with or without color, and returns both the formatted string
// for use in the log file and the formatted string with color for use in standard output, if ColorfulStdout is true.
// If UseWrapper is false, it returns the same formatted string for both outputs.
// Input: raw (any) - the values to be formatted
// Output: (string) - the formatted string for use in the log file
//
//	(string) - the formatted string with color for use in standard output, if ColorfulStdout is true
func (tl *TLogger) wrap(raw ...any) (logString string, stdString string) {
	if !tl.Options.UseWrapper {
		return fmt.Sprint(raw...), fmt.Sprint(raw...)
	}

	dump.Config(dump.WithoutColor(), dump.WithCallerSkip(tl.Options.WithCallerSkip))
	w := &bytes.Buffer{}
	dump.Std().Fprint(w, raw...)
	noColorWrappedString := w.String()

	if tl.Options.ColorfulStdout {
		return noColorWrappedString, dump.Format(raw)
	}
	return noColorWrappedString, noColorWrappedString
}

// HandleError logs the given error with the TLogger's Logger instance and, if the error is not nil,
// also dumps it to standard output and passes it on to the global errLogger.
// Input: err (error) - the error to be logged
// Output: none
func (tl *TLogger) HandleError(err error) {
	if err == nil {
		return
	}

	r1, r2 := tl.wrap(err.Error())
	tl.dumpToStdout(r2)
	tl.Logger.Println(r1)

	if tl != ErrLogger {
		ErrLogger.HandleError(err)
	}
}

// Panicln logs the given error with the TLogger's Logger instance and panics, if the error is not nil,
// after dumping it to standard output.
// Input: err (error) - the error to be logged
// Output: none
func (tl *TLogger) Panicln(err error) {
	if err == nil {
		return
	}

	r1, r2 := tl.wrap(err.Error())
	tl.dumpToStdout(r2)
	tl.Logger.Panicln(r1)

	if tl != ErrLogger {
		ErrLogger.HandleError(err)
	}
}

func (tl *TLogger) Panic(err error) {
	if err == nil {
		return
	}

	r1, r2 := tl.wrap(err.Error())
	tl.dumpToStdout(r2)
	tl.Logger.Panic(r1)

	if tl != ErrLogger {
		ErrLogger.HandleError(err)
	}
}

// Println logs the given arguments with the TLogger's Logger instance and, if PrintToStdout is true,
// also dumps them to standard output.
// Input: v (any) - the value or values to be logged
// Output: none
func (tl *TLogger) Println(v ...any) {
	r1, r2 := tl.wrap(v...)
	tl.dumpToStdout(r2)
	tl.Logger.Println(r1)
}

func (tl *TLogger) Print(v ...any) {
	r1, r2 := tl.wrap(v...)
	tl.dumpToStdout(r2)
	tl.Logger.Print(r1)
}

// Printf formats and logs the given string with the TLogger's Logger instance and, if PrintToStdout is true,
// also dumps it to standard output.
// Input: format (string) - the format string for the log message
//
//	v (any) - the values to be inserted into the format string
//
// Output: none
func (tl *TLogger) Printf(format string, v ...any) {
	result := fmt.Sprintf(format, v...)
	r1, r2 := tl.wrap(result)

	tl.dumpToStdout(r2)
	tl.Logger.Println(r1)
}

// dumpToStdout dumps the given arguments to standard output, if PrintToStdout is true.
// Input: v (any) - the value or values to be dumped to standard output
// Output: none
func (tl *TLogger) dumpToStdout(v ...any) {
	if tl.Options.PrintToStdout {
		dump.Println(v)
	}
}

// Trace returns the filename, line number, and function name of the caller at the specified number of stack frames away.
// Input: skip (int) - the number of stack frames to skip
// Output: (string) - the filename of the caller
//
//	(int)    - the line number of the caller
//	(string) - the name of the function of the caller
//
// default skip = 2
// fileN, lineN, funcN := Trace(2)
// fmt.Sprintf("file='%s',line='%d',func='%s'\n", fileN, lineN, funcN)
func Trace(skip int) (string, int, string) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.File, frame.Line, frame.Function
}

// TrackString returns a string representation of the caller's location at three stack frames away.
// Input: none
// Output: (string) - a string representation of the caller's location
func TrackString() string {
	fileN, lineN, funcN := Trace(3)
	return fmt.Sprintf("file='%s',line='%d',func='%s'", fileN, lineN, funcN)
}

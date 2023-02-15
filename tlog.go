package tlog

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	lg "github.com/tommynanny/tlog/logger"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	GlobalOptions *Options
)

func init() {
	if GlobalOptions.AutoSetupLogger {
		SetupLogger()
	}
}

// SetupLogger sets up the Main and errLogger loggers.
func SetupLogger() {
	lg.Main = CreateTLogger("Main")
	lg.ErrLogger = CreateTLogger("Errors", lg.NoStdout(), lg.WithCallSkip(5))
}

// prepareLogFile prepares the log file for use.
// Input: logfName (string) - the name of the log file to prepare
// Output: none
func prepareLogFile(logfName string) {
	_, err := os.OpenFile(logfName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		_, err := os.OpenFile(logfName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("error opening file: %v\n", err)
		}
	}
}

// CreateLogger creates a new logger with the given name.
// Input: name (string) - the name of the new logger
// Output: (*TLogger) - a pointer to the new logger
func CreateTLogger(name string, fns ...lg.OptionFunc) *lg.TLogger {
	logfName := fmt.Sprintf("./%s/%s.log", GlobalOptions.LOG_ROOT, name)

	parentDir := filepath.Dir(logfName)
	if err := os.MkdirAll(parentDir, 0755); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	prepareLogFile(logfName)

	current_logger := log.New(&lumberjack.Logger{
		Filename:   logfName,
		MaxSize:    10, // megabytes
		MaxBackups: 500,
		MaxAge:     360, //days
	}, "", log.LstdFlags)

	return &lg.TLogger{
		Logger:  *current_logger,
		Options: lg.NewDefaultOptions(),
	}
}

// CreateTLoggerWithWriter creates a new logger with the given name and a writter.
// Input: name (string) - the name of the new logger
// Input: writter (io.Writer) - the writer
// Output: (*TLogger) - a pointer to the new logger
func CreateTLoggerWithWriter(name string, writter io.Writer, fns ...lg.OptionFunc) *lg.TLogger {
	logfName := fmt.Sprintf("./%s/%s.log", GlobalOptions.LOG_ROOT, name)

	parentDir := filepath.Dir(logfName)
	if err := os.MkdirAll(parentDir, 0755); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	prepareLogFile(logfName)

	current_logger := log.New(writter, "", log.LstdFlags)

	return &lg.TLogger{
		Logger:  *current_logger,
		Options: lg.NewDefaultOptions(),
	}
}

func HandleError(err error)          { lg.Main.HandleError(err) }
func Panicln(err error)              { lg.Main.Panicln(err) }
func Panic(err error)                { lg.Main.Panic(err) }
func Println(v ...any)               { lg.Main.Println(v) }
func Print(v ...any)                 { lg.Main.Print(v) }
func Printf(format string, v ...any) { lg.Main.Printf(format, v) }

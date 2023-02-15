package tlog

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// Main is the main logger.
	Main *TLogger
	// errLogger is the global error logger.
	errLogger *TLogger
	// LOG_ROOT is the root directory for all log files, which use timestamps as sessions.
	LOG_ROOT = "logs/" + time.Now().UTC().Format("2006-01-02--15-04-05")
	// UseWrapperOnNew determines whether to use the wrapper for log files created by CreateLogger().
	UseWrapperOnNew = true
	// ColorfulStdoutOnNew determines whether to use colorful output for standard output.
	ColorfulStdoutOnNew = true
	// PrintToStdoutOnNew determines whether log messages are printed to standard output in addition to the log file.
	PrintToStdoutOnNew = true

	AutoSetupLogger = true
)

func init() {
	if AutoSetupLogger {
		SetupLogger()
	}
}

// SetupLogger sets up the Main and errLogger loggers.
func SetupLogger() {
	Main = CreateTLogger("Main")
	errLogger = CreateTLogger("Errors")
	errLogger.PrintToStdout = false
	errLogger.WithCallerSkip = 5
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
func CreateTLogger(name string) *TLogger {
	logfName := fmt.Sprintf("./%s/%s.log", LOG_ROOT, name)

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

	return &TLogger{
		Logger:         *current_logger,
		WithCallerSkip: 4,
		PrintToStdout:  PrintToStdoutOnNew,
		UseWrapper:     UseWrapperOnNew,
		ColorfulStdout: ColorfulStdoutOnNew,
	}
}

// CreateTLoggerWithWriter creates a new logger with the given name and a writter.
// Input: name (string) - the name of the new logger
// Input: writter (io.Writer) - the writer
// Output: (*TLogger) - a pointer to the new logger
func CreateTLoggerWithWriter(name string, writter io.Writer) *TLogger {
	logfName := fmt.Sprintf("./%s/%s.log", LOG_ROOT, name)

	parentDir := filepath.Dir(logfName)
	if err := os.MkdirAll(parentDir, 0755); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	prepareLogFile(logfName)

	current_logger := log.New(writter, "", log.LstdFlags)

	return &TLogger{
		Logger:         *current_logger,
		WithCallerSkip: 4,
		PrintToStdout:  PrintToStdoutOnNew,
		UseWrapper:     UseWrapperOnNew,
		ColorfulStdout: ColorfulStdoutOnNew,
	}
}

package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information
type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns a pointer to a new logger, so it can be easily shared.
// The default output is Sdout.
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}

// Debugf formats and prints a message if log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	//implementation
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelDebug {
		return
	}
	fmt.Printf(format+"\n", args...)
}

// Infof formats and prints a message if log level is debug or higher.
func (l *Logger) InfoF(format string, args ...any) {
	//implementation
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelInfo {
		return
	}
	fmt.Printf(format+"\n", args...)
}

// Warningf formats and prints a message if log level is debug or higher.
func (l *Logger) Warningf(format string, args ...any) {
	//implementation
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelWarn {
		return
	}
	fmt.Printf(format+"\n", args...)
}

// Errorf formats and prints a message if log level is debug or higher.
func (l *Logger) Errorf(format string, args ...any) {
	//implementation
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelError {
		return
	}
	fmt.Printf(format+"\n", args...)
}

// Fatalf formats and prints a message if log level is debug or higher.
func (l *Logger) Fatalf(format string, args ...any) {
	//implementation
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelFatal {
		return
	}
	fmt.Printf(format+"\n", args...)
}

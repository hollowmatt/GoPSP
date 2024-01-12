package pocketlog

import (
	"fmt"
)

// Logger is used to log information
type Logger struct {
	threshold Level
}

func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

// Debugf formats and prints a message if log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
	//implementation
	fmt.Println(format)
}

// Infof formats and prints a message if log level is debug or higher
func (l *Logger) InfoF(format string, args ...any) {
	//implementation
	fmt.Println(format)
}

// Warningf formats and prints a message if log level is debug or higher
func (l *Logger) Warningf(format string, args ...any) {
	//implementation
	fmt.Println(format)
}

// Errorf formats and prints a message if log level is debug or higher
func (l *Logger) Errorf(format string, args ...any) {
	//implementation
	fmt.Println(format)
}

// Fatalf formats and prints a message if log level is debug or higher
func (l *Logger) Fatalf(format string, args ...any) {
	//implementation
	fmt.Println(format)
}

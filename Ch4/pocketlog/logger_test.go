package pocketlog_test

import (
	"os"

	"hollowmatt.com/logger/pocketlog"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, os.Stdout)
	debugLogger.Debugf("Hello, %s", "debug")
	// Output: Hello, debug
}

package pocketlog_test

import (
	"hollowmatt.com/logger/pocketlog"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: debug:: Hello, world
}

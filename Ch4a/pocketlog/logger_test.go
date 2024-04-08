package pocketlog_test

import "hollowmatt/new_logger/pocketlog"

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, nil)
	debugLogger.Debugf("hello, %s", "world")
	// Output: hello, world
}

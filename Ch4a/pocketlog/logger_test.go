package pocketlog_test

import (
	"hollowmatt/new_logger/pocketlog"
	"testing"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("hello, %s", "world")
	// Output: hello, world
}

const (
	debugMessage = "Debug is in the ointment"
	infoMessage  = "FYI this is information"
	errorMessage = "borken"
)

func TestLogger_DebugfInfofErrorf(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testCase{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: debugMessage + "\n" + infoMessage + "\n" + errorMessage + "\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: infoMessage + "\n" + errorMessage + "\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: errorMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			testedLogger.Debugf(debugMessage)
			testedLogger.Infof(infoMessage)
			testedLogger.Errorf(errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

// testWriter is a struct that implements io.Writer
// we use it to validate we can write to a specific output
type testWriter struct {
	contents string
}

// Write implements the io.Writer interface
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}

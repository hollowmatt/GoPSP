package pocketlog

// Level represents an available log level.
type Level byte

const (
	// LevelDebug represents the lowest level of log, for debugging purposes.
	LevelDebug Level = iota
	// LevelInfo represents a level of info deemped valuable.
	LevelInfo
	// LevelWarn represents a warning, but not an error.
	LevelWarn
	// LevelError represents highest level, used to trace errors.
	LevelError
	// LevelFatal represents a fatal event.
	LevelFatal
)

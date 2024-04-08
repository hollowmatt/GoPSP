package pocketlog

// Level representa an available logging level
type Level byte

const (
	// LevelDebug represents the lowest level of logging, used for debugging purposes
	LevelDebug Level = iota
	// LevelInfo represents a level where information is passed out, along with errors (think of QA)
	LevelInfo Level
	// LevelError represents errors in the system, but not fatal (error occurs, system continues)
	LevelError Level
	// LevelFatal represents something fatal occuring in the system, causing the system to stop
	LevelFatal Level
)

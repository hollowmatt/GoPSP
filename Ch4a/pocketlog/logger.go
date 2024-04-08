package pocketlog

type Logger struct {
	threshold Level
}

func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

func (l *Logger) Debugf(format string, args ...any) {
	//implement me
}

func (l *Logger) Infof(format string, args ...any) {
	//impleent me
}

func (l *Logger) Warnf(format string, args ...any) {
	//impleent me
}

func (l *Logger) Errorf(format string, args ...any) {
	//impleent me
}

func (l *Logger) Fatalf(format string, args ...any) {
	//impleent me
}

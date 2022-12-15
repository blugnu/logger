package logger

import (
	"github.com/sirupsen/logrus"
)

var logrusLevel = map[Level]logrus.Level{
	Trace: logrus.TraceLevel,
	Debug: logrus.DebugLevel,
	Info:  logrus.InfoLevel,
	Warn:  logrus.WarnLevel,
	Error: logrus.ErrorLevel,
	Fatal: logrus.FatalLevel,
}

type LogrusAdapter struct {
	*logrus.Logger
}

func (log *LogrusAdapter) Emit(level Level, s string) {
	log.Logger.Log(logrusLevel[level], s)
}

func (log *LogrusAdapter) NewEntry() Adapter {
	entry := logrus.NewEntry(log.Logger)
	return &logrusEntryAdapter{entry}
}

func (log *LogrusAdapter) WithField(name string, value any) Adapter {
	return &logrusEntryAdapter{log.Logger.WithField(name, value)}
}

type logrusEntryAdapter struct {
	*logrus.Entry
}

func (log *logrusEntryAdapter) Emit(level Level, s string) {
	log.Entry.Log(logrusLevel[level], s)
}

func (log *logrusEntryAdapter) NewEntry() Adapter {
	entry := logrus.NewEntry(log.Logger)
	return &logrusEntryAdapter{entry}
}

func (log *logrusEntryAdapter) WithField(name string, value any) Adapter {
	return &logrusEntryAdapter{log.Entry.WithField(name, value)}
}

package logger

import (
	"context"
	"fmt"
)

type Logger struct {
	context.Context
	Adapter
}

func (log *Logger) Trace(s string)                    { log.emit(Trace, s) }
func (log *Logger) Tracef(format string, args ...any) { log.Trace(fmt.Sprintf(format, args...)) }
func (log *Logger) Debug(s string)                    { log.emit(Debug, s) }
func (log *Logger) Debugf(format string, args ...any) { log.Debug(fmt.Sprintf(format, args...)) }
func (log *Logger) Info(s string)                     { log.emit(Info, s) }
func (log *Logger) Infof(format string, args ...any)  { log.Info(fmt.Sprintf(format, args...)) }
func (log *Logger) Warn(s string)                     { log.emit(Warn, s) }
func (log *Logger) Warnf(format string, args ...any)  { log.Warn(fmt.Sprintf(format, args...)) }
func (log *Logger) Error(err error)                   { log.emit(Error, err.Error()) }
func (log *Logger) Errorf(format string, args ...any) { log.Error(fmt.Errorf(format, args...)) }
func (log *Logger) Fatal(s string)                    { log.emit(Fatal, s); exit(1) }
func (log *Logger) Fatalf(format string, args ...any) { log.Fatal(fmt.Sprintf(format, args...)) }
func (log *Logger) FatalError(err error)              { log.Fatal(err.Error()) }

func (log *Logger) emit(level Level, s string) {
	entry := log.fromContext()
	entry.Emit(level, s)
}

func (log *Logger) fromContext() *Logger {
	ctx := log.Context
	if ctx == nil {
		ctx = context.Background()
	}

	logger := &Logger{ctx, log.NewEntry()}
	for _, decorate := range enrichmentFuncs {
		logger = decorate(ctx, logger)
	}

	return logger
}

func (log *Logger) WithContext(ctx context.Context) *Logger {
	entry := log.NewEntry()
	return &Logger{ctx, entry}
}

func (log *Logger) WithField(name string, value any) *Logger {
	entry := log.Adapter.WithField(name, value)
	return &Logger{log.Context, entry}
}

package logger

import (
	"context"
	"fmt"

	"github.com/blugnu/go-errorcontext"
)

type Logger interface {
	FromContext(context.Context) LogEntry
	NewEntry() LogEntry
}

type LogEntry interface {
	Debug(s string)
	Debugf(format string, args ...any)
	Error(err error)
	Errorf(format string, args ...any)
	Fatal(s string)
	Fatalf(format string, args ...any)
	FatalError(err error)
	Info(s string)
	Infof(format string, args ...any)
	Trace(s string)
	Tracef(format string, args ...any)
	Warn(s string)
	Warnf(format string, args ...any)
	WithContext(context.Context) LogEntry
	WithField(name string, value any) LogEntry
}

type logger struct {
	context.Context
	Adapter
}

func (log *logger) Trace(s string)                    { log.emit(Trace, s) }
func (log *logger) Tracef(format string, args ...any) { log.Trace(fmt.Sprintf(format, args...)) }
func (log *logger) Debug(s string)                    { log.emit(Debug, s) }
func (log *logger) Debugf(format string, args ...any) { log.Debug(fmt.Sprintf(format, args...)) }
func (log *logger) Info(s string)                     { log.emit(Info, s) }
func (log *logger) Infof(format string, args ...any)  { log.Info(fmt.Sprintf(format, args...)) }
func (log *logger) Warn(s string)                     { log.emit(Warn, s) }
func (log *logger) Warnf(format string, args ...any)  { log.Warn(fmt.Sprintf(format, args...)) }

func (log *logger) Error(err error) {
	ctx := errorcontext.FromError(log.Context, err)
	entry := log.fromContext(ctx)
	entry.emit(Error, err.Error())
}

func (log *logger) Errorf(format string, args ...any) {
	entry := log
	for _, a := range args {
		if err, isError := a.(error); isError {
			ctx := errorcontext.FromError(log.Context, err)
			entry = entry.fromContext(ctx)
			break
		}
	}
	entry.Error(fmt.Errorf(format, args...))
}

func (log *logger) Fatal(s string) {
	log.emit(Fatal, s)
	exit(1)
}

func (log *logger) Fatalf(format string, args ...any) {
	entry := log
	for _, a := range args {
		if err, isError := a.(error); isError {
			ctx := errorcontext.FromError(log.Context, err)
			entry = entry.fromContext(ctx)
			break
		}
	}
	entry.Fatal(fmt.Sprintf(format, args...))
}

func (log *logger) FatalError(err error) {
	ctx := errorcontext.FromError(log.Context, err)
	entry := log.fromContext(ctx)
	entry.Fatal(err.Error())
}

func (log *logger) WithContext(ctx context.Context) LogEntry {
	entry := log.Adapter.NewEntry()
	return &logger{ctx, entry}
}

func (log *logger) WithField(name string, value any) LogEntry {
	entry := log.Adapter.WithField(name, value)
	return &logger{log.Context, entry}
}

func (log *logger) emit(level Level, s string) {
	entry := log.fromContext(log.Context)
	entry.Emit(level, s)
}

func (log *logger) fromContext(ctx context.Context) *logger {
	if ctx == nil {
		ctx = context.Background()
	}

	logger := &logger{ctx, log.Adapter.NewEntry()}

	var entry LogEntry = logger
	for _, decorate := range enrichmentFuncs {
		entry = decorate(ctx, entry)
	}

	return logger
}

func (log *logger) FromContext(ctx context.Context) LogEntry {
	return log.fromContext(ctx)
}

func (log *logger) NewEntry() LogEntry {
	return log.fromContext(log.Context)
}

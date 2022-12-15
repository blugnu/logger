package logger

import (
	"testing"

	logrus "github.com/sirupsen/logrus"

	log "github.com/blugnu/go-logspy"
)

func TestLogrusAdapter(t *testing.T) {
	// ARRANGE
	logger := logrus.New()
	logger.SetOutput(log.Sink())
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})
	sut := &LogrusAdapter{logger}

	testcases := []struct {
		name   string
		fn     func(string)
		output string
	}{
		{name: "trace", fn: func(s string) { sut.Emit(Trace, s) }, output: "level=trace msg=\"entry text\"\n"},
		{name: "debug", fn: func(s string) { sut.Emit(Debug, s) }, output: "level=debug msg=\"entry text\"\n"},
		{name: "info", fn: func(s string) { sut.Info(s) }, output: "level=info msg=\"entry text\"\n"},
		{name: "warn", fn: func(s string) { sut.Warn(s) }, output: "level=warning msg=\"entry text\"\n"},
		{name: "error", fn: func(s string) { sut.Emit(Error, s) }, output: "level=error msg=\"entry text\"\n"},
		{name: "fatal", fn: func(s string) { sut.Emit(Fatal, s) }, output: "level=fatal msg=\"entry text\"\n"},
		{name: "debug and error", fn: func(s string) { sut.Emit(Debug, s); sut.Emit(Error, s) }, output: "level=debug msg=\"entry text\"\nlevel=error msg=\"entry text\"\n"},
		{name: "withfield", fn: func(s string) {
			a := sut
			b := sut.WithField("field", "data")

			t.Run("returns new logger", func(t *testing.T) {
				wanted := true
				got := a != b
				if wanted != got {
					t.Errorf("wanted %v, got %v", wanted, got)
				}
			})

			a.Emit(Info, s)
			b.Emit(Info, s)
		}, output: "level=info msg=\"entry text\"\nlevel=info msg=\"entry text\" field=data\n"},
		{name: "newentry", fn: func(s string) {
			a := sut
			b := sut.NewEntry()

			t.Run("returns new logger", func(t *testing.T) {
				wanted := true
				got := a != b
				if wanted != got {
					t.Errorf("wanted %v, got %v", wanted, got)
				}
			})

			a.Emit(Info, s)
			b.Emit(Info, s)
		}, output: "level=info msg=\"entry text\"\nlevel=info msg=\"entry text\"\n"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			defer log.Reset()
			// ACT
			tc.fn("entry text")

			// ASSERT
			wanted := tc.output
			got := log.String()
			if wanted != got {
				t.Errorf("\nwanted %q\ngot    %q", wanted, got)
			}
		})
	}
}

func TestLogrusEntryAdapter(t *testing.T) {
	// ARRANGE
	logger := logrus.New()
	logger.SetOutput(log.Sink())
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})
	sut := &logrusEntryAdapter{logrus.NewEntry(logger)}

	testcases := []struct {
		name   string
		fn     func(string)
		output string
	}{
		{name: "debug", fn: func(s string) { sut.Emit(Debug, s) }, output: "level=debug msg=\"entry text\"\n"},
		{name: "info", fn: func(s string) { sut.Info(s) }, output: "level=info msg=\"entry text\"\n"},
		{name: "warn", fn: func(s string) { sut.Warn(s) }, output: "level=warning msg=\"entry text\"\n"},
		{name: "error", fn: func(s string) { sut.Emit(Error, s) }, output: "level=error msg=\"entry text\"\n"},
		{name: "debug and error", fn: func(s string) { sut.Emit(Debug, s); sut.Emit(Error, s) }, output: "level=debug msg=\"entry text\"\nlevel=error msg=\"entry text\"\n"},
		{name: "withfield", fn: func(s string) {
			a := sut
			b := sut.WithField("field", "data")

			t.Run("returns new logger", func(t *testing.T) {
				wanted := true
				got := a != b
				if wanted != got {
					t.Errorf("wanted %v, got %v", wanted, got)
				}
			})

			a.Emit(Info, s)
			b.Emit(Info, s)
		}, output: "level=info msg=\"entry text\"\nlevel=info msg=\"entry text\" field=data\n"},
		{name: "newentry", fn: func(s string) {
			a := sut
			b := sut.NewEntry()

			t.Run("returns new logger", func(t *testing.T) {
				wanted := true
				got := a != b
				if wanted != got {
					t.Errorf("wanted %v, got %v", wanted, got)
				}
			})

			a.Emit(Info, s)
			b.Emit(Info, s)
		}, output: "level=info msg=\"entry text\"\nlevel=info msg=\"entry text\"\n"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			defer log.Reset()

			// ACT
			tc.fn("entry text")

			// ASSERT
			wanted := tc.output
			got := log.String()
			if wanted != got {
				t.Errorf("\nwanted %q\ngot    %q", wanted, got)
			}
		})
	}
}

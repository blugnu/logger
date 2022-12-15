package logger

type Level int

// NOTE: though levels are modelled on logrus, Panic is not supported
const (
	Fatal Level = iota
	Error
	Warn
	Info
	Debug
	Trace
)

package logger

type Adapter interface {
	Emit(Level, string)
	NewEntry() Adapter
	WithField(string, any) Adapter
}

package logger

type NulAdapter struct{}

func noop() {}

func (*NulAdapter) Emit(Level, string)                { noop() }
func (nul *NulAdapter) NewEntry() Adapter             { return nul }
func (nul *NulAdapter) WithField(string, any) Adapter { return nul }

func Nul() Logger {
	return &logger{Adapter: &NulAdapter{}}
}

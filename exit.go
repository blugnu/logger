package logger

import "os"

var ExitFn func(int) = os.Exit

func exit(code int) {
	ExitFn(code)
}

package logger

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

var logPrefix = map[Level]string{
	Trace: "TRACE",
	Debug: "DEBUG",
	Info:  "INFO",
	Warn:  "WARN",
	Error: "ERROR",
	Fatal: "FATAL",
}

type LogAdapter struct {
	fields map[string]any
}

func (a *LogAdapter) fieldData() string {
	keys := make([]string, 0, len(a.fields))
	for k := range a.fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	data := ""
	for _, k := range keys {
		kf := "%s"
		vf := "%s"
		if strings.Contains(k, " ") {
			kf = "%q"
		}
		v := a.fields[k]
		vs := fmt.Sprintf("%s", v)
		if strings.Contains(vs, " ") {
			vf = "%q"
		}
		data = data + fmt.Sprintf(kf+"="+vf+" ", k, v)
	}
	return data
}

func (a *LogAdapter) Emit(level Level, s string) {
	log.Printf(a.fieldData() + logPrefix[level] + ": " + s)
}

func (log *LogAdapter) NewEntry() Adapter {
	fields := map[string]any{}
	for k, v := range log.fields {
		fields[k] = v
	}
	return &LogAdapter{fields}
}

func (log *LogAdapter) WithField(name string, value any) Adapter {
	entry := log.NewEntry().(*LogAdapter)
	entry.fields[name] = value
	return entry
}

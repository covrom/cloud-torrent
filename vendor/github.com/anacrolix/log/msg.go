package log

import (
	"fmt"
	"runtime"
)

type Msg struct {
	fields  map[string][]interface{}
	values  map[interface{}]struct{}
	text    string
	callers [1]uintptr
}

func Fmsg(format string, a ...interface{}) (m Msg) {
	m.text = fmt.Sprintf(format, a...)
	m.setCallers(1)
	return
}

func Str(s string) (m Msg) {
	m.text = s
	m.setCallers(1)
	return
}

func (m *Msg) setCallers(skip int) {
	runtime.Callers(skip+2, m.callers[:])
}

func (msg Msg) Add(key string, value interface{}) Msg {
	if msg.fields == nil {
		msg.fields = make(map[string][]interface{})
	}
	msg.fields[key] = append(msg.fields[key], value)
	return msg
}

func (msg Msg) Log(l *Logger) Msg {
	l.Handle(msg)
	return msg
}

func (m Msg) Values() map[interface{}]struct{} {
	return m.values
}

func (m Msg) AddValue(value interface{}) Msg {
	if m.values == nil {
		m.values = make(map[interface{}]struct{})
	}
	m.values[value] = struct{}{}
	return m
}

func (m Msg) AddValues(values ...interface{}) Msg {
	for _, v := range values {
		m = m.AddValue(v)
	}
	return m
}

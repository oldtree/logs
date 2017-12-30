package logs

import (
	"fmt"
	"io"
	"runtime"
	"time"
)

type Logs struct {
	Loglevel uint
	wc       io.WriteCloser
}

func NewLog(level uint, buffersize int64, writeTo ...*IOWrite) *Logs {
	if len(writeTo) <= 0 {
		writeTo = append(writeTo, NewStdOutWriter())
	}
	return &Logs{
		Loglevel: level,
		wc:       NewCacheBuffer(buffersize, writeTo),
	}
}

func (l *Logs) Close() error {
	return l.wc.Close()
}

func (l *Logs) Debug(format string, values ...interface{}) {
	if l.Loglevel > uint(DebugLogLevel) {
		return
	}
	pc, file, line, _ := runtime.Caller(1)
	fformat := fmt.Sprintf(string(DebugLogLevelFormat), time.Now().Format(DefaultTimeformat), file, line, runtime.FuncForPC(pc).Name(), format)
	info := fmt.Sprintf(fformat, values)
	l.wc.Write([]byte(info))
	return
}

func (l *Logs) Info(format string, values ...interface{}) {
	if l.Loglevel > uint(InfoLogLevel) {
		return
	}
	pc, file, line, _ := runtime.Caller(1)
	fformat := fmt.Sprintf(string(DebugLogLevelFormat), time.Now().Format(DefaultTimeformat), file, line, runtime.FuncForPC(pc).Name(), format)
	info := fmt.Sprintf(fformat, values)
	l.wc.Write([]byte(info))
	return
}

func (l *Logs) Warn(format string, values ...interface{}) {
	if l.Loglevel > uint(WarnLogLevel) {
		return
	}
	pc, file, line, _ := runtime.Caller(1)
	fformat := fmt.Sprintf(string(DebugLogLevelFormat), time.Now().Format(DefaultTimeformat), file, line, runtime.FuncForPC(pc).Name(), format)
	info := fmt.Sprintf(fformat, values)
	l.wc.Write([]byte(info))
	return
}

func (l *Logs) Error(format string, values ...interface{}) {
	if l.Loglevel > uint(ErrorLogLevel) {
		return
	}
	pc, file, line, _ := runtime.Caller(1)
	fformat := fmt.Sprintf(string(DebugLogLevelFormat), time.Now().Format(DefaultTimeformat), file, line, runtime.FuncForPC(pc).Name(), format)
	info := fmt.Sprintf(fformat, values)
	l.wc.Write([]byte(info))
	return
}

func (l *Logs) Panic(format string, values ...interface{}) {
	if l.Loglevel > uint(PanicLogLevel) {
		return
	}
	pc, file, line, _ := runtime.Caller(1)
	fformat := fmt.Sprintf(string(DebugLogLevelFormat), time.Now().Format(DefaultTimeformat), file, line, runtime.FuncForPC(pc).Name(), format)
	info := fmt.Sprintf(fformat, values)
	l.wc.Write([]byte(info))
	return
}

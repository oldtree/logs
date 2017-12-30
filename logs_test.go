package logs

import (
	"testing"
)

func Test_log_console(t *testing.T) {
	l := NewLog(uint(TraceLogLevel), 2, NewStdOutWriter())
	l.Info("hello")
	l.Info("hello %s", "world")
	l.Debug("hello")
	l.Debug("hello %s", "world")
	l.Warn("hello")
	l.Warn("hello %s", "world")
	l.Error("hello")
	l.Error("hello %s", "world")
	l.Panic("hello")
	l.Panic("hello %s", "world")
	l.Close()
}

func Test_log_file(t *testing.T) {
	l := NewLog(uint(TraceLogLevel), 2, NewFileIOWriter(4096, "/Users/a123/Grapes/src/github.com/FlyCynomys/logs/logs"))
	l.Info("hello")
	l.Info("hello %s", "world")
	l.Debug("hello")
	l.Debug("hello %s", "world")
	l.Warn("hello")
	l.Warn("hello %s", "world")
	l.Error("hello")
	l.Error("hello %s", "world")
	l.Panic("hello")
	l.Panic("hello %s", "world")
	l.Close()
}

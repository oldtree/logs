package logs

import (
	"io"
	"net"
	"os"
	"path/filepath"
	"time"
)

type MultiWriter struct {
	writers []io.WriteCloser
}

func (t *MultiWriter) Write(p []byte) (n int, err error) {
	for _, w := range t.writers {
		n, err = w.Write(p)
		if err != nil {
			return
		}
		if n != len(p) {
			err = io.ErrShortWrite
			return
		}
	}
	return len(p), nil
}

func (t *MultiWriter) Close() error {
	var err = make([]error, len(t.writers))
	for index, w := range t.writers {
		err[index] = w.Close()
	}
	for _, value := range err {
		if value != nil {
			return value
		}
	}
	return nil
}

func NewMultiWriter(writers ...io.WriteCloser) io.Writer {
	allWriters := make([]io.WriteCloser, 0, len(writers))
	for _, w := range writers {
		if mw, ok := w.(*MultiWriter); ok {
			allWriters = append(allWriters, mw.writers...)
		} else {
			allWriters = append(allWriters, w)
		}
	}
	return &MultiWriter{allWriters}
}

const (
	STDOUTPUT = iota
	STDERR
	STDFILE
	STDTCP
)

type IOWrite struct {
	SizeLimit int64
	Path      string
	IOType    int
	Fd        io.Writer
}

func NewStdTcpWriter(writer io.Writer) *IOWrite {
	return &IOWrite{
		IOType: STDTCP,
		Fd:     writer,
	}
}

func NewFileIOWriter(size int64, path string) *IOWrite {
	os.MkdirAll(path, 0777)
	writer, err := os.OpenFile(filepath.Join([]string{path, time.Now().Format(DefaultTimeformat) + ".log"}...), os.O_CREATE|os.O_RDWR|os.O_SYNC, 0777)
	//writer, err := os.Create(filepath.Join([]string{path, time.Now().Format(DefaultTimeformat) + ".log"}...))
	if err != nil {
		println("+++++++++++++++++++", err.Error())
	}
	return &IOWrite{
		SizeLimit: size,
		Path:      path,
		IOType:    STDFILE,
		Fd:        writer,
	}
}

func NewStdOutWriter() *IOWrite {
	return &IOWrite{
		IOType: STDOUTPUT,
		Fd:     os.Stdout,
	}
}

func NewStdErrWriter() *IOWrite {
	return &IOWrite{
		IOType: STDERR,
		Fd:     os.Stderr,
	}
}

func (w *IOWrite) Close() error {
	return w.Fd.(*os.File).Close()
}

func (w *IOWrite) Write(data []byte) (n int, err error) {
	println(w.IOType, string(data))
	switch w.IOType {
	case STDOUTPUT:
		return w.Fd.Write(data)
	case STDERR:
		return w.Fd.Write(data)
	case STDFILE:
		w.Fd.(*os.File).Sync()
		if info, err := w.Fd.(*os.File).Stat(); err == nil && info.Size() > w.SizeLimit {
			w.Fd.(*os.File).Close()
			w.Fd, err = os.Create(filepath.Join([]string{w.Path, time.Now().Format(DefaultTimeformat) + ".log"}...))
			if err != nil {
				return 0, err
			} else {
				return w.Fd.Write(data)
			}
		} else if err != nil {
			println("-------------------------------------", err.Error())
			return 0, err
		} else {
			return w.Fd.Write(data)
		}
	case STDTCP:
		n, err = w.Fd.Write(data)
		if err != nil {
			if err == net.ErrWriteToConnected {
				w.Fd = os.Stdout
			}
			return 0, err
		}
		return n, err
	}
	return
}

package logs

import (
	"bytes"
	"io"
	"sync"
)

const DefaultBufferSize = 4096

type CacheBuffer struct {
	LogChan  chan []byte
	StopChan chan struct{}
	Writer   *bytes.Buffer
	conter   int64
	sync.RWMutex
	OutSet     io.Writer
	BufferSize int64
}

func (cb *CacheBuffer) loop() {
	for {
		select {
		case data := <-cb.LogChan:
			if cb.conter > cb.BufferSize {
				cb.Writer.WriteTo(cb.OutSet)
				cb.Writer.Reset()
				cb.conter = 0
			}
			cb.Writer.Write(data)
			cb.conter++
		case <-cb.StopChan:
			return
		}
	}
}

func (cb *CacheBuffer) Write(data []byte) (number int, err error) {
	//println(string(data))
	cb.LogChan <- data
	return len(data), nil
}

func (cb *CacheBuffer) Close() error {
	close(cb.LogChan)
	close(cb.StopChan)
	err := cb.OutSet.(*MultiWriter).Close()
	return err
}

func NewCacheBuffer(buffersize int64, output []*IOWrite) *CacheBuffer {
	cb := &CacheBuffer{
		BufferSize: buffersize,
		Writer:     bytes.NewBuffer(make([]byte, DefaultBufferSize, DefaultBufferSize)),
		LogChan:    make(chan []byte, 4096),
	}
	var outputlist []io.WriteCloser
	for _, value := range output {
		outputlist = append(outputlist, value)
	}
	cb.OutSet = NewMultiWriter(outputlist...)
	cb.StopChan = make(chan struct{}, 1)
	go cb.loop()
	return cb
}

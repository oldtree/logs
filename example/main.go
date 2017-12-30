package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/FlyCynomys/logs"
)

func main() {
	/*l := logs.NewLog(uint(logs.TraceLogLevel), 4, logs.NewStdOutWriter())
	fmt.Println(time.Now())
	for index := 0; index < 1000; index++ {
		l.Info("hello %s \n", "tester")
	}*/
	//err := l.Close()
	//println("------------", err)
	l2 := logs.NewLog(uint(logs.TraceLogLevel), 4, logs.NewFileIOWriter(50, "/Users/a123/Grapes/src/github.com/FlyCynomys/logs/logs"))
	for index := 0; index < 100; index++ {
		l2.Info("hello %s \n", "tester")
	}

	//l2.Close()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-sc
	return
}

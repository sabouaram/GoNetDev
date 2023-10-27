package GLogger

import (
	"sync"
	"log"
	"os"
	"io"
)

var once sync.Once
var Glogger *GLogger

type GLogger struct {
	*log.Logger
	filename string
}

func GetInstance() *GLogger {
	once.Do(func() {
		Glogger = createLogger("/var/log/Gonetdev.log")
	})
	return Glogger
}

func createLogger(fname string) *GLogger {
	file, _ := os.OpenFile(fname,os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
	mw := io.MultiWriter(os.Stdout, file)
	return &GLogger{
		filename: fname,
		Logger:   log.New(mw, "", log.Lshortfile),
	}
}

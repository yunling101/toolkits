package logger

import (
	"log"
	"os"

	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	AppPath string
	Logs    *logrus.Logger
	File    *os.File
}

func (o *Logger) Init() {
	var err error

	o.File, err = os.OpenFile(o.AppPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	o.Logs = logrus.New()
	o.Logs.SetLevel(logrus.DebugLevel)
	o.Logs.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	newHook := filename.NewHook()
	o.Logs.AddHook(newHook)
	o.Logs.Out = o.File
}

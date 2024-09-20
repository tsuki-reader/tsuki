package core

import (
	"log"
	"os"
)

type LoggerInterface interface {
	Fatal(v ...interface{})
	Println(v ...interface{})
}

type StandardLogger struct {
	*log.Logger
}

func (l *StandardLogger) Fatal(v ...interface{}) {
	l.Logger.Fatal(v...)
}

func (l *StandardLogger) Println(v ...interface{}) {
	l.Logger.Println(v...)
}

func NewLogger() LoggerInterface {
	return &StandardLogger{Logger: log.New(os.Stderr, "", log.LstdFlags)}
}

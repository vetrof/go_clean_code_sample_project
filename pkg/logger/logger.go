package logger

import "log"

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Info(v ...any)  { log.Println(v...) }
func (l *Logger) Error(v ...any) { log.Println(v...) }
func (l *Logger) Fatal(v ...any) { log.Fatal(v...) }

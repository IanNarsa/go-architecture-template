package logger

import (
	"log"
)

func Info(v ...interface{}) {
	log.Println("[INFO]", v)
}

func Error(v ...interface{}) {
	log.Println("[ERROR]", v)
}

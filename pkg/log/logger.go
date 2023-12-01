package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"log"
	"os"
	"time"
)

func InitLogger(env string) {
	var path = "/app/logs"
	log.SetFlags(0)
	if env == "dev" {
		log.SetOutput(os.Stdout)
		return
	}

	writer, err := rotatelogs.New(
		fmt.Sprintf("%s/app-%s.log", path, "%Y-%m-%d"),
		rotatelogs.WithLinkName(fmt.Sprintf("%s/link.log", path)),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		log.Fatalf("failed to initialize log file : %v", err)
	}

	log.SetOutput(writer)
}

func write(level, msg string) {
	log.Printf("%s %s %s",
		time.Now().Format("15:04:05"),
		level,
		msg,
	)
}

func Debugf(msg string, data ...interface{}) {
	if os.Getenv("LOG_LEVEL") == "debug" {
		var m = fmt.Sprintf(msg, data...)
		write("DBG", m)
	}
}

func Infof(msg string, data ...interface{}) {
	var m = fmt.Sprintf(msg, data...)
	write("INF", m)
}

func Warnf(msg string, data ...interface{}) {
	var m = fmt.Sprintf(msg, data...)
	write("WRN", m)
}

func Errorf(msg string, data ...interface{}) {
	var m = fmt.Sprintf(msg, data...)
	write("ERR", m)
}

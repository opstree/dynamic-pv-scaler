package logger

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func LogStdout() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
	})
	mw := io.MultiWriter(os.Stdout)
	log.SetOutput(mw)
}

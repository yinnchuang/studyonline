package log

import (
	"io"
	"log"
	"os"
)

type AppLogger struct {
	LoggerFile *os.File
	Logger     *log.Logger
}

func (l *AppLogger) Init(filename string) {
	var err error
	l.LoggerFile, err = os.OpenFile(filename,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("open log file: %v", err)
	}
	l.Logger = log.New(io.MultiWriter(os.Stdout, l.LoggerFile),
		"", log.LstdFlags|log.Lmsgprefix)
}

func (l *AppLogger) Close() {
	l.LoggerFile.Close()
}

func (l *AppLogger) Log(logInfo string) {
	l.Logger.Println(logInfo)
}

var (
	DownloadLogger *AppLogger
	CommonLogger   *AppLogger
)

func Init() {
	DownloadLogger = new(AppLogger)
	DownloadLogger.Init("download.log")
	CommonLogger = new(AppLogger)
	CommonLogger.Init("common.log")
}

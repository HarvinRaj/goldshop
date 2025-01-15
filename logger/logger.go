package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	ErrorLog *log.Logger
	InfoLog  *log.Logger
)

func init() {

	logDir := "./logs"

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.Mkdir(logDir, 0755)
		if err != nil {
			log.Fatalf("Failed to make log directory, %v", err)
		}
	}

	logFilePath := filepath.Join(logDir, "app.log")

	logFile, err2 := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		log.Fatalf("Failed to open log file: %v", err2)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	ErrorLog = log.New(multiWriter, "- [ERROR]: ", log.LstdFlags|log.Lmsgprefix)
	InfoLog = log.New(multiWriter, "- [INFO]: ", log.LstdFlags|log.Lmsgprefix)
}

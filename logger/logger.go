package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
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

	currentDate := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf("%s_app.log", currentDate)

	logFilePath := filepath.Join(logDir, logFileName)

	logFile, err2 := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		log.Fatalf("Failed to open log file: %v", err2)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	ErrorLog = log.New(multiWriter, "- [ERROR]: ", log.LstdFlags|log.Lmsgprefix)
	InfoLog = log.New(multiWriter, "- [INFO]: ", log.LstdFlags|log.Lmsgprefix)
}

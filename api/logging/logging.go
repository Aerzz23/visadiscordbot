package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

// CreateLogFile creates a new log file at the specified path using a given application name and the current time & date
func CreateLogFile(logPath string, appName string) (*os.File, error) {
	f, err := os.OpenFile(fmt.Sprintf("%s/log_%s_%s", logPath, appName, getTimeFormatted()), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		return nil, err
	}
	log.SetOutput(f)
	log.Println("This is a test log entry from func")
	return f, nil
}

func getTimeFormatted() string {
	t := time.Now()
	timeFormatted := t.Format("2006-01-02T15:04:05")
	return timeFormatted
}

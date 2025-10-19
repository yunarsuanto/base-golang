package log

import (
	"fmt"
	"os"
	"time"

	"github.com/yunarsuanto/base-go/constants"

	log "github.com/sirupsen/logrus"
)

const (
	gitignore = ".gitignore"
)

// PrintTimestamp function to print timestamp for every log
func PrintTimestamp() {
	formatter := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.DateTime,
	}

	log.SetFormatter(formatter)
}

// PrintOutputToFile function to write log output to log file
// Returns *os.File
func PrintOutputToFile() *os.File {
	now := time.Now().Format(time.DateOnly)
	filename := fmt.Sprintf("%s%s", constants.LogDir, now)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)
	log.SetLevel(log.WarnLevel)

	return f
}

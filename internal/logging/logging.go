package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/log"
)

type LogConfig struct {
	Name     string
	Filepath string
}

func (l *LogConfig) Init() *log.Logger {
	// NOTE: Create .logging directory if it doesn't exist
	err := os.MkdirAll(".logging", os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create .logging directory")
		os.Exit(1)
	}

	fullPath := filepath.Join(".logging", l.Filepath)

	// NOTE: Appends to avoid overwriting the log file
	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file")
		os.Exit(1)
	}

	options := log.Options{
		ReportCaller:    false,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "[" + l.Name + "] ",
	}

	return log.NewWithOptions(file, options)
}

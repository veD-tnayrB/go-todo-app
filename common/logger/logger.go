package logger

import (
	"log/slog"
	"os"
	"time"
)

type Logger struct {
	logger      *slog.Logger
	logFolder   string
	currentfile *os.File
	filename    string
}

const (
	// Default file permissions (owner read/write, group/others read)
	filePerms = 0644
	// Default directory permissions (owner full access, group/others read/execute)
	dirPerms = 0755
)

func NewLogger(logFolder string) (*Logger, error) {
	instance := Logger{logFolder: logFolder}
	instance.generateBaseFolder()
	instance.generateDayFile()

	return &instance, nil
}

func (l *Logger) generateBaseFolder() {
	if l.logFolder == "" {
		panic("Log folder not specified, please provide a name for log folder")
	}

	err := os.MkdirAll(l.logFolder, dirPerms)
	if err != nil {
		panic(err)
	}
}

func (l *Logger) generateDayFile() {
	filename := l.getDayFileName()
	path := l.logFolder + "/" + filename + ".txt"

	if l.currentfile != nil {
		l.currentfile.Close()
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, filePerms)
	if err != nil {
		panic("Error while creating the day file" + err.Error())
	}

	l.currentfile = file
	l.filename = filename
	l.logger = slog.New(slog.NewJSONHandler(l.currentfile, nil))
}

func (l *Logger) getDayFileName() string {
	return time.Now().Format("02-01-2006")
}

func (l *Logger) ensureDailyRotation() {
	if l.getDayFileName() != l.filename {
		l.generateDayFile()
	}
}

func (l *Logger) Info(mgs string, args ...interface{}) {
	l.ensureDailyRotation()
	l.logger.Info(mgs, args...)
}

func (l *Logger) Error(mgs string, args ...interface{}) {
	l.ensureDailyRotation()
	l.logger.Error(mgs, args...)
}

func (l *Logger) Debug(mgs string, args ...interface{}) {
	l.ensureDailyRotation()
	l.logger.Debug(mgs, args...)
}

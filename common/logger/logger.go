package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"time"
)

type Logger struct {
	logger      *slog.Logger
	logFolder   string
	currentfile *os.File
}

type logEntry struct {
	Time  time.Time              `json:"time"`
	Level slog.Level             `json:"level"`
	Msg   string                 `json:"msg"`
	Attrs map[string]interface{} `json:"attrs"`
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
	if err != nil && !errors.Is(err, fs.ErrExist) {
		panic(err)
	}
}

func (l *Logger) generateDayFile() {
	filename := time.Now().Format("02-01-2006")
	path := l.logFolder + "/" + filename + ".txt"

	if l.currentfile != nil {
		l.currentfile.Close()
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, filePerms)
	if err != nil {
		panic("Error while creating the day file" + err.Error())
	}

	l.currentfile = file
	l.logger = slog.New(slog.NewJSONHandler(l.currentfile, nil))
}

func (l *Logger) Info(mgs string, args ...interface{}) {
	fmt.Printf("%v", args)
	l.log(mgs, slog.LevelInfo, args...)
}

func (l *Logger) log(msg string, level slog.Level, args ...interface{}) {
	formattedArgs := map[string]interface{}{}

	if len(args) > 0 {
		if len(args)%2 != 0 {
			panic("Arguments must be in key-value pairs")
		}
		for i := 0; i < len(args); i += 2 {
			key, ok := args[i].(string)
			if !ok {
				panic("Key must be a string")
			}
			formattedArgs[key] = args[i+1]
		}
	}

	entry := logEntry{
		Time:  time.Now(),
		Msg:   msg,
		Level: level,
		Attrs: formattedArgs,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		panic(err)
	}

	l.currentfile.WriteString(string(data) + "\n")
}

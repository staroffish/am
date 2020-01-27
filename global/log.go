package global

import (
	"fmt"
	"os"
	"time"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type FileLogger struct {
	file *os.File
}

// LogStruct struct
type LogStruct struct {
	debugOn bool
	logger  Logger
}

var Log *LogStruct

const (
	errorStr = "ERROR %s"
	warnStr  = "WARN %s"
	infoStr  = "INFO %s"
	debugStr = "DEBUG %s"
)

// newLogger 生成一个新的Log结构体
func NewLogger(debugOn bool, logger Logger) {
	Log = &LogStruct{}
	Log.logger = logger
	Log.debugOn = debugOn
}

// Errorf - output error log
func (l *LogStruct) Errorf(format string, i ...interface{}) {
	if l == nil {
		return
	}
	l.logger.Printf(errorStr, fmt.Sprintf(format, i...))
}

// Infof - output info log
func (l *LogStruct) Infof(format string, i ...interface{}) {
	if l == nil {
		return
	}
	l.logger.Printf(infoStr, fmt.Sprintf(format, i...))
}

// Debugf - output debug log
func (l *LogStruct) Debugf(format string, i ...interface{}) {
	if l == nil || l.debugOn == false {
		return
	}
	l.logger.Printf(debugStr, fmt.Sprintf(format, i...))
}

// Warnf - output warn log
func (l *LogStruct) Warnf(format string, i ...interface{}) {
	if l == nil {
		return
	}
	l.logger.Printf(warnStr, fmt.Sprintf(format, i...))
}

// TraceLog - 打印函数的trace log
func TraceLog(log string) func() {
	Log.Debugf("Start %s", log)
	return func() {
		Log.Debugf("End %s", log)
	}
}

func NewFileLogger(logPath string) (*FileLogger, error) {
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return nil, err
	}
	return &FileLogger{file: logFile}, nil
}

func (l *FileLogger) Printf(format string, v ...interface{}) {
	timeStr := time.Now().Format("2006-1-2 15:04:05")
	logMsg := fmt.Sprintf(format, v...)
	fmt.Fprintf(l.file, "%s %s\n", timeStr, logMsg)
}

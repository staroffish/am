package global

import (
	"fmt"
	"log"
	"log/syslog"
)

// LogStruct struct
type LogStruct struct {
	debugOn bool
	logger  *log.Logger
}

var Log *LogStruct

const (
	errorStr = "ERROR %s"
	warnStr  = "WARN %s"
	infoStr  = "INFO %s"
	debugStr = "DEBUG %s"
)

// newLogger 生成一个新的Log结构体
func NewLogger(debugOn bool) {
	Log = &LogStruct{}
	Log.logger, _ = syslog.NewLogger(syslog.LOG_LOCAL0, 0)
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

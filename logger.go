package logger

import "fmt"

type DebugLogger struct {
	debugLog string
}

func (dl *DebugLogger) Log(logString string, args ...interface{}) {
	dl.debugLog += fmt.Sprintf(logString+"\n", args...)
}

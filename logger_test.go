package logger

import "testing"

func Test_DebugLogger_Log(t *testing.T) {
	dl := &DebugLogger{}

	dl.Log("Hello World!")

	if dl.debugLog != "Hello World!\n" {
		t.Fatal("Wrong Value")
	}
}

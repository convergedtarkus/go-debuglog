package debuglogger

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewDebugLogger(t *testing.T) {
	dl := NewDebugLogger("!!!!")
	require.NotNil(t, dl, "DebugLogger must not be nil")
	assert.Equal(t, "!!!!", dl.prefix, "Wrong DebugLogger prefix")
	assert.Equal(t, os.Stdout, dl.ioOut, "DebugLogger ioOut should be Stdout")
}

func Test_DebugLogger_ChangePrefix(t *testing.T) {
	dl := NewDebugLogger("!!!!")
	assert.Equal(t, "!!!!", dl.prefix, "Wrong DebugLogger initial prefix")

	dl.ChangePrefix("@@@@")
	assert.Equal(t, "@@@@", dl.prefix, "Wrong DebugLogger updated prefix")
}

func Test_DebugLogger_Log(t *testing.T) {
	tw := &TestWriter{}
	dl := newDebugLoggerWithIO("!!!! ", tw)

	dl.Log("Hi!")
	tw.VerifyCaptured(t, []string{"!!!! Hi!\n"})
}

// Other test should cover behavior, this verifies a real print doesn't panic.
func Test_DebugLogger_Log_StandardOut(t *testing.T) {
	dl := NewDebugLogger("!!!! ")
	dl.Log("Two is %d", 2)
}

func Test_DebugLogger_Log_WithArguments(t *testing.T) {
	tw := &TestWriter{}
	dl := newDebugLoggerWithIO("---- ", tw)

	dl.Log("Hi %d!", 1)
	tw.VerifyCaptured(t, []string{"---- Hi 1!\n"})
}

func Test_DebugLogger_Log_Error(t *testing.T) {
	tw := &TestWriter{}
	tw.mockWriteHandler = func(_ []byte) (int, error) {
		return -1, errors.New("mockWriteHandler test error")
	}
	dl := newDebugLoggerWithIO("!!!! ", tw)

	assert.Panics(t, func() { dl.Log("Hi!") }, "Log should panic")
}

// TestWriter is for testing code that writes out.
type TestWriter struct {
	capturedOutput []string

	// If this is non-nil, it will be called rather than the normal Write code.
	mockWriteHandler func(incoming []byte) (n int, err error)
}

// Write will convert the bytes to a string and append it to capturedOutput.
func (tw *TestWriter) Write(incoming []byte) (n int, err error) {
	if tw.mockWriteHandler != nil {
		return tw.mockWriteHandler(incoming)
	}

	tw.capturedOutput = append(tw.capturedOutput, string(incoming))
	return 0, nil
}

// VerifyCaptured asserts that capturedOutput exactly matches expected.™
func (tw *TestWriter) VerifyCaptured(t *testing.T, expected []string) {
	require.Len(t, tw.capturedOutput, len(expected), "Wrong number of capturedOutput")
	assert.Equal(t, expected, tw.capturedOutput, "capturedOutput is wrong")
}

package debuglogger

import (
	"fmt"
	"io"
	"os"

	"github.com/convergedtarkus/go-debuglog/utils"
)

type DebugLogger struct {
	// The prefix that will be added to all logged lines.
	prefix string

	// Default to os.Stdout. Used for unit testing only.
	ioOut io.Writer
}

// NewDebugLogger returns a new, fully setup, DebugLogger for logging to standard out.
func NewDebugLogger(prefix string) *DebugLogger {
	return &DebugLogger{
		prefix: prefix,
		ioOut:  os.Stdout,
	}
}

// Returns a new, fully setup, DebugLogger for logging to a custom io.Writer.
// Used for unit testing.
func newDebugLoggerWithIO(prefix string, ioOut io.Writer) *DebugLogger {
	return &DebugLogger{
		prefix: prefix,
		ioOut:  ioOut,
	}
}

// Log will log the given format and arguments using Printf.
// The current prefix will be added to each newline that is logged.
func (dl *DebugLogger) Log(format string, arguments ...interface{}) {
	format = utils.AddPrefix(dl.prefix, format)

	if format[len(format)-1] != '\n' {
		// Add a newline to ensure the log prints for systems that expect a terminating newline.
		format += "\n"
	}

	dl.print(format, arguments...)
}

// ChangePrefix will update the current prefix to the new prefix.
func (dl *DebugLogger) ChangePrefix(newPrefix string) {
	dl.prefix = newPrefix
}

// An internal print function to use the capture ioOut and handle errors.
func (dl *DebugLogger) print(format string, arguments ...interface{}) {
	_, err := fmt.Fprintf(dl.ioOut, format, arguments...)
	if err != nil {
		panic(err)
	}
}

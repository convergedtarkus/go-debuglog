package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const defaultPrefix = "!!!! "

func TestAddPrefix(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		prefix   string
		expected string
	}{
		{
			name:     "One line, no newline",
			input:    "Hello World",
			prefix:   defaultPrefix,
			expected: "!!!! Hello World",
		},
	}

	for _, tc := range testCases {
		// For Scopelint
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual := AddPrefix(tc.prefix, tc.input)
			assert.Equal(t, tc.expected, actual, "Wrong prefix result")
		})
	}
}

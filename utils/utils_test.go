package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const defaultPrefix = "!!!! "

func Test_AddPrefix(t *testing.T) {
	testCases := []struct {
		name        string
		inputValue  string
		inputPrefix string
		expected    string
	}{
		{
			name:        "One line",
			inputValue:  "Hello World",
			inputPrefix: defaultPrefix,
			expected:    "!!!! Hello World",
		},
		{
			name:        "One line, with final newline",
			inputValue:  "Hello World\n",
			inputPrefix: defaultPrefix,
			expected:    "!!!! Hello World\n",
		},
		{
			name:        "Multiple Trailing Newlines",
			inputValue:  "Hello World\n\n",
			inputPrefix: defaultPrefix,
			expected:    "!!!! Hello World\n!!!! \n",
		},
		{
			name:        "One line custom prefix",
			inputValue:  "Hello World",
			inputPrefix: "Blast Off  &  ",
			expected:    "Blast Off  &  Hello World",
		},
		{
			name:        "One line prefix with newlines",
			inputValue:  "Hello World",
			inputPrefix: "Fire\nAway\n",
			expected:    "Fire\nAway\nHello World",
		},
		{
			name:        "Two lines",
			inputValue:  "Hello\nWorld",
			inputPrefix: defaultPrefix,
			expected:    "!!!! Hello\n!!!! World",
		},
		{
			name:        "Two lines, with final newline",
			inputValue:  "Hello\nWorld\n",
			inputPrefix: defaultPrefix,
			expected:    "!!!! Hello\n!!!! World\n",
		},
		{
			name:        "Many lines",
			inputValue:  "One\nTwo\nThree\nFour\nFive",
			inputPrefix: defaultPrefix,
			expected:    "!!!! One\n!!!! Two\n!!!! Three\n!!!! Four\n!!!! Five",
		},
		{
			name:        "Many lines, with final newline",
			inputValue:  "One\nTwo\nThree\nFour\nFive\n",
			inputPrefix: defaultPrefix,
			expected:    "!!!! One\n!!!! Two\n!!!! Three\n!!!! Four\n!!!! Five\n",
		},
		{
			name:        "Middle Blank Newline",
			inputValue:  "Hi\n\nBye",
			inputPrefix: defaultPrefix,
			expected:    "!!!! Hi\n!!!! \n!!!! Bye",
		},
		{
			name:        "Middle Blank Newline with final newline",
			inputValue:  "Hi\n\nBye\n",
			inputPrefix: defaultPrefix,
			expected:    "!!!! Hi\n!!!! \n!!!! Bye\n",
		},
		{
			name:        "Multiple Middle Blank Newlines",
			inputValue:  "Hi\n\n\n\nBye",
			inputPrefix: defaultPrefix,
			expected:    "!!!! Hi\n!!!! \n!!!! \n!!!! \n!!!! Bye",
		},
		{
			name:        "Multiple Middle Blank Newlines with final newline",
			inputValue:  "Hi\n\n\n\nBye\n",
			inputPrefix: defaultPrefix,
			expected:    "!!!! Hi\n!!!! \n!!!! \n!!!! \n!!!! Bye\n",
		},
	}

	for _, tc := range testCases {
		// For Scopelint
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual := AddPrefix(tc.inputPrefix, tc.inputValue)
			assert.Equal(t, tc.expected, actual, "Wrong inputPrefix result")
		})
	}
}

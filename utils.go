package utils

import (
	"strings"
)

func AddPrefix(prefix, inputString string) string {
	splitResult := strings.Split(inputString, "\n")

	hasEndingNewline := inputString[len(inputString)-1] == '\n'

	result := ""
	for idx, line := range splitResult {
		if idx == len(splitResult)-1 && !hasEndingNewline {
			result += prefix + line
		} else {
			result += prefix + line + "\n"
		}
	}
	return result
}

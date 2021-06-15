package utils

import (
	"strings"
)

func AddPrefix(prefix, inputString string) string {
	splitResult := strings.Split(inputString, "\n")

	result := ""
	for idx, curLine := range splitResult {
		ending := "\n"

		if idx == len(splitResult)-1 {
			if curLine == "" {
				// If the last line is blank, don't add it at all.
				continue
			}

			// Do not add a final new line.
			ending = ""
		}

		result += prefix + curLine + ending
	}

	return result
}

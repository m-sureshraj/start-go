package utils

import (
	"strings"
)

func Join(phrases []string) string {
	length := len(phrases)
	if length == 0 {
		return ""
	}

	if length == 1 {
		return phrases[0]
	}

	lastPhrase := phrases[length-1]
	initialPhrases := phrases[:length-1]

	if length == 2 {
		return initialPhrases[0] + " and " + lastPhrase
	}

	return strings.Join(initialPhrases, ", ") + ", and " + lastPhrase
}

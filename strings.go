package strutils

import "strings"

func Concat(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}

	var (
		totalChars int
		b          strings.Builder
	)

	// For long strings its actually faster to make it grow once
	for _, s := range strs {
		totalChars += len(s)
	}

	b.Grow(totalChars)

	for _, s := range strs {
		b.WriteString(s)
	}

	return b.String()
}

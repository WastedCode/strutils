package strutils

import (
	"math/rand"
	"strings"
	"unsafe"
)

// Concat will return a concatenation of all the given strings
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

const (
	// These are the characters we will use for our random string
	charBytes = "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// The maximum number of chars we support in random strings
	maxChars = 62
	// The number of bits we should use to get a string
	// 6 bits allow us to represent 64 characters
	bitsNeededPerChar = 6
	// This is the mask we use to generate the idx for character
	charMask = 1<<bitsNeededPerChar - 1
	// This is the maximum chars we can generate per round
	// Each int64 will have 63 usable bits
	maxCharsPerRand = 63 / bitsNeededPerChar
	// one
	one = 1
)

// GenRandomString will generate a random string of the given length
// This uses the super optimized version where we pick bits manually
// With current implementation of rand.Read we are efficient when we use 64 characters.
// Even at 64 characters, this is not sub optimal with rand.Read, just on par
// There is wastage when using less than 64 characters as a base for our string
func GenRandomString(n int) string {
	if n <= 0 {
		return ""
	}

	b := make([]byte, n)

	// A rand.Int63() generates 63 random bits, i.e. maxCharsPerRand random character can be chosen
	for i, cache, remain := n-one, rand.Int63(), maxCharsPerRand; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), maxCharsPerRand
		}

		if idx := int(cache & charMask); idx < maxChars {
			b[i] = charBytes[idx]
			i--
		}

		cache >>= bitsNeededPerChar
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

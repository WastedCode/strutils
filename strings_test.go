package strutils

import (
	"math/rand"
	"regexp"
	"testing"
	"time"
)

func TestConcat(t *testing.T) {
	testcases := map[string]struct {
		input    []string
		expected string
	}{
		"empty": {
			input:    []string{},
			expected: "",
		},
		"single": {
			input:    []string{"a"},
			expected: "a",
		},
		"multi": {
			input:    []string{"a", "bb", "ccc"},
			expected: "abbccc",
		},
	}

	for name, example := range testcases {
		t.Run(name, func(t *testing.T) {
			var value string
			if len(example.input) == 0 {
				value = Concat()
			} else {
				value = Concat(example.input...)
			}
			if value != example.expected {
				t.Fatalf("Expected Concat on %v to return %s, got %s", example.input, example.expected, value)
			}
		})
	}
}

func TestGenRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	r, err := regexp.Compile("^[a-zA-Z0-9]*$")
	if err != nil {
		panic(err)
	}

	for i := -1; i <= 64; i++ {
		result := GenRandomString(i)
		if i < 1 {
			if result != "" {
				t.Fatalf("expected GenRandomString(%d) to return empty string, got %s", i, result)
			}
		} else {
			if len(result) != i {
				t.Fatalf("expected GenRandomString(%d) to return a string of length %d, got %s of length %d", i, i, result, len(result))
			}
			if !r.MatchString(result) {
				t.Fatalf("expected GenRandomString(%d) to return an alphanumeric string, got %s instead", i, result)
			}
		}
	}
}

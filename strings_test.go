package strutils

import "testing"

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

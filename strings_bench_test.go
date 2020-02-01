package strutils

import (
	"strings"
	"testing"
	"unsafe"
)

var concatResult string

func concatStringsJoin(strs ...string) string {
	return strings.Join(strs, "")
}

func concatStringsBuilderVanilla(strs ...string) string {
	var b strings.Builder
	for _, s := range strs {
		b.WriteString(s)
	}
	return b.String()
}

func concatManual(strs ...string) string {
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}

	var totalChars int
	for _, s := range strs {
		totalChars += len(s)
	}

	if totalChars < 32 {
		buff := []byte(strs[0])
		for i := 1; i < len(strs); i++ {
			buff = append(buff, strs[i]...)
		}

		return string(buff)
	}

	buff := make([]byte, totalChars)
	var i int
	for _, s := range strs {
		i += copy(buff[i:], s[:])
	}
	return *(*string)(unsafe.Pointer(&buff))
}

var (
	concatSmall = []string{"a", "b", "c"}
	concatMid   = []string{"some", "characters", "cnt"}
	concatLong  = []string{"01234567890", "-", "01234567890", "-", "01234567890", "-", "01234567890"}
	concatLarge = []string{"justconcataveryveryverylongbufferfromonestringtoanother", "justconcataveryveryverylongbufferfromonestringtoanother"}
)

func BenchmarkConcatSmall(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = Concat(concatSmall...)
	}
	concatResult = result
}

func BenchmarkConcatJoinSmall(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = concatStringsJoin(concatSmall...)
	}
	concatResult = result
}

func BenchmarkConcatBuilderSmall(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = concatStringsBuilderVanilla(concatSmall...)
	}
	concatResult = result
}

func BenchmarkConcatManualSmall(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = concatManual(concatSmall...)
	}
	concatResult = result
}

func BenchmarkConcatMid(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = Concat(concatMid...)
	}
	concatResult = result
}

func BenchmarkConcatJoinMid(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = concatStringsJoin(concatMid...)
	}
	concatResult = result
}

func BenchmarkConcatBuilderMid(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = concatStringsBuilderVanilla(concatMid...)
	}
	concatResult = result
}

func BenchmarkConcatManualMid(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = concatManual(concatMid...)
	}
	concatResult = result
}

func BenchmarkConcatLong(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = Concat(concatLong...)
	}
	concatResult = result
}

func BenchmarkConcatJoinLong(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = concatStringsJoin(concatLong...)
	}
	concatResult = result
}

func BenchmarkConcatBuilderLong(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = concatStringsBuilderVanilla(concatLong...)
	}
	concatResult = result
}

func BenchmarkConcatManualLong(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = concatManual(concatLong...)
	}
	concatResult = result
}

func BenchmarkConcatLarge(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = Concat(concatLarge...)
	}
	concatResult = result
}

func BenchmarkConcatJoinLarge(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = concatStringsJoin(concatLarge...)
	}
	concatResult = result
}

func BenchmarkConcatBuilderLarge(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = concatStringsBuilderVanilla(concatLarge...)
	}
	concatResult = result
}

func BenchmarkConcatManualLarge(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = concatManual(concatLarge...)
	}
	concatResult = result
}

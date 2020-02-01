package strutils

import (
	"strings"
	"testing"
	"unsafe"
)

/*
Benchmark results are in concat_bench.txt
Concat - Our implementation (wrapper around strings.Builder
Join - Using strings.Join
Builder - Directly using strings.Builder
Manual - A very efficient way using manual buffer copy. As of 1.13 is barely faster than our Concat, but we prefer this since it uses standard library

❯ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/wastedcode/strutils
BenchmarkConcatSmall-16                 36458156                27.8 ns/op             3 B/op          1 allocs/op
BenchmarkConcatJoinSmall-16             34108449                34.7 ns/op             3 B/op          1 allocs/op
BenchmarkConcatBuilderSmall-16          35770252                29.8 ns/op             8 B/op          1 allocs/op
BenchmarkConcatManualSmall-16           43085545                25.7 ns/op             3 B/op          1 allocs/op
BenchmarkConcatMid-16                   29700308                35.1 ns/op            32 B/op          1 allocs/op
BenchmarkConcatJoinMid-16               25948977                44.1 ns/op            32 B/op          1 allocs/op
BenchmarkConcatBuilderMid-16            13468004                85.5 ns/op            56 B/op          3 allocs/op
BenchmarkConcatManualMid-16             33127766                33.8 ns/op            32 B/op          1 allocs/op
BenchmarkConcatLong-16                  22853175                51.6 ns/op            48 B/op          1 allocs/op
BenchmarkConcatJoinLong-16              15723216                73.3 ns/op            48 B/op          1 allocs/op
BenchmarkConcatBuilderLong-16           10799188               109 ns/op             112 B/op          3 allocs/op
BenchmarkConcatManualLong-16            22404008                50.8 ns/op            48 B/op          1 allocs/op
BenchmarkConcatLarge-16                 26872747                41.5 ns/op           112 B/op          1 allocs/op
BenchmarkConcatJoinLarge-16             21745833                47.2 ns/op           112 B/op          1 allocs/op
BenchmarkConcatBuilderLarge-16          14940462                74.7 ns/op           192 B/op          2 allocs/op
BenchmarkConcatManualLarge-16           27793780                40.3 ns/op           112 B/op          1 allocs/op

benchstat results
name                   time/op
ConcatSmall-16         27.4ns ± 5%
ConcatJoinSmall-16     36.6ns ± 8%
ConcatBuilderSmall-16  31.5ns ± 3%
ConcatManualSmall-16   27.1ns ± 4%
ConcatMid-16           38.2ns ± 5%
ConcatJoinMid-16       47.0ns ± 3%
ConcatBuilderMid-16    93.5ns ± 8%
ConcatManualMid-16     36.2ns ± 4%
ConcatLong-16          55.8ns ± 8%
ConcatJoinLong-16      79.7ns ± 8%
ConcatBuilderLong-16    114ns ± 7%
ConcatManualLong-16    54.7ns ± 3%
ConcatLarge-16         44.5ns ± 5%
ConcatJoinLarge-16     49.6ns ±10%
ConcatBuilderLarge-16  78.1ns ± 0%
ConcatManualLarge-16   42.3ns ± 1%
*/

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

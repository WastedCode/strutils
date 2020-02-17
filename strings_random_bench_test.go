package strutils

import (
	"math/rand"
	"testing"
	"time"
	"unsafe"
)

var (
	_randResult string
)

func genRandomString(n int) string {
	str := make([]byte, n)
	for i := 0; i < n; i++ {
		str[i] = charBytes[rand.Intn(maxChars)]
	}
	return *(*string)(unsafe.Pointer(&str))
}

func BenchmarkGenRandomString(b *testing.B) {
	var result string
	rand.Seed(time.Now().UnixNano())
	for n := 0; n < b.N; n++ {
		result = GenRandomString(64)
	}
	_randResult = result
}

func BenchmarkGenRandomString2(b *testing.B) {
	var result string
	rand.Seed(time.Now().UnixNano())
	for n := 0; n < b.N; n++ {
		result = genRandomString(64)
	}
	_randResult = result
}

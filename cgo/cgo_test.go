package main

import (
	"github.com/breml/mygoplayground/cgo/cgopkg"
	"testing"
)

func BenchmarkGoCGoFunc(b *testing.B) {
	// run sample b.N times
	for n := 0; n < b.N; n++ {
		cgopkg.GoBenchCall()
	}
}

func BenchmarkGoCFunc(b *testing.B) {
	// run sample b.N times
	for n := 0; n < b.N; n++ {
		cgopkg.CBenchCall()
	}
}

func BenchmarkGoFunc(b *testing.B) {
	// run sample b.N times
	for n := 0; n < b.N; n++ {
		cgopkg.GoCall()
	}
}

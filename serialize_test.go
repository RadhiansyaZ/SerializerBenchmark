package main

import (
	"testing"
)

func BenchmarkSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Serialize()
	}
}
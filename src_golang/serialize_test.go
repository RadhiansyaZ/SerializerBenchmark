package main

import (
	"testing"
)

func BenchmarkSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		Serialize()
	}
}

func BenchmarkSerializeWithStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		SerializeWithStd()
	}
}

func BenchmarkSerializeToStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		SerializeToStruct()
	}
}

func BenchmarkSerializeToStructWithStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		SerializeToStructWithStd()
	}
}

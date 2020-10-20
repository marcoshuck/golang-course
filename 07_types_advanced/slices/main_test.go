package main

import "testing"

// 3.46 ns/op
func BenchmarkSliceAndReservingMemory(b *testing.B) {
	slice := make([]int, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice = append(slice, i*3)
	}
}

// 39.3 ns/op
func BenchmarkSliceWithoutReservingMemory(b *testing.B) {
	var slice []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice = append(slice, i*3)
	}
}

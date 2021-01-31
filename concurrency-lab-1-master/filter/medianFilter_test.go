package main

import (
	"os"
	"testing"
)

func Benchmark(b *testing.B) {
	os.Stdout = nil // Disable all program output apart from benchmark results
	b.Run("Median filter benchmark", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			filter("ship.png", "out.png")
		}
	})
}

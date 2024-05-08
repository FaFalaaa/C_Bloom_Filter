package benchmark

import (
	"testing"
	"time"
)

func myFunction() {
	// Simulate some computation
	time.Sleep(2 * time.Second)
}

// Benchmark function
func BenchmarkMyFunction(b *testing.B) {
	// Run the function once
	myFunction()
}

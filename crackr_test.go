package main

import "testing"

// run using go test -bench=. -benchmem -count 3 -v -cpuprofile cpu.out -memprofile mem.mprof
// profile using go tool pprof cpu.out
// useful func
// top5
// list funcName
// more here https://blog.golang.org/profiling-go-programs
func BenchmarkSingleDictSingleHash(b *testing.B) {
	// run the Fib function b.N times
	dictionary := "passwords.txt"
	hash := "5f4dcc3b5aa765d61d8327deb882cf99"
	hashes := "nil"
	for n := 0; n < b.N; n++ {
		attackUsingSingleDictionary(&dictionary, &hash, &hashes)
	}
}

# crackr
Crackr - A CLI to crack hashes using a dictionary attack

## Exmaple Usage

### Crack 1 password against 1 dictionary:

`crackr -f passwords.txt -h 5f4dcc3b5aa765d61d8327deb882cf99`

### Crack multiple passwords against 1 dictionary:

`crackr -f passwords.txt -hf hashes.txt`

### Crack 1 password against multiple dictionaries:

`crackr -h 5f4dcc3b5aa765d61d8327deb882cf99 -f dictionaries`

### Crack multiple passwords against multiple dictionaries

`crackr -hf hashes.txt -d dicitonaries`

## Benchmark
Run using `go test -bench=. -benchmem -count 3 -v -cpuprofile cpu.out -memprofile mem.mprof`
```
PARALLEL

BEST SEQUENTIAL

BenchmarkSingleDictSingleHash-8   	     300	   5558339 ns/op	 2924959 B/op	   28674 allocs/op
BenchmarkSingleDictSingleHash-8   	     300	   5779599 ns/op	 2978536 B/op	   29297 allocs/op
BenchmarkSingleDictSingleHash-8   	     300	   5666784 ns/op	 2863379 B/op	   28382 allocs/op

SEQUENTIAL
BenchmarkSingleDictSingleHash-8   	     100	  12989579 ns/op	 4356895 B/op	   40286 allocs/op
BenchmarkSingleDictSingleHash-8   	     100	  11808874 ns/op	 4366221 B/op	   40488 allocs/op
BenchmarkSingleDictSingleHash-8   	     100	  15402018 ns/op	 4375962 B/op	   40690 allocs/op
```




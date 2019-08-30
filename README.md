Quick and dirty hamming distance benchmark showing performance improvement potential of a full-optimized parameterized array generic function over a slice based alternative.

Output:
```
Running tool: /usr/local/bin/go test -benchmem -run=^$ github.com/vsivsi/hdtest -bench ^(BenchmarkHD)$

goos: darwin
goarch: amd64
pkg: github.com/vsivsi/hdtest
BenchmarkHD/256-bit_unrolled_array_HD-20         	2000000000	         1.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkHD/256-bit_unrolled_array_sum_HD-20     	2000000000	         1.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkHD/256-bit_slice_HD-20                  	300000000	         4.99 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/vsivsi/hdtest	6.386s
Success: Benchmarks passed.
```
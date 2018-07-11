# Solution

Using DP bottom-up fashion.

Memory table dimension is reduced.

# Benchmarking

## Simple

`go test -bench=.`

```
goos: darwin
goarch: amd64
pkg: lcs
BenchmarkComparator_Evaluate-8   	 2000000	       790 ns/op
PASS
ok  	lcs	2.400s
```

## With Memory allocation

`go test -bench=. -benchmem`

```
goos: darwin
goarch: amd64
pkg: lcs
BenchmarkComparator_Evaluate-8   	 2000000	       803 ns/op	     152 B/op	       5 allocs/op
PASS
ok  	lcs	2.426s
```

# Reference

1. https://en.wikipedia.org/wiki/Longest_common_subsequence_problem
2. http://www.cs.cmu.edu/afs/cs/academic/class/15451-s15/LectureNotes/lecture04.pdf
3. http://www.techiedelight.com/longest-common-subsequence-lcs-space-optimized-version/

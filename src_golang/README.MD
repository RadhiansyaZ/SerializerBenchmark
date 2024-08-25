# Modification

## Goal

Use a defined struct to store the value of the unmarshalled JSON instead of using a `map[string]interface{}`. This is to avoid type inference and runtime reflection. On the original code, the map is used to store the value of the struct, which is can be done by using the struct itself. Since `Nim` implementation is using a struct, it is better to use a struct for the `Go` implementation as well, because both are staticly typed and compiled language. 

## Methodology

### Execution

Use `make test` to run the test. The test will run the benchmark and the escape analysis. The escape analysis is done by running `go build -gcflags="-m"`. The benchmark is done by running `go test -benchtime=3000x -bench=. -benchmem`. The result of the benchmark is stored in `../results/go.bench.txt` and the escape analysis is stored in `../results/go.escape.txt`.

### Changes

#### Timing

The timing used in the flag `-benchtime=` is changed from `1s` in second to `3000x` iteration. The reason is to even up the same sample size regardless of test bench used.

#### Comparator

Will use `benchstat` of [goperf](https://cs.opensource.google/go/x/perf) to compare the benchmark result difference. Since the the implementation difference is split into two function, the result from `../results/go.bench` will be split into two different file, `../results/go.bench.old.txt` and `../results/go_struct.bench.new.txt`. Then, the function name of compared function will be renamed with the same name so it can be compared by bench stat. Installation of `benchstat` can be done by running `go install golang.org/x/perf/cmd/benchstat@latest`.

For example, if the original file is 
```zsh
go test -benchtime=3000x -bench=. -benchmem
goos: linux
goarch: amd64
pkg: benchmark
cpu: AMD Ryzen 5 7600 6-Core Processor              
BenchmarkSerialize-12                   	    3000	    459060 ns/op	   55791 B/op	     472 allocs/op
BenchmarkSerializeWithStd-12            	    3000	    462582 ns/op	   51302 B/op	     359 allocs/op
BenchmarkSerializeToStruct-12           	    3000	    424353 ns/op	   40820 B/op	     113 allocs/op
BenchmarkSerializeToStructWithStd-12    	    3000	    441454 ns/op	   41410 B/op	     122 allocs/op
PASS
```
it will be split into two files, `go.bench.old.txt` and `go_struct.bench.new.txt` with the following content:
```zsh
# go.bench.old.txt
go test -benchtime=3000x -bench=. -benchmem
goos: linux
goarch: amd64
pkg: benchmark
cpu: AMD Ryzen 5 7600 6-Core Processor   
BenchmarkSerialize-12                   	    3000	    459060 ns/op	   55791 B/op	     472 allocs/op
BenchmarkSerializeWithStd-12            	    3000	    462582 ns/op	   51302 B/op	     359 allocs/op
PASS
```
Then, the function name of the `go_struct.bench.new.txt` will be renamed to `BenchmarkSerialize` and `BenchmarkSerializeWithStd`:
```zsh
# go.bench.new.txt
go test -benchtime=3000x -bench=. -benchmem
goos: linux
goarch: amd64
pkg: benchmark
cpu: AMD Ryzen 5 7600 6-Core Processor
BenchmarkSerialize-12           	            3000	    424353 ns/op	   40820 B/op	     113 allocs/op
BenchmarkSerializeWithStd-12    	            3000	    441454 ns/op	   41410 B/op	     122 allocs/op
PASS
```
so it can be compared by `benchstat`. The command used to compare the result is `benchstat ./results/go.bench.old.txt ./results/go.bench.new.txt` in root directory. The comparison stored in `./results/go.bench.compared.txt`.

### Removal of HTTP Server

Since the goal is to compare the performance of the serialization, the HTTP server will be removed from the benchmark. The HTTP server will be removed from the benchmark and the benchmark will be run on the main function. It will reduce the scope of the testing and make the result more focused on the serialization performance. It also give cleaner result when we do the escape analysis to pinpoint which part goes into heap allocation.

### Machine Specification

1. Ryzen 5 7600 6-core processor, 32 GB 5200Mhz DDR5 128bit*, NixOS 24.04 on WSL.
2. MacBook Pro M2 2022, 8-core processor, 8 GB 6400Mhz LPDDR5 128bit, macOS Ventura 13.04.

\* 128bit from dual channel memory configuration.

## Result and Analysis

### Escape analysis

The escape analysis is done by running `go build -gcflags="-m"`. The difference between the original and the new implementation is between line 43-52 and line 63-72 in `../results/go.escape.txt`. The main difference is line 57-58 in the original implementation:
```zsh
./serialize.go:85:90: i + 1 escapes to heap
./serialize.go:86:96: len(objmap["discussion_results"].([]interface {})[i].(map[string]interface {})["replies"].([]interface {})) escapes to heap
``` 
and line 63 in the new implementation:
```zsh
./serialize.go:115:7: c does not escape
```
The difference is there are more heap allocation in the original implementation compared to the new implementation.
The reasons are:
- Since the compiler cannot determined `i` is not used after the function returns, the compiler takes a safety approach by assuming the value could possibly referenced after the function returns.
- The use of `map[string]interface{}` in the original implementation makes the compiler store the `interface{}` pointer to another memory, then use the type inference to access it in runtime. 

You may notice that in both scenario, the result of `objmap` is always gone to heap, shown by following:
```zsh
./serialize.go:20:6: moved to heap: objmap
```
This is because the compiler doesn't know the size of `objmap` in compile time, thus allocating it in heap. It applies to both original implementation with `map[string]interface{}` but also to the new implementation with struct. The struct is having a unknown size slice of `D` and inside the `D`, there is an `R` slice which is also unknown size.

### Benchmark

The result used for comparison is the first benchmark used in `go.bench.txt`. The result is as follows:
```zsh
goos: linux
goarch: amd64
pkg: benchmark
cpu: AMD Ryzen 5 7600 6-Core Processor              
                    │ ./results/go.bench.old.txt │      ./results/go.bench.new.txt      │
                    │           sec/op           │    sec/op     vs base                │
Serialize-12                        459.1µ ± ∞ ¹   424.4µ ± ∞ ¹       ~ (p=1.000 n=1) ²
SerializeWithStd-12                 462.6µ ± ∞ ¹   441.5µ ± ∞ ¹       ~ (p=1.000 n=1) ²
geomean                             460.8µ         432.8µ        -6.08%
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05

                    │ ./results/go.bench.old.txt │       ./results/go.bench.new.txt       │
                    │            B/op            │     B/op       vs base                 │
Serialize-12                       54.48Ki ± ∞ ¹   39.86Ki ± ∞ ¹        ~ (p=1.000 n=1) ²
SerializeWithStd-12                50.10Ki ± ∞ ¹   40.44Ki ± ∞ ¹        ~ (p=1.000 n=1) ²
geomean                            52.25Ki         40.15Ki        -23.15%
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05

                    │ ./results/go.bench.old.txt │      ./results/go.bench.new.txt      │
                    │         allocs/op          │  allocs/op   vs base                 │
Serialize-12                         472.0 ± ∞ ¹   113.0 ± ∞ ¹        ~ (p=1.000 n=1) ²
SerializeWithStd-12                  359.0 ± ∞ ¹   122.0 ± ∞ ¹        ~ (p=1.000 n=1) ²
geomean                              411.6         117.4        -71.48%
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05
```
From the result above, there is small difference in the time taken to serialize the struct. The difference is around 6.08% faster than the original implementation. The memory allocation is also reduced by 23.15% and the number of allocation is reduced by 71.48%. The result is consistent with the goal of the modification, which is to reduce the number of allocation and memory usage.

## Conclusion

The new implementation is clearly faster than the original and also more memory efficient, which is consistent with the goal of the modification. Notice that in terms of execution speed, it is not much of a difference. For a small JSON file, **the difference is not that significant**. It is still not known for a larger JSON file and need a larger test file. It need testing with a larger JSON file to see further impact on the memory (or even the execution speed), thus it may make a big difference.

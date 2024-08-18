# Performance Benchmark
## Stack
- Golang
  - Net/Http
  - Jsoniter
- Javascript
  - Elysia
  - Typia
- Python
  - Falcon
  - orjson
- Nim
  - Mummy

## Scope
- Single http API call speed
- Deserialization speed
- Serialization speed

## Use Case
1. Call a localhost json file using http request http://localhost:8080/sample.json
2. Deserialize the response to respective programming language object
3. Transform the deserialized into needed response
4. Serialize the transformed data into response

## How to measure
### Prerequisite
1. Have npm installed on your local machine
2. Start the server by running `sample_data/start_server.(sh|bat)` script
3. Run the benchmark by running each programming language script

### Golang
1. using test by running `go test -benchtime=1s -bench=. -benchmem`. but in windows, you must specify the test function in `-bench` option, in this case `-bench=BenchmarkSerialize`
2. using duration measured by `time.since` function and write the result to file (this is faster than using fmt.Println ± 10-15x)

### Nim
1. using timeit

# Modification

## Goal

- Originally use a nil map[string]interface{}, want to test out by using a known struct to avoiding type inference and runtime reflection.

## Methodology

### Changes

- Change from 1s timing to 3000 iteration for same sample size regardless of test bench used.
- Will use `benchstat` of [goperf](https://cs.opensource.google/go/x/perf) to compare the benchmark result difference.

### Removal of HTTP Server

- Serve no purpose.
- Use the function directly to do escape analysis.

## Result and Analysis

To be written.

## Conclusion

To be written.

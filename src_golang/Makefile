ESCAPE=CGO_ENABLED=0 go build -o bin/serialize_escape -gcflags "-m" serialize.go
BENCH=go test -benchtime=3000x -bench=. -benchmem -cpu=6

.PHONY: all bench

all: build serve

build:
	CGO_ENABLED=0 go build -o bin/serialize serialize.go

serve:
	./bin/serialize

clean:
	echo "Removing Go binary and benchmark results"
	rm -rf bin/*
	rm ../results/go.escape.txt
	rm ../results/go.bench.txt

test: ../results/go.escape.txt ../results/go.bench.txt

../results/go.escape.txt:
	echo "Running escape analysis on Go code"
	$(ESCAPE) &> ../results/go.escape.txt

../results/go.bench.txt: 
	echo "Running benchmarks on Go code"
	$(BENCH) >> ../results/go.bench.txt
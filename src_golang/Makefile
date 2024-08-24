.PHONY: all bench

all: build serve

bench: 
	go test -benchtime=1s -bench=. -benchmem

build: 
	CGO_ENABLED=0 GOOS=linux go build -o bin/serialize serialize.go

serve:
	./bin/serialize

escape:
	CGO_ENABLED=0 GOOS=linux go build -o bin/serialize_escape -gcflags "-m" serialize.go

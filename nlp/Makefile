all:
	$(error please pick a target)

test:
	go test -v ./...


nlpd: cmd/nlpd
	go build \
	    -ldflags="-X main.Version=$(shell git rev-parse --short HEAD)" \
	    ./cmd/nlpd

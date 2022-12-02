tag=ardan/nlpd

all:
	$(error please pick a target)

test:
	go test -v ./...

nlpd: cmd/nlpd
	go build \
	    -ldflags="-X main.Version=$(shell git rev-parse --short HEAD)" \
	    ./cmd/nlpd

build-docker:
	docker build -t $(tag) .

run-docker:
	docker run --rm -p 8080:8080 $(tag)

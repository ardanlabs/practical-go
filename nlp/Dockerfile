FROM golang:1.19-bullseye AS build
WORKDIR /nlpd
# Faster builds - will run only if go.mod or go.sum changes
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Generate static executable
ENV CGO_ENABLED=0
RUN go build -o nlpd ./cmd/nlpd

FROM debian:bullseye-slim
# Security: Don't run as root
RUN groupadd -r nlp && useradd --no-log-init -r -g nlp nlp
USER nlp
COPY --from=build /nlpd/nlpd /usr/local/bin
CMD nlpd

# Practical Go for Developers

This repo contains the material for the "Practical Go Foundations" class.
The code & links are synced with the [online class](for Developers).

This is an assorted collection of exercises for teaching, not a real Go project.

---

## Day 1

### Agenda

- Strings & formatted output
    - What is a string?
    - Unicode basics
    - Using fmt package for formatted output
- Calling REST APIs
    - Making HTTP calls with net/http
    - Defining structs
    - Serializing JSON
- Working with files
    - Handling errors
    - Using defer to manage resources
    - Working with io.Reader & io.Writer interfaces

### Code

- [hw.go](hw/hw.go) - Hello World
    - `GOOS=drawin go build` (also `GOARCH`)
- [banner.go](banner/banner.go) - Strings & printing
- [github.go](github/github.go) - Calling REST APIs
- [sha1.go](sha1/sha1.go) - Working with `io.Reader` & `io.Writer`

[Terminal Log](_extra/day-1.log)


### Links

- [HTTP status cats](https://http.cat/)
- [errors](https://pkg.go.dev/errors/) package ([Go 1.13](https://go.dev/blog/go1.13-errors))
- [encoding/json](https://pkg.go.dev/encoding/json)
- [net/http](https://pkg.go.dev/net/http)
- [Let's talk about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging) by Dave Cheney
- Numbers
    - [math/big](https://pkg.go.dev/math/big/) - Big numbers
    - [Numeric types](https://go.dev/ref/spec#Numeric_types)
- Strings
    - [Unicode table](https://unicode-table.com/)
    - [strings](https://pkg.go.dev/strings/) package - string utilities
    - [Go strings](https://go.dev/blog/strings)
- [Annotated "Hello World"](https://www.353solutions.com/annotated-go)
- [Effective Go](https://go.dev/doc/effective_go.html) - Read this!
- [Go standard library](https://pkg.go.dev/) - official documentation
- [A Tour of Go](https://do.dev/tour/)
- Setting Up
    - [The Go SDK](https://go.dev/dl/)
    - [git](https://git-scm.com/)
    - IDE's: [Visual Studio Code](https://code.visualstudio.com/) + [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) or [Goland](https://www.jetbrains.com/go/) (paid)

### Data & Other

- `G☺`
- `♡`
- [http.log.gz](_extra/http.log.gz)

---

## Day 2

### Agenda

- Sorting
    - Working with slices
    - Writing methods
    - Understanding interfaces
- Catching panics
    - The built-in recover function
    - Named return values
- Processing text
    - Reading line by line with bufio.Scanner
    - Using regular expressions
    - Working with maps

### Code

- [slices.go](slices/slices.go) - Working with slices
- [game.go](game/game.go) - Structs, methods & interfaces
- [empty.go](empty/empty.go) - The empty interface, type assertions
- [div.go](div/div.go) - Catching panics
- [freq.go](freq/freq.go) - Most common word (files, regular expressions, maps)

[Terminal Log](_extra/day-2.log)

### Exercises

- Read and understand the [sort package examples](https://pkg.go.dev/sort/#pkg-examples)
- Implement `sortByDistance(players []Player, x, y int)` in `game.go`
- Change `mostCommon` to return the most common `n` words (e.g. `func mostCommon(r io.Reader, n int) ([]string, error)`)

### Links

- [regex101](https://regex101.com/) - Regular expression builder
- [Go Proverbs](https://go-proverbs.github.io/) - Think about them ☺
- [sort examples](https://pkg.go.dev/sort/#pkg-examples) - Read and try to understand
- [When to use generics](https://go.dev/blog/when-generics)
- [Generics tutorial](https://go.dev/doc/tutorial/generics)
- [Methods, interfaces & embedded types in Go](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html)
- [Methods & Interfaces](https://go.dev/tour/methods/1) in the Go tour
- Slices
    - [Slices](https://go.dev/blog/slices) & [Slice internals](https://go.dev/blog/go-slices-usage-and-internals) on the Go blog
    - [Slice tricks](https://github.com/golang/go/wiki/SliceTricks)
- Error Handling
    - [Defer, Panic and Recover](https://go.dev/blog/defer-panic-and-recover)
    - [errors](https://pkg.go.dev/errors/) package ([Go 1.13](https://go.dev/blog/go1.13-errors))
    - [pkg/errors](https://github.com/pkg/errors)


### Data & Other

- [sherlock.txt](data/sherlock.txt)

---

## Day 3

### Agenda

- Distributing work
    - Using goroutines & channels
    - Using the sync package to coordinate work
- Timeouts & cancellation
    - Working with multiple channels using select
    - Using context for timeouts & cancellations
    - Standard library support for context

### Code

- [go_chan.go](go_chan/go_chan.go) - Goroutines & channels
    - [sleep_sort.sh](go_chan/sleep_sort.sh) - Sleep sort in bash
- [taxi_check.go](taxi/taxi_check.go) - Turn sequential code to parallel
- [sites_time.go](sites_time/sites_time.go) - Using sync.WaitGroup
- [payment.go](payment/payment.go) - Using sync.Once & sync.WaitGroup
- [counter.go](counter/counter.go) - Using the race detector, sync.Mutex and sync/atomic
- [select.go](select/select.go) - Using `select`
- [rtb.go](rtb/rtb.go) - Using `context` for cancellations

[Terminal Log](_extra/day-3.log)

### Exercise

In `taxi_check.go`
- Limit the number of goroutines to "n". Which "n" yields the best results?
- Cancel all goroutines once there's an error or mismatch in signature

### Links

- [The race detector](https://go.dev/doc/articles/race_detector)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
- [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)
- [Data Race Patterns in Go](https://eng.uber.com/data-race-patterns-in-go/)
- [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
- [Go Concurrency Patterns: Context](https://go.dev/blog/context)
- [Curious Channels](https://dave.cheney.net/2013/04/30/curious-channels)
- [The Behavior of Channels](https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html)
- [Channel Semantics](https://www.353solutions.com/channel-semantics)
- [Why are there nil channels in Go?](https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308)
- [Amdahl's Law](https://en.wikipedia.org/wiki/Amdahl%27s_law) - Limits of concurrency
- [Computer Latency at Human Scale](https://twitter.com/jordancurve/status/1108475342468120576/photo/1)
- [Concurrency is not Parallelism](https://www.youtube.com/watch?v=cN_DpYBzKso) by Rob Pike
- [Scheduling in Go](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)

### Data & Other

- [rtb.go](_extra/rtb.go)
- [site_times.go](_extra/site_time.go)
- [taxi_check.go](_extra/taxi_check.go)
    - [taxi-sha256.zip](_extra/taxi-sha256.zip)

---

## Day 4

### Agenda

- Testing your code
    - Working with the testing package
    - Using testify
    - Managing dependencies with go mod
- Structuring your code
    - Writing sub-packages
- Writing an HTTP server
    - Writing handlers
    - Using gorilla/mux for routing
Adding metrics & logging
    - Using expvar for metrics
    - Using the log package and a look at user/zap
- Configuration patterns
    - Reading environment variables and a look at external packages
    - Using the flag package for command line processing

### Code

`nlp` project

<pre>
├── <a href="html/nlp/go.mod.html">go.mod</a> - Project & dependencies
├── <a href="html/nlp/nlp.go.html">nlp.go</a> - Package code
├── <a href="html/nlp/doc.go.html">doc.go</a> - Package level documentation
├── <a href="html/nlp/nlp_test.go.html">nlp_test.go</a> - Test & benchmark file
├── <a href="html/nlp/example_test.go.html">example_test.go</a> - Testable example
├── stemmer - Sub module
│   ├── <a href="html/nlp/stemmer/stemmer.go.html">stemmer.go</a>
│   └── <a href="html/nlp/stemmer/stemmer_test.go.html">stemmer_test.go</a>
├── testdata - Test data
│      └── <a href="html/nlp/testdata/tokenize_cases.yml.html">tokenize_cases.yml</a> - Test cases
└── cmd  - Executables
    └── nlpd - HTTP server
        ├── <a href="html/nlp/cmd/nlpd/main.go.html">main.go</a>
        └── <a href="html/nlp/cmd/nlpd/main_test.go.html">main_test.go</a>
</pre>


[Terminal Log](_extra/day-4.log)


### Links

- Configuration
    - [envconfig](https://github.com/kelseyhightower/envconfig)
    - [viper](https://github.com/spf13/viper) & [cobra](https://github.com/spf13/cobra)
- Logging 
    - Built-in [log](https://pkg.go.dev/log/)
    - [uber/zap](https://pkg.go.dev/go.uber.org/zap)
    - [logrus](https://github.com/sirupsen/logrus)
    - [zerolog](https://github.com/rs/zerolog)
- Metrics
    - Built-in [expvar](https://pkg.go.dev/expvar/)
    - [Open Telemetry](https://opentelemetry.io/)
    - [Prometheus](https://pkg.go.dev/github.com/prometheus/client_golang/prometheus)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)
- [Example Project Structure](https://github.com/ardanlabs/service)
- [How to Write Go Code](https://go.dev/doc/code.html)
- Documentation
    - [Godoc: documenting Go code](https://go.dev/blog/godoc)
    - [Testable examples in Go](https://go.dev/blog/examples)
    - [Go documentation tricks](https://godoc.org/github.com/fluhus/godoc-tricks)
    - [gob/doc.go](https://github.com/golang/go/blob/master/src/encoding/gob/doc.go) of the `gob` package. Generates [this documentation](https://pkg.go.dev/encoding/gob/)
    - `go install golang.org/x/pkgsite/cmd/pkgsite@9ffe8b928e4fbd3ff7dcf984254629a47f8b6e63` (require go 1.18)
    - `pkgsite -http=:8080` (open browser on http://localhost:8080/${module name})
- [Out Software Dependency Problem](https://research.swtch.com/deps) - Good read on dependencies by Russ Cox
- Linters (static analysis)
    - [staticcheck](https://staticcheck.io/)
    - [gosec](https://github.com/securego/gosec) - Security oriented
    - [golang.org/x/tools/go/analysis](https://pkg.go.dev/golang.org/x/tools/go/analysis) - Helpers to write analysis tools (see [example](https://arslan.io/2019/06/13/using-go-analysis-to-write-a-custom-linter/))
- Testing
    - [testing](https://pkg.go.dev/testing/)
    - [testify](https://pkg.go.dev/github.com/stretchr/testify) - Many test utilities (including suites & mocking)
    - [Tutorial: Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)
        - [testing/quick](https://pkg.go.dev/testing/quick) - Initial fuzzing library
    - [test containers](https://golang.testcontainers.org/)
- HTTP Servers
    - [net/http](https://pkg.go.dev/net/http/)
    - [net/http/httptest](https://pkg.go.dev/net/http/httptest)
    - [gorilla/mux](https://github.com/gorilla/mux) - HTTP router with more frills
    - [chi](https://github.com/go-chi/chi) - A nice web framework

### Data & Other

- [nlp.go](_extra/nlp.go)
- [stemmer.go](_extra/stemmer.go)
- [tokenize_cases.toml](_extra/tokenize_cases.toml)
    - `github.com/BurntSushi/toml`

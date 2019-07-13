# Splitex (Split & execute)

[![Go Report Card](https://goreportcard.com/badge/github.com/symfony-doge/splitex)](https://goreportcard.com/report/github.com/symfony-doge/splitex)
[![GoDoc](https://godoc.org/github.com/symfony-doge/splitex?status.svg)](https://godoc.org/github.com/symfony-doge/splitex)
[![GitHub](https://img.shields.io/github/license/symfony-doge/splitex.svg)](LICENSE)

Splitex is a Go package that helps to balance some heavy work across all requests
by splitting a task between multiple execution flows.
A splitting (and partial results merge) algorithm should be provided by the user according to the task context.

## Installation

```
$ go get -u -d github.com/symfony-doge/splitex
```

## Usage

### DefaultWorkerPool

[DefaultWorkerPool](default_worker_pool.go) acts like a subscriber and uses 

See [example](example/concurrent_slice_sum.go) code snippet:

```go
listenerSession := event.MustListen(cssConsumeFunc)
defer func() {
	listenerSession.Close()

	fmt.Println("Sum:", cssSum)
}()

var workerPool = splitex.DefaultWorkerPoolWith(NewExampleSplitter(), NewExampleWorkerFactory())

var data = generateData()
var notifyChannel chan<- event.Event = listenerSession.NotifyChannel()

var waitGroup, distributeErr = workerPool.Distribute(data, notifyChannel)
if nil != distributeErr {
	fmt.Println("An error has been occurred during Distribute call:", distributeErr)

	os.Exit(1)
}

waitGroup.Wait()
```

Partial data handler example:

```go
...

// will be executed by workers (see example worker)
func (s *ExampleService) DoSomeWork(data []int) int {
	fmt.Printf("Partial data has been received: %v\n", data)

	var sum int

	for index := range data {
		sum += data[index]
	}

	return sum
}
```

Partial results merge example:

```go
var cssSum int

// will be executed in a single "collecting" flow
func cssConsumeFunc(e event.Event) {
	fmt.Printf("An event has been received. Type: %d, Payload: %v\n", e.Type, e.Payload)

	var partialSum, isDataTypeExpected = e.Payload.(int)
	if !isDataTypeExpected {
		panic("example: event payload misuse, invalid partial result format")
	}

	cssSum += partialSum
}
```

Example output:

```
Partial data has been received: [1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1]
Partial data has been received: [1 1 1 1 1 1 1 1 1 1 1 1 1 1]
Partial data has been received: [1 1 1 1 1 1 1 1 1 1 1 1 1 1]
An event has been received. Type: 0, Payload: 16
An event has been received. Type: 0, Payload: 14
An event has been received. Type: 0, Payload: 14
Partial data has been received: [1 1 1 1 1 1 1 1 1 1 1 1 1 1]
Partial data has been received: [1 1 1 1 1 1 1 1 1 1 1 1 1 1]
An event has been received. Type: 0, Payload: 14
An event has been received. Type: 0, Payload: 14
Partial data has been received: [1 1 1 1 1 1 1 1 1 1 1 1 1 1]
An event has been received. Type: 0, Payload: 14
Partial data has been received: [1 1 1 1 1 1 1 1 1 1 1 1 1 1]
An event has been received. Type: 0, Payload: 14
Sum: 100
```

## See also

- [panjf2000/ants](https://github.com/panjf2000/ants) — A high-performance goroutine pool for Go, inspired by fasthttp.
- [Jeffail/tunny](https://github.com/Jeffail/tunny) — A goroutine pool for Go.
- [gammazero/workerpool](https://github.com/gammazero/workerpool) — Concurrency limiting goroutine pool.

## Changelog

All notable changes to this project will be documented in [CHANGELOG.md](CHANGELOG.md).

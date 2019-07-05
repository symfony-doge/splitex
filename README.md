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

## See also

- [panjf2000/ants](https://github.com/panjf2000/ants) — A high-performance goroutine pool for Go, inspired by fasthttp.
- [Jeffail/tunny](https://github.com/Jeffail/tunny) — A goroutine pool for Go.
- [gammazero/workerpool](https://github.com/gammazero/workerpool) — Concurrency limiting goroutine pool.

## Changelog

All notable changes to this project will be documented in [CHANGELOG.md](CHANGELOG.md).

# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Changed

- TODO: shortcut `WorkerPool.DistributeAndMerge(interface{}, MergeFunc)`.

## [v0.1.0] - 2019-07-17
### Added

- `Worker`, `WorkerFactory`, `WorkerPool` interfaces.
- `TaskSplitter` interface.
- `DefaultWorkerPool` implementation & usage example (based on `GOMAXPROCS`).

[Unreleased]: https://github.com/symfony-doge/splitex/compare/v0.1.0...v0.x
[v0.1.0]: https://github.com/symfony-doge/splitex/releases/tag/v0.1.0
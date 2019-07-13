// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package example

import (
	"github.com/symfony-doge/splitex"
)

// ExampleWorkerFactory implements splitex.WorkerFactory interface.
type ExampleWorkerFactory struct{}

// CreateFor creates a worker instance for the given task or a data set.
func (wf *ExampleWorkerFactory) CreateFor(task interface{}) (splitex.Worker, error) {
	switch task.(type) {
	case []int:
		return NewExampleWorker(), nil
	default:
		return nil, splitex.UndefinedWorkerError{task}
	}
}

// NewExampleWorkerFactory creates a new worker factory instance.
func NewExampleWorkerFactory() *ExampleWorkerFactory {
	return &ExampleWorkerFactory{}
}

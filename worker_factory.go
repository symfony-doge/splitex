// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package splitex

import (
	"fmt"
)

// WorkerFactory is responsible for Worker instantiation. It receives a task
// instance and should return a valid Worker for it.
type WorkerFactory interface {
	CreateFor(interface{}) (Worker, error)
}

// UndefinedWorkerError can be returned by the WorkerFactory implementation if
// an applicable Worker is not found for task execution.
type UndefinedWorkerError struct {
	task interface{}
}

// UndefinedWorkerError implements the error interface.
func (err UndefinedWorkerError) Error() string {
	return fmt.Sprintf("Worker for task is not defined (task=%T)", err.task)
}

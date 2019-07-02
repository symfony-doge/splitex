// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package splitex

import (
	"fmt"
	"sync"

	"github.com/symfony-doge/event"
)

// WorkerPool splits a task to separate parts and delegates their execution
// among all available workers; related task splitter and a factory method
// to construct appropriate worker should be implemented by the end-user.
type WorkerPool interface {
	// Receives a task to be splitted into parts and a channel for events
	// from workers. Returns a wait group instance if workers are successfully
	// started.
	Distribute(interface{}, chan<- event.Event) (*sync.WaitGroup, error)
}

// WorkerNotPreparedError can be returned by the WorkerPool implementation if
// workers could not be prepared for execution.
type WorkerNotPreparedError struct {
	reason string
	task   interface{}
}

// WorkerNotPreparedError implements the error interface.
func (err WorkerNotPreparedError) Error() string {
	return fmt.Sprintf(
		"Unable to prepare workers for execution (task=%T, reason=%q)",
		err.task,
		err.reason,
	)
}

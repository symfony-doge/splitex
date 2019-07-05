// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package splitex

import (
	"context"
	"fmt"
)

// TaskSplitter encapsulates an algorithm for splitting a task into parts for
// concurrent execution.
type TaskSplitter interface {
	// IsSplittable receives a task instance and should return negatively
	// whenever it cannot be divided into separate and independent parts.
	IsSplittable(interface{}) (bool, error)

	// Split receives a task and the number of parts for splitting; returns
	// a list of contexts with partial data for concurrent execution by workers.
	// Component that implements TaskSplitter interface should provide
	// the task-specific splitting algorithm, otherwise UndefinedSplitAlgorithmError
	// must be returned.
	Split(interface{}, int) ([]context.Context, error)
}

// UndefinedSplitAlgorithmError can be returned by the TaskSplitter implementation
// if an applicable splitting algorithm is not defined for the received task
// instance.
type UndefinedSplitAlgorithmError struct {
	Task interface{}
}

// UndefinedSplitAlgorithmError implements the error interface.
func (err UndefinedSplitAlgorithmError) Error() string {
	return fmt.Sprintf(
		"Split algorithm for the task is not defined (task=%T)",
		err.Task,
	)
}

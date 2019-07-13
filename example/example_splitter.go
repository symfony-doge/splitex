// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package example

import (
	"context"
	"errors"

	"github.com/symfony-doge/splitex"
)

// ExampleSplitter splits a given slice into specified number of parts.
// Implements splitex.TaskSplitter interface.
type ExampleSplitter struct{}

// IsSplittable checks if a given task is supported by the splitter or not.
func (s *ExampleSplitter) IsSplittable(task interface{}) (bool, error) {
	_, isDataTypeExpected := task.([]int)

	if !isDataTypeExpected {
		return false, splitex.UndefinedSplitAlgorithmError{task}
	}

	return true, nil
}

// Split divides a task into partsCount separate tasks.
func (s *ExampleSplitter) Split(task interface{}, partsCount int) ([]context.Context, error) {
	if partsCount < 1 {
		var err = errors.New("example: parts count for task splitting should be 1 or greater")

		return []context.Context{}, err
	}

	data, isDataTypeExpected := task.([]int)
	if !isDataTypeExpected {
		return []context.Context{}, splitex.UndefinedSplitAlgorithmError{task}
	}

	// A single execution flow case, no splitting actually required.
	var partLen = len(data) / partsCount
	if partsCount < 2 || partLen < 1 {
		return []context.Context{NewDataContext(data)}, nil
	}

	var contexts []context.Context

	for i := 0; i < partsCount-1; i++ {
		var context = s.extractByRange(data, i*partLen, (i+1)*partLen)

		contexts = append(contexts, context)
	}

	var last = s.extractByRange(data, (partsCount-1)*partLen, len(data))
	contexts = append(contexts, last)

	return contexts, nil
}

func (s *ExampleSplitter) extractByRange(data []int, lowerBound, upperBound int) context.Context {
	var dataPart = data[lowerBound:upperBound]
	var context = NewDataContext(dataPart)

	return context
}

// NewExampleSplitter creates a new splitter instance.
func NewExampleSplitter() *ExampleSplitter {
	return &ExampleSplitter{}
}

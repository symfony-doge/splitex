// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package example

import (
	"fmt"
)

// ExampleService represents a handler with some data processing logic.
type ExampleService struct{}

// DoSomeWork performs some work with partial data from worker.
func (s *ExampleService) DoSomeWork(data []int) int {
	fmt.Printf("Partial data has been received: %v\n", data)

	var sum int

	for index := range data {
		sum += data[index]
	}

	return sum
}

// NewExampleService creates a new data handler instance.
func NewExampleService() *ExampleService {
	return &ExampleService{}
}

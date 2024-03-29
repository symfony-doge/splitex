// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package example

import (
	"context"

	"github.com/symfony-doge/event"
)

// Worker implements splitex.Worker interface.
type ExampleWorker struct {
	// Context with partial data and any additional settings for worker.
	context context.Context

	// Channels for worker events.
	channelsToNotify []chan<- event.Event

	// Service performs an actual work with partial data.
	service *ExampleService
}

// SetContext sets the worker's context for task execution.
func (w *ExampleWorker) SetContext(context context.Context) {
	w.context = context
}

// AddNotifyChannel adds a channel for worker events.
func (w *ExampleWorker) AddNotifyChannel(notifyChannels ...chan<- event.Event) {
	w.channelsToNotify = append(w.channelsToNotify, notifyChannels...)
}

// Run extracts partial data from the context and passes it to the service
// for processing.
func (w *ExampleWorker) Run() {
	var data, isValidData = DataFromContext(w.context)
	if !isValidData {
		panic("example: worker context misuse, invalid data format")
	}

	var sum = w.service.DoSomeWork(data)
	var event = NewPartialSumCalculatedEvent(sum)

	for channelIndex := range w.channelsToNotify {
		w.channelsToNotify[channelIndex] <- event
	}
}

// NewExampleWorker creates a new worker instance.
func NewExampleWorker() *ExampleWorker {
	return &ExampleWorker{
		service: NewExampleService(),
	}
}

// Key for storing a task instance as a value within context.Context.
type taskKey int

var exampleDataKey taskKey

// NewDataContext creates a new context instance with specified data.
func NewDataContext(data []int) context.Context {
	return context.WithValue(context.Background(), exampleDataKey, data)
}

// DataFromContext extracts a data set from the context.
func DataFromContext(context context.Context) ([]int, bool) {
	data, isDataTypeExpected := context.Value(exampleDataKey).([]int)

	return data, isDataTypeExpected
}

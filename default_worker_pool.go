// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package splitex

import (
	"runtime"
	"sync"

	"github.com/symfony-doge/event"
)

const (
	// Number of execution flows reserved (for environment with 2 and more CPU
	// cores), e.g. for processing results from workers in parallel or limiting
	// a goroutines count yielded per single Distribute call.
	dwpExecutionFlowsLimiting = 1
)

// DefaultWorkerPool starts a worker with part of the specified task for each
// available CPU core (any limiting logic is taken into account).
type DefaultWorkerPool struct {
	// Encapsulates the task splitting algorithm.
	taskSplitter TaskSplitter

	// Instantiates workers.
	workerFactory WorkerFactory

	workers []Worker
}

// Receives a task to be splitted into parts and a channel for events
// from workers. Returns a wait group instance if workers are successfully
// started.
func (wp *DefaultWorkerPool) Distribute(
	task interface{},
	notifyChannel chan<- event.Event,
) (*sync.WaitGroup, error) {
	if prepareErr := wp.prepareWorkers(task, notifyChannel); nil != prepareErr {
		return nil, WorkerNotPreparedError{task, prepareErr.Error()}
	}

	return wp.runWorkers()
}

// Creates workers and sets their execution contexts.
func (wp *DefaultWorkerPool) prepareWorkers(
	task interface{},
	notifyChannel chan<- event.Event,
) error {
	var workerCount, resolveErr = wp.resolveWorkerCount(task)
	if nil != resolveErr {
		return resolveErr
	}

	wp.workers = make([]Worker, workerCount)

	var contexts, taskSplitErr = wp.taskSplitter.Split(task, workerCount)
	if nil != taskSplitErr {
		return taskSplitErr
	}

	for workerNumber := range wp.workers {
		var worker, createErr = wp.workerFactory.CreateFor(task)
		if nil != createErr {
			return createErr
		}

		worker.SetContext(contexts[workerNumber])
		worker.AddNotifyChannel(notifyChannel)

		wp.workers[workerNumber] = worker
	}

	return nil
}

// Runs all prepared workers and returns a wait group to track their activity.
func (wp *DefaultWorkerPool) runWorkers() (*sync.WaitGroup, error) {
	var workerCount = len(wp.workers)

	var waitGroup sync.WaitGroup
	waitGroup.Add(workerCount)

	for workerNumber := range wp.workers {
		// We should not capture loop variables in closure, goroutine will
		// see only the last assigned value; instead, we pass a copy as an argument.
		go func(wn int) {
			defer waitGroup.Done()
			wp.workers[wn].Run()
		}(workerNumber)
	}

	return &waitGroup, nil
}

// Determines concurrent workers count for the specified task instance.
func (wp *DefaultWorkerPool) resolveWorkerCount(task interface{}) (int, error) {
	var workerCount int = runtime.GOMAXPROCS(0) - dwpExecutionFlowsLimiting

	// There is no reason to gain a splitting and communication overhead,
	// if only one execution flow is available.
	if workerCount < 2 {
		return 1, nil
	}

	// There can be a set of task-specific conditions, when it should be
	// splitted and when not (e.g. small data amount).
	var isTaskSplittable, checkErr = wp.taskSplitter.IsSplittable(task)
	if nil != checkErr {
		return 0, checkErr
	}

	if !isTaskSplittable {
		return 1, nil
	}

	return workerCount, nil
}

// Returns a new DefaultWorkerPool instance with specified TaskSplitter and
// WorkerFactory implementations.
func DefaultWorkerPoolWith(ts TaskSplitter, wf WorkerFactory) *DefaultWorkerPool {
	return &DefaultWorkerPool{
		taskSplitter:  ts,
		workerFactory: wf,
	}
}

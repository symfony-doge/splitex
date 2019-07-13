// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package example

import (
	"fmt"
	"os"

	"github.com/symfony-doge/event"
	"github.com/symfony-doge/splitex"
)

func cssConsumeFunc(e event.Event) {
	fmt.Printf("An event has been received. Type: %d, Payload: %v\n", e.Type, e.Payload)
}

func generateData() []int {
	var data = make([]int, 100)

	for index := range data {
		data[index] = 1
	}

	return data
}

// ConcurrentSliceSum is a demo code snippet that represents a splitex use case.
func ConcurrentSliceSum() {
	fmt.Println("Calculating sum of slice elements using multiple execution flows...")

	fmt.Println("Starting a listening session...")

	listenerSession := event.MustListen(cssConsumeFunc)
	defer listenerSession.Close()

	var notifyChannel chan<- event.Event = listenerSession.NotifyChannel()

	fmt.Println("Distributing sum calculation between multiple execution flows...")

	var workerPool = splitex.DefaultWorkerPoolWith(NewExampleSplitter(), NewExampleWorkerFactory())
	var data = generateData()

	var waitGroup, distributeErr = workerPool.Distribute(data, notifyChannel)
	if nil != distributeErr {
		fmt.Println("An error has been occurred during Distribute call:", distributeErr)

		os.Exit(1)
	}

	waitGroup.Wait()
}

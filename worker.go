// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package splitex

import (
	"context"

	"github.com/symfony-doge/event"
)

// Worker performs task processing routine and emits events, e.g. with result
// payload, exceptional state alerts, etc.
type Worker interface {
	// Sets context of the partial task for processing.
	SetContext(context.Context)

	// Adds a channel or group of channels for events pushing. Worker should
	// send an event to each notification channel registered that way.
	AddNotifyChannel(...chan<- event.Event)

	// Starts routine execution.
	Run()
}

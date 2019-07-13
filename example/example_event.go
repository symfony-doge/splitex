// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package example

import (
	"github.com/symfony-doge/event"
)

const (
	// Whenever a sum of elements from slice has been calculated by the service.
	PartialSumCalculatedEvent event.EventType = iota
)

// PartialSumCalculatedContext contains an event-specific data.
type PartialSumCalculatedContext struct {
	// Holds a sum of elements from slice.
	Value int
}

// NewPartialSumCalculatedEvent creates an event instance.
func NewPartialSumCalculatedEvent(context PartialSumCalculatedContext) event.Event {
	return event.WithTypeAndPayload(PartialSumCalculatedEvent, context)
}

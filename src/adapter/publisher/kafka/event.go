package kafka

import (
	"gitlab.com/golibs-starter/golib/web/event"
)

type Event[T any] struct {
	*event.AbstractEvent
	PayloadData T `json:"payload"`
}

func (e Event[T]) Payload() T {
	return e.PayloadData
}

func (e Event[T]) String() string {
	return e.ToString(e)
}

package kafka

import (
	"context"
	"errors"
)

type EventHandler func(ctx context.Context, payload []byte) error

type EventRouter struct {
	handlers map[string]EventHandler
}

func NewRouter() *EventRouter {
	return &EventRouter{
		handlers: make(map[string]EventHandler),
	}
}

func (r *EventRouter) Register(eventType string, handler EventHandler) {
	r.handlers[eventType] = handler
}

func (r *EventRouter) Route(ctx context.Context, eventType string, payload []byte) error {
	handler, ok := r.handlers[eventType]
	if !ok {
		return errors.New("no handler for event type: " + eventType)
	}
	return handler(ctx, payload)
}

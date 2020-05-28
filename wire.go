//+build wireinject

package main

import (
	"github.com/google/wire"
)

func initializeEvent(msg string) (event, error) {
	wire.Build(newMessage, newGreeter, newEvent)
	return event{}, nil
}

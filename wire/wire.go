//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreater, NewMessage)
	return Event{}
}

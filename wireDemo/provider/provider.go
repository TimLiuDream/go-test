//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"
	"github.com/timliudream/go-test/wireDemo/types"
)

var superProvider = wire.NewSet(types.NewMessage, types.NewGreeter, types.NewEvent)

func InitializeEvent(phrase string) (types.Event, error) {
	wire.Build(superProvider)
	return types.Event{}, nil
}

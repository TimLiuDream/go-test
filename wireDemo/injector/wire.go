//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/timliudream/go-test/wireDemo/types"
)

func InitializeEvent(phrase string) (types.Event, error) {
	wire.Build(types.NewEvent, types.NewGreeter, types.NewMessage)
	return types.Event{}, nil
}

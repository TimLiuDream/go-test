package interview

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/timliudream/go-test/serviceWeaverDemo/model"
)

type Service interface {
	MakeInterview(ctx context.Context, golang model.Golang) error
}

type implementation struct {
	weaver.Implements[Service]
}

func (s *implementation) MakeInterview(ctx context.Context, golang model.Golang) error {
	defer s.Logger(ctx).Info(
		"talk about golang",
		"channel: ", golang.Channel,
		"goroutine: ", golang.Goroutine,
	)
	return nil
}

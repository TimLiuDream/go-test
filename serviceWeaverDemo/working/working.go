package working

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/timliudream/go-test/serviceWeaverDemo/model"
)

type Service interface {
	Working(ctx context.Context, golang model.Golang) error
}

type implementation struct {
	weaver.Implements[Service]
}

func (s *implementation) Working(ctx context.Context, golang model.Golang) error {
	defer s.Logger(ctx).Info(
		"working, fixing golang  bug",
		"channel bug: ", golang.Channel,
		"goroutines bug: ", golang.Goroutine,
	)
	return nil
}

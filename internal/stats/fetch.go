package stats

import (
	"context"

	"google.golang.org/protobuf/types/known/anypb"
)

type Fetch interface {
	Fetch(ctx context.Context) (anypb.Any, error)
}

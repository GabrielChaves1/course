package context

import (
	"context"
	"time"
)

const (
	CurrentTimeContextKey = "currentTime"
)

func NewContextWithCurrentTime(ctx context.Context, currentTime time.Time) context.Context {
	return context.WithValue(ctx, CurrentTimeContextKey, currentTime)
}

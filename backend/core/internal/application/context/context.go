package context

import (
	"context"
	"fmt"
	"time"

	"github.com/GabrielChaves1/course/internal/domain/types"
)

const (
	CurrentTimeContextKey = "currentTime"
	UserIDContextKey      = "userId"
)

func NewContextFieldNotFoundError(fieldName interface{}) error {
	return fmt.Errorf("couldn't find context field with name %q", fieldName)
}

func NewContextWithCurrentTime(ctx context.Context, currentTime time.Time) context.Context {
	return context.WithValue(ctx, CurrentTimeContextKey, currentTime)
}

func NewContextWithUserID(ctx context.Context, id types.UserID) context.Context {
	return context.WithValue(ctx, UserIDContextKey, id)
}

func ExtractUserIDFromContext(ctx context.Context) (types.UserID, error) {
	userID, ok := ctx.Value(UserIDContextKey).(types.UserID)
	if !ok {
		return types.UserID{}, NewContextFieldNotFoundError(UserIDContextKey)
	}

	return userID, nil
}

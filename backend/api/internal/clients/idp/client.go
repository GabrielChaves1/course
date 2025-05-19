package idp

import "context"

type Client interface {
	Authenticate(ctx context.Context, username, password string) error
}

package middleware_basic

import "context"

type BasicMiddleware struct {
}

func (b *BasicMiddleware) BasicValidateToken(ctx context.Context, credentials string) error {
	return nil
}

package middleware_jwt

import "context"

type JWTMiddleware struct {
}

func (j *JWTMiddleware) JWTValidateToken(ctx context.Context, credentials string) error {
	return nil
}

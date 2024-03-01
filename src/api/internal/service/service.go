package service_api

import (
	"context"
	"menti/pkg/types"
)

type ServiceCounter interface {
	Concate(a, b int) int
}

type Repo interface {
	CreateUser(ctx context.Context, data types.UserPayload) error
	GetUser(ctx context.Context, ID int) (user types.User, err error)
	GetToken()
	CreateToken()
}

type MiddleWare interface {
	BasicValidateToken(ctx context.Context, credentials string) error
	JWTValidateToken(ctx context.Context, credentials string) error
}

type Service struct {
	middleware MiddleWare
	counter    ServiceCounter
	repo       Repo
}

func NewService(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) ConcateLogic(a, b int) int {
	result := s.counter.Concate(a, b)
	return result
}

func (s *Service) MiddleWareBasic(ctx context.Context, credentials string) error {
	if err := s.middleware.BasicValidateToken(ctx, credentials); err != nil {
		return err
	}
	return nil
}

func (s *Service) MiddleWareJWT(ctx context.Context, credentials string) error {
	if err := s.middleware.JWTValidateToken(ctx, credentials); err != nil {
		return err
	}
	return nil
}

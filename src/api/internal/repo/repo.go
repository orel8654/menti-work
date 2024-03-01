package repo_api

import (
	"context"
	"menti/pkg/types"
)

type RepoBasic interface {
	CreateUser(ctx context.Context, user types.UserPayload) error
	GetUser(ctx context.Context, username string) (res types.User, err error)
	GetBasicToken()
	CreateBasicToken()
}

type Repo struct {
	repo RepoBasic
}

func NewRepo(repo RepoBasic) *Repo {
	return &Repo{
		repo: repo,
	}
}

func (r *Repo) CreateUser(ctx context.Context, data types.UserPayload) error {
	if err := r.repo.CreateUser(ctx, types.UserPayload{}); err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetUser(ctx context.Context, ID int) (user types.User, err error) {
	return user, err
}

func (r *Repo) GetToken() {
	r.repo.GetBasicToken()
}

func (r *Repo) CreateToken() {
	r.repo.CreateBasicToken()
}

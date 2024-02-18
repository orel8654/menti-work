package psql_pkg

import (
	"context"
	"menti/pkg/types"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db sqlx.DB) *Repo {
	return &Repo{
		db: &db,
	}
}

func (r *Repo) CreateUser(ctx context.Context, user types.User) error {
	querry := `
		INSERT INTO user (uuid, password, username)
		VALUES (:uuid, :password, :username)
	`
	_, err := r.db.NamedExecContext(ctx, querry, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetUser(ctx context.Context, username string) (res types.User, err error) {
	query := `
		SELECT id_key, uuid, password, username
		WHERE username = $1
	`
	row := r.db.QueryRowContext(ctx, query, username)
	if err = row.Scan(&res); err != nil {
		return res, err
	}
	return res, nil
}

func (r *Repo) GetBasicToken() {
	querry := ``
	_ = querry
}

func (r *Repo) CreateBasicToken() {
	querry := ``
	_ = querry
}

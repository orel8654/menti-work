package repo

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

func (r *Repo) CreateUser(ctx context.Context, user types.UserPayloadCreat) error {
	querry := `
		INSERT INTO "user" (uuid, password, username)
		VALUES (:uuid, :password, :username)
	`
	_, err := r.db.NamedExecContext(ctx, querry, user)
	return err
}

func (r *Repo) GetUser(ctx context.Context, username string) (res types.User, err error) {
	query := `
		SELECT id_key, uuid, password, username
		FROM "user"
		WHERE username = $1
	`
	row := r.db.QueryRowContext(ctx, query, username)
	if err = row.Scan(&res.IDKey, &res.UUID, &res.Password, &res.Username); err != nil {
		return res, err
	}
	return res, nil
}

func (r *Repo) GetBasicToken(ctx context.Context, IDkey int) (res types.UserToken, err error) {
	querry := `
		SELECT token, id_key
		FROM "user_token"
		WHERE id_key = $1
	`
	row := r.db.QueryRowxContext(ctx, querry, IDkey)
	if err = row.Scan(&res.Token, &res.IDKey); err != nil {
		return res, err
	}
	return res, nil
}

func (r *Repo) CreateBasicToken(ctx context.Context, token string, IDkey int) (res types.UserToken, err error) {
	query := `
		INSERT INTO "user_token" (token, id_key)
		VALUES ($1, $2)
		RETURNING token, id_key
	`
	err = r.db.QueryRowxContext(ctx, query, token, IDkey).Scan(&res.Token, &res.IDKey)
	return res, err
}

func (r *Repo) GetUserByBasicToken(ctx context.Context, token string) (user types.User, err error) {
	query := `
		SELECT u.id_key, u.uuid, u.password, u.username
		FROM "user_token" ut
		JOIN "user" u ON ut.id_key = u.id_key
		WHERE ut.token = $1
	`
	err = r.db.QueryRowxContext(ctx, query, token).Scan(&user.IDKey, &user.UUID, &user.Password, &user.Username)
	return user, err
}

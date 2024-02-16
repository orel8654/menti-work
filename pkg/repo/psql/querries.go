package psql_pkg

import "github.com/jmoiron/sqlx"

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db sqlx.DB) *Repo {
	return &Repo{
		db: &db,
	}
}

func (r *Repo) GetBasicToken() {
	querry := ``
	_ = querry
}

func (r *Repo) CreateBasicToken() {
	querry := ``
	_ = querry
}

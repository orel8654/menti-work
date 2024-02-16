package repo_api

type RepoBasic interface {
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

func (r *Repo) GetToken() {
	r.repo.GetBasicToken()
}

func (r *Repo) CreateToken() {
	r.repo.CreateBasicToken()
}

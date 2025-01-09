package helperrepo

import (
	"desabiller/repositories"
)

type helper struct {
	repo repositories.Repositories
}

func NewHelperRepo(repo repositories.Repositories) helper {
	return helper{
		repo: repo,
	}
}

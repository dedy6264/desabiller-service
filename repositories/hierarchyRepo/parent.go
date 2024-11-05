package hierarchyrepo

import (
	"desabiller/repositories"
)

type hierarchy struct {
	repo repositories.Repositories
}

func NewHierarcyRepo(repo repositories.Repositories) hierarchy {
	return hierarchy{
		repo: repo,
	}
}

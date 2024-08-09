package hierarchyrepo

import (
	"desabiller/repositories"
)

type nHierarchy struct {
	repo repositories.Repositories
}

func NewNHierarcyRepo(repo repositories.Repositories) nHierarchy {
	return nHierarchy{
		repo: repo,
	}
}

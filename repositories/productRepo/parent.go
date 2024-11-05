package productrepo

import (
	"desabiller/repositories"
)

type product struct {
	repo repositories.Repositories
}

func NewProductRepo(repo repositories.Repositories) product {
	return product{
		repo: repo,
	}
}

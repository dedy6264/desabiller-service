package trxrepo

import "desabiller/repositories"

type trxRepository struct {
	repo repositories.Repositories
}

func NewTrxRepo(repo repositories.Repositories) trxRepository {
	return trxRepository{
		repo: repo,
	}
}

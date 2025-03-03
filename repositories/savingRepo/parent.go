package savingrepo

import "desabiller/repositories"

type savingRepository struct {
	repo repositories.Repositories
}

func NewSavingRepo(repo repositories.Repositories) savingRepository {
	return savingRepository{
		repo: repo,
	}
}

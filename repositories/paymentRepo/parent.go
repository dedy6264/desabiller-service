package paymentrepo

import "desabiller/repositories"

type paymentRepo struct {
	repo repositories.Repositories
}

func NewPaymentRepo(repo repositories.Repositories) paymentRepo {
	return paymentRepo{
		repo: repo,
	}
}

package trxservice

import "desabiller/services"

type trxService struct {
	services services.UsecaseService
}

func NewRepoTrxService(services services.UsecaseService) trxService {
	return trxService{
		services: services,
	}
}

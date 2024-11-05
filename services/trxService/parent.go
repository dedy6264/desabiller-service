package trxservice

import "desabiller/services"

type trxService struct {
	services services.UsecaseService
}

func NewApiTrxService(services services.UsecaseService) trxService {
	return trxService{
		services: services,
	}
}

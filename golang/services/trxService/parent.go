package trxservice

import "desabiller/services"

type TrxServices struct {
	service services.UsecaseService
}

func NewApiTrxService(service services.UsecaseService) TrxServices {
	return TrxServices{
		service: service,
	}
}

package savingservices

import "desabiller/services"

type savingServices struct {
	services services.UsecaseService
}

func NewApiSavingServices(service services.UsecaseService) savingServices {
	return savingServices{services: service}
}

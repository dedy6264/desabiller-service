package providerservices

import "desabiller/services"

type providerServices struct {
	services services.UsecaseService
}

func NewApiProviderServices(service services.UsecaseService) providerServices {
	return providerServices{services: service}
}

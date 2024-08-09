package vueservice

import "desabiller/services"

type VueService struct {
	services services.UsecaseService
}

func NewApiVueService(service services.UsecaseService) VueService {
	return VueService{
		services: service,
	}
}

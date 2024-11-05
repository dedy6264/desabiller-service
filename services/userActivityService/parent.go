package useractivityservice

import "desabiller/services"

type userActivity struct {
	services services.UsecaseService
}

func NewApiUserActivityService(services services.UsecaseService) userActivity {
	return userActivity{
		services: services,
	}
}

package administrationservice

import "desabiller/services"

type AdministrationService struct {
	service services.UsecaseService
}

func ApiAdministration(service services.UsecaseService) AdministrationService {
	return AdministrationService{
		service: service,
	}
}

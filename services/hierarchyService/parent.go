package hierarchyservice

import "desabiller/services"

type HierarcyService struct {
	service services.UsecaseService
}

func RepoHierarchy(service services.UsecaseService) HierarcyService {
	return HierarcyService{
		service: service,
	}
}

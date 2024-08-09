package productservice

import "desabiller/services"

type ProductService struct {
	service services.UsecaseService
}

func ApiProduct(service services.UsecaseService) ProductService {
	return ProductService{
		service: service,
	}
}

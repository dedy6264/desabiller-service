package services

import (
	"database/sql"
	"desabiller/repositories"
)

type UsecaseService struct {
	RepoDB *sql.DB
	// cron   *cron.Cron
	// GenAutonumRepo                        genautonum.GenerateAutonumberRepository
	ApiHierarchy repositories.Hierarchyrepo
	ApiProduct   repositories.ProductRepo
	ApiTrx       repositories.TrxRepo
}

func NewUsecaseService(
	repoDB *sql.DB,
	ApiHierarchyRepo repositories.Hierarchyrepo,
	ApiProduct repositories.ProductRepo,
	ApiTrx repositories.TrxRepo,

) UsecaseService {
	return UsecaseService{
		RepoDB:       repoDB,
		ApiHierarchy: ApiHierarchyRepo,
		ApiProduct:   ApiProduct,
		ApiTrx:       ApiTrx,
	}
}

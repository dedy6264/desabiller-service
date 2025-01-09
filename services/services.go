package services

import (
	"database/sql"
	"desabiller/repositories"
)

type UsecaseService struct {
	RepoDB *sql.DB
	// cron   *cron.Cron
	// GenAutonumRepo                        genautonum.GenerateAutonumberRepository
	RepoHierarchy repositories.Hierarchyrepo
	RepoProduct   repositories.ProductRepo
	RepoTrx       repositories.TrxRepo
	HelperRepo    repositories.HelperRepo
}

func NewUsecaseService(
	repoDB *sql.DB,
	RepoHierarchyRepo repositories.Hierarchyrepo,
	RepoProduct repositories.ProductRepo,
	RepoTrx repositories.TrxRepo,
	HelperRepo repositories.HelperRepo,

) UsecaseService {
	return UsecaseService{
		RepoDB:        repoDB,
		RepoHierarchy: RepoHierarchyRepo,
		RepoProduct:   RepoProduct,
		RepoTrx:       RepoTrx,
		HelperRepo:    HelperRepo,
	}
}

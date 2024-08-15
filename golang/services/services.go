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
	// ApiProduct   repositories.ProductRepo
	// ApiTrx       repositories.TrxRepo
	// ApiNoTrx     repositories.NoTrxGenerator
	// ApiPayment   repositories.PaymentRepo
	// ApiNHierarchy     repositories.NHierarchy
	// ApiNUserDashboard repositories.NUserDashboard
	// ApiNFeatures      repositories.NFeatures
}

func NewUsecaseService(
	repoDB *sql.DB,
	ApiHierarchyRepo repositories.Hierarchyrepo,
	// ApiProduct repositories.ProductRepo,
	// ApiTrx repositories.TrxRepo,
	// ApiNoTrx repositories.NoTrxGenerator,
	// ApiPayment repositories.PaymentRepo,
	// ApiNHierarchy repositories.NHierarchy,
	// // ApiMongo repositories.ApiMongoRepository,
	// ApiNUserDashboard repositories.NUserDashboard,
	// ApiNFeatures repositories.NFeatures,

) UsecaseService {
	return UsecaseService{
		RepoDB:       repoDB,
		ApiHierarchy: ApiHierarchyRepo,
		// ApiProduct:    ApiProduct,
		// ApiTrx:        ApiTrx,
		// ApiNoTrx:      ApiNoTrx,
		// ApiPayment:    ApiPayment,
		// ApiNHierarchy: ApiNHierarchy,
		// ApiMongo:     ApiMongo,,
		// 	ApiNUserDashboard: ApiNUserDashboard,
		// 	ApiNFeatures:      ApiNFeatures,
	}
}

package apps

import (
	"database/sql"
	"desabiller/repositories"
	hierarchyrepo "desabiller/repositories/hierarchyRepo"
	productrepo "desabiller/repositories/productRepo"
	"desabiller/services"
)

func SetupApp(DB *sql.DB, repo repositories.Repositories) services.UsecaseService {

	// transactionRepo := transactionRepository.NewTransactionRepository(repo)
	hierarchyRepo := hierarchyrepo.NewHierarcyRepo(repo)
	productRepo := productrepo.NewProductRepo(repo)
	// trxRepo := trxrepo.NewTrxRepo(repo)
	// trxNoRepo := trxgeneratorrepo.NewTrxNoGenerator(repo)
	// paymentRepo := paymentrepo.NewPaymentRepo(repo)
	// // mongoRepo := mongorepo.NewApiMongoRepository(repo)
	// nHierarchyRepo := nhierarchyrepo.NewNHierarcyRepo(repo)
	// nUserDashboardsRepo := nUserDashboardsRepo.NewNUserDashboardsRepo(repo)
	// nFeaturesRepo := nfeaturesrepo.NewNFeaturesRepo((repo))

	usecaseSvc := services.NewUsecaseService(
		DB,
		hierarchyRepo,
		productRepo,
		// trxRepo,
		// trxNoRepo,
		// paymentRepo,
		// // mongoRepo,
		// nHierarchyRepo,
		// nUserDashboardsRepo,
		// nFeaturesRepo,
	)

	return usecaseSvc
}

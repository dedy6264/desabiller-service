package apps

import (
	"database/sql"
	"desabiller/repositories"
	helperrepo "desabiller/repositories/helperRepo"
	hierarchyrepo "desabiller/repositories/hierarchyRepo"
	productrepo "desabiller/repositories/productRepo"
	savingrepo "desabiller/repositories/savingRepo"
	trxrepo "desabiller/repositories/trxRepo"
	"desabiller/services"
)

func SetupApp(DB *sql.DB, repo repositories.Repositories) services.UsecaseService {

	// transactionRepo := transactionRepository.NewTransactionRepository(repo)
	hierarchyRepo := hierarchyrepo.NewHierarcyRepo(repo)
	productRepo := productrepo.NewProductRepo(repo)
	trxRepo := trxrepo.NewTrxRepo(repo)
	helperRepo := helperrepo.NewHelperRepo(repo)
	savingRepo := savingrepo.NewSavingRepo(repo)
	usecaseSvc := services.NewUsecaseService(
		DB,
		hierarchyRepo,
		productRepo,
		trxRepo,
		helperRepo,
		savingRepo,
	)

	return usecaseSvc
}

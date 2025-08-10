package repositories

import (
	"database/sql"
	"desabiller/models"
)

type Hierarchyrepo interface {
	AddUserApp(req models.ReqGetUserApp, tx *sql.Tx) (models.UserApp, error)
	DropUserApp(id int, tx *sql.Tx) (err error)
	UpdateUserApp(req models.ReqGetUserApp, tx *sql.Tx) (err error)
	GetUserAppCount(req models.ReqGetUserApp) (result int, err error)
	GetUserApps(req models.ReqGetUserApp) (result []models.UserApp, err error)
	GetUserApp(req models.ReqGetUserApp) (result models.UserApp, err error)

	AddOtp(req models.ReqGetOtp, tx *sql.Tx) (result models.Otp, err error)
	DropOtp(id int, tx *sql.Tx) (err error)
	UpdateOtp(req models.ReqGetOtp, tx *sql.Tx) (err error)
	GetOtp(req models.ReqGetOtp) (result models.Otp, err error)
}
type ProductRepo interface {
	AddProductType(req models.ReqGetProductType) (result models.ProductType, err error)
	GetProductTypes(req models.ReqGetProductType) (result []models.ProductType, err error)
	UpdateProductType(req models.ReqGetProductType) (result models.ProductType, err error)
	DropProductType(req models.ReqGetProductType) (err error)
	GetProductType(req models.ReqGetProductType) (result models.ProductType, err error)
	GetProductTypeCount(req models.ReqGetProductType) (result int, err error)
	// AddProvider(req models.ReqGetProvider) (err error)
	// GetProviders(req models.ReqGetProvider) (result []models.RespGetProvider, err error)
	// UpdateProvider(req models.ReqGetProvider) (result models.RespGetProvider, err error)
	// DropProvider(req models.ReqGetProvider) (err error)
	// GetProvider(req models.ReqGetProvider) (result models.RespGetProvider, err error)
	// GetProviderCount(req models.ReqGetProvider) (result int, err error)

	// AddProductClan(req models.ReqGetProductClan) (err error)
	// GetProductClans(req models.ReqGetProductClan) (result []models.RespGetProductClan, err error)
	// UpdateProductClan(req models.ReqGetProductClan) (result models.RespGetProductClan, err error)
	// DropProductClan(req models.ReqGetProductClan) (err error)
	// GetProductClanCount(req models.ReqGetProductClan) (result int, err error)
	// GetProductClan(req models.ReqGetProductClan) (result models.RespGetProductClan, err error)

	AddProductCategory(req models.ReqGetProductCategory) (result models.ProductCategory, err error)
	GetProductCategories(req models.ReqGetProductCategory) (result []models.ProductCategory, err error)
	UpdateProductCategory(req models.ReqGetProductCategory) (result models.ProductCategory, err error)
	DropProductCategory(req models.ReqGetProductCategory) (err error)
	GetProductCategory(req models.ReqGetProductCategory) (result models.ProductCategory, err error)
	GetProductCategoryCount(req models.ReqGetProductCategory) (result int, err error)

	AddProductReference(req models.ReqGetProductReference) (result models.ProductReference, err error)
	GetProductReferences(req models.ReqGetProductReference) (result []models.ProductReference, err error)
	GetProductReferenceCount(req models.ReqGetProductReference) (result int, err error)
	GetProductReference(req models.ReqGetProductReference) (result models.ProductReference, err error)
	UpdateProductReference(req models.ReqGetProductReference) (result models.ProductReference, err error)
	DropProductReference(req models.ReqGetProductReference) (err error)

	// AddProductProvider(req models.ReqGetProductProvider) (result models.RespGetProductProvider, err error)
	// GetProductProviders(req models.ReqGetProductProvider) (result []models.RespGetProductProvider, err error)
	// GetProductProviderCount(req models.ReqGetProductProvider) (result int, err error)
	// UpdateProductProvider(req models.ReqGetProductProvider) (result models.RespGetProductProvider, err error)
	// DropProductProvider(req models.ReqGetProductProvider) (err error)
	// GetProductProvider(req models.ReqGetProductProvider) (result models.RespGetProductProvider, err error)

	AddProduct(req models.ReqGetProduct) (result models.Product, err error)
	GetProducts(req models.ReqGetProduct) (result []models.Product, err error)
	GetProductCount(req models.ReqGetProduct) (result int, err error)
	GetProduct(req models.ReqGetProduct) (result models.Product, err error)
	UpdateProduct(req models.ReqGetProduct) (result models.Product, err error)
	DropProduct(req models.ReqGetProduct) (err error)
}

type TrxRepo interface {
	GetTrx(req models.ReqGetTransaction) (result models.RespGetTrx, err error)
	GetTrxCount(req models.ReqGetTransaction) (result int, err error)
	GetTrxs(req models.ReqGetTransaction) (result []models.RespGetTrx, err error)
	GetPaymentTrxs(req models.ReqGetTransaction) (result []models.RespGetTrx, err error)
	InsertTrx(req models.ReqGetTransaction, tx *sql.Tx) (err error)
	UpdateTrx(req models.ReqGetTransaction, tx *sql.Tx) (err error)

	// InsertTrxStatus(req models.ReqGetTrxStatus, tx *sql.Tx) (err error)

	GenerateNo(datatype string, prefix string, leadingZero ...int) (code string, err error)
}
type TrxNoGenerator interface {
	GetLastTrxNo() (noTrx string, status bool)
	InsertTrxNo(noTrx string) (id int, status bool)
}
type HelperRepo interface {
	GetProductReferenceById(subscriberId string) (result models.RespGetPrefix, err error)
}
type SavingRepo interface {
	DropCif(id int, tx *sql.Tx) (err error)
	UpdateCif(req models.ReqGetCIF, tx *sql.Tx) (err error)
	AddCif(req models.ReqGetCIF, tx *sql.Tx) (result models.CIF, err error)
	GetCif(req models.ReqGetCIF) (result models.CIF, err error)
	GetCifs(req models.ReqGetCIF) (result []models.CIF, err error)
	GetCifCount(req models.ReqGetCIF) (result int, err error)

	GetSavingTypeCount(req models.ReqGetSavingType) (result int, err error)
	DropSavingType(id int, tx *sql.Tx) (err error)
	UpdateSavingType(req models.ReqGetSavingType, tx *sql.Tx) (err error)
	AddSavingType(req models.ReqGetSavingType, tx *sql.Tx) (result models.SavingType, err error)
	GetSavingType(req models.ReqGetSavingType) (result models.SavingType, err error)
	GetSavingTypes(req models.ReqGetSavingType) (result []models.SavingType, err error)

	GetSavingSegmentCount(req models.ReqGetSavingSegment) (result int, err error)
	DropSavingSegment(id int, tx *sql.Tx) (err error)
	UpdateSavingSegment(req models.ReqGetSavingSegment, tx *sql.Tx) (err error)
	AddSavingSegment(req models.ReqGetSavingSegment, tx *sql.Tx) (result models.SavingSegment, err error)
	GetSavingSegment(req models.ReqGetSavingSegment) (result models.SavingSegment, err error)
	GetSavingSegments(req models.ReqGetSavingSegment) (result []models.SavingSegment, err error)

	GetAccountCount(req models.ReqGetAccountSaving) (result int, err error)
	DropAccount(id int, tx *sql.Tx) (err error)
	UpdateAccount(req models.ReqGetAccountSaving, tx *sql.Tx) (err error)
	AddAccount(req models.ReqGetAccountSaving, tx *sql.Tx) (result models.RespGetAccount, err error)
	GetAccount(req models.ReqGetAccountSaving) (result models.RespGetAccount, err error)
	GetAccounts(req models.ReqGetAccountSaving) (result []models.RespGetAccount, err error)

	GetSavingTransactionCount(req models.ReqGetSavingTransaction) (result int, err error)
	DropSavingTransaction(id int, tx *sql.Tx) (err error)
	UpdateSavingTransaction(req models.ReqGetSavingTransaction, tx *sql.Tx) (err error)
	AddSavingTransaction(req models.ReqGetSavingTransaction, tx *sql.Tx) (result models.SavingTransaction, err error)
	GetSavingTransaction(req models.ReqGetSavingTransaction) (result models.SavingTransaction, err error)
	GetSavingTransactions(req models.ReqGetSavingTransaction) (result []models.SavingTransaction, err error)
}

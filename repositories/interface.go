package repositories

import (
	"database/sql"
	"desabiller/models"
)

type Hierarchyrepo interface {
	//client
	AddClient(req models.ReqGetClient, tx *sql.Tx) (err error)
	DropClient(id int, tx *sql.Tx) (err error)
	UpdateClient(req models.ReqGetClient, tx *sql.Tx) (err error)
	GetCount(req models.ReqGetClient) (result int, err error)
	GetClients(req models.ReqGetClient) (result []models.RespGetClient, err error)
	GetClient(req models.ReqGetClient) (result models.RespGetClient, err error)

	DropGroup(id int) (err error)
	UpdateGroup(req models.ReqGetGroup) (result models.RespGetGroup, err error)
	GetGroupCount(req models.ReqGetGroup) (result int, err error)
	GetGroups(req models.ReqGetGroup) (result []models.RespGetGroup, err error)
	AddGroup(req models.ReqGetGroup, tx *sql.Tx) (err error)
	GetGroup(req models.ReqGetGroup) (result models.RespGetGroup, err error)

	DropMerchant(req models.ReqGetMerchant) (err error)
	UpdateMerchant(req models.ReqGetMerchant) (result models.RespGetMerchant, err error)
	GetMerchantCount(req models.ReqGetMerchant) (result int, err error)
	GetMerchants(req models.ReqGetMerchant) (result []models.RespGetMerchant, err error)
	AddMerchant(req models.ReqGetMerchant) (err error)
	GetMerchant(req models.ReqGetMerchant) (result models.RespGetMerchant, err error)

	DropMerchantOutlet(req models.ReqGetMerchantOutlet) (err error)
	UpdateMerchantOutlet(req models.ReqGetMerchantOutlet) (result models.RespGetMerchantOutlet, err error)
	GetMerchantOutletCount(req models.ReqGetMerchantOutlet) (result int, err error)
	GetMerchantOutlets(req models.ReqGetMerchantOutlet) (result []models.RespGetMerchantOutlet, err error)
	AddMerchantOutlet(req models.ReqGetMerchantOutlet) (err error)
	GetMerchantOutlet(req models.ReqGetMerchantOutlet) (result models.RespGetMerchantOutlet, err error)
}
type ProductRepo interface {
	AddProvider(req models.ReqGetProvider) (err error)
	GetProviders(req models.ReqGetProvider) (result []models.RespGetProvider, err error)
	UpdateProvider(req models.ReqGetProvider) (result models.RespGetProvider, err error)
	DropProvider(req models.ReqGetProvider) (err error)
	GetProvider(req models.ReqGetProvider) (result models.RespGetProvider, err error)
	GetProviderCount(req models.ReqGetProvider) (result int, err error)

	// AddProductClan(req models.ReqGetProductClan) (err error)
	// GetProductClans(req models.ReqGetProductClan) (result []models.RespGetProductClan, err error)
	// UpdateProductClan(req models.ReqGetProductClan) (result models.RespGetProductClan, err error)
	// DropProductClan(req models.ReqGetProductClan) (err error)
	// GetProductClanCount(req models.ReqGetProductClan) (result int, err error)
	// GetProductClan(req models.ReqGetProductClan) (result models.RespGetProductClan, err error)

	AddProductCategory(req models.ReqGetProductCategory) (result models.RespGetProductCategory, err error)
	GetProductCategories(req models.ReqGetProductCategory) (result []models.RespGetProductCategory, err error)
	UpdateProductCategory(req models.ReqGetProductCategory) (result models.RespGetProductCategory, err error)
	DropProductCategory(req models.ReqGetProductCategory) (err error)
	GetProductCategory(req models.ReqGetProductCategory) (result models.RespGetProductCategory, err error)
	GetProductCategoryCount(req models.ReqGetProductCategory) (result int, err error)

	AddProductType(req models.ReqGetProductType) (result models.RespGetProductType, err error)
	GetProductTypes(req models.ReqGetProductType) (result []models.RespGetProductType, err error)
	UpdateProductType(req models.ReqGetProductType) (result models.RespGetProductType, err error)
	DropProductType(req models.ReqGetProductType) (err error)
	GetProductType(req models.ReqGetProductType) (result models.RespGetProductType, err error)
	GetProductTypeCount(req models.ReqGetProductType) (result int, err error)

	AddProductProvider(req models.ReqGetProductProvider) (result models.RespGetProductProvider, err error)
	GetProductProviders(req models.ReqGetProductProvider) (result []models.RespGetProductProvider, err error)
	GetProductProviderCount(req models.ReqGetProductProvider) (result int, err error)
	UpdateProductProvider(req models.ReqGetProductProvider) (result models.RespGetProductProvider, err error)
	DropProductProvider(req models.ReqGetProductProvider) (err error)
	GetProductProvider(req models.ReqGetProductProvider) (result models.RespGetProductProvider, err error)

	AddProduct(req models.ReqGetProduct) (result models.RespGetProduct, err error)
	GetProducts(req models.ReqGetProduct) (result []models.RespGetProduct, err error)
	GetProductCount(req models.ReqGetProduct) (result int, err error)
	GetProduct(req models.ReqGetProduct) (result models.RespGetProduct, err error)
	UpdateProduct(req models.ReqGetProduct) (result models.RespGetProduct, err error)
	DropProduct(req models.ReqGetProduct) (err error)
}

type TrxRepo interface {
	GetTrx(req models.ReqGetTrx) (result models.RespGetTrx, err error)
	GetTrxCount(req models.ReqGetTrx) (result int, err error)
	GetTrxs(req models.ReqGetTrx) (result []models.RespGetTrx, err error)
	InsertTrx(req models.ReqGetTrx, tx *sql.Tx) (err error)
	UpdateTrx(req models.ReqGetTrx, tx *sql.Tx) (err error)

	InsertTrxStatus(req models.ReqGetTrxStatus, tx *sql.Tx) (err error)

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
	UpdateCif(req models.ReqGetCif, tx *sql.Tx) (err error)
	AddCif(req models.ReqGetCif, tx *sql.Tx) (result models.RespGetCif, err error)
	GetCif(req models.ReqGetCif) (result models.RespGetCif, err error)
	GetCifs(req models.ReqGetCif) (result []models.RespGetCif, err error)
	GetCifCount(req models.ReqGetCif) (result int, err error)

	GetSavingTypeCount(req models.ReqGetSavingType) (result int, err error)
	DropSavingType(id int, tx *sql.Tx) (err error)
	UpdateSavingType(req models.ReqGetSavingType, tx *sql.Tx) (err error)
	AddSavingType(req models.ReqGetSavingType, tx *sql.Tx) (result models.RespGetSavingType, err error)
	GetSavingType(req models.ReqGetSavingType) (result models.RespGetSavingType, err error)
	GetSavingTypes(req models.ReqGetSavingType) (result []models.RespGetSavingType, err error)

	GetSavingSegmentCount(req models.ReqGetSavingSegment) (result int, err error)
	DropSavingSegment(id int, tx *sql.Tx) (err error)
	UpdateSavingSegment(req models.ReqGetSavingSegment, tx *sql.Tx) (err error)
	AddSavingSegment(req models.ReqGetSavingSegment, tx *sql.Tx) (result models.RespGetSavingSegment, err error)
	GetSavingSegment(req models.ReqGetSavingSegment) (result models.RespGetSavingSegment, err error)
	GetSavingSegments(req models.ReqGetSavingSegment) (result []models.RespGetSavingSegment, err error)

	GetAccountCount(req models.ReqGetAccount) (result int, err error)
	DropAccount(id int, tx *sql.Tx) (err error)
	UpdateAccount(req models.ReqGetAccount, tx *sql.Tx) (err error)
	AddAccount(req models.ReqGetAccount, tx *sql.Tx) (result models.RespGetAccount, err error)
	GetAccount(req models.ReqGetAccount) (result models.RespGetAccount, err error)
	GetAccounts(req models.ReqGetAccount) (result []models.RespGetAccount, err error)

	GetSavingTransactionCount(req models.ReqGetSavingTransaction) (result int, err error)
	DropSavingTransaction(id int, tx *sql.Tx) (err error)
	UpdateSavingTransaction(req models.ReqGetSavingTransaction, tx *sql.Tx) (err error)
	AddSavingTransaction(req models.ReqGetSavingTransaction, tx *sql.Tx) (result models.RespGetSavingTransaction, err error)
	GetSavingTransaction(req models.ReqGetSavingTransaction) (result models.RespGetSavingTransaction, err error)
	GetSavingTransactions(req models.ReqGetSavingTransaction) (result []models.RespGetSavingTransaction, err error)
}

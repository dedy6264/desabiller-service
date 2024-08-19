package repositories

import (
	"database/sql"
	"desabiller/models"
)

type Hierarchyrepo interface {
	//client
	AddClient(req models.ReqGetClient, tx *sql.DB) (err error)
	DropClient(id int, tx *sql.DB) (err error)
	UpdateClient(req models.ReqGetClient, tx *sql.DB) (err error)
	GetCount(req models.ReqGetClient) (result int, err error)
	GetClients(req models.ReqGetClient) (result []models.RespGetClient, err error)
	GetClient(req models.ReqGetClient) (result models.RespGetClient, err error)

	DropGroup(id int) (err error)
	UpdateGroup(req models.ReqGetGroup) (result models.RespGetGroup, err error)
	GetGroupCount(req models.ReqGetGroup) (result int, err error)
	GetGroups(req models.ReqGetGroup) (result []models.RespGetGroup, err error)
	AddGroup(req models.ReqGetGroup, tx *sql.DB) (err error)
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

	AddProductClan(req models.ReqGetProductClan) (err error)
	GetProductClans(req models.ReqGetProductClan) (result []models.RespGetProductClan, err error)
	UpdateProductClan(req models.ReqGetProductClan) (result models.RespGetProductClan, err error)
	DropProductClan(req models.ReqGetProductClan) (err error)
	GetProductClanCount(req models.ReqGetProductClan) (result int, err error)
	GetProductClan(req models.ReqGetProductClan) (result models.RespGetProductClan, err error)

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
}
type PaymentRepo interface {
}
type TrxRepo interface {
}
type NoTrxGenerator interface {
	GetLastTrxNo() (noTrx string, status bool)
	InsertTrxNo(noTrx string) (id int, status bool)
	GenerateNo(datatype string, prefix string, leadingZero ...int) (code string, err error)
}

// type NHierarchy interface {
// 	NCreateClient(req models.ReqGetListNClient) (id int, err error)
// 	NReadClient(req models.ReqGetListNClient) (result []models.ResGetNClient, err error)
// 	NReadSingleClient(req models.ReqGetListNClient) (result models.ResGetNClient, err error)
// 	NDropClient(id int) (status bool, err error)
// 	NUpdateClient(req models.ReqUpdateNClient) (result models.ResGetNClient, err error)
// 	///merchant
// 	NCreateMerchant(req models.ReqGetListNMerchant) (id int, err error)
// 	NReadMerchant(req models.ReqGetListNMerchant) (result []models.ResGetNMerchant, err error)
// 	NReadSingleMerchant(req models.ReqGetListNMerchant) (result models.ResGetNMerchant, err error)
// 	NDropMerchant(id int) (status bool, err error)
// 	NUpdateMerchant(req models.ReqUpdateNMerchant) (result models.ResGetNMerchant, err error)
// 	//merchantOutlet
// 	NCreateMerchantOutlet(req models.ReqGetListNMerchantOutlet) (id int, err error)
// 	NReadMerchantOutlet(req models.ReqGetListNMerchantOutlet) (result []models.ResGetNMerchantOutlet, err error)
// 	NReadSingleMerchantOutlet(req models.ReqGetListNMerchantOutlet) (result models.ResGetNMerchantOutlet, err error)
// 	NDropMerchantOutlet(id int) (status bool, err error)
// 	NUpdateMerchantOutlet(req models.ReqUpdateNMerchantOutlet) (result models.ResGetNMerchantOutlet, err error)
// 	//userOutlet
// 	NCreateUserOutlet(req models.ReqGetListNUserOutlet) (id int, err error)
// 	NReadUserOutlet(req models.ReqGetListNUserOutlet) (result []models.ResGetNUserOutlet, err error)
// 	NReadSingleUserOutlet(req models.ReqGetListNUserOutlet) (result models.ResGetNUserOutlet, err error)
// 	NDropUserOutlet(id int) (status bool, err error)
// 	NUpdateUserOutlet(req models.ReqUpdateNUserOutlet) (result models.ResGetNUserOutlet, err error)
// }
// type NUserDashboard interface {
// 	NCreateUserDashboard(req models.ReqCreateNUserDashboard) (id int, err error)
// 	NReadUserDashboard(req models.ReqGetListNUserDashboard) (result []models.RespGetListNUserDashboard, err error)
// 	NReadSingleUserDashboard(req models.ReqCreateNUserDashboard) (result models.RespGetListNUserDashboard, err error)
// 	NDropUserDashboard(id int) (status bool, err error)
// 	NUpdateUserDashboard(req models.ReqCreateNUserDashboard) (result models.RespGetListNUserDashboard, err error)
// }
// type NFeatures interface {
// 	NCreateFeature(req models.ReqCreateNFeature) (id int, err error)
// 	NReadFeature(req models.ReqGetListNFeature) (result []models.RespGetListNFeature, err error)
// 	NReadSingleFeature(req models.ReqCreateNFeature) (result models.RespGetListNFeature, err error)
// 	NDropFeature(id int) (status bool, err error)
// 	NUpdateFeature(req models.ReqCreateNFeature) (result models.RespGetListNFeature, err error)
// 	//
// 	NCreateFeatureAssignment(req models.ReqCreateNFeatureAssignment) (id int, err error)
// 	NReadFeatureAssignment(req models.ReqGetListNFeatureAssignment) (result []models.RespGetListNFeatureAssignment, err error)
// 	NReadSingleFeatureAssignment(req models.ReqCreateNFeatureAssignment) (result models.RespGetListNFeatureAssignment, err error)
// 	NDropFeatureAssignment(id int) (status bool, err error)
// 	NUpdateFeatureAssignment(req models.ReqCreateNFeatureAssignment) (result models.RespGetListNFeatureAssignment, err error)
// }

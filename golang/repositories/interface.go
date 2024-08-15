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
	//merchant
	// AddMerchant(req models.ReqGetListMerchant) (status bool)
	// DropMerchant(req models.ReqGetListMerchant) (status bool)
	// UpdateMerchant(req models.ReqGetListMerchant) (result models.ResGetMerchant, status bool)
	// GetListMerchant(req models.ReqGetListMerchant) (result []models.ResGetMerchant, status bool)
	// GetListMerchantCount(req models.ReqGetListMerchant) (result int, status bool)
	// //merchantOutlet
	// DropMerchantOutlet(req models.ReqGetListMerchantOutlet) (status bool)
	// UpdateMerchantOutlet(req models.ReqGetListMerchantOutlet) (result models.ResGetMerchantOutlet, status bool)
	// GetListMerchantOutletCount(req models.ReqGetListMerchantOutlet) (result int, status bool)
	// GetListMerchantOutlet(req models.ReqGetListMerchantOutlet) (result []models.ResGetMerchantOutlet, status bool)
	// AddMerchantOutlet(req models.ReqGetListMerchantOutlet) (status bool)
	// //useroutlet
	// DropUserOutlet(req models.ReqGetListUserOutlet) (status bool)
	// UpdateUserOutlet(req models.ReqGetListUserOutlet) (result models.ResGetUserOutlet, status bool)
	// GetListUserOutletCount(req models.ReqGetListUserOutlet) (result int, status bool)
	// GetListUserOutlet(req models.ReqGetListUserOutlet) (result []models.ResGetUserOutlet, status bool)
	// AddUserOutlet(req models.ReqGetListUserOutlet) (status bool)
	// //outletDevice
	// DropOutletDevice(req models.ReqGetListOutletDevice) (status bool)
	// UpdateOutletDevice(req models.ReqGetListOutletDevice) (result models.ResGetOutletDevice, status bool)
	// GetListOutletDeviceCount(req models.ReqGetListOutletDevice) (result int, status bool)
	// GetListOutletDevice(req models.ReqGetListOutletDevice) (result []models.ResGetOutletDevice, status bool)
	// AddOutletDevice(req models.ReqGetListOutletDevice) (status bool)

	// GetListUser(req models.ReqUserList) (resp []models.RespUserList, err error, status bool)
	// GetHierarchy(mID int) (result models.RespHierarchy, status bool)
	// GetHierarchyByOutlet(oUID int) (result models.RespHierarchy, status bool)
}
type ProductRepo interface {
	//productType
	// GetListProductType() (result []models.ListProductType, err error)
	// //productCategory
	// GetListProductCategory(req models.ReqGetListProductCategory) (result []models.ResGetProductCategory, status bool)
	// AddProductCategory(req models.ReqGetListProductCategory) (result models.ResGetProductCategory, status bool)
	// UpdateProductCategory(req models.ReqGetListProductCategory) (result models.ResGetProductCategory, status bool)
	// DropProductCategory(req models.ReqGetListProductCategory) (status bool)
	// //productBillerProvider
	// AddProductBillerProvider(req models.ReqGetListProductBillerProvider) (result models.ResGetProductBillerProvider, status bool)
	// GetListProductBillerProvider(req models.ReqGetListProductBillerProvider) (result []models.ResGetProductBillerProvider, status bool)
	// UpdateProductBillerProvider(req models.ReqGetListProductBillerProvider) (result models.ResGetProductBillerProvider, status bool)
	// DropProductBillerProvider(req models.ReqGetListProductBillerProvider) (status bool)
	// //productBiller
	// AddProductBiller(req models.ReqGetListProductBiller) (result models.ResGetProductBiller, status bool)
	// GetListProductBiller(req models.ReqGetListProductBiller) (result []models.ResGetProductBiller, status bool)
	// UpdateProductBiller(req models.ReqGetListProductBiller) (result models.ResGetProductBiller, status bool)
	// DropProductBiller(req models.ReqGetListProductBiller) (status bool)
	// //productPos
	// AddProductPos(req models.ReqGetListProductPos) (result models.ResGetProductPos, status bool)
	// GetListProductPos(req models.ReqGetListProductPos) (result []models.ResGetProductPos, status bool)
	// UpdateProductPos(req models.ReqGetListProductPos) (result models.ResGetProductPos, status bool)
	// DropProductPos(req models.ReqGetListProductPos) (status bool)
	// GetListProductPosMany(productId []int, merchantId int) (result []models.ResGetProductPos, status bool)
	// //segment
	// AddSegment(req models.ReqListSegment) (result models.ResListSegment, status bool)
	// GetListSegment(req models.ReqListSegment) (result []models.ResListSegment, status bool)
	// UpdateSegment(req models.ReqListSegment) (result models.ResListSegment, status bool)
	// DropSegment(req models.ReqListSegment) (status bool)
	// // segmentproduct
	// AddSegmentProduct(req models.ReqListSegmentProduct) (result models.ResListSegmentProduct, status bool)
	// GetListSegmentProduct(req models.ReqListSegmentProduct) (result []models.ResListSegmentProduct, status bool)
	// UpdateSegmentProduct(req models.ReqListSegmentProduct) (result models.ResListSegmentProduct, status bool)
	// DropSegmentProduct(req models.ReqListSegmentProduct) (status bool)
}
type PaymentRepo interface {
	//payment
	// GetPaymentMethod(req models.ReqGetListPaymentMethod) (result models.ResPaymentMethod, status bool)
	// AddPaymentMethod(req models.ReqGetListPaymentMethod) (result models.ResPaymentMethod, status bool)
	// UpdatePaymentMethod(req models.ReqGetListPaymentMethod) (result models.ResPaymentMethod, status bool)
	// DropPaymentMethod(id int) (status bool)
	// GetListPaymentMethod(req models.ReqGetListPaymentMethod) (result []models.ResPaymentMethod, status bool)

	// //payment category
	// GetPaymentMethodCategory(req models.ReqGetListPaymentMethodCategory) (result models.ResPaymentMethodCategory, status bool)
	// AddPaymentMethodCategory(req models.ReqGetListPaymentMethodCategory) (result models.ResPaymentMethodCategory, status bool)
	// UpdatePaymentMethodCategory(req models.ReqGetListPaymentMethodCategory) (result models.ResPaymentMethodCategory, status bool)
	// DropPaymentMethodCategory(id int) (status bool)
	// GetListPaymentMethodCategory(req models.ReqGetListPaymentMethodCategory) (result []models.ResPaymentMethodCategory, status bool)
}
type TrxRepo interface {
	// GetTrxListPos(req models.ReqTrx, table string) (result []models.RespTrxList, status bool)
	// GetTrxPos(req models.ReqTrx, table string) (result models.RespTrxList, status bool)
	// InsertTrxPos(req models.ReqInsertTrx, table string, tx *sql.Tx) (id int, status bool)
	// UpdateTrxPos(req models.ReqUpdateTrx, table string) (status bool)
	// GetTrxListBiller(req models.ReqTrx, table string) (result []models.RespTrxList, status bool)
	// GetTrxBiller(req models.ReqTrx, table string) (result models.RespTrxList, status bool)
	// InsertTrxBiller(req models.ReqInsertTrx, table string) (id int, status bool)
	// UpdateTrxBiller(req models.ReqUpdateTrx, table string) (status bool)
	// InsertTrxDetails(req []models.ReqInsertTrxDetails, tx *sql.Tx) (status bool)
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

package models

type (
	RespNHierarchy struct {
		MerchantOutletId   int    `json:"merchantOutletId"`
		MerchantOutletName string `json:"merchantOutletName"`
		MerchantId         int    `json:"merchantId"`
		MerchantName       string `json:"merchantName"`
		ClientId           int    `json:"clientId"`
		ClientName         string `json:"clientName"`
	}
	//client
	ReqGetListNClient struct {
		ID         int    `json:"id"`
		ClientName string `json:"clientName"`
		Limit      int    `json:"limit"`
		Offset     int    `json:"offset"`
		OrderBy    string `json:"orderBy"`
		StartDate  string `json:"startDate"`
		EndDate    string `json:"endDate"`
		Username   string `json:"username"`
	}

	ReqUpdateNClient struct {
		ID         int    `json:"id"`
		ClientName string `json:"clientName"`
		CreatedAt  string `json:"createdAt"`
		UpdatedAt  string `json:"updatedAt"`
		CreatedBy  string `json:"createdBy"`
		UpdatedBy  string `json:"updatedBy"`
	}
	ResGetNClient struct {
		ID         int    `json:"id"`
		ClientName string `json:"clientName"`
		CreatedAt  string `json:"createdAt"`
		UpdatedAt  string `json:"updatedAt"`
		CreatedBy  string `json:"createdBy"`
		UpdatedBy  string `json:"updatedBy"`
	}
	//merchant
	ReqGetListNMerchant struct {
		ID           int    `json:"id"`
		MerchantName string `json:"merchantName"`
		ClientId     int    `json:"clientId"`
		ClientName   string `json:"clientName"`
		Limit        int    `json:"limit"`
		Offset       int    `json:"offset"`
		OrderBy      string `json:"orderBy"`
		StartDate    string `json:"startDate"`
		EndDate      string `json:"endDate"`
		Username     string `json:"username"`
	}
	ResGetNMerchant struct {
		ID           int    `json:"id"`
		MerchantName string `json:"merchantName"`
		ClientId     int    `json:"clientId"`
		ClientName   string `json:"clientName"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
		CreatedBy    string `json:"createdBy"`
		UpdatedBy    string `json:"updatedBy"`
	}
	//merchantOutlet
	ReqGetListNMerchantOutlet struct {
		ID                 int    `json:"id"`
		MerchantOutletName string `json:"merchantOutletName"`
		MerchantId         int    `json:"merchantId"`
		MerchantName       string `json:"merchantName"`
		ClientId           int    `json:"clientId"`
		Limit              int    `json:"limit"`
		Offset             int    `json:"offset"`
		OrderBy            string `json:"orderBy"`
		StartDate          string `json:"startDate"`
		EndDate            string `json:"endDate"`
		Username           string `json:"username"`
	}
	ResGetNMerchantOutlet struct {
		ID                 int    `json:"id"`
		MerchantOutletName string `json:"merchantOutletName"`
		MerchantId         int    `json:"merchantId"`
		MerchantName       string `json:"merchantName"`
		ClientId           int    `json:"clientId"`
		ClientName         string `json:"clientName"`
		CreatedAt          string `json:"createdAt"`
		UpdatedAt          string `json:"updatedAt"`
		CreatedBy          string `json:"createdBy"`
		UpdatedBy          string `json:"updatedBy"`
	}
	ReqUpdateNMerchant struct {
		ID           int    `json:"id"`
		ClientId     int    `json:"clientId"`
		ClientName   string `json:"clientName"`
		MerchantName string `json:"merchantName"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
		CreatedBy    string `json:"createdBy"`
		UpdatedBy    string `json:"updatedBy"`
	}
	//userOutlet
	ReqUpdateNMerchantOutlet struct {
		ID                 int    `json:"id"`
		MerchantId         int    `json:"merchantId"`
		MerchantName       string `json:"merchantName"`
		MerchantOutletName string `json:"merchantOutletName"`
		CreatedAt          string `json:"createdAt"`
		UpdatedAt          string `json:"updatedAt"`
		CreatedBy          string `json:"createdBy"`
		UpdatedBy          string `json:"updatedBy"`
	}
	ReqUpdateNUserOutlet struct {
		ID             int    `json:"id"`
		Nickname       string `json:"nickname"`
		OutletUsername string `json:"outletUsername"`
		OutletPassword string `json:"outletPassword"`

		MerchantOutletId   int    `json:"merchantOutletId"`
		MerchantOutletName string `json:"merchantOutletName"`
		MerchantId         int    `json:"merchantId"`
		MerchantName       string `json:"merchantName"`
		ClientId           int    `json:"clientId"`
		CreatedAt          string `json:"createdAt"`
		UpdatedAt          string `json:"updatedAt"`
		CreatedBy          string `json:"createdBy"`
		UpdatedBy          string `json:"updatedBy"`
	}
	ReqGetListNUserOutlet struct {
		ID                 int    `json:"id"`
		Nickname           string `json:"nickname"`
		OutletUsername     string `json:"outletUsername"`
		OutletPassword     string `json:"outletPassword"`
		MerchantOutletId   int    `json:"merchantOutletId"`
		MerchantOutletName string `json:"merchantOutletName"`
		MerchantId         int    `json:"merchantId"`
		MerchantName       string `json:"merchantName"`
		ClientId           int    `json:"clientId"`
		Limit              int    `json:"limit"`
		Offset             int    `json:"offset"`
		OrderBy            string `json:"orderBy"`
		StartDate          string `json:"startDate"`
		EndDate            string `json:"endDate"`
		Username           string `json:"username"`
	}
	ResGetNUserOutlet struct {
		ID                 int    `json:"id"`
		Nickname           string `json:"nickname"`
		OutletUsername     string `json:"outletUsername"`
		OutletPassword     string `json:"outletPassword"`
		MerchantOutletId   int    `json:"merchantOutletId"`
		MerchantOutletName string `json:"merchantOutletName"`
		MerchantId         int    `json:"merchantId"`
		MerchantName       string `json:"merchantName"`
		ClientId           int    `json:"clientId"`
		ClientName         string `json:"clientName"`
		CreatedAt          string `json:"createdAt"`
		UpdatedAt          string `json:"updatedAt"`
		CreatedBy          string `json:"createdBy"`
		UpdatedBy          string `json:"updatedBy"`
	}
	ResLoginNUserOutlet struct {
		ID                 int      `json:"id"`
		Nickname           string   `json:"nickname"`
		OutletUsername     string   `json:"outletUsername"`
		OutletPassword     string   `json:"outletPassword"`
		MerchantOutletId   int      `json:"merchantOutletId"`
		MerchantOutletName string   `json:"merchantOutletName"`
		MerchantId         int      `json:"merchantId"`
		MerchantName       string   `json:"merchantName"`
		ClientId           int      `json:"clientId"`
		ClientName         string   `json:"clientName"`
		Token              string   `json:"token"`
		DeviceType         string   `json:"deviceType"`
		DeviceSn           string   `json:"deviceSn"`
		Role               []string `json:"role"`
	}
	//OutletDevice
	// ReqGetListOutletDevice struct {
	// 	ID int `json:"id"`
	// 	// Nickname           string `json:"nickname"`
	// 	DeviceType         string `json:"deviceType"`
	// 	DeviceSn           string `json:"deviceSn"`
	// 	MerchantOutletId   int    `json:"merchantOutletId"`
	// 	MerchantOutletName string `json:"merchantOutletName"`
	// 	MerchantId         int    `json:"merchantId"`
	// 	MerchantName       string `json:"merchantName"`
	// 	ClientId           int    `json:"clientId"`
	// 	Limit              int    `json:"limit"`
	// 	Offset             int    `json:"offset"`
	// 	OrderBy            string `json:"orderBy"`
	// 	StartDate          string `json:"startDate"`
	// 	EndDate            string `json:"endDate"`
	// 	Username           string `json:"username"`
	// }
	// ResGetOutletDevice struct {
	// 	ID int `json:"id"`
	// 	// Nickname           string `json:"nickname"`
	// 	DeviceType         string `json:"deviceType"`
	// 	DeviceSn           string `json:"deviceSn"`
	// 	MerchantOutletId   string `json:"merchantOutletId"`
	// 	MerchantOutletName string `json:"merchantOutletName"`
	// 	MerchantId         int    `json:"merchantId"`
	// 	MerchantName       string `json:"merchantName"`
	// 	ClientId           int    `json:"clientId"`
	// 	ClientName         string `json:"clientName"`
	// 	CreatedAt          string `json:"createdAt"`
	// 	UpdatedAt          string `json:"updatedAt"`
	// 	CreatedBy          string `json:"createdBy"`
	// 	UpdatedBy          string `json:"updatedBy"`
	// }
)
type (
	ReqLoginUserDashboard struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	RespLoginUserDashboard struct {
		ID               int    `json:"id"`
		Username         string `json:"username"`
		Email            string `json:"email"`
		Password         string `json:"password"`
		Role             string `json:"role"`
		Token            string `json:"token"`
		ClientId         int    `json:"clientId"`
		MerchantId       int    `json:"merchantId"`
		MerchantOutletId int    `json:"merchantOutletId"`
	}
	ReqCreateNUserDashboard struct {
		ID               int    `json:"id"`
		Username         string `json:"username"`
		Email            string `json:"email"`
		Password         string `json:"password"`
		Role             string `json:"role"`
		ClientId         int    `json:"clientId"`
		MerchantId       int    `json:"merchantId"`
		MerchantOutletId int    `json:"merchantOutletId"`
		CreatedAt        string `json:"createdAt"`
		UpdatedAt        string `json:"updatedAt"`
		CreatedBy        string `json:"createdBy"`
		UpdatedBy        string `json:"updatedBy"`
	}
	ReqGetListNUserDashboard struct {
		Draw      int                     `json:"draw"`
		Limit     int                     `json:"limit"`
		Offset    int                     `json:"offset"`
		OrderBy   string                  `json:"orderBy"`
		AscDesc   string                  `json:"ascDesc"`
		StartDate string                  `json:"startDate"`
		EndDate   string                  `json:"endDate"`
		Data      ReqCreateNUserDashboard `json:"data"`
	}

	RespGetListNUserDashboard struct {
		ID               int    `json:"id"`
		Username         string `json:"username"`
		Email            string `json:"email"`
		Password         string `json:"password"`
		Role             string `json:"role"`
		ClientId         int    `json:"clientId"`
		MerchantId       int    `json:"merchantId"`
		MerchantOutletId int    `json:"merchantOutletId"`
		CreatedAt        string `json:"createdAt"`
		UpdatedAt        string `json:"updatedAt"`
		CreatedBy        string `json:"createdBy"`
		UpdatedBy        string `json:"updatedBy"`
	}
	ReqCreateNFeature struct {
		ID          int    `json:"id"`
		FeatureName string `json:"featureName"`
		MerchantId  int    `json:"merchantId"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
		CreatedBy   string `json:"createdBy"`
		UpdatedBy   string `json:"updatedBy"`
	}
	ReqGetListNFeature struct {
		Draw      int               `json:"draw"`
		Limit     int               `json:"limit"`
		Offset    int               `json:"offset"`
		OrderBy   string            `json:"orderBy"`
		AscDesc   string            `json:"ascDesc"`
		StartDate string            `json:"startDate"`
		EndDate   string            `json:"endDate"`
		Data      ReqCreateNFeature `json:"data"`
	}
	RespGetListNFeature struct {
		ID          int    `json:"id"`
		FeatureName string `json:"featureName"`
		MerchantId  int    `json:"merchantId"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
		CreatedBy   string `json:"createdBy"`
		UpdatedBy   string `json:"updatedBy"`
	}
	ReqCreateNFeatureAssignment struct {
		ID         int    `json:"id"`
		FeatureId  int    `json:"featureId"`
		MerchantId int    `json:"merchantId"`
		CreatedAt  string `json:"createdAt"`
		UpdatedAt  string `json:"updatedAt"`
		CreatedBy  string `json:"createdBy"`
		UpdatedBy  string `json:"updatedBy"`
		Index      int    `json:"index"`
	}
	ReqUpdateNFeatureAssignment struct {
		ID         int    `json:"id"`
		FeatureId  int    `json:"featureId"`
		MerchantId int    `json:"merchantId"`
		CreatedAt  string `json:"createdAt"`
		UpdatedAt  string `json:"updatedAt"`
		CreatedBy  string `json:"createdBy"`
		UpdatedBy  string `json:"updatedBy"`
	}
	ReqGetListNFeatureAssignment struct {
		Draw      int                         `json:"draw"`
		Limit     int                         `json:"limit"`
		Offset    int                         `json:"offset"`
		OrderBy   string                      `json:"orderBy"`
		AscDesc   string                      `json:"ascDesc"`
		StartDate string                      `json:"startDate"`
		EndDate   string                      `json:"endDate"`
		Data      ReqCreateNFeatureAssignment `json:"data"`
	}
	RespGetListNFeatureAssignment struct {
		ID           int    `json:"id"`
		FeatureId    int    `json:"featureId"`
		FeatureName  string `json:"featureName"`
		MerchantId   int    `json:"merchantId"`
		MerchantName string `json:"merchantName"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
		CreatedBy    string `json:"createdBy"`
		UpdatedBy    string `json:"updatedBy"`
	}
)

package models

type (
	RespHierarchy struct {
		MerchantOutletId   int    `json:"merchantOutletId"`
		MerchantOutletName string `json:"merchantOutletName"`
		MerchantId         int    `json:"merchantId"`
		MerchantName       string `json:"merchantName"`
		ClientId           int    `json:"clientId"`
		ClientName         string `json:"clientName"`
	}
	//client
	ReqGetListClient struct {
		ID         int    `json:"id"`
		ClientName string `json:"clientName"`
		Limit      int    `json:"limit"`
		Offset     int    `json:"offset"`
		OrderBy    string `json:"orderBy"`
		StartDate  string `json:"startDate"`
		EndDate    string `json:"endDate"`
		Username   string `json:"username"`
	}

	ResGetClient struct {
		ID         int    `json:"id"`
		ClientName string `json:"clientName"`
		CreatedAt  string `json:"createdAt"`
		UpdatedAt  string `json:"updatedAt"`
		CreatedBy  string `json:"createdBy"`
		UpdatedBy  string `json:"updatedBy"`
	}
	//merchant
	ReqGetListMerchant struct {
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
	ResGetMerchant struct {
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
	ReqGetListMerchantOutlet struct {
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
	ResGetMerchantOutlet struct {
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
	//userOutlet
	ReqGetListUserOutlet struct {
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
	ResGetUserOutlet struct {
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
	ResLoginUserOutlet struct {
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
	ReqGetListOutletDevice struct {
		ID int `json:"id"`
		// Nickname           string `json:"nickname"`
		DeviceType         string `json:"deviceType"`
		DeviceSn           string `json:"deviceSn"`
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
	ResGetOutletDevice struct {
		ID int `json:"id"`
		// Nickname           string `json:"nickname"`
		DeviceType         string `json:"deviceType"`
		DeviceSn           string `json:"deviceSn"`
		MerchantOutletId   string `json:"merchantOutletId"`
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
)

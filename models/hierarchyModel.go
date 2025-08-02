package models

type (
	FilterReq struct {
		// Limit     int    `json:"limit"`
		// Offset    int    `json:"offset"`
		AscDesc   string `json:"ascDesc"`
		OrderBy   string `json:"orderBy"`
		CreatedAt string `json:"createdAt"`
		CreatedBy string `json:"createdBy"`
		UpdatedAt string `json:"updatedAt"`
		UpdatedBy string `json:"updatedBy"`
		Start     int    `json:"start"`
		Length    int    `json:"length"`
		Draw      int    `json:"draw"`
		Order     string `json:"order"`
		Search    string `json:"search"`
	}
	ReqGetClient struct {
		ID         int       `json:"id"`
		ClientName string    `json:"clientName"`
		Filter     FilterReq `json:"filter"`
	}
	RespGetClient struct {
		ID         int    `json:"id"`
		ClientName string `json:"clientName"`
		CreatedAt  string `json:"createdAt"`
		UpdatedAt  string `json:"updatedAt"`
		CreatedBy  string `json:"createdBy"`
		UpdatedBy  string `json:"updatedBy"`
	}
	RespAttribute struct {
		DataCount     int             `json:"dataCount"`
		DataSummaries int             `json:"dataSummaries"`
		Data          []RespGetClient `json:"data"`
	}
	// group
	ReqGetGroup struct {
		ID         int       `json:"id"`
		GroupName  string    `json:"groupName"`
		ClientId   int       `json:"clientId"`
		ClientName string    `json:"clientName"`
		Filter     FilterReq `json:"filter"`
	}
	RespGetGroup struct {
		ID         int    `json:"id"`
		GroupName  string `json:"groupName"`
		ClientId   int    `json:"clientId"`
		ClientName string `json:"clientName"`
		CreatedAt  string `json:"createdAt"`
		UpdatedAt  string `json:"updatedAt"`
		CreatedBy  string `json:"createdBy"`
		UpdatedBy  string `json:"updatedBy"`
	}
	//merchant
	ReqGetMerchant struct {
		ID           int       `json:"id"`
		MerchantName string    `json:"merchantName"`
		GroupId      int       `json:"groupId"`
		GroupName    string    `json:"groupName"`
		ClientId     int       `json:"clientId"`
		ClientName   string    `json:"clientName"`
		Filter       FilterReq `json:"filter"`
	}
	RespGetMerchant struct {
		ID           int    `json:"id"`
		MerchantName string `json:"merchantName"`
		GroupId      int    `json:"groupId"`
		GroupName    string `json:"groupName"`
		ClientId     int    `json:"clientId"`
		ClientName   string `json:"clientName"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
		CreatedBy    string `json:"createdBy"`
		UpdatedBy    string `json:"updatedBy"`
	}
	//merchantOutlet
	ReqGetMerchantOutlet struct {
		ID                     int       `json:"id"`
		MerchantOutletName     string    `json:"merchantOutletName"`
		MerchantOutletUsername string    `json:"merchantOutletUsername"`
		MerchantOutletPassword string    `json:"merchantOutletPassword"`
		MerchantId             int       `json:"merchantId"`
		MerchantName           string    `json:"merchantName"`
		GroupId                int       `json:"groupId"`
		GroupName              string    `json:"groupName"`
		ClientId               int       `json:"clientId"`
		ClientName             string    `json:"clientName"`
		Filter                 FilterReq `json:"filter"`
	}
	RespGetMerchantOutlet struct {
		ID                     int    `json:"id"`
		MerchantOutletName     string `json:"merchantOutletName"`
		MerchantOutletUsername string `json:"merchantOutletUsername"`
		MerchantOutletPassword string `json:"merchantOutletPassword"`
		MerchantId             int    `json:"merchantId"`
		MerchantName           string `json:"merchantName"`
		GroupId                int    `json:"groupId"`
		GroupName              string `json:"groupName"`
		ClientId               int    `json:"clientId"`
		ClientName             string `json:"clientName"`
		CreatedAt              string `json:"createdAt"`
		UpdatedAt              string `json:"updatedAt"`
		CreatedBy              string `json:"createdBy"`
		UpdatedBy              string `json:"updatedBy"`
	}
	DataToken struct {
		// MerchantOutletId       int    `json:"merchantOutletId"`
		// MerchantOutletUsername string `json:"merchantOutletUsername"`
		// MerchantOutletName     string `json:"merchantOutletName"`
		UserAppId int `json:"userAppId"`
	}
	// //userOutlet
	// ReqGetListUserOutlet struct {
	// 	ID                 int    `json:"id"`
	// 	Nickname           string `json:"nickname"`
	// 	OutletUsername     string `json:"outletUsername"`
	// 	OutletPassword     string `json:"outletPassword"`
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
	// ResGetUserOutlet struct {
	// 	ID                 int    `json:"id"`
	// 	Nickname           string `json:"nickname"`
	// 	OutletUsername     string `json:"outletUsername"`
	// 	OutletPassword     string `json:"outletPassword"`
	// 	MerchantOutletId   int    `json:"merchantOutletId"`
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
	// ResLoginUserOutlet struct {
	// 	ID                 int      `json:"id"`
	// 	Nickname           string   `json:"nickname"`
	// 	OutletUsername     string   `json:"outletUsername"`
	// 	OutletPassword     string   `json:"outletPassword"`
	// 	MerchantOutletId   int      `json:"merchantOutletId"`
	// 	MerchantOutletName string   `json:"merchantOutletName"`
	// 	MerchantId         int      `json:"merchantId"`
	// 	MerchantName       string   `json:"merchantName"`
	// 	ClientId           int      `json:"clientId"`
	// 	ClientName         string   `json:"clientName"`
	// 	Token              string   `json:"token"`
	// 	DeviceType         string   `json:"deviceType"`
	// 	DeviceSn           string   `json:"deviceSn"`
	// 	Role               []string `json:"role"`
	// }
	// //OutletDevice
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

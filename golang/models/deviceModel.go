package models

type (
	ReqLogin struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		DeviceSn string `json:"deviceSn" `
	}
	ResLogin struct {
		UserInfo  UserInfo `json:"userInfo"`
		TokenInfo string   `json:"tokenInfo"`
	}
	UserInfo struct {
		Username           string `json:"username"`
		DeviceSn           string `json:"deviceSn"`
		Nickname           string `json:"nickname"`
		MerchantOutletId   string `json:"merchantOutletId"`
		MerchantOutletName string `json:"merchantOutletName"`
		MerchantId         int    `json:"merchantId"`
		MerchantName       string `json:"merchantName"`
		ClientId           int    `json:"clientId"`
		ClientName         string `json:"clientName"`
	}
)

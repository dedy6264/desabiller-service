package models

type (
	ReqAviceTrx struct {
		ReferenceNumber string `json:"referenceNumber" validate:"required"`
	}
	ReqGetTrx struct {
		Id                         int       `json:"id"` //
		ProductClanId              int       `json:"productClanId"`
		ProductClanName            string    `json:"productClanName"`
		ProductCategoryId          int       `json:"productCategoryId"` //
		ProductCategoryName        string    `json:"productCategoryName"`
		ProductTypeId              int       `json:"productTypeId"`
		ProductTypeName            string    `json:"productTypeName"`
		ProductId                  int       `json:"productId"`
		ProductName                string    `json:"productName"` //
		ProductCode                string    `json:"productCode"`
		ProductPrice               float64   `json:"productPrice"`
		ProductAdminFee            float64   `json:"productAdminFee"`
		ProductMerchantFee         float64   `json:"productMerchantFee"`
		ProviderId                 int       `json:"providerId"`
		ProviderName               string    `json:"providerName"`
		ProductProviderId          int       `json:"productProviderId"`
		ProductProviderName        string    `json:"productProviderName"`
		ProductProviderCode        string    `json:"productProviderCode"`
		ProductProviderPrice       float64   `json:"productProviderPrice"`
		ProductProviderAdminFee    float64   `json:"productProviderAdminFee"`
		ProductProviderMerchantFee float64   `json:"productProviderMerchantFee"`
		StatusCode                 string    `json:"statusCode"` //
		StatusMessage              string    `json:"statusMessage"`
		StatusDesc                 string    `json:"statusDesc"`
		ReferenceNumber            string    `json:"referenceNumber"` //
		ProviderStatusCode         string    `json:"providerStatusCode"`
		ProviderStatusMessage      string    `json:"providerStatusMessage"`
		ProviderStatusDesc         string    `json:"providerStatusDesc"`
		ProviderReferenceNumber    string    `json:"providerReferenceNumber"`
		ClientId                   int       `json:"clientId"` //
		ClientName                 string    `json:"clientName"`
		GroupId                    int       `json:"groupId"` //
		GroupName                  string    `json:"groupName"`
		MerchantId                 int       `json:"merchantId"` //
		MerchantName               string    `json:"merchantName"`
		MerchantOutletId           int       `json:"merchantOutletId"` //
		MerchantOutletName         string    `json:"merchantOutletName"`
		MerchantOutletUsername     string    `json:"merchantOutletUsername"`
		CustomerId                 string    `json:"customerId"` //
		OtherMsg                   string    `json:"otherMsg"`
		TotalTrxAmount             float64   `json:"totalTrxAmount"`
		Filter                     FilterReq `json:"filter"`
	}
	RespGetTrx struct {
		Index                      int     `json:"index"`
		Id                         int     `json:"id"`
		ProductClanId              int     `json:"productClanId"`
		ProductClanName            string  `json:"productClanName"`
		ProductCategoryId          int     `json:"productCategoryId"`
		ProductCategoryName        string  `json:"productCategoryName"`
		ProductTypeId              int     `json:"productTypeId"`
		ProductTypeName            string  `json:"productTypeName"`
		ProductId                  int     `json:"productId"`
		ProductName                string  `json:"productName"`
		ProductCode                string  `json:"productCode"`
		ProductPrice               float64 `json:"productPrice"`
		ProductAdminFee            float64 `json:"productAdminFee"`
		ProductMerchantFee         float64 `json:"productMerchantFee"`
		ProductProviderId          int     `json:"productProviderId"`
		ProductProviderName        string  `json:"productProviderName"`
		ProductProviderCode        string  `json:"productProviderCode"`
		ProductProviderPrice       float64 `json:"productProviderPrice"`
		ProductProviderAdminFee    float64 `json:"productProviderAdminFee"`
		ProductProviderMerchantFee float64 `json:"productProviderMerchantFee"`
		ProviderId                 int     `json:"providerId"`
		ProviderName               string  `json:"providerName"`
		StatusCode                 string  `json:"statusCode"`
		StatusMessage              string  `json:"statusMessage"`
		StatusDesc                 string  `json:"statusDesc"`
		ReferenceNumber            string  `json:"referenceNumber"`
		ProviderStatusCode         string  `json:"providerStatusCode"`
		ProviderStatusMessage      string  `json:"providerStatusMessage"`
		ProviderStatusDesc         string  `json:"providerStatusDesc"`
		ProviderReferenceNumber    string  `json:"providerReferenceNumber"`
		ClientId                   int     `json:"clientId"`
		ClientName                 string  `json:"clientName"`
		GroupId                    int     `json:"groupId"`
		GroupName                  string  `json:"groupName"`
		MerchantId                 int     `json:"merchantId"`
		MerchantName               string  `json:"merchantName"`
		MerchantOutletId           int     `json:"merchantOutletId"`
		MerchantOutletName         string  `json:"merchantOutletName"`
		MerchantOutletUsername     string  `json:"merchantOutletUsername"`
		CustomerId                 string  `json:"customerId"`
		OtherMsg                   string  `json:"otherMsg"`
		TotalTrxAmount             float64 `json:"totalTrxAmount"`
		CreatedAt                  string  `json:"createdAt"`
		CreatedBy                  string  `json:"createdBy"`
		UpdatedAt                  string  `json:"updatedAt"`
		UpdatedBy                  string  `json:"updatedBy"`
	}
	ReqGetTrxStatus struct {
		// ClientId                int       `json:"clientId"`
		ReferenceNumber         string    `json:"referenceNumber"`
		ProviderReferenceNumber string    `json:"providerReferenceNumber"`
		StatusCode              string    `json:"statusCode"`
		StatusMessage           string    `json:"statusMessage"`
		Filter                  FilterReq `json:"filter"`
	}
	RespGetTrxStatus struct {
		// ClientId                int    `json:"clientId"`
		ReferenceNumber         string `json:"referenceNumber"`
		ProviderReferenceNumber string `json:"providerReferenceNumber"`
		StatusCode              string `json:"statusCode"`
		StatusMessage           string `json:"statusMessage"`
		CreatedAt               string `json:"createdAt"`
		CreatedBy               string `json:"createdBy"`
		UpdatedAt               string `json:"updatedAt"`
		UpdatedBy               string `json:"updatedBy"`
	}
	ReqPaymentTrx struct {
		PaymentMethodId         string `json:"paymentMethodId"`
		PaymentMethodName       string `json:"paymentMethodName"`
		ReferenceNumber         string `json:"referenceNumber"`
		ProviderReferenceNumber string `json:"providerReferenceNumber"`
	}
	RespPayment struct { //jelasin produknya
		Id                     int         `json:"id"`
		CustomerId             string      `json:"customerId"`
		ProductName            string      `json:"productName"`
		ProductCode            string      `json:"productCode"`
		ProductPrice           float64     `json:"productPrice"`
		ProductAdminFee        float64     `json:"productAdminFee"`
		ProductMerchantFee     float64     `json:"productMerchantFee"`
		TotalTrxAmount         float64     `json:"totalTrxAmount"`
		MerchantOutletId       int         `json:"merchantOutletId"`
		MerchantOutletName     string      `json:"merchantOutletName"`
		MerchantOutletUsername string      `json:"merchantOutletUsername"`
		ReferenceNumber        string      `json:"referenceNumber"`
		CreatedAt              string      `json:"createdAt"`
		BillInfo               interface{} `json:"billInfo"` //jelasin tagihan/detil produk
		// ProductId          int     `json:"productId"`
		// StatusCode      string `json:"statusCode"`
		// StatusMessage   string `json:"statusMessage"`
		// StatusDesc      string `json:"statusDesc"`
		// ProviderStatusCode      string `json:"providerStatusCode"`
		// ProviderStatusMessage   string `json:"providerStatusMessage"`
		// ProviderStatusDesc      string `json:"providerStatusDesc"`
		// ProviderReferenceNumber string `json:"providerReferenceNumber"`
		// UpdatedAt string `json:"updatedAt"`

		// ClientId               int    `json:"clientId"`
		// ClientName             string `json:"clientName"`
		// GroupId                int    `json:"groupId"`
		// GroupName              string `json:"groupName"`
		// MerchantId             int    `json:"merchantId"`
		// MerchantName           string `json:"merchantName"`

	}
	RespInquiry struct {
		Id                     int         `json:"id"`
		CustomerId             string      `json:"customerId"`
		ProductName            string      `json:"productName"`
		ProductCode            string      `json:"productCode"`
		ProductPrice           float64     `json:"productPrice"`
		ProductAdminFee        float64     `json:"productAdminFee"`
		ProductMerchantFee     float64     `json:"productMerchantFee"`
		TotalTrxAmount         float64     `json:"totalTrxAmount"`
		MerchantOutletId       int         `json:"merchantOutletId"`
		MerchantOutletName     string      `json:"merchantOutletName"`
		MerchantOutletUsername string      `json:"merchantOutletUsername"`
		ReferenceNumber        string      `json:"referenceNumber"`
		CreatedAt              string      `json:"createdAt"`
		BillInfo               interface{} `json:"billInfo"` //jelasin tagihan/detil produk
	}
)

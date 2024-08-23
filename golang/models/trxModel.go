package models

type (
	ReqGetTrx struct {
		Id                         int       `json:"id"` //
		ProductClanId              int       `json:"productClan_id"`
		ProductClanName            string    `json:"productClanName"`
		ProductCategoryId          int       `json:"productCategory_id"` //
		ProductCategoryName        string    `json:"productCategoryName"`
		ProductTypeId              int       `json:"productType_id"`
		ProductTypeName            string    `json:"productTypeName"`
		ProductId                  int       `json:"product_id"`
		ProductName                string    `json:"productName"` //
		ProductCode                string    `json:"productCode"`
		ProductPrice               float64   `json:"productPrice"`
		ProductAdminFee            float64   `json:"productAdminFee"`
		ProductMerchantFee         float64   `json:"productMerchantFee"`
		ProductProviderId          int       `json:"productProvider_id"`
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
		ClientId                   int       `json:"client_id"` //
		ClientName                 string    `json:"clientName"`
		GroupId                    int       `json:"group_id"` //
		GroupName                  string    `json:"groupName"`
		MerchantId                 int       `json:"merchant_id"` //
		MerchantName               string    `json:"merchantName"`
		MerchantOutletId           int       `json:"merchantOutlet_id"` //
		MerchantOutletName         string    `json:"merchantOutletName"`
		MerchantOutletUsername     string    `json:"merchantOutletUsername"`
		CustomerId                 string    `json:"customer_id"` //
		OtherMsg                   string    `json:"otherMsg"`
		Filter                     FilterReq `json:"filter"`
	}
	RespGetTrx struct {
		Index                      int     `json:"index"`
		Id                         int     `json:"id"`
		ProductClanId              int     `json:"productClan_id"`
		ProductClanName            string  `json:"productClanName"`
		ProductCategoryId          int     `json:"productCategory_id"`
		ProductCategoryName        string  `json:"productCategoryName"`
		ProductTypeId              int     `json:"productType_id"`
		ProductTypeName            string  `json:"productTypeName"`
		ProductId                  int     `json:"product_id"`
		ProductName                string  `json:"productName"`
		ProductCode                string  `json:"productCode"`
		ProductPrice               float64 `json:"productPrice"`
		ProductAdminFee            float64 `json:"productAdminFee"`
		ProductMerchantFee         float64 `json:"productMerchantFee"`
		ProductProviderId          int     `json:"productProvider_id"`
		ProductProviderName        string  `json:"productProviderName"`
		ProductProviderCode        string  `json:"productProviderCode"`
		ProductProviderPrice       float64 `json:"productProviderPrice"`
		ProductProviderAdminFee    float64 `json:"productProviderAdminFee"`
		ProductProviderMerchantFee float64 `json:"productProviderMerchantFee"`
		StatusCode                 string  `json:"statusCode"`
		StatusMessage              string  `json:"statusMessage"`
		StatusDesc                 string  `json:"statusDesc"`
		ReferenceNumber            string  `json:"referenceNumber"`
		ProviderStatusCode         string  `json:"providerStatusCode"`
		ProviderStatusMessage      string  `json:"providerStatusMessage"`
		ProviderStatusDesc         string  `json:"providerStatusDesc"`
		ProviderReferenceNumber    string  `json:"providerReferenceNumber"`
		ClientId                   int     `json:"client_id"`
		ClientName                 string  `json:"clientName"`
		GroupId                    int     `json:"group_id"`
		GroupName                  string  `json:"groupName"`
		MerchantId                 int     `json:"merchant_id"`
		MerchantName               string  `json:"merchantName"`
		MerchantOutletId           int     `json:"merchantOutlet_id"`
		MerchantOutletName         string  `json:"merchantOutletName"`
		MerchantOutletUsername     string  `json:"merchantOutletUsername"`
		CustomerId                 string  `json:"customer_id"`
		OtherMsg                   string  `json:"otherMsg"`
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
)

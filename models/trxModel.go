package models

type (
	ProviderInqRequest struct {
		ProviderId           int    `json:"providerId"`
		ProviderName         string `json:"providerName"`
		ProductReferenceCode string `json:"productReferenceCode"`
		ProductCode          string `json:"productCode"`
		SubscriberNumber     string `json:"subscriberNumber"`
		SubscriberName       string `json:"subscriberName"`
		ReferenceNumber      string `json:"referenceNumber"`
		Url                  string `json:"url"`
		Periode              int    `json:"periode"`
	}
	ProviderPayRequest struct {
		ProviderName            string `json:"providerName"`
		ProviderReferenceNumber string `json:"ProviderReferenceNumber"`
		Url                     string `json:"url"`
		ProductReferenceCode    string `json:"productReferenceCode"`
	}
	ReqAviceTrx struct {
		ReferenceNumber string `json:"referenceNumber" validate:"required"`
	}
	ReqGetTrx struct {
		Id                         int       `json:"id"` //
		ProductReferenceId         int       `json:"productReferenceId"`
		ProductReferenceCode       string    `json:"productReferenceCode"`
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
		ProductReferenceId         int     `json:"productReferenceId"`
		ProductReferenceCode       string  `json:"productReferenceCode"`
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
		AccountNumber           string `json:"accountNumber"`
		AccountPIN              string `json:"accounrPin"`
	}
	RespPayment struct { //jelasin produknya
		CreatedAt              string      `json:"createdAt"`
		MerchantOutletName     string      `json:"merchantOutletName"`
		MerchantOutletUsername string      `json:"merchantOutletUsername"`
		ReferenceNumber        string      `json:"referenceNumber"`
		ProductName            string      `json:"productName"`
		ProductCode            string      `json:"productCode"`
		SubscriberNumber       string      `json:"subscriberNumber"`
		ProductPrice           float64     `json:"productPrice"`
		ProductAdminFee        float64     `json:"productAdminFee"`
		ProductMerchantFee     float64     `json:"productMerchantFee"`
		TotalTrxAmount         float64     `json:"totalTrxAmount"`
		BillInfo               interface{} `json:"billInfo"` //jelasin tagihan/detil produk

	}
	RespInquiry struct {
		CreatedAt              string      `json:"createdAt"`
		MerchantOutletName     string      `json:"merchantOutletName"`
		MerchantOutletUsername string      `json:"merchantOutletUsername"`
		ReferenceNumber        string      `json:"referenceNumber"`
		ProductName            string      `json:"productName"`
		ProductCode            string      `json:"productCode"`
		SubscriberNumber       string      `json:"subscriberNumber"`
		ProductPrice           float64     `json:"productPrice"`
		ProductAdminFee        float64     `json:"productAdminFee"`
		ProductMerchantFee     float64     `json:"productMerchantFee"`
		TotalTrxAmount         float64     `json:"totalTrxAmount"`
		BillInfo               interface{} `json:"billInfo"` //jelasin tagihan/detil produk
	}
	ReqInquiry struct {
		ProductCode     string          `json:"productCode"`
		AdditionalField AdditionalField `json:"additionalField"`
	}
	AdditionalField struct {
		Periode          int    `json:"periode"`
		SubscriberNumber string `json:"SubscriberNumber"`
		SubscriberName   string `json:"SubscriberName"`
	}

	BillInfoBPJS struct {
		BillDesc struct {
			CustomerID   string `json:"customerId"`
			CustomerName string `json:"customerName"`
			Detail       []struct {
				Periode    string `json:"periode"`
				Admin      int    `json:"admin"`
				Denda      int    `json:"denda"`
				Tagihan    int    `json:"tagihan"`
				JmlPeserta string `json:"jmlPeserta"`
			} `json:"detail"`
		} `json:"billDesc"`
		Sn string `json:"sn"`
	}
	// BillInfoBPJS struct {
	// 	BillDesc struct {
	// 		CustomerID   string `json:"customerId"`
	// 		CustomerName string `json:"customerName"`
	// 		Detail       []struct {
	// 			Periode    string `json:"periode"`
	// 			Admin      int    `json:"admin"`
	// 			Denda      int    `json:"denda"`
	// 			Tagihan    int    `json:"tagihan"`
	// 			JmlPeserta string `json:"jmlPeserta"`
	// 		} `json:"detail"`
	// 	} `json:"billDesc"`
	// 	Sn string `json:"sn"`
	// }
)

package models

type (
	ReqGetTransaction struct {
		Start     int64       `json:"start" `
		Lenght    int64       `json:"length"`
		Columns   string      `json:"columns"`
		Search    string      `json:"search"`
		Order     string      `json:"order" `
		Sort      string      `json:"sort" `
		StartDate string      `json:"startDate"`
		EndDate   string      `json:"endDate"`
		Draw      int         `json:"draw"`
		Filter    Transaction `json:"filter"`
	}
	Transaction struct {
		ID                         int64   `json:"id"`
		ProductProviderName        string  `json:"productProviderName"`
		ProductProviderCode        string  `json:"productProviderCode"`
		ProductProviderPrice       float64 `json:"productProviderPrice"`
		ProductProviderAdminFee    float64 `json:"productProviderAdminFee"`
		ProductProviderMerchantFee float64 `json:"productProviderMerchantFee"`
		ProductID                  int64   `json:"productId"`
		ProductName                string  `json:"productName"`
		ProductCode                string  `json:"productCode"`
		ProductPrice               float64 `json:"productPrice"`
		ProductAdminFee            float64 `json:"productAdminFee"`
		ProductMerchantFee         float64 `json:"productMerchantFee"`
		ProductCategoryID          int64   `json:"productCategoryId"`
		ProductCategoryName        string  `json:"productCategoryName"`
		ProductTypeID              int64   `json:"productTypeId"`
		ProductTypeName            string  `json:"productTypeName"`
		ReferenceNumber            string  `json:"referenceNumber"`
		ProviderReferenceNumber    string  `json:"providerReferenceNumber"`
		StatusCode                 string  `json:"statusCode"`
		StatusMessage              string  `json:"statusMessage"`
		StatusDesc                 string  `json:"statusDesc"`
		StatusCodeDetail           string  `json:"statusCodeDetail"`
		StatusMessageDetail        string  `json:"statusMessageDetail"`
		StatusDescDetail           string  `json:"statusDescDetail"`
		ProductReferenceID         int64   `json:"productReferenceId"`
		ProductReferenceCode       string  `json:"productReferenceCode"`
		CustomerID                 string  `json:"customerId"`
		OtherReff                  string  `json:"otherReff"`
		OtherCustomerInfo          string  `json:"otherCustomerInfo"`
		SavingAccountName          string  `json:"savingAccountName"`
		SavingAccountID            int64   `json:"savingAccountId"`
		SavingAccountNumber        string  `json:"savingAccountNumber"`
		TransactionTotalAmount     float64 `json:"transactionTotalAmount"`
		UserAppID                  int64   `json:"userAppId"`
		Username                   string  `json:"username"`
		CreatedAt                  string  `json:"createdAt"`
		CreatedBy                  string  `json:"createdBy"`
		UpdatedAt                  string  `json:"updatedAt"`
		UpdatedBy                  string  `json:"updatedBy"`
	}
	RespGetTrx struct {
		Index                      int     `json:"index"`
		Id                         int     `json:"id"`
		ProductProviderName        string  `json:"productProviderName"`
		ProductProviderCode        string  `json:"productProviderCode"`
		ProductProviderPrice       float64 `json:"productProviderPrice"`
		ProductProviderAdminFee    float64 `json:"productProviderAdminFee"`
		ProductProviderMerchantFee float64 `json:"productProviderMerchantFee"`
		ProductID                  int64   `json:"productId"`
		ProductName                string  `json:"productName"`
		ProductCode                string  `json:"productCode"`
		ProductPrice               float64 `json:"productPrice"`
		ProductAdminFee            float64 `json:"productAdminFee"`
		ProductMerchantFee         float64 `json:"productMerchantFee"`
		ProductCategoryID          int64   `json:"productCategoryId"`
		ProductCategoryName        string  `json:"productCategoryName"`
		ProductTypeID              int64   `json:"productTypeId"`
		ProductTypeName            string  `json:"productTypeName"`
		ReferenceNumber            string  `json:"referenceNumber"`
		ProviderReferenceNumber    string  `json:"providerReferenceNumber"`
		StatusCode                 string  `json:"statusCode"`
		StatusMessage              string  `json:"statusMessage"`
		StatusDesc                 string  `json:"statusDesc"`
		StatusCodeDetail           string  `json:"statusCodeDetail"`
		StatusMessageDetail        string  `json:"statusMessageDetail"`
		StatusDescDetail           string  `json:"statusDescDetail"`
		ProductReferenceID         int64   `json:"productReferenceId"`
		ProductReferenceCode       string  `json:"productReferenceCode"`
		CustomerID                 string  `json:"customerId"`
		OtherReff                  string  `json:"otherReff"`
		OtherCustomerInfo          string  `json:"otherCustomerInfo"`
		SavingAccountName          string  `json:"savingAccountName"`
		SavingAccountID            int64   `json:"savingAccountId"`
		SavingAccountNumber        string  `json:"savingAccountNumber"`
		TransactionTotalAmount     float64 `json:"transactionTotalAmount"`
		UserAppID                  int64   `json:"userAppId"`
		Username                   string  `json:"username"`
		CreatedAt                  string  `json:"createdAt"`
		CreatedBy                  string  `json:"createdBy"`
		UpdatedAt                  string  `json:"updatedAt"`
		UpdatedBy                  string  `json:"updatedBy"`
	}
)
type (
	ReqInquiry struct {
		ProductCode     string          `json:"productCode"`
		AdditionalField AdditionalField `json:"additionalField"`
	}
	AdditionalField struct {
		Periode          int     `json:"periode"`
		SubscriberNumber string  `json:"subscriberNumber"`
		SubscriberName   string  `json:"subscriberName"`
		Amount           float64 `json:"amount"`
	}
	ProviderInqRequest struct {
		ProductReferenceId   int    `json:"productReferenceId"`
		ProductReferenceCode string `json:"productReferenceCode"`
		ProductCode          string `json:"productCode"`
		SubscriberNumber     string `json:"subscriberNumber"`
		SubscriberName       string `json:"subscriberName"`
		ReferenceNumber      string `json:"referenceNumber"`
		Url                  string `json:"url"`
		Periode              int    `json:"periode"`
	}
	RespInquiry struct {
		StatusMessage string `json:"statusMessage"`
		CreatedAt     string `json:"createdAt"`
		// MerchantOutletName     string      `json:"merchantOutletName"`
		// MerchantOutletUsername string      `json:"merchantOutletUsername"`
		ReferenceNumber string `json:"referenceNumber"`
		ProductName     string `json:"productName"`
		// ProductCode        string      `json:"productCode"`
		SubscriberNumber   string      `json:"subscriberNumber"`
		ProductPrice       float64     `json:"productPrice"`
		ProductAdminFee    float64     `json:"productAdminFee"`
		ProductMerchantFee float64     `json:"productMerchantFee"`
		TotalTrxAmount     float64     `json:"totalTrxAmount"`
		BillInfo           interface{} `json:"billInfo"` //jelasin tagihan/detil produk
	}
	//////////////////
	ReqPaymentTrx struct {
		PaymentMethodId         string `json:"paymentMethodId"`
		PaymentMethodName       string `json:"paymentMethodName"`
		ReferenceNumber         string `json:"referenceNumber"`
		ProviderReferenceNumber string `json:"providerReferenceNumber"`
		AccountNumber           string `json:"accountNumber"`
		AccountPIN              string `json:"accounrPin"`
	}
	RespPayment struct { //jelasin produknya
		CreatedAt       string `json:"createdAt"`
		ReferenceNumber string `json:"referenceNumber"`
		ProductName     string `json:"productName"`
		// ProductCode         string      `json:"productCode"`
		ProductCategoryId   int         `json:"productCategoryId"`
		ProductCategoryName string      `json:"productCategoryName"`
		SubscriberNumber    string      `json:"subscriberNumber"`
		ProductPrice        float64     `json:"productPrice"`
		ProductAdminFee     float64     `json:"productAdminFee"`
		ProductMerchantFee  float64     `json:"productMerchantFee"`
		TotalTrxAmount      float64     `json:"totalTrxAmount"`
		BillInfo            interface{} `json:"billInfo"` //jelasin tagihan/detil produk
	}
)
type (
	ReqInquiryProvider struct {
		ProductCode             string  `json:"product_code"`
		ReferenceNumber         string  `json:"reference_number"`
		ReferenceNumberMerchant string  `json:"reference_number_merchant"`
		CustomerID              string  `json:"customer_id"`
		Periode                 string  `json:"periode"`
		Amount                  float64 `json:"amount"`
	}
	ReqPaymentProvider struct {
		ReferenceNumber         string `json:"referenceNumber"`
		ReferenceNumberMerchant string `json:"referenceNumber_merchant"`
	}
)
type (
	ReqAviceTrx struct {
		ReferenceNumber string `json:"referenceNumber" validate:"required"`
	}
)

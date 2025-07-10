package models

type (
	ReqGetTransaction struct {
		Start     int64       `json:"start" `
		Lenght    int64       `json:"length"`
		Columns   string      `json:"columns"`
		Search    string      `json:"search"`
		Order     string      `json:"order" `
		Sort      string      `json:"sort" `
		StartDate string      `json:"start_date"`
		EndDate   string      `json:"end_date"`
		Draw      int         `json:"draw"`
		Filter    Transaction `json:"filter"`
	}
	Transaction struct {
		ID                         int64   `json:"id"`
		ProductProviderName        string  `json:"product_provider_name"`
		ProductProviderCode        string  `json:"product_provider_code"`
		ProductProviderPrice       float64 `json:"product_provider_price"`
		ProductProviderAdminFee    float64 `json:"product_provider_admin_fee"`
		ProductProviderMerchantFee float64 `json:"product_provider_merchant_fee"`
		ProductID                  int64   `json:"product_id"`
		ProductName                string  `json:"product_name"`
		ProductCode                string  `json:"product_code"`
		ProductPrice               float64 `json:"product_price"`
		ProductAdminFee            float64 `json:"product_admin_fee"`
		ProductMerchantFee         float64 `json:"product_merchant_fee"`
		ProductCategoryID          int64   `json:"product_category_id"`
		ProductCategoryName        string  `json:"product_category_name"`
		ProductTypeID              int64   `json:"product_type_id"`
		ProductTypeName            string  `json:"product_type_name"`
		ReferenceNumber            string  `json:"reference_number"`
		ProviderReferenceNumber    string  `json:"provider_reference_number"`
		StatusCode                 string  `json:"status_code"`
		StatusMessage              string  `json:"status_message"`
		StatusDesc                 string  `json:"status_desc"`
		StatusCodeDetail           string  `json:"status_code_detail"`
		StatusMessageDetail        string  `json:"status_message_detail"`
		StatusDescDetail           string  `json:"status_desc_detail"`
		ProductReferenceID         int64   `json:"product_reference_id"`
		ProductReferenceCode       string  `json:"product_reference_code"`
		CustomerID                 string  `json:"customer_id"`
		OtherReff                  string  `json:"other_reff"`
		OtherCustomerInfo          string  `json:"other_customer_info"`
		SavingAccountName          string  `json:"saving_account_name"`
		SavingAccountID            int64   `json:"saving_account_id"`
		SavingAccountNumber        string  `json:"saving_account_number"`
		TransactionTotalAmount     string  `json:"transaction_total_amount"`
		UserAppID                  int64   `json:"user_app_id"`
		Username                   string  `json:"username"`
		CreatedAt                  string  `json:"createdAt"`
		CreatedBy                  string  `json:"createdBy"`
		UpdatedAt                  string  `json:"updatedAt"`
		UpdatedBy                  string  `json:"updatedBy"`
	}
	RespGetTrx struct {
		Index                      int     `json:"index"`
		Id                         int     `json:"id"`
		ProductProviderName        string  `json:"product_provider_name"`
		ProductProviderCode        string  `json:"product_provider_code"`
		ProductProviderPrice       float64 `json:"product_provider_price"`
		ProductProviderAdminFee    float64 `json:"product_provider_admin_fee"`
		ProductProviderMerchantFee float64 `json:"product_provider_merchant_fee"`
		ProductID                  int64   `json:"product_id"`
		ProductName                string  `json:"product_name"`
		ProductCode                string  `json:"product_code"`
		ProductPrice               float64 `json:"product_price"`
		ProductAdminFee            float64 `json:"product_admin_fee"`
		ProductMerchantFee         float64 `json:"product_merchant_fee"`
		ProductCategoryID          int64   `json:"product_category_id"`
		ProductCategoryName        string  `json:"product_category_name"`
		ProductTypeID              int64   `json:"product_type_id"`
		ProductTypeName            string  `json:"product_type_name"`
		ReferenceNumber            string  `json:"reference_number"`
		ProviderReferenceNumber    string  `json:"provider_reference_number"`
		StatusCode                 string  `json:"status_code"`
		StatusMessage              string  `json:"status_message"`
		StatusDesc                 string  `json:"status_desc"`
		StatusCodeDetail           string  `json:"status_code_detail"`
		StatusMessageDetail        string  `json:"status_message_detail"`
		StatusDescDetail           string  `json:"status_desc_detail"`
		ProductReferenceID         int64   `json:"product_reference_id"`
		ProductReferenceCode       string  `json:"product_reference_code"`
		CustomerID                 string  `json:"customer_id"`
		OtherReff                  string  `json:"other_reff"`
		OtherCustomerInfo          string  `json:"other_customer_info"`
		SavingAccountName          string  `json:"saving_account_name"`
		SavingAccountID            int64   `json:"saving_account_id"`
		SavingAccountNumber        string  `json:"saving_account_number"`
		TransactionTotalAmount     string  `json:"transaction_total_amount"`
		UserAppID                  int64   `json:"user_app_id"`
		Username                   string  `json:"username"`
		CreatedAt                  string  `json:"createdAt"`
		CreatedBy                  string  `json:"createdBy"`
		UpdatedAt                  string  `json:"updatedAt"`
		UpdatedBy                  string  `json:"updatedBy"`
	}
)

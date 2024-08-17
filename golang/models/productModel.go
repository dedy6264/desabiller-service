package models

type (
	ReqGetProvider struct {
		ID           int       `json:"id"`
		ProviderName string    `json:"providerName"`
		Filter       FilterReq `json:"filter"`
	}
	/////
	RespGetProvider struct {
		ID           int    `json:"id"`
		ProviderName string `json:"providerName"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
		CreatedBy    string `json:"createdBy"`
		UpdatedBy    string `json:"updatedBy"`
	}
	ReqGetProductClan struct {
		ID           int       `json:"id"`
		ProductClanName string    `json:"productClanName"`
		Filter       FilterReq `json:"filter"`
	}
	/////
	RespGetProductClan struct {
		ID           int    `json:"id"`
		ProductClanName string `json:"productClanName"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
		CreatedBy    string `json:"createdBy"`
		UpdatedBy    string `json:"updatedBy"`
	}
	// ResGetProductCategory struct {
	// 	ID                  int    `json:"id"`
	// 	ProductCategoryName string `json:"productCategoryName"`
	// 	ProductCategoryCode string `json:"productCategoryCode"`
	// 	MerchantOutletId    int    `json:"merchantOutletId"`
	// 	MerchantOutletName  string `json:"merchantOutletName"`
	// 	Updateable          bool   `json:"updateable"`

	// 	MerchantName string `json:"merchantName"`
	// 	MerchantId   int    `json:"merchantId"`
	// 	ClientName   string `json:"clientName"`
	// 	ClientId     int    `json:"clientId"`
	// 	CreatedAt    string `json:"createdAt"`
	// 	UpdatedAt    string `json:"updatedAt"`
	// 	CreatedBy    string `json:"createdBy"`
	// 	UpdatedBy    string `json:"updatedBy"`
	// }
	// /////
	// ReqGetListProductBillerProvider struct {
	// 	ID                         int     `json:"id"`
	// 	ProductProviderName        string  `json:"productProviderName"`
	// 	ProductProviderCode        string  `json:"productProviderCode"`
	// 	ProductProviderPrice       float64 `json:"productProviderPrice"`
	// 	ProductProviderAdminFee    float64 `json:"productProviderAdminFee"`
	// 	ProductProviderMerchantFee float64 `json:"productProviderMerchantFee"`
	// 	IsOpen                     bool    `json:"isOpen"`
	// 	ProductTypeId              int     `json:"productTypeId"`
	// 	ProductCategoryId          int     `json:"productCategoryId"`
	// 	Limit                      int     `json:"limit"`
	// 	Offset                     int     `json:"offset"`
	// 	OrderBy                    string  `json:"orderBy"`
	// 	StartDate                  string  `json:"startDate"`
	// 	EndDate                    string  `json:"endDate"`
	// 	Username                   string  `json:"username"`
	// }
	// ResGetProductBillerProvider struct {
	// 	ID                         int     `json:"id"`
	// 	ProductProviderName        string  `json:"productProviderName"`
	// 	ProductProviderCode        string  `json:"productProviderCode"`
	// 	ProductProviderPrice       float64 `json:"productProviderPrice"`
	// 	ProductProviderAdminFee    float64 `json:"productProviderAdminFee"`
	// 	ProductProviderMerchantFee float64 `json:"productProviderMerchantFee"`
	// 	IsOpen                     bool    `json:"isOpen"`
	// 	ProductTypeId              int     `json:"productTypeId"`
	// 	ProductCategoryId          int     `json:"productCategoryId"`
	// 	CreatedAt                  string  `json:"createdAt"`
	// 	UpdatedAt                  string  `json:"updatedAt"`
	// 	CreatedBy                  string  `json:"createdBy"`
	// 	UpdatedBy                  string  `json:"updatedBy"`
	// }
	// /////
	// ReqGetListProductBiller struct {
	// 	ID                int    `json:"id"`
	// 	ProductName       string `json:"productName"`
	// 	ProductCode       string `json:"productCode"`
	// 	ProductProviderId int    `json:"productProviderId"`
	// 	IsOpen            bool   `json:"isOpen"`
	// 	ProductTypeId     int    `json:"productTypeId"`
	// 	ProductCategoryId int    `json:"productCategoryId"`
	// 	Limit             int    `json:"limit"`
	// 	Offset            int    `json:"offset"`
	// 	OrderBy           string `json:"orderBy"`
	// 	StartDate         string `json:"startDate"`
	// 	EndDate           string `json:"endDate"`
	// 	Username          string `json:"username"`
	// }
	// ResGetProductBiller struct {
	// 	ID                int    `json:"id"`
	// 	ProductName       string `json:"productName"`
	// 	ProductCode       string `json:"productCode"`
	// 	ProductProviderId int    `json:"productProviderId"`
	// 	IsOpen            bool   `json:"isOpen"`
	// 	ProductTypeId     int    `json:"productTypeId"`
	// 	ProductCategoryId int    `json:"productCategoryId"`
	// 	CreatedAt         string `json:"createdAt"`
	// 	UpdatedAt         string `json:"updatedAt"`
	// 	CreatedBy         string `json:"createdBy"`
	// 	UpdatedBy         string `json:"updatedBy"`
	// }
	// /////
	// ReqGetListProductPos struct {
	// 	ID                   int     `json:"id"`
	// 	ProductName          string  `json:"productName"`
	// 	ProductCode          string  `json:"productCode"`
	// 	ProductPriceProvider float64 `json:"productPriceProvider"`
	// 	MerchantId           int     `json:"merchantId"`
	// 	MerchantName         string  `json:"merchantName"`
	// 	ProductPrice         float64 `json:"productPrice"`
	// 	IsOpen               bool    `json:"isOpen"`
	// 	ProductTypeId        int     `json:"productTypeId"`
	// 	ProductCategoryId    int     `json:"productCategoryId"`
	// 	Limit                int     `json:"limit"`
	// 	Offset               int     `json:"offset"`
	// 	OrderBy              string  `json:"orderBy"`
	// 	StartDate            string  `json:"startDate"`
	// 	EndDate              string  `json:"endDate"`
	// 	Username             string  `json:"username"`
	// 	Draw                 int     `json:"draw"`
	// }
	// ResGetProductPos struct {
	// 	ID                   int     `json:"id"`
	// 	ProductName          string  `json:"productName"`
	// 	ProductCode          string  `json:"productCode"`
	// 	ProductPriceProvider float64 `json:"productPriceProvider"`
	// 	MerchantId           int     `json:"merchantId"`
	// 	MerchantName         string  `json:"merchantName"`
	// 	ProductPrice         float64 `json:"productPrice"`
	// 	IsOpen               bool    `json:"isOpen"`
	// 	ProductTypeId        int     `json:"productTypeId"`
	// 	ProductCategoryId    int     `json:"productCategoryId"`
	// 	CreatedAt            string  `json:"createdAt"`
	// 	UpdatedAt            string  `json:"updatedAt"`
	// 	CreatedBy            string  `json:"createdBy"`
	// 	UpdatedBy            string  `json:"updatedBy"`
	// }
)

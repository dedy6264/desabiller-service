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

	ReqGetProductCategory struct {
		ID                  int       `json:"id"`
		ProductCategoryName string    `json:"productCategoryName"`
		Filter              FilterReq `json:"filter"`
	}
	/////
	RespGetProductCategory struct {
		ID                  int    `json:"id"`
		ProductCategoryName string `json:"productCategoryName"`
		CreatedAt           string `json:"createdAt"`
		UpdatedAt           string `json:"updatedAt"`
		CreatedBy           string `json:"createdBy"`
		UpdatedBy           string `json:"updatedBy"`
	}
	ReqGetProductType struct {
		ID              int       `json:"id"`
		ProductTypeName string    `json:"productTypeName"`
		Filter          FilterReq `json:"filter"`
	}
	/////
	RespGetProductType struct {
		ID              int    `json:"id"`
		ProductTypeName string `json:"productTypeName"`
		CreatedAt       string `json:"createdAt"`
		UpdatedAt       string `json:"updatedAt"`
		CreatedBy       string `json:"createdBy"`
		UpdatedBy       string `json:"updatedBy"`
	}
	ReqGetProductProvider struct {
		ID                         int       `json:"id"`
		ProviderName               string    `json:"providerName"`
		ProductProviderId          int       `json:"productProviderId"`
		ProductProviderName        string    `json:"productProviderName"`
		ProductProviderCode        string    `json:"productProviderCode"`
		ProviderId                 int       `json:"providerId"`
		ProductProviderAdminFee    float64   `json:"productProviderAdminFee"`
		ProductProviderPrice       float64   `json:"productProviderPrice"`
		ProductProviderMerchantFee float64   `json:"productProviderMerchantFee"`
		Filter                     FilterReq `json:"filter"`
	}
	/////
	RespGetProductProvider struct {
		ID                         int     `json:"id"`
		ProviderName               string  `json:"providerName"`
		ProductProviderName        string  `json:"productProviderName"`
		ProductProviderCode        string  `json:"productProviderCode"`
		ProviderId                 int     `json:"providerId"`
		ProductProviderAdminFee    float64 `json:"productProviderAdminFee"`
		ProductProviderPrice       float64 `json:"productProviderPrice"`
		ProductProviderMerchantFee float64 `json:"productProviderMerchantFee"`
		CreatedAt                  string  `json:"createdAt"`
		UpdatedAt                  string  `json:"updatedAt"`
		CreatedBy                  string  `json:"createdBy"`
		UpdatedBy                  string  `json:"updatedBy"`
	}
	ReqGetProduct struct {
		ID                   int       `json:"id"`
		ProviderId           int       `json:"providerId"`
		ProductProviderId    int       `json:"productProviderId"`
		ProductCategoryId    int       `json:"productCategoryId"`
		ProductTypeId        int       `json:"productTypeId"`
		ProductName          string    `json:"productName"`
		ProductCode          string    `json:"productCode"`
		ProductAdminFee      float64   `json:"productAdminFee"`
		ProductPrice         float64   `json:"productPrice"`
		ProductMerchantFee   float64   `json:"productMerchantFee"`
		ProductReferenceId   int       `json:"productReferenceId"`
		ProductReferenceCode string    `json:"productReferenceCode"`
		Filter               FilterReq `json:"filter"`
	}
	/////
	RespGetProduct struct {
		ID                         int     `json:"id"`
		ProviderId                 int     `json:"providerId"`
		ProviderName               string  `json:"providerName"`
		ProductCode                string  `json:"productCode"`
		ProductProviderId          int     `json:"productProviderId"`
		ProductProviderName        string  `json:"productProviderName"`
		ProductProviderCode        string  `json:"productProviderCode"`
		ProductProviderAdminFee    float64 `json:"productProviderAdminFee"`
		ProductProviderPrice       float64 `json:"productProviderPrice"`
		ProductProviderMerchantFee float64 `json:"productProviderMerchantFee"`
		ProductCategoryId          int     `json:"productCategoryId"`
		ProductCategoryName        string  `json:"productCategoryName"`
		ProductTypeId              int     `json:"productTypeId"`
		ProductTypeName            string  `json:"productTypeName"`
		ProductName                string  `json:"productName"`
		ProductAdminFee            float64 `json:"productAdminFee"`
		ProductPrice               float64 `json:"productPrice"`
		ProductMerchantFee         float64 `json:"productMerchantFee"`
		CreatedAt                  string  `json:"createdAt"`
		UpdatedAt                  string  `json:"updatedAt"`
		CreatedBy                  string  `json:"createdBy"`
		UpdatedBy                  string  `json:"updatedBy"`
		ProductReferenceId         int     `json:"productReferenceId"`
		ProductReferenceCode       string  `json:"productReferenceCode"`
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

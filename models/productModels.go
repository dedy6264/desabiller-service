package models

type (
	ReqGetProductType struct {
		Start     int64       `json:"start" `
		Lenght    int64       `json:"length"`
		Columns   string      `json:"columns"`
		Search    string      `json:"search"`
		Order     string      `json:"order" `
		Sort      string      `json:"sort" `
		StartDate string      `json:"startDate"`
		EndDate   string      `json:"endDate"`
		Draw      int         `json:"draw"`
		Filter    ProductType `json:"filter"`
	}
	ProductType struct {
		ID              int64  `json:"id"`
		ProductTypeName string `json:"productTypeName"`
		CreatedBy       string `json:"createdBy"`
		UpdatedBy       string `json:"updatedBy"`
		CreatedAt       string `json:"createdAt"`
		UpdatedAt       string `json:"updatedAt"`
	}
	ReqGetProductCategory struct {
		Start     int64           `json:"start" `
		Lenght    int64           `json:"length"`
		Columns   string          `json:"columns"`
		Search    string          `json:"search"`
		Order     string          `json:"order" `
		Sort      string          `json:"sort" `
		StartDate string          `json:"startDate"`
		EndDate   string          `json:"endDate"`
		Draw      int             `json:"draw"`
		Filter    ProductCategory `json:"filter"`
	}
	ProductCategory struct {
		ID                  int64  `json:"id"`
		ProductCategoryName string `json:"productCategoryName"`
		CreatedBy           string `json:"createdBy"`
		UpdatedBy           string `json:"updatedBy"`
		CreatedAt           string `json:"createdAt"`
		UpdatedAt           string `json:"updatedAt"`
	}
	ReqGetProductReference struct {
		Start     int64            `json:"start" `
		Lenght    int64            `json:"length"`
		Columns   string           `json:"columns"`
		Search    string           `json:"search"`
		Order     string           `json:"order" `
		Sort      string           `json:"sort" `
		StartDate string           `json:"startDate"`
		EndDate   string           `json:"endDate"`
		Draw      int              `json:"draw"`
		Filter    ProductReference `json:"filter"`
	}
	ProductReference struct {
		ID                   int64  `json:"id"`
		ProductReferenceName string `json:"productCategoryName"`
		ProductReferenceCode string `json:"productCategoryCode"`
		CreatedBy            string `json:"createdBy"`
		UpdatedBy            string `json:"updatedBy"`
		CreatedAt            string `json:"createdAt"`
		UpdatedAt            string `json:"updatedAt"`
	}
	ReqGetProduct struct {
		Start     int64   `json:"start" `
		Lenght    int64   `json:"length"`
		Columns   string  `json:"columns"`
		Search    string  `json:"search"`
		Order     string  `json:"order" `
		Sort      string  `json:"sort" `
		StartDate string  `json:"startDate"`
		EndDate   string  `json:"endDate"`
		Draw      int     `json:"draw"`
		Filter    Product `json:"filter"`
	}
	Product struct {
		ID                         int     `json:"id"`
		ProductProviderName        string  `json:"productProviderName"`
		ProductProviderCode        string  `json:"productProviderCode"`
		ProductProviderPrice       float64 `json:"productProviderPrice"`
		ProductProviderAdminFee    float64 `json:"productProviderAdminFee"`
		ProductProviderMerchantFee float64 `json:"productProviderMerchantFee"`
		ProductCategoryID          int     `json:"productCategoryId"`
		ProductCategoryName        string  `json:"productCategoryName"`
		ProductTypeID              int     `json:"productTypeId"`
		ProductTypeName            string  `json:"productTypeName"`
		ProductName                string  `json:"productName"`
		ProductCode                string  `json:"productCode"`
		ProductAdminFee            float64 `json:"productAdminFee"`
		ProductPrice               float64 `json:"productPrice"`
		ProductMerchantFee         float64 `json:"productMerchantFee"`
		ProductReferenceID         int     `json:"productReferenceId"`
		ProductReferenceCode       string  `json:"productReferenceCode"`
		ProductDenom               string  `json:"productDenom"`
		CreatedBy                  string  `json:"createdBy"`
		UpdatedBy                  string  `json:"updatedBy"`
		CreatedAt                  string  `json:"createdAt"`
		UpdatedAt                  string  `json:"updatedAt"`
	}
)

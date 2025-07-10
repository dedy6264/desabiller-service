package models

type (
	ReqGetProductType struct {
		Start     int64       `json:"start" `
		Lenght    int64       `json:"length"`
		Columns   string      `json:"columns"`
		Search    string      `json:"search"`
		Order     string      `json:"order" `
		Sort      string      `json:"sort" `
		StartDate string      `json:"start_date"`
		EndDate   string      `json:"end_date"`
		Draw      int         `json:"draw"`
		Filter    ProductType `json:"filter"`
	}
	ProductType struct {
		ID              int64  `json:"id"`
		ProductTypeName string `json:"product_type_name"`
		CreatedBy       string `json:"created_by"`
		UpdatedBy       string `json:"updated_by"`
		CreatedAt       string `json:"created_at"`
		UpdatedAt       string `json:"updated_at"`
	}
	ReqGetProductCategory struct {
		Start     int64           `json:"start" `
		Lenght    int64           `json:"length"`
		Columns   string          `json:"columns"`
		Search    string          `json:"search"`
		Order     string          `json:"order" `
		Sort      string          `json:"sort" `
		StartDate string          `json:"start_date"`
		EndDate   string          `json:"end_date"`
		Draw      int             `json:"draw"`
		Filter    ProductCategory `json:"filter"`
	}
	ProductCategory struct {
		ID                  int64  `json:"id"`
		ProductCategoryName string `json:"product_category_name"`
		CreatedBy           string `json:"created_by"`
		UpdatedBy           string `json:"updated_by"`
		CreatedAt           string `json:"created_at"`
		UpdatedAt           string `json:"updated_at"`
	}
	ReqGetProductReference struct {
		Start     int64            `json:"start" `
		Lenght    int64            `json:"length"`
		Columns   string           `json:"columns"`
		Search    string           `json:"search"`
		Order     string           `json:"order" `
		Sort      string           `json:"sort" `
		StartDate string           `json:"start_date"`
		EndDate   string           `json:"end_date"`
		Draw      int              `json:"draw"`
		Filter    ProductReference `json:"filter"`
	}
	ProductReference struct {
		ID                   int64  `json:"id"`
		ProductReferenceName string `json:"product_category_name"`
		ProductReferenceCode string `json:"product_category_code"`
		CreatedBy            string `json:"created_by"`
		UpdatedBy            string `json:"updated_by"`
		CreatedAt            string `json:"created_at"`
		UpdatedAt            string `json:"updated_at"`
	}
)

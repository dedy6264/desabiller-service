package models

type (
	ReqListSegment struct {
		ID          int    `json:"id"`
		SegmentName string `json:"segmentName"`
		Limit       int    `json:"limit"`
		Offset      int    `json:"offset"`
		OrderBy     string `json:"orderBy"`
		StartDate   string `json:"startDate"`
		EndDate     string `json:"endDate"`
		Username    string `json:"username"`
	}
	ResListSegment struct {
		ID          int    `json:"id"`
		SegmentName string `json:"segmentName"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
		CreatedBy   string `json:"createdBy"`
		UpdatedBy   string `json:"updatedBy"`
	}
	///
	ReqListSegmentProduct struct {
		ID                      int     `json:"id"`
		SegmentProductPrefix    string  `json:"segmentProductPrefix"`
		SegmentId               int     `json:"segmentId"`
		ProductBillerId         int     `json:"productBillerId"`
		ProductBillerProviderId int     `json:"productBillerProviderId"`
		ProductPrice            float64 `json:"productPrice"`
		ProductAdminFee         float64 `json:"productAdminFee"`
		ProductMerchantFee      float64 `json:"productMerchantFee"`
		IsOpen                  bool    `json:"isOpen"`
		Limit                   int     `json:"limit"`
		Offset                  int     `json:"offset"`
		OrderBy                 string  `json:"orderBy"`
		StartDate               string  `json:"startDate"`
		EndDate                 string  `json:"endDate"`
		Username                string  `json:"username"`
	}
	ResListSegmentProduct struct {
		ID                      int     `json:"id"`
		SegmentProductPrefix    string  `json:"segmentProductPrefix"`
		SegmentId               int     `json:"segmentId"`
		ProductBillerId         int     `json:"productBillerId"`
		ProductBillerProviderId int     `json:"productBillerProviderId"`
		ProductPrice            float64 `json:"productPrice"`
		ProductAdminFee         float64 `json:"productAdminFee"`
		ProductMerchantFee      float64 `json:"productMerchantFee"`
		IsOpen                  bool    `json:"isOpen"`
		CreatedAt               string  `json:"createdAt"`
		UpdatedAt               string  `json:"updatedAt"`
		CreatedBy               string  `json:"createdBy"`
		UpdatedBy               string  `json:"updatedBy"`
	}
)

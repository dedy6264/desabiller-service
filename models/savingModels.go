package models

type (
	RespGetAccount struct {
		ID              int     `json:"id"`
		CifID           int     `json:"cifID"`
		CifName         string  `json:"cifName"`
		CifEmail        string  `json:"cifEmail"`
		AccountNumber   string  `json:"accountNumber"`
		Balance         float64 `json:"balance"`
		SavingSegmentID int     `json:"savingSegmentId"`
		AccountPin      string  `json:"accountPin"`
		CreatedAt       string  `json:"createdAt"`
		CreatedBy       string  `json:"createdBy"`
		UpdatedAt       string  `json:"updatedAt"`
		UpdatedBy       string  `json:"updatedBy"`
	}
	ReqGetSavingType struct {
		Start     int64      `json:"start" `
		Lenght    int64      `json:"length"`
		Columns   string     `json:"columns"`
		Search    string     `json:"search"`
		Order     string     `json:"order" `
		Sort      string     `json:"sort" `
		StartDate string     `json:"startDate"`
		EndDate   string     `json:"endDate"`
		Draw      int        `json:"draw"`
		Filter    SavingType `json:"filter"`
	}
	SavingType struct {
		ID             int    `json:"id"`
		SavingTypeName string `json:"savingTypeName"`
		SavingTypeDesc string `json:"savingTypeDesc"`
		CreatedAt      string `json:"createdAt"`
		CreatedBy      string `json:"createdBy"`
		UpdatedAt      string `json:"updatedAt"`
		UpdatedBy      string `json:"updatedBy"`
	}
	ReqGetSavingSegment struct {
		Start     int64         `json:"start" `
		Lenght    int64         `json:"length"`
		Columns   string        `json:"columns"`
		Search    string        `json:"search"`
		Order     string        `json:"order" `
		Sort      string        `json:"sort" `
		StartDate string        `json:"startDate"`
		EndDate   string        `json:"endDate"`
		Draw      int           `json:"draw"`
		Filter    SavingSegment `json:"filter"`
	}
	SavingSegment struct {
		ID                int     `json:"id"`
		SavingSegmentName string  `json:"savingSegmentName"`
		LimitAmount       float64 `json:"limitAmount"`
		SavingTypeID      int     `json:"savingTypeId"`
		SavingTypeName    string  `json:"savingTypeName"`
		CreatedAt         string  `json:"createdAt"`
		CreatedBy         string  `json:"createdBy"`
		UpdatedAt         string  `json:"updatedAt"`
		UpdatedBy         string  `json:"updatedBy"`
	}
	ReqGetSavingTransaction struct {
		Start     int64             `json:"start" `
		Lenght    int64             `json:"length"`
		Columns   string            `json:"columns"`
		Search    string            `json:"search"`
		Order     string            `json:"order" `
		Sort      string            `json:"sort" `
		StartDate string            `json:"startDate"`
		EndDate   string            `json:"endDate"`
		Draw      int               `json:"draw"`
		Filter    SavingTransaction `json:"filter"`
	}
	SavingTransaction struct {
		ID                     int     `json:"id"`
		ReferenceNumber        string  `json:"referenceNumber"`
		ReferenceNumberPartner string  `json:"referenceNumberPartner"`
		DcType                 string  `json:"dcType"`
		TransactionAmount      float64 `json:"transactionAmount"`
		TransactionCode        string  `json:"transactionCode"`
		AccountID              int     `json:"accountId"`
		AccountNumber          string  `json:"accountNumber"`
		LastBalance            float64 `json:"lastBalance"`
		CreatedAt              string  `json:"createdAt"`
		CreatedBy              string  `json:"createdBy"`
		UpdatedAt              string  `json:"updatedAt"`
		UpdatedBy              string  `json:"updatedBy"`
	}
)

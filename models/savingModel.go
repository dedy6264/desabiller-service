package models

type (
	ReqGetCif struct {
		ID         int       `json:"id"`
		CifName    string    `json:"cifName"`
		CifNik     string    `json:"cifNik"`
		CifPhone   string    `json:"cifPhone"`
		CifEmail   string    `json:"cifEmail"`
		CifAddress string    `json:"cifAddress"`
		Filter     FilterReq `json:"filter"`
	}
	RespGetCif struct {
		ID         int    `json:"id"`
		CifName    string `json:"cifName"`
		CifNik     string `json:"cifNik"`
		CifPhone   string `json:"cifPhone"`
		CifEmail   string `json:"cifEmail"`
		CifAddress string `json:"cifAddress"`
		CreatedAt  string `json:"createdAt"`
		UpdatedAt  string `json:"updatedAt"`
		CreatedBy  string `json:"createdBy"`
		UpdatedBy  string `json:"updatedBy"`
	}
)
type (
	ReqGetSavingType struct {
		ID             int       `json:"id"`
		SavingTypeName string    `json:"savingTypeName"`
		SavingTypeDesc string    `json:"savingTypeDesc"`
		Filter         FilterReq `json:"filter"`
	}
	RespGetSavingType struct {
		ID             int    `json:"id"`
		SavingTypeName string `json:"savingTypeName"`
		SavingTypeDesc string `json:"savingTypeDesc"`
		CreatedAt      string `json:"createdAt"`
		CreatedBy      string `json:"createdBy"`
		UpdatedAt      string `json:"updatedAt"`
		UpdatedBy      string `json:"updatedBy"`
	}
)
type (
	ReqGetSavingSegment struct {
		ID                int       `json:"id"`
		SavingSegmentName string    `json:"savingSegmentName"`
		LimitAmount       float64   `json:"limitAmount"`
		SavingTypeID      int       `json:"savingTypeID"`
		Filter            FilterReq `json:"filter"`
	}
	RespGetSavingSegment struct {
		ID                int     `json:"id"`
		SavingSegmentName string  `json:"savingSegmentName"`
		LimitAmount       float64 `json:"limitAmount"`
		SavingTypeID      int     `json:"savingTypeID"`
		SavingTypeName    string  `json:"savingTypeName"`
		CreatedAt         string  `json:"createdAt"`
		CreatedBy         string  `json:"createdBy"`
		UpdatedAt         string  `json:"updatedAt"`
		UpdatedBy         string  `json:"updatedBy"`
	}
)
type (
	ReqGetAccount struct {
		ID              int       `json:"id"`
		CifID           int       `json:"cifId"`
		AccountNumber   string    `json:"accountNumber"`
		Balance         float64   `json:"balance"`
		SavingSegmentID int       `json:"savingSegmentId"`
		AccountPin      string    `json:"accountPin"`
		Filter          FilterReq `json:"filter"`
	}
	RespGetAccount struct {
		ID              int     `json:"id"`
		CifID           int     `json:"cifId"`
		AccountNumber   string  `json:"accountNumber"`
		Balance         float64 `json:"balance"`
		SavingSegmentID int     `json:"savingSegmentId"`
		AccountPin      string  `json:"accountPin"`
		CreatedAt       string  `json:"createdAt"`
		CreatedBy       string  `json:"createdBy"`
		UpdatedAt       string  `json:"updatedAt"`
		UpdatedBy       string  `json:"updatedBy"`
	}
)
type (
	ReqGetSavingTransaction struct {
		ID                    int       `json:"id"`
		ReferenceNumber       string    `json:"referenceNumber"`
		SavingReferenceNumber string    `json:"savingReferenceNumber"`
		DcType                string    `json:"dcType"`
		TransactionAmount     float64   `json:"transactionAmount"`
		TransactionCode       string    `json:"transactionCode"`
		AccountID             int       `json:"accountId"`
		AccountNumber         string    `json:"accountNumber"`
		LastBalance           float64   `json:"lastBalance"`
		Filter                FilterReq `json:"filter"`
	}
	RespGetSavingTransaction struct {
		ID                    int     `json:"id"`
		ReferenceNumber       string  `json:"referenceNumber"`
		SavingReferenceNumber string  `json:"savingReferenceNumber"`
		DcType                string  `json:"dcType"`
		TransactionAmount     float64 `json:"transactionAmount"`
		TransactionCode       string  `json:"transactionCode"`
		AccountID             int     `json:"accountId"`
		AccountNumber         string  `json:"accountNumber"`
		LastBalance           float64 `json:"lastBalance"`
		CreatedAt             string  `json:"createdAt"`
		CreatedBy             string  `json:"createdBy"`
		UpdatedAt             string  `json:"updatedAt"`
		UpdatedBy             string  `json:"updatedBy"`
	}
)

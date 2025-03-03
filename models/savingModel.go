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

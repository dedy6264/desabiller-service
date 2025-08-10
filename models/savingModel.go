package models

type (
	RespGetCif struct {
		ID         int    `json:"id"`
		CifName    string `json:"cifName"`
		CifNik     string `json:"cifNik"`
		CifEmail   string `json:"cifEmail"`
		CifAddress string `json:"cifAddress"`
		CreatedAt  string `json:"createdAt"`
		UpdatedAt  string `json:"updatedAt"`
		CreatedBy  string `json:"createdBy"`
		UpdatedBy  string `json:"updatedBy"`
	}
)

type ()
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
)

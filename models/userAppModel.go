package models

type (
	ReqGetUserApp struct {
		Start     int64   `json:"start" `
		Lenght    int64   `json:"length"`
		Columns   string  `json:"columns"`
		Search    string  `json:"search"`
		Order     string  `json:"order" `
		Sort      string  `json:"sort" `
		StartDate string  `json:"start_date"`
		EndDate   string  `json:"end_date"`
		Draw      int     `json:"draw"`
		Filter    UserApp `json:"filter"`
	}
	UserApp struct {
		ID              int64  `json:"id" `
		Username        string `json:"username"`
		Password        string `json:"password"`
		Name            string `json:"name"`
		Identity_type   string `json:"identity_type"`
		Identity_number string `json:"identity_number"`
		Phone           string `json:"phone"`
		Email           string `json:"email"`
		Gender          string `json:"gender"`
		Province        string `json:"province"`
		City            string `json:"city"`
		Address         string `json:"address"`
		Account_id      int64  `json:"account_id"`
		Status          string `json:"status"`
		CreatedBy       string `json:"created_by"`
		UpdatedBy       string `json:"updated_by"`
		CreatedAt       string `json:"created_at"`
		UpdatedAt       string `json:"updated_at"`
	}
	RespGetUserApp struct {
		RecordsTotal    int64     `json:"records_total"`
		RecordsFiltered int64     `json:"records_filtered"`
		Data            []UserApp `json:"data"`
	}
)

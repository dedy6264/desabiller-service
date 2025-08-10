package models

type (
	ReqGetUserApp struct {
		Start     int64   `json:"start" `
		Lenght    int64   `json:"length"`
		Columns   string  `json:"columns"`
		Search    string  `json:"search"`
		Order     string  `json:"order" `
		Sort      string  `json:"sort" `
		StartDate string  `json:"startDate"`
		EndDate   string  `json:"endDate"`
		Draw      int     `json:"draw"`
		Filter    UserApp `json:"filter"`
	}
	UserApp struct {
		ID             int64  `json:"id" `
		Username       string `json:"username"`
		Password       string `json:"password"`
		Name           string `json:"name"`
		IdentityType   string `json:"identityType"`
		IdentityNumber string `json:"identityNumber"`
		Phone          string `json:"phone"`
		Email          string `json:"email"`
		Gender         string `json:"gender"`
		Province       string `json:"province"`
		City           string `json:"city"`
		Address        string `json:"address"`
		CifID          int64  `json:"cifID"`
		Status         string `json:"status"`
		CreatedBy      string `json:"createdBy"`
		UpdatedBy      string `json:"updatedBy"`
		CreatedAt      string `json:"createdAt"`
		UpdatedAt      string `json:"updatedAt"`
		// AccountNumber   string  `json:"accountNumber"`
		// Balance         float64 `json:"balance"`
		// SavingSegmentID int64   `json:"savingSegmentId"`
	}

	RespGetUserApp struct {
		RecordsTotal    int64     `json:"recordsTotal"`
		RecordsFiltered int64     `json:"recordsFiltered"`
		Data            []UserApp `json:"data"`
	}
	ReqGetOtp struct {
		Start     int64  `json:"start" `
		Lenght    int64  `json:"length"`
		Columns   string `json:"columns"`
		Search    string `json:"search"`
		Order     string `json:"order" `
		Sort      string `json:"sort" `
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		Draw      int    `json:"draw"`
		Filter    Otp    `json:"filter"`
	}
	Otp struct {
		ID              int64  `json:"id" `
		CifID           int64  `json:"cifID"`
		Username        string `json:"username"`
		Otp             string `json:"otp"`
		ExpiredDuration int64  `json:"expiredDuration"`
		Phone           string `json:"phone"`
		CreatedBy       string `json:"createdBy"`
		UpdatedBy       string `json:"updatedBy"`
		CreatedAt       string `json:"createdAt"`
		UpdatedAt       string `json:"updatedAt"`
	}

	RespGetOtp struct {
		RecordsTotal    int64 `json:"recordsTotal"`
		RecordsFiltered int64 `json:"recordsFiltered"`
		Data            []Otp `json:"data"`
	}
	ReqGetCIF struct {
		Start     int64  `json:"start" `
		Lenght    int64  `json:"length"`
		Columns   string `json:"columns"`
		Search    string `json:"search"`
		Order     string `json:"order" `
		Sort      string `json:"sort" `
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		Draw      int    `json:"draw"`
		Filter    CIF    `json:"filter"`
	}
	CIF struct {
		ID         int64  `json:"id" `
		CifName    string `json:"cifName"`
		CifNoID    string `json:"cifNoId"`
		CifTypeID  string `json:"cifTypeId"`
		CifIDIndex string `json:"cifIDIndex"`
		CifEmail   string `json:"cifEmail"`
		CifAddress string `json:"cifAddress"`
		CreatedBy  string `json:"createdBy"`
		UpdatedBy  string `json:"updatedBy"`
		CreatedAt  string `json:"createdAt"`
		UpdatedAt  string `json:"updatedAt"`
	}
	ReqGetAccountSaving struct {
		Start     int64   `json:"start" `
		Lenght    int64   `json:"length"`
		Columns   string  `json:"columns"`
		Search    string  `json:"search"`
		Order     string  `json:"order" `
		Sort      string  `json:"sort" `
		StartDate string  `json:"startDate"`
		EndDate   string  `json:"endDate"`
		Draw      int     `json:"draw"`
		Filter    Account `json:"filter"`
	}
	Account struct {
		ID              int64   `json:"id" `
		CifID           int64   `json:"cifID"`
		AccountNumber   string  `json:"accountNumber"`
		AccountPin      string  `json:"accountPin"`
		Balance         float64 `json:"balance"`
		SavingSegmentID int64   `json:"savingSegmentId"`
		CreatedBy       string  `json:"createdBy"`
		UpdatedBy       string  `json:"updatedBy"`
		CreatedAt       string  `json:"createdAt"`
		UpdatedAt       string  `json:"updatedAt"`
	}
	ReqSetPin struct {
		Pin   string `json:"pin"`
		Phone string `json:"phone"`
	}
)
type (
	ReqLogin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	RespLogin struct {
		Data  Data   `json:"data"`
		Token string `json:"token"`
	}
	Data struct {
		//userAppId
		ID             int    `json:"id"`
		Username       string `json:"username"`
		Password       string `json:"password"`
		Name           string `json:"name"`
		IdentityType   string `json:"identityType"`
		IdentityNumber string `json:"identityNumber"`
		Phone          string `json:"phone"`
		Email          string `json:"email"`
		Gender         string `json:"gender"`
		Province       string `json:"province"`
		City           string `json:"city"`
		Address        string `json:"address"`
		CifID          int64  `json:"cifID"`
		Status         string `json:"status"`
		AccountNumber  string `json:"accountNumber"`
		IsSetPin       string `json:"isSetPin"`
	}
)

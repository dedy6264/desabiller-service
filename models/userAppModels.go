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
		IdentityType   string `json:"identity_type"`
		IdentityNumber string `json:"identity_number"`
		Phone          string `json:"phone"`
		Email          string `json:"email"`
		Gender         string `json:"gender"`
		Province       string `json:"province"`
		City           string `json:"city"`
		Address        string `json:"address"`
		AccountID      int64  `json:"account_id"`
		Status         string `json:"status"`
		CreatedBy      string `json:createdBy`
		UpdatedBy      string `json:updatedBy`
		CreatedAt      string `json:createdAt`
		UpdatedAt      string `json:updatedAt`
	}

	RespGetUserApp struct {
		RecordsTotal    int64     `json:"records_total"`
		RecordsFiltered int64     `json:"records_filtered"`
		Data            []UserApp `json:"data"`
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
		CifName    string `json:"cif_name"`
		CifNoID    string `json:"cif_no_id"`
		CifTypeID  string `json:"cif_type_id"`
		CifIDIndex string `json:"cif_id_index"`
		CifPhone   string `json:"cif_phone"`
		CifEmail   string `json:"cif_email"`
		CifAddress string `json:"cif_address"`
		CreatedBy  string `json:createdBy`
		UpdatedBy  string `json:updatedBy`
		CreatedAt  string `json:createdAt`
		UpdatedAt  string `json:updatedAt`
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
		CifID           int64   `json:"cif_id"`
		AccountNumber   string  `json:"account_number"`
		AccountPin      string  `json:"account_pin"`
		Balance         float64 `json:"balance"`
		SavingSegmentID int64   `json:"saving_segment_id"`
		CreatedBy       string  `json:createdBy`
		UpdatedBy       string  `json:updatedBy`
		CreatedAt       string  `json:createdAt`
		UpdatedAt       string  `json:updatedAt`
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
		IdentityType   string `json:"identity_type"`
		IdentityNumber string `json:"identity_number"`
		Phone          string `json:"phone"`
		Email          string `json:"email"`
		Gender         string `json:"gender"`
		Province       string `json:"province"`
		City           string `json:"city"`
		Address        string `json:"address"`
		AccountID      int64  `json:"account_id"`
		Status         string `json:"status"`
		AccountNumber  string `json:"account_number"`
		IsSetPin       string `json:"isSetPin"`
	}
)

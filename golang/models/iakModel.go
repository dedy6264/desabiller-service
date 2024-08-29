package models

type (
	IakCallback struct {
		Data struct {
			RefID       string `json:"ref_id"`
			Status      string `json:"status"`
			ProductCode string `json:"product_code"`
			CustomerID  string `json:"customer_id"`
			Price       string `json:"price"`
			Message     string `json:"message"`
			Sn          string `json:"sn"`
			Pin         string `json:"pin"`
			Balance     string `json:"balance"`
			TrID        string `json:"tr_id"`
			Rc          string `json:"rc"`
			Sign        string `json:"sign"`
		} `json:"data"`
	}

	ReqPaymentPrepaidIak struct {
		CustomerId  string `json:"customer_id"`
		ProductCode string `json:"product_code"`
		RefId       string `json:"ref_id"`
		Username    string `json:"username"`
		Sign        string `json:"sign"`
	}
	ReqInqIak struct {
		ProductCode string `json:"product_code"`
		CustomerId  string `json:"customer_id"`
		RefId       string `json:"ref_id"`
		Url         string `json:"url"`
	}
	ReqPaymentIak struct {
		ProductCode string `json:"product_code"`
		CustomerId  string `json:"customer_id"`
		RefId       string `json:"ref_id"`
	}
	ReqInquiryPostpaidIak struct {
		Commands string `json:"commands"`
		Hp       string `json:"hp"`
		Code     string `json:"code"`
		RefId    string `json:"ref_id"`
		Username string `json:"username"`
		Sign     string `json:"sign"`
	}
	RespPaymentPrepaidIak struct {
		Data struct {
			RefID       string `json:"ref_id"`
			Status      int    `json:"status"`
			ProductCode string `json:"product_code"`
			CustomerID  string `json:"customer_id"`
			Price       int    `json:"price"`
			Message     string `json:"message"`
			Balance     int    `json:"balance"`
			TrID        int    `json:"tr_id"`
			Rc          string `json:"rc"`
			Sn          string `json:"sn"`
		} `json:"data"`
	}
	RespInquiryPLNPostpaidIak struct {
		Data struct {
			TrID         int    `json:"tr_id"`
			Code         string `json:"code"`
			Hp           string `json:"hp"`
			TrName       string `json:"tr_name"`
			Period       string `json:"period"`
			Nominal      int    `json:"nominal"`
			Admin        int    `json:"admin"`
			RefID        string `json:"ref_id"`
			ResponseCode string `json:"response_code"`
			Message      string `json:"message"`
			Price        int    `json:"price"`
			SellingPrice int    `json:"selling_price"`
			Desc         struct {
				Tarif         string `json:"tarif"`
				Daya          int    `json:"daya"`
				LembarTagihan string `json:"lembar_tagihan"`
				Tagihan       struct {
					Detail []struct {
						Periode      string `json:"periode"`
						NilaiTagihan string `json:"nilai_tagihan"`
						Admin        string `json:"admin"`
						Denda        string `json:"denda"`
						Total        int    `json:"total"`
					} `json:"detail"`
				} `json:"tagihan"`
			} `json:"desc"`
		} `json:"data"`
		Meta []interface{} `json:"meta"`
	}
)

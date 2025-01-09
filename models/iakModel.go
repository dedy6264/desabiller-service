package models

type (
	RespWorkerUndefined struct {
		ResponseCode string `json:"response_code"`
		Message      string `json:"message"`
	}
	RespWorkerUndefinedI struct {
		Data struct {
			ResponseCode string `json:"response_code"`
			Message      string `json:"message"`
		} `json:"data"`
	}
	RespWorkerUndefinedII struct {
		Data struct {
			Rc      string `json:"rc"`
			Message string `json:"message"`
			Status  int    `json:"status"`
		} `json:"data"`
	}
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
		Commands             string `json:"commands"`
		ProductCode          string `json:"product_code"`
		CustomerId           string `json:"customer_id"`
		RefId                string `json:"ref_id"`
		Url                  string `json:"url"`
		Month                string `json:"month"`
		ProductReferenceCode string `json:"productReferenceCode"`
	}
	ReqPaymentIak struct {
		ProductCode string `json:"product_code"`
		CustomerId  string `json:"customer_id"`
		RefId       string `json:"ref_id"`
	}
	ReqCheckStatusPostpaidIak struct {
		Commands string `json:"commands"`
		Username string `json:"username"`
		Sign     string `json:"sign"`
		RefId    string `json:"ref_id"`
	}
	ReqInquiryPostpaidIak struct {
		Commands string `json:"commands"`
		Hp       string `json:"hp"`
		Code     string `json:"code"`
		RefId    string `json:"ref_id"`
		Username string `json:"username"`
		Sign     string `json:"sign"`
		Month    string `json:"month"`
	}
	ReqPaymentPostpaidIak struct {
		Commands string `json:"commands"`
		Username string `json:"username"`
		TrID     string `json:"tr_id"`
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
	RespPaymentPLNPostpaidIak struct {
		Data struct {
			TrID         int    `json:"tr_id"`
			Code         string `json:"code"`
			Datetime     string `json:"datetime"`
			Hp           string `json:"hp"`
			TrName       string `json:"tr_name"`
			Period       string `json:"period"`
			Nominal      int    `json:"nominal"`
			Admin        int    `json:"admin"`
			ResponseCode string `json:"response_code"`
			Message      string `json:"message"`
			Price        int    `json:"price"`
			SellingPrice int    `json:"selling_price"`
			Balance      int    `json:"balance"`
			Noref        string `json:"noref"`
			RefID        string `json:"ref_id"`
			Desc         struct {
				Tarif             string `json:"tarif"`
				Daya              int    `json:"daya"`
				LembarTagihan     string `json:"lembar_tagihan"`
				LembarTagihanSisa int    `json:"lembar_tagihan_sisa"`
				Tagihan           struct {
					Detail []struct {
						MeterAwal    string `json:"meter_awal"`
						MeterAkhir   string `json:"meter_akhir"`
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
	RespCheckStatusPostpaidIak struct {
		Data struct {
			TrID         int         `json:"tr_id"`
			Code         string      `json:"code"`
			Datetime     string      `json:"datetime"`
			Hp           string      `json:"hp"`
			TrName       string      `json:"tr_name"`
			Period       string      `json:"period"`
			Nominal      int         `json:"nominal"`
			Admin        int         `json:"admin"`
			Status       int         `json:"status"`
			ResponseCode string      `json:"response_code"`
			Message      string      `json:"message"`
			Price        int         `json:"price"`
			SellingPrice int         `json:"selling_price"`
			Balance      int         `json:"balance"`
			Noref        string      `json:"noref"`
			RefID        string      `json:"ref_id"`
			Desc         interface{} `json:"desc"`
		} `json:"data"`
		Meta []interface{} `json:"meta"`
	}
	RespCheckStatusPrepaidIak struct {
		Data struct {
			RefID       string `json:"ref_id"`
			Status      int    `json:"status"`
			ProductCode string `json:"product_code"`
			CustomerID  string `json:"customer_id"`
			Price       int    `json:"price"`
			Message     string `json:"message"`
			Sn          string `json:"sn"`
			Balance     int    `json:"balance"`
			TrID        int    `json:"tr_id"`
			Rc          string `json:"rc"`
		} `json:"data"`
	}

	//bpjs
	RespInquiryBPJSIak struct {
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
				KodeCabang     string `json:"kode_cabang"`
				NamaCabang     string `json:"nama_cabang"`
				SisaPembayaran string `json:"sisa_pembayaran"`
				JumlahPeserta  string `json:"jumlah_peserta"`
			} `json:"desc"`
		} `json:"data"`
		Meta []interface{} `json:"meta"`
	}
	RespPaymentBPJSIak struct {
		Data struct {
			TrID         int    `json:"tr_id"`
			Code         string `json:"code"`
			Datetime     string `json:"datetime"`
			Hp           string `json:"hp"`
			TrName       string `json:"tr_name"`
			Period       string `json:"period"`
			Nominal      int    `json:"nominal"`
			Admin        int    `json:"admin"`
			ResponseCode string `json:"response_code"`
			Message      string `json:"message"`
			Price        int    `json:"price"`
			SellingPrice int    `json:"selling_price"`
			Balance      int    `json:"balance"`
			Noref        string `json:"noref"`
			RefID        string `json:"ref_id"`
			Desc         struct {
				KodeCabang     string `json:"kode_cabang"`
				NamaCabang     string `json:"nama_cabang"`
				SisaPembayaran string `json:"sisa_pembayaran"`
				JumlahPeserta  string `json:"jumlah_peserta"`
			} `json:"desc"`
		} `json:"data"`
		Meta []interface{} `json:"meta"`
	}
)

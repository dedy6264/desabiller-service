package models

type (
	ReqPaymentPrepaidIak struct {
		CustomerId  string `json:"customer_id"`
		ProductCode string `json:"product_code"`
		RefId       string `json:"ref_id"`
		Username    string `json:"username"`
		Sign        string `json:"sign"`
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
		} `json:"data"`
	}
)

package models

type ProviderResponse struct {
	ResponseCode     string `json:"responseCode"`
	ResponseMessage  string `json:"responseMessage"`
	ResponseDateTime string `json:"responseDateTime"`
	Result           struct {
		ProductID               int     `json:"product_id"`
		ProductName             string  `json:"product_name"`
		ProductCode             string  `json:"product_code"`
		ProductPrice            float64 `json:"product_price"`
		ProductAdminFee         int     `json:"product_admin_fee"`
		ProductMerchantFee      int     `json:"product_merchant_fee"`
		TransactionTotalAmount  int     `json:"transaction_total_amount"`
		ReferenceNumber         string  `json:"reference_number"`
		ReferenceNumberMerchant string  `json:"reference_number_merchant"`
		StatusCode              string  `json:"status_code"`
		StatusMessage           string  `json:"status_message"`
		StatusDesc              string  `json:"status_desc"`
		CustomerID              string  `json:"customer_id"`
		MerchantID              int     `json:"merchant_id"`
		MerchantName            string  `json:"merchant_name"`
		MerchantOutletName      string  `json:"merchant_outlet_name"`
		CreatedAt               string  `json:"createdAt"`
		UpdatedAt               string  `json:"updatedAt"`
		BillInfo                struct {
			Details  string `json:"details"`
			BillDesc string `json:"bill_desc"`
			Sn       string `json:"sn"`
			LembTag  int    `json:"lemb_tag"`
		} `json:"bill_info"`
	} `json:"result"`
}

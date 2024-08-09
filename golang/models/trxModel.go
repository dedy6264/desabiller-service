package models

type (
	//inq
	ReqInqTrx struct {
		PaymentMethodId         int              `json:"paymentMethodId"`
		PaymentMethodCategoryId int              `json:"paymentMethodCategoryId"`
		ProductTypeId           int              `json:"productTypeId"`
		ProductDetails          []ProductDetails `json:"productDetails"`
	}
	ProductDetails struct {
		ProductId    int     `json:"productId"`
		ProductName  string  `json:"productName"`
		ProductCode  string  `json:"productCode"`
		ProductPrice float64 `json:"productPrice"`
		Qty          int     `json:"qty"`
		CustomerId   string  `json:"customerId"`
	}
	ReqPaymentTrx struct {
		TrxNumber         string `json:"trxNumber"`
		PaymentMethodId   int    `json:"paymentMethodId"`
		PaymentMethodName string `json:"paymentMethodName"`
	}
	ResPaymentTrx struct {
		TrxNumber            string  `json:"trxNumber"`
		TrxNumberPartner     string  `json:"trxNumberPartner"`
		PaymentAt            string  `json:"paymentAt"`
		CreatedBy            string  `json:"createdBy"`
		CreatedAt            string  `json:"createdAt"`
		UpdatedBy            string  `json:"updatedBy"`
		UpdatedAt            string  `json:"updatedAt"`
		StatusCode           string  `json:"statusCode"`
		StatusMessage        string  `json:"statusMessage"`
		StatusDesc           string  `json:"statusDesc"`
		StatusCodePartner    string  `json:"statusCodePartner"`
		StatusMessagePartner string  `json:"statusMessagePartner"`
		StatusDescPartner    string  `json:"statusDescPartner"`
		SegmentId            int     `json:"segmentId"`
		ProductTypeId        int     `json:"productTypeId"`
		ProductTypeName      string  `json:"productTypeName"`
		ProductCategoryId    int     `json:"productCategoryId"`
		ProductCategoryName  string  `json:"productCategoryName"`
		ProductId            int     `json:"productId"`
		ProductCode          string  `json:"productCode"`
		ProductName          string  `json:"productName"`
		ProductPrice         float64 `json:"productPrice"`
		ProductAdminFee      float64 `json:"productAdminFee"`
		ProductMerchantFee   float64 `json:"productMerchantFee"`
		Quantity             int     `json:"quantity"`
		SubTotal             float64 `json:"subTotal"`
		GrandTotal           float64 `json:"grandTotal"`
		CustomerId           string  `json:"customerId"`
		BillInfo             string  `json:"billInfo"`
		ClientId             int     `json:"clientId"`
		ClientName           string  `json:"clientName"`
		MerchantId           int     `json:"merchantId"`
		MerchantName         string  `json:"merchantName"`
		MerchantOutletId     int     `json:"merchantOutletId"`
		MerchantOutletName   string  `json:"merchantOutletName"`
		UserOutletId         int     `json:"userOutletId"`
		UserOutletName       string  `json:"userOutletName"`
		OutletDeviceId       int     `json:"outletDeviceId"`
		OutletDeviceType     string  `json:"outletDeviceType"`
		OutletDeviceSn       string  `json:"outletDeviceSn"`
		PaymentMethodId      int     `json:"paymentMethodId"`
		PaymentMethodName    string  `json:"paymentMethodName"`
	}
	ReqUpdateTrx struct {
		Id                   int    `json:"Id"`
		TrxNumber            string `json:"trxNumber"`
		TrxNumberPartner     string `json:"trxNumberPartner"`
		PaymentAt            string `json:"paymentAt"`
		PaymentMethodId      int    `json:"paymentMethodId"`
		PaymentMethodName    string `json:"paymentMethodName"`
		UpdatedBy            string `json:"updatedBy"`
		UpdatedAt            string `json:"updatedAt"`
		StatusCode           string `json:"statusCode"`
		StatusMessage        string `json:"statusMessage"`
		StatusDesc           string `json:"statusDesc"`
		StatusCodePartner    string `json:"statusCodePartner"`
		StatusMessagePartner string `json:"statusMessagePartner"`
		StatusDescPartner    string `json:"statusDescPartner"`
		BillInfo             string `json:"billInfo"`
	}
	ReqInsertTrxDetails struct {
		ID                  int     `json:"id"`
		ProductTypeId       int     `json:"productTypeId"`
		ProductTypeName     string  `json:"productTypeName"`
		ProductCategoryId   int     `json:"productCategoryId"`
		ProductCategoryName string  `json:"productCategoryName"`
		ProductId           int     `json:"productId"`
		ProductCode         string  `json:"productCode"`
		ProductName         string  `json:"productName"`
		ProductPrice        float64 `json:"productPrice"`
		Quantity            int     `json:"quantity"`
		CustomerId          string  `json:"customerId"`
		BillInfo            string  `json:"billInfo"`
	}
	ReqInsertTrx struct {
		TrxNumber            string  `json:"trxNumber"`
		TrxNumberPartner     string  `json:"trxNumberPartner"`
		PaymentAt            string  `json:"paymentAt"`
		CreatedBy            string  `json:"createdBy"`
		CreatedAt            string  `json:"createdAt"`
		UpdatedBy            string  `json:"updatedBy"`
		UpdatedAt            string  `json:"updatedAt"`
		StatusCode           string  `json:"statusCode"`
		StatusMessage        string  `json:"statusMessage"`
		StatusDesc           string  `json:"statusDesc"`
		StatusCodePartner    string  `json:"statusCodePartner"`
		StatusMessagePartner string  `json:"statusMessagePartner"`
		StatusDescPartner    string  `json:"statusDescPartner"`
		SegmentId            int     `json:"segmentId"`
		ProductTypeId        int     `json:"productTypeId"`
		ProductTypeName      string  `json:"productTypeName"`
		ProductCategoryId    int     `json:"productCategoryId"`
		ProductCategoryName  string  `json:"productCategoryName"`
		ProductId            int     `json:"productId"`
		ProductCode          string  `json:"productCode"`
		ProductName          string  `json:"productName"`
		ProductPrice         float64 `json:"productPrice"`
		ProductAdminFee      float64 `json:"productAdminFee"`
		ProductMerchantFee   float64 `json:"productMerchantFee"`
		SubTotal             float64 `json:"subTotal"`
		GrandTotal           float64 `json:"grandTotal"`
		CustomerId           string  `json:"customerId"`
		BillInfo             string  `json:"billInfo"`
		ClientId             int     `json:"clientId"`
		ClientName           string  `json:"clientName"`
		MerchantId           int     `json:"merchantId"`
		MerchantName         string  `json:"merchantName"`
		MerchantOutletId     int     `json:"merchantOutletId"`
		MerchantOutletName   string  `json:"merchantOutletName"`
		UserOutletId         int     `json:"userOutletId"`
		UserOutletName       string  `json:"userOutletName"`
		OutletDeviceId       int     `json:"outletDeviceId"`
		OutletDeviceType     string  `json:"outletDeviceType"`
		OutletDeviceSn       string  `json:"outletDeviceSn"`
		PaymentMethodId      int     `json:"paymentMethodId"`
		PaymentMethodName    string  `json:"paymentMethodName"`
		// ProductDetails       string  `json:"productDetails"`
	}

	ReqTrx struct {
		Limit     int    `json:"limit"`
		Offset    int    `json:"offset"`
		OrderBy   string `json:"orderBy"`
		SortBy    string `json:"sortBy"` //asc desc
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`

		Id                 int    `json:"Id"`
		ProductCode        string `json:"productCode"`
		TrxNumber          string `json:"trxNumber"`
		TrxNumberPartner   string `json:"trxNumberPartner"`
		PaymentAt          string `json:"paymentAt"`
		CreatedAt          string `json:"createdAt"`
		UpdatedAt          string `json:"updatedAt"`
		StatusCode         string `json:"statusCode"`
		ClientId           int    `json:"clientId"`
		ClientName         string `json:"clientName"`
		MerchantId         int    `json:"merchantId"`
		MerchantName       string `json:"merchantName"`
		MerchantOutletId   int    `json:"merchantOutletId"`
		MerchantOutletName string `json:"merchantOutletName"`
		UserOutletId       int    `json:"userOutletId"`
		UserOutletName     string `json:"userOutletName"`
		OutletDeviceId     int    `json:"outletDeviceId"`
		OutletDeviceType   string `json:"outletDeviceType"`
		OutletDeviceSn     string `json:"outletDeviceSn"`
		PaymentMethodId    int    `json:"paymentMethodId"`
		PaymentMethodName  string `json:"paymentMethodName"`
		CustomerId         string `json:"customerId"`
	}
	RespTrxList struct {
		Index                int     `json:"Index"`
		Id                   int     `json:"Id"`
		TrxNumber            string  `json:"trxNumber"`
		TrxNumberPartner     string  `json:"trxNumberPartner"`
		PaymentAt            string  `json:"paymentAt"`
		CreatedBy            string  `json:"createdBy"`
		CreatedAt            string  `json:"createdAt"`
		UpdatedBy            string  `json:"updatedBy"`
		UpdatedAt            string  `json:"updatedAt"`
		StatusCode           string  `json:"statusCode"`
		StatusMessage        string  `json:"statusMessage"`
		StatusDesc           string  `json:"statusDesc"`
		StatusCodePartner    string  `json:"statusCodePartner"`
		StatusMessagePartner string  `json:"statusMessagePartner"`
		StatusDescPartner    string  `json:"statusDescPartner"`
		SegmentId            int     `json:"segmentId"`
		ProductTypeId        int     `json:"productTypeId"`
		ProductTypeName      string  `json:"productTypeName"`
		ProductCategoryId    int     `json:"productCategoryId"`
		ProductCategoryName  string  `json:"productCategoryName"`
		ProductId            int     `json:"productId"`
		ProductCode          string  `json:"productCode"`
		ProductName          string  `json:"productName"`
		ProductPrice         float64 `json:"productPrice"`
		ProductAdminFee      float64 `json:"productAdminFee"`
		ProductMerchantFee   float64 `json:"productMerchantFee"`
		Quantity             int     `json:"quantity"`
		SubTotal             float64 `json:"subTotal"`
		GrandTotal           float64 `json:"grandTotal"`
		CustomerId           string  `json:"customerId"`
		BillInfo             string  `json:"billInfo"`
		ClientId             int     `json:"clientId"`
		ClientName           string  `json:"clientName"`
		MerchantId           int     `json:"merchantId"`
		MerchantName         string  `json:"merchantName"`
		MerchantOutletId     int     `json:"merchantOutletId"`
		MerchantOutletName   string  `json:"merchantOutletName"`
		UserOutletId         int     `json:"userOutletId"`
		UserOutletName       string  `json:"userOutletName"`
		OutletDeviceId       int     `json:"outletDeviceId"`
		OutletDeviceType     string  `json:"outletDeviceType"`
		OutletDeviceSn       string  `json:"outletDeviceSn"`
		PaymentMethodId      int     `json:"paymentMethodId"`
		PaymentMethodName    string  `json:"paymentMethodName"`
	}
)

// id
// trx_number
// trx_number_partner
// payment_at
// created_by
// created_at
// updated_by
// updated_at
// status_code
// status_message
// status_desc
// status_code_partner
// status_message_partner
// status_desc_partner
// segment_id
// product_type_id
// product_type_name
// product_category_id
// product_category_name
// product_id
// product_name
// product_price
// product_admin_fee
// product_merchant_fee
// quantity
// sub_total
// grand_total
// customer_id
// bill_info
// client_id
// client_name
// merchant_id
// merchant_name
// merchant_outlet_id
// merchant_outlet_name
// user_outlet_id
// user_outlet_name
// outlet_device_id
// outlet_device_type
// outlet_device_sn
// payment_method_id
// payment_method_name

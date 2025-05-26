package models

type (
	ResponseWorkerPayment struct {
		PaymentStatus              string                 `json:"paymentStatus"`
		PaymentStatusDesc          string                 `json:"paymentStatusDesc"`
		PaymentStatusMsg           string                 `json:"paymentStatusMsg"`
		PaymentStatusDetail        string                 `json:"paymentStatusDetail"`
		PaymentStatusDescDetail    string                 `json:"paymentStatusDescDetail"`
		TrxReferenceNumber         string                 `json:"trxReferenceNumber"`
		TrxProviderReferenceNumber string                 `json:"trxProviderReferenceNumber"`
		TotalTrxAmount             float64                `json:"totalTrxAmount"`
		TrxAmount                  float64                `json:"trxAmount"`
		AdminFee                   float64                `json:"adminFee"`
		BillInfo                   map[string]interface{} `json:"billInfo"`
		// PaymentDetail              PaymentDetails         `json:"paymentDetail"`
	}
	ResponseWorkerInquiry struct {
		InquiryStatus              string                 `json:"inquiryStatus"`
		InquiryStatusDesc          string                 `json:"inquiryStatusDesc"`
		InquiryStatusMsg           string                 `json:"inquiryStatusMsg"`
		InquiryStatusDetail        string                 `json:"inquiryStatusDetail"`
		InquiryStatusDescDetail    string                 `json:"inquiryStatusDescDetail"`
		TrxReferenceNumber         string                 `json:"trxReferenceNumber"`
		TrxProviderReferenceNumber string                 `json:"trxProviderReferenceNumber"`
		TotalTrxAmount             float64                `json:"totalTrxAmount"`
		TrxAmount                  float64                `json:"trxAmount"`
		SubscriberName             string                 `json:"subscriberName"`
		SubscriberNumber           string                 `json:"subscriberNumber"`
		AdminFee                   float64                `json:"adminFee"`
		BillInfo                   map[string]interface{} `json:"billInfo"`
		// InquiryDetail              InquiryDetail          `json:"inquiryDetail"`
	}
	// PaymentDetails struct {
	// 	Price       float64 `json:"price"`
	// 	AdminFee    float64 `json:"adminFee"`
	// 	MerchantFee float64 `json:"merchantFee"`
	// }

	// BillDescPulsa struct {
	// 	CustomerId string `json:"customerId"`
	// 	Sn         string `json:"sn"`
	// }

	// BillDesc struct {
	// 	Total
	// }

	// InquiryDetail struct {
	// 	Price       float64 `json:"price"`
	// 	AdminFee    float64 `json:"adminFee"`
	// 	MerchantFee float64 `json:"merchantFee"`
	// }
	BillDescPLN struct {
		SubscriberName   string              `json:"subscriberName"`
		SubscriberNumber string              `json:"subscriberNumber"`
		MeterNo          string              `json:"meterNo"`
		LembarTagihan    int                 `json:"lembarTagihan"`
		Detail           []DetailBillDescPLN `json:"detail"`
	}
	DetailBillDescPLN struct {
		Periode    string  `json:"periode"`
		Admin      float64 `json:"admin"`
		Denda      float64 `json:"denda"`
		Tagihan    float64 `json:"tagihan"`
		MeterAwal  string  `json:"meterAwal"`
		MeterAkhir string  `json:"meterAkhir"`
		Tarif      string  `json:"tarif"`
		Daya       string  `json:"daya"`
	}
	//bpjs
	BillDescBPJS struct {
		CustomerId   string               `json:"customerId"`
		CustomerName string               `json:"customerName"`
		Detail       []DetailBillDescBPJS `json:"detail"`
	}
	DetailBillDescBPJS struct {
		Periode    string  `json:"periode"`
		Admin      float64 `json:"admin"`
		Denda      float64 `json:"denda"`
		Tagihan    float64 `json:"tagihan"`
		JmlPeserta string  `json:"jmlPeserta"`
	}
)

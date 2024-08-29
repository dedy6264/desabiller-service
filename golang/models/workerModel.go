package models

type (
	ResponseWorkerPayment struct {
		PaymentStatus              string                 `json:"paymentStatus"`
		PaymentStatusDesc          string                 `json:"paymentStatusDesc"`
		PaymentStatusDetail        string                 `json:"paymentStatusDetail"`
		PaymentStatusDescDetail    string                 `json:"paymentStatusDescDetail"`
		TrxReferenceNumber         string                 `json:"trxReferenceNumber"`
		TrxProviderReferenceNumber string                 `json:"trxProviderReferenceNumber"`
		TotalTrxAmount             float64                `json:"totalTrxAmount"`
		BillInfo                   map[string]interface{} `json:"billInfo"`
		PaymentDetail              PaymentDetails         `json:"paymentDetail"`
	}
	ResponseWorkerInquiry struct {
		InquiryStatus              string                 `json:"inquiryStatus"`
		InquiryStatusDesc          string                 `json:"inquiryStatusDesc"`
		InquiryStatusDetail        string                 `json:"inquiryStatusDetail"`
		InquiryStatusDescDetail    string                 `json:"inquiryStatusDescDetail"`
		TrxReferenceNumber         string                 `json:"trxReferenceNumber"`
		TrxProviderReferenceNumber string                 `json:"trxProviderReferenceNumber"`
		TotalTrxAmount             float64                `json:"totalTrxAmount"`
		BillInfo                   map[string]interface{} `json:"billInfo"`
		InquiryDetail              InquiryDetail          `json:"inquiryDetail"`
	}
	PaymentDetails struct {
		Price       float64 `json:"price"`
		AdminFee    float64 `json:"adminFee"`
		MerchantFee float64 `json:"merchantFee"`
	}

	// BillDescPulsa struct {
	// 	CustomerId string `json:"customerId"`
	// 	Sn         string `json:"sn"`
	// }

	// BillDesc struct {
	// 	Total
	// }

	InquiryDetail struct {
		Price       float64 `json:"price"`
		AdminFee    float64 `json:"adminFee"`
		MerchantFee float64 `json:"merchantFee"`
	}
	BillDescPLN struct {
		CustomerId    string              `json:"customerId"`
		Tarif         float64             `json:"tarif"`
		Daya          string              `json:"daya"`
		LembarTagihan int                 `json:"lembar_tagihan"`
		Detail        []DetailBillDescPLN `json:"detail"`
	}
	DetailBillDescPLN struct {
		Periode string  `json:"periode"`
		Admin   float64 `json:"admin"`
		Denda   float64 `json:"denda"`
		Tagihan float64 `json:"tagihan"`
	}
)

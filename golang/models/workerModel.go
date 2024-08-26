package models

type (
	ResponseWorkerPayment struct {
		PaymentStatus              string  `json:"paymentStatus"`
		PaymentStatusDesc          string  `json:"paymentStatusDesc"`
		PaymentStatusDetail        string  `json:"paymentStatusDetail"`
		PaymentStatusDescDetail    string  `json:"paymentStatusDescDetail"`
		TrxReferenceNumber         string  `json:"trxReferenceNumber"`
		TrxProviderReferenceNumber string  `json:"trxProviderReferenceNumber"`
		TotalAmount                float64 `json:"totalAmount"`
		BillDesc                   string  `json:"billDesc"`
	}
	BillDescPulsa struct {
		CustomerId string `json:"customerId"`
		Sn         string `json:"sn"`
	}
	// BillDesc struct {
	// 	Total
	// }
)

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
	// BillDesc struct {
	// 	Total
	// }
)

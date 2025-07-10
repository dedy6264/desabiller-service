package models

type (
	ReqGetListPaymentMethodCategory struct {
		Id                        int    `json:"id"`
		PaymentMethodCategoryName string `json:"paymentMethodCategoryName"`
		Limit                     int    `json:"limit"`
		OrderBy                   string `json:"orderBy"`
		Offset                    int    `json:"offset"`
		AscDesc                   string `json:"ascDesc"`
		Draw                      int    `json:"draw"`
	}
	ReqGetListPaymentMethod struct {
		Id                      int    `json:"id"`
		PaymentMethodName       string `json:"paymentMethodName"`
		PaymentMethodCategoryId int    `json:"paymentMethodCategoryId"`
		Limit                   int    `json:"limit"`
		OrderBy                 string `json:"orderBy"`
		Offset                  int    `json:"offset"`
		AscDesc                 string `json:"ascDesc"`
		Draw                    int    `json:"draw"`
	}
	ResPaymentMethod struct {
		Id                      int    `json:"id"`
		PaymentMethodName       string `json:"paymentMethodName"`
		PaymentMethodCategoryId int    `json:"paymentMethodCategoryId"`
	}
	ResPaymentMethodCategory struct {
		Id                        int    `json:"id"`
		PaymentMethodCategoryName string `json:"paymentMethodCategoryName"`
	}
)

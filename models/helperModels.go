package models

type (
	GetPrefix struct {
		SubscriberId string `json:"subscriberId" validate:"required"`
	}
	RespGetPrefix struct {
		ProductReferenceId   int    `json:"productReferenceId"`
		ProductReferenceCode string `json:"productReferenceCode"`
		ProductReferenceName string `json:"productReferenceName"`
	}
)

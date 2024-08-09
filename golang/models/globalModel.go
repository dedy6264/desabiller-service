package models

type Response struct {
	StatusCode       string      `json:"statusCode"`
	Success          bool        `json:"success"`
	ResponseDatetime string      `json:"responseDatetime"`
	Result           interface{} `json:"result"`
	Message          string      `json:"message"`
}
type ResponseList struct {
	TotalData int         `json:"totalData"`
	TotalRow  int         `json:"totalRow"`
	Data      interface{} `json:"result"`
}

package models

type Response struct {
	StatusCode       string      `json:"statusCode"`
	StatusMessage    string      `json:"statusMessage"`
	ResponseDatetime string      `json:"responseDatetime"`
	Result           interface{} `json:"result"`
}
type ResponseList struct {
	TotalData int         `json:"totalData"`
	TotalRow  int         `json:"totalRow"`
	Data      interface{} `json:"result"`
}

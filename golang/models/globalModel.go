package models

type Response struct {
	StatusCode       string      `json:"statusCode"`
	ResponseDatetime string      `json:"responseDatetime"`
	Result           interface{} `json:"result"`
	StatusMessage    string      `json:"statusMessage"`
}
type ResponseList struct {
	TotalData int         `json:"totalData"`
	TotalRow  int         `json:"totalRow"`
	Data      interface{} `json:"result"`
}

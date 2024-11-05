package models

type Response struct {
	StatusCode       string      `json:"statusCode"`
	StatusMessage    string      `json:"statusMessage"`
	ResponseDatetime string      `json:"responseDatetime"`
	Result           interface{} `json:"result"`
}
type ResponseList struct {
	Draw            int         `json:"draw"`
	RecordsTotal    int         `json:"recordsTotal"`
	RecordsFiltered int         `json:"recordsFiltered"`
	Data            interface{} `json:"data"`
}

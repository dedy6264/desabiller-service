package models

import "time"

type Response struct {
	StatusCode       string      `json:"statusCode"`
	StatusMessage    string      `json:"statusMessage"`
	StatusDesc       string      `json:"statusDesc"`
	ResponseDatetime string      `json:"responseDatetime"`
	Result           interface{} `json:"result"`
}
type ResponseList struct {
	Draw            int         `json:"draw"`
	RecordsTotal    int         `json:"recordsTotal"`
	RecordsFiltered int         `json:"recordsFiltered"`
	Data            interface{} `json:"data"`
}
type GetToken struct {
	ResponseCode     string    `json:"responseCode"`
	ResponseMessage  string    `json:"responseMessage"`
	ResponseDateTime time.Time `json:"responseDateTime"`
	Result           struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   string `json:"expires_in"`
	} `json:"result"`
}

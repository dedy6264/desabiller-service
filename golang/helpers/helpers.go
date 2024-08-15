package helpers

import (
	"desabiller/models"
	"time"
)

func ResponseJSON(success bool, code string, msg string, result interface{}) models.Response {
	dbTime := time.Now().Format(time.RFC3339)
	response := models.Response{
		StatusCode:       code,
		Result:           result,
		StatusMessage:    msg,
		ResponseDatetime: dbTime,
	}
	return response
}

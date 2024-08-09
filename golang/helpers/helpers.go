package helpers

import (
	"desabiller/models"
	"time"
)

func ResponseJSON(success bool, code string, msg string, result interface{}) models.Response {
	dbTime := time.Now().Format(time.RFC3339)
	response := models.Response{
		Success:          success,
		StatusCode:       code,
		Result:           result,
		Message:          msg,
		ResponseDatetime: dbTime,
	}
	return response
}

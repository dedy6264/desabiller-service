package helpers

import (
	"desabiller/models"
	"reflect"
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
func InArray(v interface{}, in interface{}) (ok bool, i int) {
	val := reflect.Indirect(reflect.ValueOf(in))
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for ; i < val.Len(); i++ {
			if ok = v == val.Index(i).Interface(); ok {
				return
			}
		}
	}
	return
}

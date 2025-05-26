package helpers

import (
	"desabiller/models"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
)

func ErrLogger(svc, desc string, err error) {
	log.Println("Error: ", svc, " | ", desc, " | ", err)
}
func ResponseJSON(success bool, code string, msg, desc string, result interface{}) models.Response {
	dbTime := time.Now().Format(time.RFC3339)
	response := models.Response{
		StatusCode:       code,
		Result:           result,
		StatusMessage:    msg,
		StatusDesc:       desc,
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
func JsonDescape(str string) string {

	stringifiedData := strings.ReplaceAll(str, `\n`, ``)
	stringifiedData = strings.ReplaceAll(stringifiedData, `\\"`, `\"`)
	return stringifiedData
}
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

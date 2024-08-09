package helpers

import (
	"log"

	"github.com/labstack/echo"
)

//	func Binding(req interface{}, ctx echo.Context) (status bool) {
//		err := ctx.Bind(req)
//		if err != nil {
//			log.Println("ERROR Binding :: ", err.Error())
//			return false
//		}
//		return true
//	}
func BindValidate(req interface{}, ctx echo.Context) (status bool, err error) {

	err = ctx.Bind(req)
	if err != nil {
		log.Println("ERROR Binding :: ", err.Error())
		return false, err
	}
	err = ctx.Validate(req)
	if err != nil {
		log.Println("ERROR Validate :: ", err.Error())
		return false, err
	}
	return true, nil
}

package vueservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"net/http"

	"github.com/labstack/echo"
)

func (svc VueService) Home(ctx echo.Context) error {
	var (
	// svcName = " Home "
	// errSvc  = " Error " + svcName
	)
	type carouselFoto struct {
		Id       int    `json"id"`
		PhotoUrl string `json"photoUrl"`
		Tittle   string `json"tittle"`
	}
	// type alur struct {
	// 	Id
	// 	PhotoUrl
	// 	Tittle
	// }
	// type carouselProject struct {
	// 	Id
	// 	PhotoList
	// 	Tittle
	// 	Desc
	// }
	// type galeries struct {
	// 	Id
	// 	PhotoUrl

	// }
	// type footer struct {
	// }

	// for i := 0; i < 3; i++ {
	// 	resp := carouselFoto{
	// 		Id:       i,
	// 		PhotoUrl: "",
	// 		Tittle:   "",
	// 	}
	// }
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", nil)
	return ctx.JSON(http.StatusOK, result)
}

// "statusCode": "00",
// "success": true,
// "responseDatetime": "2023-12-14T12:28:24+07:00",
// "result": {
// 	"carouselFoto":{[
// 		"id":"",
// 		"photoUrl":"",
// 		"tittle":"",
// 	],[
// 		"id":"",
// 		"photoUrl":"",
// 		"tittle":"",
// 	]},
// 	"alur":{[
// 		"id":"",
// 		"photoUrl":"",
// 		"tittle":"",
// 	],[
// 		"id":"",
// 		"photoUrl":"",
// 		"tittle":"",
// 	]}
// }

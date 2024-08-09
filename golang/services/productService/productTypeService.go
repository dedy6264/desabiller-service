package productservice

import (
	"desabiller/configs"
	"desabiller/models"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (svc ProductService) GetListProductType(ctx echo.Context) error {
	var (
		svcName    = "GetListProductType"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	resProTy, _ := svc.service.ApiProduct.GetListProductType()
	if len(resProTy) == 0 {
		log.Println("Error " + svcName + " // not found :: ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}

	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Message = "SUCCESS"
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	respGlobal.Result = resProTy
	return ctx.JSON(http.StatusOK, respGlobal)
}

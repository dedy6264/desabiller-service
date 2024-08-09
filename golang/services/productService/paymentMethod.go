package productservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

func (svc ProductService) AddPaymentMethod(ctx echo.Context) error {
	var (
		svcName    = "AddPaymentMethod"
		logErr     = "Error " + svcName
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListPaymentMethod)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.PaymentMethodName == "" {
		log.Println(logErr + " PaymentMethodName ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "PaymentMethodName"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.PaymentMethodCategoryId == 0 {
		log.Println(logErr + " PaymentMethodCategoryId ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "PaymentMethodCategoryId"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	req.PaymentMethodName = strings.ToUpper(req.PaymentMethodName)

	res, status := svc.service.ApiPayment.AddPaymentMethod(*req)
	if !status {
		log.Println(logErr + " AddPaymentMethod ")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "Add Failed"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	res.PaymentMethodCategoryId = req.PaymentMethodCategoryId
	res.PaymentMethodName = res.PaymentMethodName
	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	respGlobal.Message = "Success"
	respGlobal.Result = res
	return ctx.JSON(http.StatusOK, respGlobal)
}
func (svc ProductService) GetListPaymentMethod(ctx echo.Context) error {
	var (
		svcName    = "GetListPaymentMethod"
		logErr     = "Error " + svcName
		respGlobal models.Response
		resSvc     models.ResponseList
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListPaymentMethod)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	res, status := svc.service.ApiPayment.GetListPaymentMethod(*req)
	if !status {
		log.Println(logErr + " GetListPaymentMethod")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "Failed"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	resSvc.Data = res
	respGlobal.Result = resSvc
	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	respGlobal.Message = "Success"
	return ctx.JSON(http.StatusOK, respGlobal)
}
func (svc ProductService) UpdatePaymentMethod(ctx echo.Context) error {
	var (
		svcName    = "UpdatePaymentMethod"
		logErr     = "Error " + svcName
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListPaymentMethod)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.PaymentMethodName == "" {
		log.Println(logErr + " PaymentMethodName ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "PaymentMethodName"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.PaymentMethodCategoryId == 0 {
		log.Println(logErr + " PaymentMethodCategoryId ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "PaymentMethodCategoryId"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	req.PaymentMethodName = strings.ToUpper(req.PaymentMethodName)
	res, status := svc.service.ApiPayment.UpdatePaymentMethod(*req)
	if !status {
		log.Println(logErr + " UpdatePaymentMethod")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "Failed"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	respGlobal.Result = res
	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	respGlobal.Message = "Success"
	return ctx.JSON(http.StatusOK, respGlobal)
}
func (svc ProductService) DropPaymentMethod(ctx echo.Context) error {
	var (
		svcName    = "UpdatePaymentMethod"
		logErr     = "Error " + svcName
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListPaymentMethod)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.Id == 0 {
		log.Println(logErr + " ID ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "Failed"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	status = svc.service.ApiPayment.DropPaymentMethod(req.Id)
	if !status {
		log.Println(logErr + " DropPaymentMethod")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "Failed"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	respGlobal.Message = "Success"
	return ctx.JSON(http.StatusOK, respGlobal)
}

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

func (svc ProductService) AddPaymentMethodCategory(ctx echo.Context) error {
	var (
		svcName    = "AddPaymentMethodCategory"
		logErr     = "Error " + svcName
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListPaymentMethodCategory)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.PaymentMethodCategoryName == "" {
		log.Println(logErr + " PaymentMethodCategoryName ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "PaymentMethodCategoryName"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	req.PaymentMethodCategoryName = strings.ToUpper(req.PaymentMethodCategoryName)
	res, status := svc.service.ApiPayment.AddPaymentMethodCategory(*req)
	if !status {
		log.Println(logErr + " Add Failed ")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "Add Failed"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	res.PaymentMethodCategoryName = req.PaymentMethodCategoryName
	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	respGlobal.Message = "Success"
	respGlobal.Result = res
	return ctx.JSON(http.StatusOK, respGlobal)
}
func (svc ProductService) GetListPaymentMethodCategory(ctx echo.Context) error {
	var (
		svcName    = "GetListPaymentMethodCategory"
		logErr     = "Error " + svcName
		respGlobal models.Response
		// respList   models.ResponseList
		dbTime = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListPaymentMethodCategory)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	// if req.Draw == 1 {
	// 	req.Offset = (req.Limit * req.Draw) - req.Limit
	// }
	res, status := svc.service.ApiPayment.GetListPaymentMethodCategory(*req)
	if !status {
		log.Println(logErr + " GetListPaymentMethodCategory")
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
func (svc ProductService) UpdatePaymentMethodCategory(ctx echo.Context) error {
	var (
		svcName    = "UpdatePaymentMethodCategory"
		logErr     = "Error " + svcName
		respGlobal models.Response
		// respList   models.ResponseList
		dbTime = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListPaymentMethodCategory)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.Id == 0 {
		log.Println(logErr + " Id")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "Failed"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.PaymentMethodCategoryName == "" {
		log.Println(logErr + " PaymentMethodCategoryName")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "Failed"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	req.PaymentMethodCategoryName = strings.ToUpper(req.PaymentMethodCategoryName)

	res, status := svc.service.ApiPayment.UpdatePaymentMethodCategory(*req)
	if !status {
		log.Println(logErr + " UpdatePaymentMethodCategory")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "Failed"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	res.Id = req.Id
	res.PaymentMethodCategoryName = req.PaymentMethodCategoryName
	respGlobal.Result = res
	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	respGlobal.Message = "Success"
	return ctx.JSON(http.StatusOK, respGlobal)
}
func (svc ProductService) DropPaymentMethodCategory(ctx echo.Context) error {
	var (
		svcName    = "DropPaymentMethodCategory"
		logErr     = "Error " + svcName
		respGlobal models.Response
		// respList   models.ResponseList
		dbTime = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListPaymentMethodCategory)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	status = svc.service.ApiPayment.DropPaymentMethodCategory(req.Id)
	if !status {
		log.Println(logErr + " DropPaymentMethodCategory")
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

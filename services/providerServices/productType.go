package providerservices

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func (svc providerServices) AddProductType(ctx echo.Context) error {
	var (
		svcName = "AddProductType"
	)
	req := new(models.ReqGetProductType)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ProductTypeName == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Product Type name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ProductTypeName = strings.ToUpper(req.ProductTypeName)
	_, err = svc.services.ApiProduct.AddProductType(*req)
	if err != nil {
		log.Println("Err ", svcName, "AddProductType", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"failed",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		"Success",
		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc providerServices) GetProductTypes(ctx echo.Context) error {
	var (
		svcName = "GetProductTypes"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetProductType)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	count, err := svc.services.ApiProduct.GetProductTypeCount(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetProductTypeCount", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.ApiProduct.GetProductTypes(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetProductTypes", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respSvc.Data = resp
	respSvc.RecordsTotal = count
	respSvc.RecordsFiltered = count
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		"Success",
		respSvc)
	return ctx.JSON(http.StatusOK, result)
}
func (svc providerServices) DropProductType(ctx echo.Context) error {
	var (
		svcName = "DropProductType"
	)
	req := new(models.ReqGetProductType)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.ApiProduct.DropProductType(*req)
	if err != nil {
		log.Println("Err ", svcName, "DropProductType", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		"Success",
		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc providerServices) UpdateProductType(ctx echo.Context) error {
	var (
		svcName = "UpdateProductType"
	)
	req := new(models.ReqGetProductType)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ProductTypeName = strings.ToUpper(req.ProductTypeName)
	_, err = svc.services.ApiProduct.UpdateProductType(*req)
	if err != nil {
		log.Println("Err ", svcName, "UpdateProductType", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		"Success",
		nil)
	return ctx.JSON(http.StatusOK, result)
}

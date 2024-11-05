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

func (svc providerServices) AddProvider(ctx echo.Context) error {
	var (
		svcName = "AddProvider"
	)
	req := new(models.ReqGetProvider)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ProviderName == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Provider name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ProviderName = strings.ToUpper(req.ProviderName)
	err = svc.services.ApiProduct.AddProvider(*req)
	if err != nil {
		log.Println("Err ", svcName, "AddProvider", err)
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
func (svc providerServices) DropProvider(ctx echo.Context) error {
	var (
		svcName = "DropProvider"
	)
	req := new(models.ReqGetProvider)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.ApiProduct.DropProvider(*req)
	if err != nil {
		log.Println("Err ", svcName, "DropProvider", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		"Success",
		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc providerServices) UpdateProvider(ctx echo.Context) error {
	var (
		svcName = "UpdateProvider"
	)
	req := new(models.ReqGetProvider)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ProviderName = strings.ToUpper(req.ProviderName)
	_, err = svc.services.ApiProduct.UpdateProvider(*req)
	if err != nil {
		log.Println("Err ", svcName, "UpdateProvider", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		"Success",
		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc providerServices) GetProviders(ctx echo.Context) error {
	var (
		svcName = "GetProviders"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetProvider)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	count, err := svc.services.ApiProduct.GetProviderCount(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetProviderCount", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.ApiProduct.GetProviders(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetProviders", err)
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

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

func (svc providerServices) AddProductClan(ctx echo.Context) error {
	var (
		svcName = "AddProductClan"
	)
	req := new(models.ReqGetProductClan)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ProductClanName == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Product clan name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ProductClanName = strings.ToUpper(req.ProductClanName)
	err = svc.services.ApiProduct.AddProductClan(*req)
	if err != nil {
		log.Println("Err ", svcName, "AddProductClan", err)
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
func (svc providerServices) GetProductClans(ctx echo.Context) error {
	var (
		svcName = "GetProductClans"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetProductClan)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	count, err := svc.services.ApiProduct.GetProductClanCount(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetProductClanCount", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.ApiProduct.GetProductClans(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetProductClans", err)
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
func (svc providerServices) DropProductClan(ctx echo.Context) error {
	var (
		svcName = "DropProductClan"
	)
	req := new(models.ReqGetProductClan)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.ApiProduct.DropProductClan(*req)
	if err != nil {
		log.Println("Err ", svcName, "DropProductClan", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		"Success",
		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc providerServices) UpdateProductClan(ctx echo.Context) error {
	var (
		svcName = "UpdateProductClan"
	)
	req := new(models.ReqGetProductClan)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ProductClanName = strings.ToUpper(req.ProductClanName)
	_, err = svc.services.ApiProduct.UpdateProductClan(*req)
	if err != nil {
		log.Println("Err ", svcName, "UpdateProductClan", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		"Success",
		nil)
	return ctx.JSON(http.StatusOK, result)
}

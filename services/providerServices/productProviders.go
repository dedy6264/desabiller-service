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

func (svc providerServices) AddProductProvider(ctx echo.Context) error {
	var (
		svcName = "AddProductProvider"
	)
	req := new(models.ReqGetProductProvider)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ProviderName == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Provider name is empty",
			"Provider name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ProviderName = strings.ToUpper(req.ProviderName)
	_, err = svc.services.RepoProduct.AddProductProvider(*req)
	if err != nil {
		log.Println("Err ", svcName, "AddProductProvider", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.DB_ERROR,
			"failed",
			"failed",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		configs.SUCCESS_MSG,
		configs.SUCCESS_MSG,

		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc providerServices) DropProductProvider(ctx echo.Context) error {
	var (
		svcName = "DropProductProvider"
	)
	req := new(models.ReqGetProductProvider)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.RepoProduct.DropProductProvider(*req)
	if err != nil {
		log.Println("Err ", svcName, "DropProductProvider", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		configs.SUCCESS_MSG,
		configs.SUCCESS_MSG,

		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc providerServices) UpdateProductProvider(ctx echo.Context) error {
	var (
		svcName = "UpdateProductProvider"
	)
	req := new(models.ReqGetProductProvider)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ProviderName = strings.ToUpper(req.ProviderName)
	_, err = svc.services.RepoProduct.UpdateProductProvider(*req)
	if err != nil {
		log.Println("Err ", svcName, "UpdateProductProvider", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		configs.SUCCESS_MSG,
		configs.SUCCESS_MSG,

		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc providerServices) GetProductProviders(ctx echo.Context) error {
	var (
		svcName = "GetProductProviders"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetProductProvider)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	count, err := svc.services.RepoProduct.GetProductProviderCount(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetProductProviderCount", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.RepoProduct.GetProductProviders(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetProductProviders", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respSvc.Data = resp
	respSvc.RecordsTotal = count
	respSvc.RecordsFiltered = count
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		configs.SUCCESS_MSG,
		configs.SUCCESS_MSG,

		respSvc)
	return ctx.JSON(http.StatusOK, result)
}

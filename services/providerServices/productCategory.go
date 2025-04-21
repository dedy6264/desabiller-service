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

func (svc providerServices) AddProductCategory(ctx echo.Context) error {
	var (
		svcName = "AddProductCategory"
	)
	req := new(models.ReqGetProductCategory)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ProductCategoryName == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Product Category name is empty",
			"Product Category name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ProductCategoryName = strings.ToUpper(req.ProductCategoryName)
	_, err = svc.services.RepoProduct.AddProductCategory(*req)
	if err != nil {
		log.Println("Err ", svcName, "AddProductCategory", err)
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
func (svc providerServices) GetProductCategories(ctx echo.Context) error {
	var (
		svcName = "GetProductCategorys"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetProductCategory)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	count, err := svc.services.RepoProduct.GetProductCategoryCount(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetProductCategoryCount", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.RepoProduct.GetProductCategories(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetProductCategorys", err)
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
func (svc providerServices) DropProductCategory(ctx echo.Context) error {
	var (
		svcName = "DropProductCategory"
	)
	req := new(models.ReqGetProductCategory)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.RepoProduct.DropProductCategory(*req)
	if err != nil {
		log.Println("Err ", svcName, "DropProductCategory", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", "failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		configs.SUCCESS_MSG,
		configs.SUCCESS_MSG,

		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc providerServices) UpdateProductCategory(ctx echo.Context) error {
	var (
		svcName = "UpdateProductCategory"
	)
	req := new(models.ReqGetProductCategory)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ProductCategoryName = strings.ToUpper(req.ProductCategoryName)
	_, err = svc.services.RepoProduct.UpdateProductCategory(*req)
	if err != nil {
		log.Println("Err ", svcName, "UpdateProductCategory", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", "failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		configs.SUCCESS_MSG,
		configs.SUCCESS_MSG,

		nil)
	return ctx.JSON(http.StatusOK, result)
}

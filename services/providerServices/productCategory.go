package providerservices

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/utils"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

func (svc providerServices) AddProductCategory(ctx echo.Context) error {
	var (
		svcName = "AddProductCategory"
		t       = time.Now()
		dbTime  = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqGetProductCategory)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_VALIDATION_FAILED[0], configs.RC_VALIDATION_FAILED[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.ProductCategoryName == "" {
		utils.Log(" ", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.RC_INVALID_PARAM[0],
			"Product Category name is empty",
			"Product Category name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Filter.ProductCategoryName = strings.ToUpper(req.Filter.ProductCategoryName)
	req.Filter.CreatedAt = dbTime
	req.Filter.UpdatedAt = dbTime
	req.Filter.CreatedBy = "sys"
	req.Filter.UpdatedBy = "sys"
	_, err = svc.services.RepoProduct.AddProductCategory(*req)
	if err != nil {
		utils.Log(" AddProductCategory", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			configs.RC_SYSTEM_ERROR[1], nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0],
		configs.RC_SUCCESS[1],
		configs.RC_SUCCESS[1],
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
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_VALIDATION_FAILED[0], configs.RC_VALIDATION_FAILED[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	count, err := svc.services.RepoProduct.GetProductCategoryCount(*req)
	if err != nil {
		utils.Log(" GetProductCategoryCount", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			configs.RC_SYSTEM_ERROR[1], nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.RepoProduct.GetProductCategories(*req)
	if err != nil {
		utils.Log(" GetProductCategories", svcName, err)
		if err == sql.ErrNoRows {
			result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1], respSvc)
			return ctx.JSON(http.StatusOK, result)
		}
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			configs.RC_SYSTEM_ERROR[1], nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respSvc.Data = resp
	respSvc.RecordsTotal = count
	respSvc.RecordsFiltered = count
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0],
		configs.RC_SUCCESS[1],
		configs.RC_SUCCESS[1],
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
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_VALIDATION_FAILED[0], configs.RC_VALIDATION_FAILED[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.RepoProduct.DropProductCategory(*req)
	if err != nil {
		utils.Log(" DropProductCategory", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			configs.RC_SYSTEM_ERROR[1], nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0],
		configs.RC_SUCCESS[1],
		configs.RC_SUCCESS[1],
		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc providerServices) UpdateProductCategory(ctx echo.Context) error {
	var (
		svcName = "UpdateProductCategory"
		t       = time.Now()
		dbTime  = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqGetProductCategory)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_VALIDATION_FAILED[0], configs.RC_VALIDATION_FAILED[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Filter.ProductCategoryName = strings.ToUpper(req.Filter.ProductCategoryName)
	req.Filter.UpdatedAt = dbTime
	req.Filter.UpdatedBy = "sys"
	_, err = svc.services.RepoProduct.UpdateProductCategory(*req)
	if err != nil {
		utils.Log(" UpdateProductCategory", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			configs.RC_SYSTEM_ERROR[1], nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0],
		configs.RC_SUCCESS[1],
		configs.RC_SUCCESS[1],
		nil)
	return ctx.JSON(http.StatusOK, result)
}

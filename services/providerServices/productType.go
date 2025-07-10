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

func (svc providerServices) AddProductType(ctx echo.Context) error {
	var (
		svcName = "AddProductType"
		t       = time.Now()
		dbTime  = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqGetProductType)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_VALIDATION_FAILED[0], configs.RC_VALIDATION_FAILED[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.ProductTypeName == "" {
		utils.Log(" ", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.RC_INVALID_PARAM[0],
			"Product Type name is empty",
			"Product Type name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Filter.ProductTypeName = strings.ToUpper(req.Filter.ProductTypeName)
	req.Filter.CreatedAt = dbTime
	req.Filter.CreatedBy = "sys"
	req.Filter.UpdatedAt = dbTime
	req.Filter.UpdatedBy = "sys"
	_, err = svc.services.RepoProduct.AddProductType(*req)
	if err != nil {
		utils.Log(" AddProductType", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			configs.RC_SYSTEM_ERROR[1],
			nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0],
		configs.RC_SUCCESS[1],
		configs.RC_SUCCESS[1],
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
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_VALIDATION_FAILED[0], configs.RC_VALIDATION_FAILED[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	count, err := svc.services.RepoProduct.GetProductTypeCount(*req)
	if err != nil {
		utils.Log(" GetProductTypeCount", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			configs.RC_SYSTEM_ERROR[1], nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.RepoProduct.GetProductTypes(*req)
	if err != nil {
		utils.Log(" GetProductTypes", svcName, err)
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
	respSvc.RecordsFiltered = len(resp)
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0],
		configs.RC_SUCCESS[1],
		configs.RC_SUCCESS[1],
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
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_VALIDATION_FAILED[0], configs.RC_VALIDATION_FAILED[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.RepoProduct.DropProductType(*req)
	if err != nil {
		utils.Log(" DropProductType", svcName, err)
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
func (svc providerServices) UpdateProductType(ctx echo.Context) error {
	var (
		svcName = "UpdateProductType"
		t       = time.Now()
		dbTime  = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqGetProductType)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_VALIDATION_FAILED[0], configs.RC_VALIDATION_FAILED[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Filter.ProductTypeName = strings.ToUpper(req.Filter.ProductTypeName)
	req.Filter.UpdatedAt = dbTime
	req.Filter.UpdatedBy = "sys"
	_, err = svc.services.RepoProduct.UpdateProductType(*req)
	if err != nil {
		utils.Log(" UpdateProductType", svcName, err)
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

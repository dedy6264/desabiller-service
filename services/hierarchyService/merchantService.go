package hierarchyservice

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func (svc HierarcyService) GetMerchants(ctx echo.Context) error {
	var (
		svcName = "GetMerchants"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetMerchant)
	//binding request
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ClientName = strings.ToUpper(req.ClientName)
	req.GroupName = strings.ToUpper(req.GroupName)
	req.MerchantName = strings.ToUpper(req.MerchantName)
	count, err := svc.service.RepoHierarchy.GetMerchantCount(*req)
	if err != nil {
		log.Println("Err "+svcName+" GetMerchantCount ", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resMerchant, err := svc.service.RepoHierarchy.GetMerchants(*req)
	if err != nil {
		log.Println("Err ", svcName, " GetMerchants ", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respSvc.RecordsTotal = count
	respSvc.RecordsFiltered = count
	respSvc.Data = resMerchant
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "SUCCESS", respSvc)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) AddMerchant(ctx echo.Context) error {
	var (
		svcName = "AddMerchant"
	)
	req := new(models.ReqGetMerchant)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ClientName == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"client name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.GroupName == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"group name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.MerchantName == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"merchant name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ClientName = strings.ToUpper(req.ClientName)
	req.GroupName = strings.ToUpper(req.GroupName)
	req.MerchantName = strings.ToUpper(req.MerchantName)

	_, err = svc.service.RepoHierarchy.GetMerchant(*req)
	if err != nil {
		if err == sql.ErrNoRows {
			err = svc.service.RepoHierarchy.AddMerchant(*req)
			if err != nil {
				log.Println("Err ", svcName, "AddMerchant", err)
				result := helpers.ResponseJSON(configs.FALSE_VALUE,
					configs.VALIDATE_ERROR_CODE,
					"failed",
					nil)
				return ctx.JSON(http.StatusOK, result)
			}
		} else {
			log.Println("Err ", svcName, "GetMerchant", err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE,
				configs.VALIDATE_ERROR_CODE,
				"failed",
				nil)
			return ctx.JSON(http.StatusOK, result)
		}
	} else {
		log.Println("Err ", svcName, "GetMerchant", " Merchant is exist")
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Merchant name is exist",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		"Success",
		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) DropMerchant(ctx echo.Context) error {
	var (
		svcName = "DropMerchant"
	)
	req := new(models.ReqGetMerchant)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.service.RepoHierarchy.DropMerchant(*req)
	if err != nil {
		log.Println("Err ", svcName, "DropMerchant", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"failed",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) UpdateMerchant(ctx echo.Context) error {
	var (
		svcName = "UpdateMerchant"
		// respSvc    models.ResponseList
	)
	req := new(models.ReqGetMerchant)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ClientName = strings.ToUpper(req.ClientName)
	req.GroupName = strings.ToUpper(req.GroupName)
	req.MerchantName = strings.ToUpper(req.MerchantName)

	_, err = svc.service.RepoHierarchy.UpdateMerchant(*req)
	if err != nil {
		log.Println("Err ", svcName, "UpdateMerchant", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"failed",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "SUCCESS", nil)
	return ctx.JSON(http.StatusOK, result)
}

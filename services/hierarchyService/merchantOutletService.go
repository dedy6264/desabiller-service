package hierarchyservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func (svc HierarcyService) GetMerchantOutlets(ctx echo.Context) error {
	var (
		svcName = "GetMerchantOutlets"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetMerchantOutlet)
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
	req.MerchantOutletName = strings.ToUpper(req.MerchantOutletName)
	count, err := svc.service.ApiHierarchy.GetMerchantOutletCount(*req)
	if err != nil {
		log.Println("Err "+svcName+" GetMerchantOutletCount ", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resMerchantOutlet, err := svc.service.ApiHierarchy.GetMerchantOutlets(*req)
	if err != nil {
		log.Println("Err ", svcName, " GetMerchantOutlets ", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respSvc.RecordsTotal = count
	respSvc.RecordsFiltered = count
	respSvc.Data = resMerchantOutlet
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "SUCCESS", respSvc)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) AddMerchantOutlet(ctx echo.Context) error {
	var (
		svcName = "AddMerchantOutlet"
	)
	req := new(models.ReqGetMerchantOutlet)
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
	if req.MerchantOutletName == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"merchantOutlet name is empty",
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
	req.MerchantOutletName = strings.ToUpper(req.MerchantOutletName)
	req.MerchantName = strings.ToUpper(req.MerchantName)
	req.MerchantOutletPassword, err = helpers.PswEnc(req.MerchantOutletPassword)
	if err != nil {
		log.Println("Err ", svcName, "PswEnc", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"password not generated",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	request := models.ReqGetMerchantOutlet{
		MerchantOutletName:     req.MerchantOutletName,
		MerchantOutletUsername: "",
		MerchantOutletPassword: "",
		MerchantId:             req.MerchantId,
		MerchantName:           req.MerchantName,
		GroupId:                req.GroupId,
		GroupName:              req.GroupName,
		ClientId:               req.ClientId,
		ClientName:             req.ClientName,
	}
	resp, err := svc.service.ApiHierarchy.GetMerchantOutlets(request)
	if err != nil {
		log.Println("Err ", svcName, "GetMerchantOutlet", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"failed",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if len(resp) != 0 {
		log.Println("Err ", svcName, "GetMerchantOutlet", " MerchantOutlet or username is exist")
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Merchant Outlet name or username is exist",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.service.ApiHierarchy.AddMerchantOutlet(*req)
	if err != nil {
		log.Println("Err ", svcName, "AddMerchantOutlet", err)
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
func (svc HierarcyService) DropMerchantOutlet(ctx echo.Context) error {
	var (
		svcName = "DropMerchantOutlet"
	)
	req := new(models.ReqGetMerchantOutlet)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.service.ApiHierarchy.DropMerchantOutlet(*req)
	if err != nil {
		log.Println("Err ", svcName, "DropMerchantOutlet", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"failed",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) UpdateMerchantOutlet(ctx echo.Context) error {
	var (
		svcName = "UpdateMerchantOutlet"
		// respSvc    models.ResponseList
	)
	req := new(models.ReqGetMerchantOutlet)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ClientName = strings.ToUpper(req.ClientName)
	req.GroupName = strings.ToUpper(req.GroupName)
	req.MerchantOutletName = strings.ToUpper(req.MerchantOutletName)
	req.MerchantName = strings.ToUpper(req.MerchantName)
	if req.MerchantOutletPassword != "" {
		req.MerchantOutletPassword, err = helpers.PswEnc(req.MerchantOutletPassword)
		if err != nil {
			log.Println("Err ", svcName, "PswEnc", err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE,
				configs.VALIDATE_ERROR_CODE,
				"password not generated",
				nil)
			return ctx.JSON(http.StatusOK, result)
		}
	} else {
		resp, err := svc.service.ApiHierarchy.GetMerchantOutlet(models.ReqGetMerchantOutlet{
			ID: req.ID,
		})
		if err != nil {
			log.Println("Err ", svcName, "UpdateMerchantOutlet", err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE,
				configs.VALIDATE_ERROR_CODE,
				"failed",
				nil)
			return ctx.JSON(http.StatusOK, result)
		}
		req.MerchantOutletPassword = resp.MerchantOutletPassword
	}
	_, err = svc.service.ApiHierarchy.UpdateMerchantOutlet(*req)
	if err != nil {
		log.Println("Err ", svcName, "UpdateMerchantOutlet", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"failed",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "SUCCESS", nil)
	return ctx.JSON(http.StatusOK, result)
}
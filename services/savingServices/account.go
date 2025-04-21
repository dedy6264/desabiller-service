package savingservices

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (svc savingServices) AddAccount(ctx echo.Context) error {
	var (
		svcName = "AddAccount"
	)
	req := new(models.ReqGetAccount)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILED", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.CifID == 0 {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"CIF is empty", "CIF is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.AccountNumber == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Account Number is empty", "Account Number is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.SavingSegmentID == 0 {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Segment is empty", "Segment is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.AccountPin != "" {
		req.AccountPin, err = helpers.PassEncrypt(req.AccountPin)
		if err != nil {
			log.Println("Err ", svcName, "AddAccount", err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE,
				configs.VALIDATE_ERROR_CODE,
				"failed", "failed",
				err)
			return ctx.JSON(http.StatusOK, result)
		}
	}
	_, err = svc.services.SavingRepo.AddAccount(*req, nil)
	if err != nil {
		log.Println("Err ", svcName, "AddAccount", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"failed", "failed",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		configs.SUCCESS_MSG, configs.SUCCESS_MSG,
		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc savingServices) GetAccounts(ctx echo.Context) error {
	var (
		svcName = "GetAccount"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetAccount)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	count, err := svc.services.SavingRepo.GetAccountCount(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetAccountCount", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.SavingRepo.GetAccounts(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetAccount", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respSvc.Data = resp
	respSvc.RecordsTotal = count
	respSvc.RecordsFiltered = count
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		configs.SUCCESS_MSG, configs.SUCCESS_MSG,
		respSvc)
	return ctx.JSON(http.StatusOK, result)
}
func (svc savingServices) DropAccount(ctx echo.Context) error {
	var (
		svcName = "DropAccount"
	)
	req := new(models.ReqGetAccount)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.SavingRepo.DropAccount(req.ID, nil)
	if err != nil {
		log.Println("Err ", svcName, "DropAccount", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "Failed", "failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		configs.SUCCESS_MSG,
		configs.SUCCESS_MSG,
		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc savingServices) UpdateAccount(ctx echo.Context) error {
	var (
		svcName = "UpdateAccount"
	)
	req := new(models.ReqGetAccount)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.AccountPin != "" {
		req.AccountPin, err = helpers.PassEncrypt(req.AccountPin)
		if err != nil {
			log.Println("Err ", svcName, "AddAccount", err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE,
				configs.VALIDATE_ERROR_CODE,
				"failed", "Failed",
				nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}
	err = svc.services.SavingRepo.UpdateAccount(*req, nil)
	if err != nil {
		log.Println("Err ", svcName, "UpdateAccount", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "Failed", "failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		configs.SUCCESS_MSG,
		configs.SUCCESS_MSG,
		nil)
	return ctx.JSON(http.StatusOK, result)
}

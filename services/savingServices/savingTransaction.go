package savingservices

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (svc savingServices) GetSavingTransactions(ctx echo.Context) error {
	var (
		svcName = "GetSavingTransaction"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetSavingTransaction)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	count, err := svc.services.SavingRepo.GetSavingTransactionCount(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetSavingTransactionCount", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.SavingRepo.GetSavingTransactions(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetSavingTransaction", err)
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

func (svc savingServices) UpdateSavingTransaction(ctx echo.Context) error {
	var (
		svcName = "UpdateSavingTransaction"
	)
	req := new(models.ReqGetSavingTransaction)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.SavingRepo.UpdateSavingTransaction(*req, nil)
	if err != nil {
		log.Println("Err ", svcName, "UpdateSavingTransaction", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		"Success",
		nil)
	return ctx.JSON(http.StatusOK, result)
}

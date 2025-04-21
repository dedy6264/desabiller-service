package savingservices

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func (svc savingServices) AddSavingSegment(ctx echo.Context) error {
	var (
		svcName = "AddSavingSegment"
	)
	req := new(models.ReqGetSavingSegment)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.SavingSegmentName == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Segment name is empty",
			"Segment name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.LimitAmount == 0 {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Limit Amount is empty",
			"Limit Amount is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.SavingTypeID == 0 {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Saving Type is empty",
			"Saving Type is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}

	req.SavingSegmentName = strings.ToUpper(req.SavingSegmentName)
	_, err = svc.services.SavingRepo.AddSavingSegment(*req, nil)
	if err != nil {
		log.Println("Err ", svcName, "AddSavingSegment", err)
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
func (svc savingServices) GetSavingSegments(ctx echo.Context) error {
	var (
		svcName = "GetSavingSegment"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetSavingSegment)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	count, err := svc.services.SavingRepo.GetSavingSegmentCount(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetSavingSegmentCount", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.SavingRepo.GetSavingSegments(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetSavingSegment", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "Failed", err.Error(), nil)
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
func (svc savingServices) DropSavingSegment(ctx echo.Context) error {
	var (
		svcName = "DropSavingSegment"
	)
	req := new(models.ReqGetSavingSegment)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.SavingRepo.DropSavingSegment(req.ID, nil)
	if err != nil {
		log.Println("Err ", svcName, "DropSavingSegment", err)
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
func (svc savingServices) UpdateSavingSegment(ctx echo.Context) error {
	var (
		svcName = "UpdateSavingSegment"
	)
	req := new(models.ReqGetSavingSegment)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.SavingSegmentName = strings.ToUpper(req.SavingSegmentName)
	err = svc.services.SavingRepo.UpdateSavingSegment(*req, nil)
	if err != nil {
		log.Println("Err ", svcName, "UpdateSavingSegment", err)
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

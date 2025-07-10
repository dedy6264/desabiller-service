package savingservices

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

func (svc savingServices) AddSavingSegment(ctx echo.Context) error {
	var (
		svcName = "AddSavingSegment"
		t       = time.Now()
		dbTime  = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqGetSavingSegment)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.SavingSegmentName == "" {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Segment name is empty",
			"Segment name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.LimitAmount == 0 {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Limit Amount is empty",
			"Limit Amount is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.SavingTypeID == 0 {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Saving Type is empty",
			"Saving Type is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}

	req.Filter.SavingSegmentName = strings.ToUpper(req.Filter.SavingSegmentName)
	req.Filter.CreatedAt = dbTime
	req.Filter.UpdatedAt = dbTime
	req.Filter.CreatedBy = "sys"
	req.Filter.UpdatedBy = "sys"
	_, err = svc.services.SavingRepo.AddSavingSegment(*req, nil)
	if err != nil {
		utils.Log(" AddSavingSegment", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1],
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
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	count, err := svc.services.SavingRepo.GetSavingSegmentCount(*req)
	if err != nil {
		utils.Log(" GetSavingSegmentCount", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.SavingRepo.GetSavingSegments(*req)
	if err != nil {
		utils.Log(" GetSavingSegments", svcName, err)
		if err == sql.ErrNoRows {
			result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1], respSvc)
			return ctx.JSON(http.StatusOK, result)
		}
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respSvc.Data = resp
	respSvc.RecordsTotal = count
	respSvc.RecordsFiltered = count
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1],
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
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.SavingRepo.DropSavingSegment(req.Filter.ID, nil)
	if err != nil {
		utils.Log(" DropSavingSegment", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1],
		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc savingServices) UpdateSavingSegment(ctx echo.Context) error {
	var (
		svcName = "UpdateSavingSegment"
		t       = time.Now()
		dbTime  = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqGetSavingSegment)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Filter.SavingSegmentName = strings.ToUpper(req.Filter.SavingSegmentName)
	req.Filter.UpdatedAt = dbTime
	req.Filter.UpdatedBy = "sys"
	err = svc.services.SavingRepo.UpdateSavingSegment(*req, nil)
	if err != nil {
		utils.Log(" UpdateSavingSegment", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1],
		nil)
	return ctx.JSON(http.StatusOK, result)
}

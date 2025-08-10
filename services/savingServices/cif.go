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

func (svc savingServices) AddCif(ctx echo.Context) error {
	var (
		svcName = "AddCif"
		t       = time.Now()
		dbTime  = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqGetCIF)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.CifName == "" {

		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Cif name is empty",
			"Cif name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.CifNoID == "" {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Cif NIK is empty",
			"Cif NIK is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.CifTypeID == "" {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Cif TYPE is empty",
			"Cif TYPE is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.CifEmail == "" {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Cif Email is empty",
			"Cif Email is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Filter.CifName = strings.ToUpper(req.Filter.CifName)
	req.Filter.CreatedAt = dbTime
	req.Filter.UpdatedAt = dbTime
	req.Filter.CreatedBy = "sys"
	req.Filter.UpdatedBy = "sys"
	_, err = svc.services.SavingRepo.AddCif(*req, nil)
	if err != nil {
		utils.Log("AddCif ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			configs.RC_SYSTEM_ERROR[1],
			nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1],
		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc savingServices) GetCifs(ctx echo.Context) error {
	var (
		svcName = "GetCif"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetCIF)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	count, err := svc.services.SavingRepo.GetCifCount(*req)
	if err != nil {
		utils.Log("GetCifCount", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.SavingRepo.GetCifs(*req)
	if err != nil {
		utils.Log(" GetCif", svcName, err)
		if err == sql.ErrNoRows {
			result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1], respSvc)
			return ctx.JSON(http.StatusOK, result)
		}
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			err.Error(), nil)
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
func (svc savingServices) DropCif(ctx echo.Context) error {
	var (
		svcName = "DropCif"
	)
	req := new(models.ReqGetCIF)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.SavingRepo.DropCif(int(req.Filter.ID), nil)
	if err != nil {
		utils.Log(" DropCif", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			configs.RC_SYSTEM_ERROR[1], nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1],
		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc savingServices) UpdateCif(ctx echo.Context) error {
	var (
		svcName = "UpdateCif"
		t       = time.Now()
		dbTime  = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqGetCIF)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Filter.CifName = strings.ToUpper(req.Filter.CifName)
	req.Filter.UpdatedAt = dbTime
	req.Filter.UpdatedBy = "sys"
	err = svc.services.SavingRepo.UpdateCif(*req, nil)
	if err != nil {
		utils.Log(" UpdateCif", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			configs.RC_SYSTEM_ERROR[1], nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1],
		nil)
	return ctx.JSON(http.StatusOK, result)
}

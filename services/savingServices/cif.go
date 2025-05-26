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

func (svc savingServices) AddCif(ctx echo.Context) error {
	var (
		svcName = "AddCif"
	)
	req := new(models.ReqGetCif)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.CifName == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Cif name is empty",
			"Cif name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.CifNik == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Cif NIK is empty",
			"Cif NIK is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.CifPhone == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Cif Phone is empty",
			"Cif Phone is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.CifEmail == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Cif Email is empty",
			"Cif Email is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.CifName = strings.ToUpper(req.CifName)
	_, err = svc.services.SavingRepo.AddCif(*req, nil)
	if err != nil {
		log.Println("Err ", svcName, "AddCif", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.DB_NOT_FOUND,
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
func (svc savingServices) GetCifs(ctx echo.Context) error {
	var (
		svcName = "GetCif"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetCif)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	count, err := svc.services.SavingRepo.GetCifCount(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetCifCount", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.SavingRepo.GetCifs(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetCif", err)
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
func (svc savingServices) DropCif(ctx echo.Context) error {
	var (
		svcName = "DropCif"
	)
	req := new(models.ReqGetCif)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.SavingRepo.DropCif(req.ID, nil)
	if err != nil {
		log.Println("Err ", svcName, "DropCif", err)
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
func (svc savingServices) UpdateCif(ctx echo.Context) error {
	var (
		svcName = "UpdateCif"
	)
	req := new(models.ReqGetCif)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.CifName = strings.ToUpper(req.CifName)
	err = svc.services.SavingRepo.UpdateCif(*req, nil)
	if err != nil {
		log.Println("Err ", svcName, "UpdateCif", err)
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

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

func (svc HierarcyService) GetUserApps(ctx echo.Context) error {
	var (
		svcName = "GetUserApps"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetUserApp)
	//binding request
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Filter.Name = strings.ToUpper(req.Filter.Name)
	count, err := svc.service.RepoHierarchy.GetUserAppCount(*req)
	if err != nil {
		log.Println("Err "+svcName+" GetUserAppCount ", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resUserApp, err := svc.service.RepoHierarchy.GetUserApps(*req)
	if err != nil {
		log.Println("Err ", svcName, " GetUserApps ", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respSvc.RecordsTotal = count
	respSvc.RecordsFiltered = count
	respSvc.Data = resUserApp
	respSvc.Draw = req.Draw
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, configs.SUCCESS_MSG, configs.SUCCESS_MSG, respSvc)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) AddUserApp(ctx echo.Context) error {
	var (
		svcName = "AddUserApp"
	)
	req := new(models.ReqGetUserApp)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.Name == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Name is empty",
			"Name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.Username == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Username is empty",
			"Username is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.Password == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Password is empty",
			"Password is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Filter.Name = strings.ToUpper(req.Filter.Name)

	_, err = svc.service.RepoHierarchy.GetUserApp(*req)
	if err != nil {
		if err == sql.ErrNoRows {
			err = svc.service.RepoHierarchy.AddUserApp(*req, nil)
			if err != nil {
				log.Println("Err ", svcName, "AddUserApp", err)
				result := helpers.ResponseJSON(configs.FALSE_VALUE,
					configs.DB_NOT_FOUND,
					"failed",
					"failed",
					nil)
				return ctx.JSON(http.StatusOK, result)
			}
		} else {
			log.Println("Err ", svcName, "GetUserApp", err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE,
				configs.DB_NOT_FOUND,
				"failed",
				"failed",
				nil)
			return ctx.JSON(http.StatusOK, result)
		}
	} else {
		log.Println("Err ", svcName, "GetUserApp", " Merchant is exist")
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.DB_NOT_FOUND,
			"user is exist",
			"user is exist",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0],
		configs.RC_SUCCESS[1], configs.RC_SUCCESS[1],
		req.Filter)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) DropUserApp(ctx echo.Context) error {
	var (
		svcName = "DropUserApp"
	)
	req := new(models.ReqGetUserApp)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.service.RepoHierarchy.DropUserApp(int(req.Filter.ID), nil)
	if err != nil {
		log.Println("Err ", svcName, "DropUserApp", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.DB_NOT_FOUND,
			"failed",
			"failed",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1], nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) UpdateUserApp(ctx echo.Context) error {
	var (
		svcName = "UpdateUserApp"
		// respSvc    models.ResponseList
	)
	req := new(models.ReqGetUserApp)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Filter.Name = strings.ToUpper(req.Filter.Name)
	err = svc.service.RepoHierarchy.UpdateUserApp(*req, nil)
	if err != nil {
		log.Println("Err ", svcName, "UpdateUserApp", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"failed",
			"failed",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1], req.Filter)
	return ctx.JSON(http.StatusOK, result)
}

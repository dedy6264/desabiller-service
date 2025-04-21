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

func (svc HierarcyService) GetGroups(ctx echo.Context) error {
	var (
		svcName = "GetGroups"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetGroup)
	//binding *req
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ClientName = strings.ToUpper(req.ClientName)
	req.GroupName = strings.ToUpper(req.GroupName)
	count, err := svc.service.RepoHierarchy.GetGroupCount(*req)
	if err != nil {
		log.Println("Err "+svcName+" GetGroupCount ", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resGroup, err := svc.service.RepoHierarchy.GetGroups(*req)
	if err != nil {
		log.Println("Err ", svcName, " GetGroups ", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respSvc.RecordsTotal = count
	respSvc.RecordsFiltered = count
	respSvc.Data = resGroup
	respSvc.Draw = req.Filter.Draw
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, configs.SUCCESS_MSG, configs.SUCCESS_MSG, respSvc)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) AddGroup(ctx echo.Context) error {
	var (
		svcName = "AddGroup"
	)
	req := new(models.ReqGetGroup)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ClientName == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"client name is empty",
			"client name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.GroupName == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"group name is empty",
			"group name is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ClientName = strings.ToUpper(req.ClientName)
	req.GroupName = strings.ToUpper(req.GroupName)

	_, err = svc.service.RepoHierarchy.GetGroup(*req)
	if err != nil {
		if err == sql.ErrNoRows {
			err = svc.service.RepoHierarchy.AddGroup(*req, nil)
			if err != nil {
				log.Println("Err ", svcName, "AddGroup", err)
				result := helpers.ResponseJSON(configs.FALSE_VALUE,
					configs.DB_ERROR,
					"failed",
					"failed",
					nil)
				return ctx.JSON(http.StatusOK, result)
			}
		} else {
			log.Println("Err ", svcName, "GetGroup", err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE,
				configs.DB_ERROR,
				"failed",
				"failed",
				nil)
			return ctx.JSON(http.StatusOK, result)
		}
	} else {
		log.Println("Err ", svcName, "GetGroup", " client name is exist")
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"client name is exist",
			"client name is exist",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		configs.SUCCESS_MSG, configs.SUCCESS_MSG,
		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) DropGroup(ctx echo.Context) error {
	var (
		svcName = "DropGroup"
	)
	req := new(models.ReqGetGroup)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.service.RepoHierarchy.DropGroup(req.ID)
	if err != nil {
		log.Println("Err ", svcName, "DropGroup", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.DB_NOT_FOUND,
			"failed",
			"failed",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, configs.SUCCESS_MSG, configs.SUCCESS_MSG, nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) UpdateGroup(ctx echo.Context) error {
	var (
		svcName = "UpdateGroup"
		// respSvc    models.ResponseList
	)
	req := new(models.ReqGetGroup)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ClientName = strings.ToUpper(req.ClientName)
	req.GroupName = strings.ToUpper(req.GroupName)

	_, err = svc.service.RepoHierarchy.UpdateGroup(*req)
	if err != nil {
		log.Println("Err ", svcName, "UpdateGroup", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.DB_NOT_FOUND,
			"failed",
			"failed",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, configs.SUCCESS_MSG, configs.SUCCESS_MSG, nil)
	return ctx.JSON(http.StatusOK, result)
}

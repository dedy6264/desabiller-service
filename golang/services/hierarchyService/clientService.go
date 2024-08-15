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

func (svc HierarcyService) GetClients(ctx echo.Context) error {
	var (
		svcName = "GetListClient"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetClient)
	//binding request
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ClientName = strings.ToUpper(req.ClientName)
	request := models.ReqGetClient{
		ID:         req.ID,
		ClientName: req.ClientName,
		Filter: models.FilterReq{
			Limit:     req.Filter.Limit,
			Offset:    req.Filter.Offset,
			OrderBy:   req.Filter.OrderBy,
			CreatedAt: req.Filter.CreatedAt,
			CreatedBy: req.Filter.CreatedBy,
			UpdatedAt: req.Filter.UpdatedAt,
			UpdatedBy: req.Filter.UpdatedBy,
		},
	}
	count, err := svc.service.ApiHierarchy.GetCount(request)
	if err != nil {
		log.Println("Err "+svcName+" GetCount ", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resClient, err := svc.service.ApiHierarchy.GetClients(request)
	if err != nil {
		log.Println("Err ", svcName, " GetClients ", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respSvc.TotalData = count
	respSvc.Data = resClient
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "SUCCESS", respSvc)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) AddClient(ctx echo.Context) error {
	var (
		svcName = "AddClient"
	)
	req := new(models.ReqGetClient)
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
	req.ClientName = strings.ToUpper(req.ClientName)
	request := models.ReqGetClient{
		ClientName: req.ClientName,
		Filter:     models.FilterReq{},
	}
	_, err = svc.service.ApiHierarchy.GetClient(request)
	if err != nil {
		if err == sql.ErrNoRows {
			err = svc.service.ApiHierarchy.AddClient(request, nil)
			if err != nil {
				log.Println("Err ", svcName, "AddClient", err)
				result := helpers.ResponseJSON(configs.FALSE_VALUE,
					configs.VALIDATE_ERROR_CODE,
					"failed",
					nil)
				return ctx.JSON(http.StatusOK, result)
			}
		} else {
			log.Println("Err ", svcName, "GetClient", err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE,
				configs.VALIDATE_ERROR_CODE,
				"failed",
				nil)
			return ctx.JSON(http.StatusOK, result)
		}
	} else {
		log.Println("Err ", svcName, "GetClient", " client name is exist")
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"client name is exist",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.SUCCESS_CODE,
		"Success",
		nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) DropClient(ctx echo.Context) error {
	var (
		svcName = "DropClient"
	)
	req := new(models.ReqGetClient)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.service.ApiHierarchy.DropClient(req.ID, nil)
	if err != nil {
		log.Println("Err ", svcName, "DropClient", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"failed",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) UpdateClient(ctx echo.Context) error {
	var (
		svcName = "UpdateClient"
		// respSvc    models.ResponseList
	)
	req := new(models.ReqGetClient)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ClientName = strings.ToUpper(req.ClientName)
	err = svc.service.ApiHierarchy.UpdateClient(*req, nil)
	if err != nil {
		log.Println("Err ", svcName, "UpdateClient", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"failed",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "SUCCESS", nil)
	return ctx.JSON(http.StatusOK, result)
}

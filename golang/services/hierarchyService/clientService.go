package hierarchyservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func (svc HierarcyService) GetListClient(ctx echo.Context) error {
	var (
		svcName = "GetListClient"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetListClient)
	//binding request
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	req.ClientName = strings.ToUpper(req.ClientName)
	request := models.ReqGetListClient{
		ID:         req.ID,
		ClientName: req.ClientName,
		Limit:      req.Limit,
		Offset:     req.Offset,
		OrderBy:    req.OrderBy,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Username:   req.Username,
	}
	count, _ := svc.service.ApiHierarchy.GetListClientCount(request)
	if count == 0 {
		log.Println("Error " + svcName + " // GetListClient :: Not found")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, "34", "Data :: empty", nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	resClient, status := svc.service.ApiHierarchy.GetListClient(request)
	if !status {
		log.Println("Error " + svcName + " // GetListClient :: Not found")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, "34", "Data :: empty", nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	respSvc.TotalData = count
	respSvc.Data = resClient
	fmt.Println(":::RESULT::: ", resClient)
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "SUCCESS", respSvc)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) AddClient(ctx echo.Context) error {
	var (
		svcName = "AddClient"
	)
	req := new(models.ReqGetListClient)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED Validate"+err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	if req.ClientName == "" {
		log.Println("Error " + svcName + " // ClientName :: empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "ClientName :: empty", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	req.ClientName = strings.ToUpper(req.ClientName)
	request := models.ReqGetListClient{
		ID:         req.ID,
		ClientName: req.ClientName,
		Limit:      req.Limit,
		Offset:     0,
		OrderBy:    req.OrderBy,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Username:   req.Username,
	}
	resClient, status := svc.service.ApiHierarchy.GetListClient(request)
	fmt.Println("::::::", resClient)
	if status {
		for _, data := range resClient {
			fmt.Println(":::", data.ClientName, req.ClientName)
			if data.ClientName == req.ClientName {
				log.Println("Error " + svcName + " // ClientName :: client name rejected")
				result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "ClientName :: rejected", nil)
				return ctx.JSON(http.StatusNotFound, result)
			}
		}
	}
	status = svc.service.ApiHierarchy.AddClient(request)
	if !status {
		log.Println("Error " + svcName + " // AddClient :: failed")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Data :: empty", nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) DropClient(ctx echo.Context) error {
	var (
		svcName = "DropClient"
	)
	req := new(models.ReqGetListClient)
	status, err := helpers.BindValidate(req, ctx)
	if !status {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED Validate"+err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	status = svc.service.ApiHierarchy.DropClient(*req)
	if !status {
		log.Println("Error " + svcName + "// DropClient :: Failed")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED Validate", nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) UpdateClient(ctx echo.Context) error {
	var (
		svcName = "UpdateClient"
		// respSvc    models.ResponseList
	)
	req := new(models.ReqGetListClient)
	status, err := helpers.BindValidate(req, ctx)
	if !status {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED Validate"+err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	req.ClientName = strings.ToUpper(req.ClientName)
	update := models.ReqGetListClient{
		ID:         req.ID,
		ClientName: req.ClientName,
		Limit:      req.Limit,
		Offset:     req.Offset,
		OrderBy:    req.OrderBy,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Username:   req.Username,
	}
	resUpd, status := svc.service.ApiHierarchy.UpdateClient(update)
	if !status {
		log.Println("Error " + svcName + " // UpdateClient :: ")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED Validate", nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "SUCCESS", resUpd)
	return ctx.JSON(http.StatusOK, result)

}

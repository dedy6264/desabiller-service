package clientservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func (svc nHierarchyClientServices) GetClientService(ctx echo.Context) error {
	var (
		svcName  = "GetClientService"
		logErr   = "Error " + svcName
		req      models.ReqGetListNClient
		response models.ResGetNClient
		// dbTime   = time.Now().Format(time.RFC3339)
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println(logErr+"BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ClientName == "" {
		log.Println(logErr + " Client Name cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Client Name cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ClientName = strings.ToUpper(req.ClientName)
	resp, err := svc.services.ApiNHierarchy.NReadSingleClient(req)
	if err != nil {
		log.Println(logErr+"NReadSingleClient", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Not Found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	response.ID = resp.ID
	response.ClientName = resp.ClientName
	response.CreatedAt = resp.CreatedAt
	response.CreatedBy = resp.CreatedBy
	response.UpdatedAt = resp.UpdatedAt
	response.UpdatedBy = resp.UpdatedBy
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", response)
	return ctx.JSON(http.StatusOK, result)
}

func (svc nHierarchyClientServices) GetListClientService(ctx echo.Context) error {
	var (
		svcName = "GetListClientService"
		logErr  = "Error " + svcName
		req     models.ReqGetListNClient
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println(logErr+"BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ClientName = strings.ToUpper(req.ClientName)
	resP, err := svc.services.ApiNHierarchy.NReadClient(req)
	if err != nil {
		log.Println(logErr+"NReadClient", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if len(resP) == 0 {
		log.Println(logErr + " Not Found")
		result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", resP)
	return ctx.JSON(http.StatusOK, result)
}

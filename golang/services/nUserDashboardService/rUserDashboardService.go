package nuserdashboardservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (svc nUserDashboardServices) GetUserDashboard(ctx echo.Context) error {
	var (
		svcName = "GetUserDashboard"
		logErr  = "Error " + svcName
		req     models.ReqCreateNUserDashboard
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println(logErr+"BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.ApiNUserDashboard.NReadSingleUserDashboard(req)
	if err != nil {
		log.Println(logErr+"CreateUserDashboard", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp.Password = "##############"
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", resp)
	return ctx.JSON(http.StatusOK, result)
}

func (svc nUserDashboardServices) GetListUserDashboard(ctx echo.Context) error {
	var (
		svcName = "GetListUserDashboard"
		logErr  = "Error " + svcName
		req     models.ReqGetListNUserDashboard
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println(logErr+"BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Draw == 0 {
		req.Draw = 1
	}
	req.Offset = (req.Draw * req.Limit) - req.Limit + 1
	resp, err := svc.services.ApiNUserDashboard.NReadUserDashboard(req)
	if err != nil {
		log.Println(logErr+"GetListUserDashboard", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", resp)
	return ctx.JSON(http.StatusOK, result)
}

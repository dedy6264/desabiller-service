package nuserdashboardservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (svc nUserDashboardServices) DropUserDashboard(ctx echo.Context) error {
	var (
		svcName = "DropUserDashboard"
		logErr  = "Error " + svcName
		req     models.ReqGetListNUserDashboard
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println(logErr+"BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Data.ID == 0 {
		log.Println(logErr + " Id cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Id cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.ApiNUserDashboard.NDropUserDashboard(req.Data.ID)
	if err != nil {
		log.Println(logErr+"DropUserDashboard", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", resp)
	return ctx.JSON(http.StatusOK, result)
}

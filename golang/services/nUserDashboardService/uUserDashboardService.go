package nuserdashboardservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (svc nUserDashboardServices) UpdateUserDashboard(ctx echo.Context) error {
	var (
		svcName  = "UpdateUserDashboard"
		logErr   = "Error " + svcName
		req      models.ReqCreateNUserDashboard
		response models.RespGetListNUserDashboard
		dbTime   = time.Now().Format(time.RFC3339)
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println(logErr+"BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ID == 0 {
		log.Println(logErr + " ID cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "ID cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Email == "" {
		log.Println(logErr + " Email cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Email cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Username == "" {
		log.Println(logErr + " Username cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Username cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Role == "" {
		log.Println(logErr + " Role cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Role cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resGet, err := svc.services.ApiNUserDashboard.NReadSingleUserDashboard(models.ReqCreateNUserDashboard{
		ID: req.ID,
	})
	if err != nil {
		log.Println(logErr+"CreateUserDashboard", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Password != "" {
		req.Password, err = helpers.PswEnc(req.Password)
		if err != nil {
			log.Println(logErr+"CreateUserDashboard", err.Error())
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		}
	} else {
		req.Password = resGet.Password
	}
	req.UpdatedAt = dbTime
	req.UpdatedBy = "sys"
	response, err = svc.services.ApiNUserDashboard.NUpdateUserDashboard(req)
	if err != nil {
		log.Println(logErr+"CreateUserDashboard", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	response.Password = "###############"
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", response)
	return ctx.JSON(http.StatusOK, result)
}

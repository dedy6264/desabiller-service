package nuserdashboardservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/services"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type nUserDashboardServices struct {
	services services.UsecaseService
}

func NewApiNUserDashboardServices(services services.UsecaseService) nUserDashboardServices {
	return nUserDashboardServices{
		services: services,
	}
}
func (svc nUserDashboardServices) CreateUserDashboard(ctx echo.Context) error {
	var (
		svcName  = "CreateUserDashboard"
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
	if req.Password == "" {
		log.Println(logErr + " Password cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Password cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Role == "" {
		log.Println(logErr + " Role cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Role cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Password, err = helpers.PswEnc(req.Password)
	req.CreatedAt = dbTime
	req.UpdatedAt = dbTime
	if err != nil {
		log.Println(logErr+"CreateUserDashboard", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	id, err := svc.services.ApiNUserDashboard.NCreateUserDashboard(req)
	if err != nil {
		log.Println(logErr+"CreateUserDashboard", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	response.ID = id
	response.Username = req.Username
	response.Email = req.Email
	response.CreatedAt = dbTime
	response.CreatedBy = "sys"
	response.UpdatedAt = dbTime
	response.UpdatedBy = "sys"
	response.Password = "###############"
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", response)
	return ctx.JSON(http.StatusOK, result)
}

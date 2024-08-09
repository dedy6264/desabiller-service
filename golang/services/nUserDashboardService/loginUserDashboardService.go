package nuserdashboardservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (svc nUserDashboardServices) LoginUserDashboard(ctx echo.Context) error {
	var (
		svcName = "LoginUserDashboard"
		logErr  = "Error " + svcName
		req     models.ReqLoginUserDashboard
		resSvc  models.RespLoginUserDashboard
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
	if req.Password == "" {
		log.Println(logErr + " Password cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Password cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respEnc, err := helpers.PswEnc(req.Password)
	if err != nil {
		log.Println(logErr+"Encrypt", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.ApiNUserDashboard.NReadSingleUserDashboard(models.ReqCreateNUserDashboard{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		log.Println(logErr+"LoginUserDashboard", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, " Email, username or password not match", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if resp.Password != respEnc {
		log.Println(logErr+"LoginUserDashboard ", "Email, username or password not match")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, " Email, username or password not match", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	token, err := helpers.TokenJwtGenerateDashboard(resp.ID)
	if err != nil {
		log.Println(logErr+"LoginUserDashboard", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resSvc.ID = resp.ID
	resSvc.Username = resp.Username
	resSvc.Email = resp.Email
	resSvc.Role = resp.Role
	resSvc.Token = token
	resSvc.ClientId = resp.ClientId
	resSvc.MerchantId = resp.MerchantId
	resSvc.MerchantOutletId = resp.MerchantOutletId
	resSvc.Password = "##########"
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", resSvc)
	return ctx.JSON(http.StatusOK, result)
}

// func (svc nUserDashboardServices) LogoutUserDashboard(ctx echo.Context) error {
// 	var (
// 		svcName = "LogoutUserDashboard"
// 		logErr  = "Error " + svcName
// 		req     models.ReqLoginUserDashboard
// 		resSvc  models.RespLoginUserDashboard
// 	)
// 	token := ctx.Get("user")
// 	if token == nil {
// 		log.Println(logErr + "LoginUserDashboard")
// 		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED ", nil)
// 		return ctx.JSON(http.StatusOK, result)
// 	}
// 	// user := ctx.Get("user").(*jwt.Token)
// 	// accessToken := user.Raw

// 	// claims := user.Claims.(jwt.MapClaims)

// 	// result := models.AccessToken{
// 	// 	UserPartnerID:   claims["userPartnerID"].(string),
// 	// 	UserPartnerCode: claims["userPartnerCode"].(string),
// 	// 	UserPartnerName: claims["userPartnerName"].(string),
// 	// 	AccessToken:     accessToken,
// 	// }
// }

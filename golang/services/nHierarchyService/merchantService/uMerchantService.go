package merchantservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func (svc nHierarchyMerchantServices) UpdateMerchantService(ctx echo.Context) error {
	var (
		svcName  = "UpdateMerchantService"
		logErr   = "Error " + svcName
		req      models.ReqGetListMerchant
		response models.ResGetNMerchant
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
	if req.ClientId == 0 {
		token := ctx.Get("user")
		if token == nil {
			log.Println(logErr + "get Token failed")
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED ", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		user := ctx.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userDashboardId := claims["userDashboardId"].(float64)
		clientId := int(userDashboardId)
		resUser, err := svc.services.ApiNUserDashboard.NReadSingleUserDashboard(models.ReqCreateNUserDashboard{
			ID: clientId,
		})
		if err != nil {
			log.Println(logErr+"NReadSingleUserDashboard ", err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED ", err)
			return ctx.JSON(http.StatusOK, result)
		}
		if resUser.ClientId == (-1) {
			log.Println(logErr + " Client ID cannot be empty")
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Client ID cannot be empty ", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		req.ClientId = resUser.ClientId
	}
	if req.MerchantName == "" {
		log.Println(logErr + " Merchant Name cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Merchant Name cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.MerchantName = strings.ToUpper(req.MerchantName)
	response, err = svc.services.ApiNHierarchy.NUpdateMerchant(models.ReqUpdateNMerchant{
		ID:           req.ID,
		ClientId:     req.ClientId,
		MerchantName: req.MerchantName,
		UpdatedAt:    dbTime,
		UpdatedBy:    "sys",
	})
	if err != nil {
		log.Println(logErr+"NCreateMerchant", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	response.ID = req.ID
	response.MerchantName = req.MerchantName
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", response)
	return ctx.JSON(http.StatusOK, result)
}

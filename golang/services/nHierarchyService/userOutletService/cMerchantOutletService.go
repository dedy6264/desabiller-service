package useroutletservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/services"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
)

type nHierarchyUserOutletServices struct {
	services services.UsecaseService
}

func NewApiNHierarchyUserOutletServices(services services.UsecaseService) nHierarchyUserOutletServices {
	return nHierarchyUserOutletServices{
		services: services,
	}
}
func (svc nHierarchyUserOutletServices) CreateUserOutletService(ctx echo.Context) error {
	var (
		svcName  = "CreateUserOutletService"
		logErr   = "Error " + svcName
		req      models.ReqGetListNUserOutlet
		response models.ResGetNUserOutlet
		dbTime   = time.Now().Format(time.RFC3339)
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println(logErr+"BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
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
		// accessToken := user.Raw
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
	if req.MerchantId == 0 {
		log.Println(logErr + " Merchant ID cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Merchant ID cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.MerchantOutletId == 0 {
		log.Println(logErr + " Merchant Outlet ID cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Merchant Outlet ID cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Nickname == "" {
		log.Println(logErr + " Nickname cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Nickname cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.OutletUsername == "" {
		log.Println(logErr + " OutletUsername cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "OutletUsername cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Username == "" {
		log.Println(logErr + " Username cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Username cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.OutletPassword == "" {
		log.Println(logErr + " OutletPassword cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "OutletPassword cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.OutletPassword, _ = helpers.PswEnc(req.OutletPassword)
	_, err = svc.services.ApiNHierarchy.NReadSingleMerchantOutlet(models.ReqGetListNMerchantOutlet{
		ID:           req.MerchantOutletId,
		MerchantName: "",
		ClientId:     req.ClientId,
		MerchantId:   req.MerchantId,
		Limit:        0,
		Offset:       0,
		OrderBy:      "",
		StartDate:    "",
		EndDate:      "",
		Username:     "",
	})
	if err != nil {
		log.Println(logErr+" NReadSingleMerchantOutlet ", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.MerchantName = strings.ToUpper(req.MerchantName)
	req.MerchantOutletName = strings.ToUpper(req.MerchantOutletName)
	id, err := svc.services.ApiNHierarchy.NCreateUserOutlet(models.ReqGetListNUserOutlet{
		Nickname:           req.Nickname,
		OutletUsername:     req.OutletUsername,
		OutletPassword:     req.OutletPassword,
		MerchantOutletName: req.MerchantOutletName,
		ClientId:           req.ClientId,
		MerchantId:         req.MerchantId,
		MerchantOutletId:   req.MerchantOutletId,
	})
	if err != nil {
		log.Println(logErr+"NCreateUserOutlet", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	response.ID = id
	response.ClientId = req.ClientId
	response.MerchantId = req.MerchantId
	response.MerchantOutletId = req.MerchantOutletId
	response.MerchantName = req.MerchantName
	response.MerchantOutletName = req.MerchantOutletName
	response.OutletUsername = req.OutletUsername
	response.OutletPassword = "#############"
	response.CreatedAt = dbTime
	response.CreatedBy = "sys"
	response.UpdatedAt = dbTime
	response.UpdatedBy = "sys"
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", response)
	return ctx.JSON(http.StatusOK, result)
}

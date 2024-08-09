package merchantoutletservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func (svc nHierarchyMerchantOutletServices) GetMerchantOutletService(ctx echo.Context) error {
	var (
		svcName  = "GetMerchantOutletService"
		logErr   = "Error " + svcName
		req      models.ReqGetListNMerchantOutlet
		response models.ResGetNMerchantOutlet
		// dbTime   = time.Now().Format(time.RFC3339)
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
	if req.MerchantOutletName == "" {
		log.Println(logErr + " Merchant Outlet Name cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Merchant Outlet Name cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.ApiNHierarchy.NReadSingleMerchantOutlet(req)
	if err != nil {
		log.Println(logErr+"NReadSingleMerchantOutlet", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Not Found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	response.ID = resp.ID
	response.ClientId = resp.ClientId
	response.ClientName = resp.ClientName
	response.MerchantName = resp.MerchantName
	response.MerchantId = resp.MerchantId
	response.MerchantOutletName = resp.MerchantOutletName
	response.CreatedAt = resp.CreatedAt
	response.CreatedBy = resp.CreatedBy
	response.UpdatedAt = resp.UpdatedAt
	response.UpdatedBy = resp.UpdatedBy
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", response)
	return ctx.JSON(http.StatusOK, result)
}

func (svc nHierarchyMerchantOutletServices) GetListMerchantOutletService(ctx echo.Context) error {
	var (
		svcName = "GetListMerchantOutletService"
		logErr  = "Error " + svcName
		req     models.ReqGetListNMerchantOutlet
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
		req.ClientId = resUser.ClientId
	}
	resP, err := svc.services.ApiNHierarchy.NReadMerchantOutlet(req)
	if err != nil {
		log.Println(logErr+"NReadMerchantOutlet", err.Error())
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

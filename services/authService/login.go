package administrationservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (svc AdministrationService) Login(ctx echo.Context) error {
	var (
		svcName   = "Login"
		resultSvc models.RespLogin
		// uID       int
		// oID       int
		// mID       int
		// cID       int
	)
	//binding n validate required value
	req := new(models.ReqLogin)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.MerchantOutletUsername == "" {
		log.Println("Err ", svcName, "username cannot emptyEnc")
		result := helpers.ResponseJSON(false, configs.VALIDATE_ERROR_CODE, "username cannot empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.MerchantOutletPassword == "" {
		log.Println("Err ", svcName, "MerchantOutletPassword cannot emptyEnc")
		result := helpers.ResponseJSON(false, configs.VALIDATE_ERROR_CODE, "password cannot empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.MerchantOutletPassword, err = helpers.PswEnc(req.MerchantOutletPassword)
	if err != nil {
		log.Println("Err ", svcName, "PswEnc", err)
		result := helpers.ResponseJSON(false, configs.VALIDATE_ERROR_CODE, "wrong username or password", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respGet, err := svc.service.RepoHierarchy.GetMerchantOutlet(models.ReqGetMerchantOutlet{
		MerchantOutletUsername: req.MerchantOutletUsername,
	})
	if err != nil {
		log.Println("Err ", svcName, "GetMerchantOutlet", err)
		result := helpers.ResponseJSON(false, configs.VALIDATE_ERROR_CODE, "wrong username or password", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if respGet.MerchantOutletPassword != req.MerchantOutletPassword {
		log.Println("Err ", svcName, "MerchantOutletPassword cannot emptyEnc")
		result := helpers.ResponseJSON(false, configs.VALIDATE_ERROR_CODE, "wrong username or password", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	token, err := helpers.TokenJwtGenerate(respGet.MerchantId, respGet.ID, respGet.MerchantOutletName)
	if err != nil {
		log.Println("Err ", svcName, "PswEnc", err)
		result := helpers.ResponseJSON(false, configs.VALIDATE_ERROR_CODE, "wrong username or password", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	resultSvc.ID = respGet.ID
	resultSvc.MerchantOutletName = respGet.MerchantOutletName
	resultSvc.MerchantOutletUsername = respGet.MerchantOutletUsername
	resultSvc.MerchantOutletPassword = ""
	resultSvc.MerchantId = respGet.MerchantId
	resultSvc.MerchantName = respGet.MerchantName
	resultSvc.GroupId = respGet.GroupId
	resultSvc.GroupName = respGet.GroupName
	resultSvc.ClientId = respGet.ClientId
	resultSvc.ClientName = respGet.ClientName
	resultSvc.Token = token
	result := helpers.ResponseJSON(false, configs.SUCCESS_CODE, "success", resultSvc)
	return ctx.JSON(http.StatusOK, result)
}

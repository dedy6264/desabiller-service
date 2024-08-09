package trxservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func (svc TrxServices) GetListReportPos(ctx echo.Context) error {
	var (
		svcName = "GetListReportPos"
		logErr  = "ERROR " + svcName
	)
	req := new(models.ReqTrx)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println(logErr+" FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	a := ctx.Get("user").(*jwt.Token)
	claim := a.Claims.(jwt.MapClaims)
	snDevice := claim["snDevice"].(string)
	uID := claim["userId"].(float64)
	oID := claim["outletId"].(float64)
	mID := claim["merchantId"].(float64)
	cID := claim["clientId"].(float64)
	res, status := svc.service.ApiTrx.GetTrxListPos(models.ReqTrx{
		MerchantId:       int(mID),
		MerchantOutletId: int(oID),
		UserOutletId:     int(uID),
		OutletDeviceSn:   snDevice,
		ClientId:         int(cID),
	}, "trx_poses")
	if !status {
		log.Println(logErr + " GetTrxListPos")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "SUCCESS", res)
	return ctx.JSON(http.StatusOK, result)
}

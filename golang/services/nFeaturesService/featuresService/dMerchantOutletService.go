package featuresservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (svc nFeaturesServices) DropFeaturesService(ctx echo.Context) error {
	var (
		svcName = "DropFeaturesService"
		logErr  = "Error " + svcName
		req     models.ReqGetListNFeature
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println(logErr+"BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Data.ID == 0 {
		log.Println(logErr + " id cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "id cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	_, err = svc.services.ApiNFeatures.NDropFeature(req.Data.ID)
	if err != nil {
		log.Println(logErr+"NDropFeatures", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Check your data", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", true)
	return ctx.JSON(http.StatusOK, result)
}

package featuresservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

func (svc nFeaturesServices) UpdateFeaturesService(ctx echo.Context) error {
	var (
		svcName  = "UpdateFeaturesService"
		logErr   = "Error " + svcName
		req      models.ReqGetListNFeature
		response models.RespGetListNFeature
		dbTime   = time.Now().Format(time.RFC3339)
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println(logErr+"BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Data.ID == 0 {
		log.Println(logErr + " ID cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "ID cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Data.FeatureName = strings.ToUpper(req.Data.FeatureName)
	response, err = svc.services.ApiNFeatures.NUpdateFeature(models.ReqCreateNFeature{
		ID:          req.Data.ID,
		FeatureName: req.Data.FeatureName,
		MerchantId:  req.Data.MerchantId,
		CreatedAt:   dbTime,
		UpdatedAt:   dbTime,
		CreatedBy:   "sys",
		UpdatedBy:   "sys",
	})
	if err != nil {
		log.Println(logErr+"NUpdateFeatures", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", response)
	return ctx.JSON(http.StatusOK, result)
}

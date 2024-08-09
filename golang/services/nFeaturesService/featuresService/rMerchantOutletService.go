package featuresservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func (svc nFeaturesServices) GetFeaturesService(ctx echo.Context) error {
	var (
		svcName  = "GetFeaturesService"
		logErr   = "Error " + svcName
		req      models.ReqGetListNFeature
		response models.RespGetListNFeature
		// dbTime   = time.Now().Format(time.RFC3339)
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println(logErr+"BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Data.FeatureName == "" {
		log.Println(logErr + " Feature Name cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Feature Name cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Data.FeatureName = strings.ToUpper(req.Data.FeatureName)
	resp, err := svc.services.ApiNFeatures.NReadSingleFeature(req.Data)
	if err != nil {
		log.Println(logErr+"NReadSingleFeatures", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Not Found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	response.ID = resp.ID
	response.FeatureName = resp.FeatureName
	response.CreatedAt = resp.CreatedAt
	response.CreatedBy = resp.CreatedBy
	response.UpdatedAt = resp.UpdatedAt
	response.UpdatedBy = resp.UpdatedBy
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", response)
	return ctx.JSON(http.StatusOK, result)
}

func (svc nFeaturesServices) GetListFeaturesService(ctx echo.Context) error {
	var (
		svcName = "GetListFeaturesService"
		logErr  = "Error " + svcName
		req     models.ReqGetListNFeature
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println(logErr+"BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resP, err := svc.services.ApiNFeatures.NReadFeature(req)
	if err != nil {
		log.Println(logErr+"NReadFeatures", err.Error())
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

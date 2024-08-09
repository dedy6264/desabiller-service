package featuresservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/services"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

type nFeaturesServices struct {
	services services.UsecaseService
}

func NewApiNFeaturesServices(services services.UsecaseService) nFeaturesServices {
	return nFeaturesServices{
		services: services,
	}
}
func (svc nFeaturesServices) CreateFeaturesService(ctx echo.Context) error {
	var (
		svcName  = "CreateFeaturesService"
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
	if req.Data.FeatureName == "" {
		log.Println(logErr + " Feature Name cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Feature Name cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Data.FeatureName = strings.ToUpper(req.Data.FeatureName)
	id, err := svc.services.ApiNFeatures.NCreateFeature(models.ReqCreateNFeature{
		ID:          0,
		FeatureName: req.Data.FeatureName,
		MerchantId:  0,
		CreatedAt:   dbTime,
		UpdatedAt:   dbTime,
		CreatedBy:   "sys",
		UpdatedBy:   "sys",
	})
	if err != nil {
		log.Println(logErr+"NCreateFeature", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	response.ID = id
	response.FeatureName = req.Data.FeatureName
	response.CreatedAt = dbTime
	response.CreatedBy = "sys"
	response.UpdatedAt = dbTime
	response.UpdatedBy = "sys"
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", response)
	return ctx.JSON(http.StatusOK, result)
}

package featuresassignmentservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (svc nFeaturesAssignmentServices) UpdateFeaturesAssignmentService(ctx echo.Context) error {
	var (
		svcName  = "UpdateFeaturesAssignmentService"
		logErr   = "Error " + svcName
		req      models.ReqGetListNFeatureAssignment
		response models.RespGetListNFeatureAssignment
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
	response, err = svc.services.ApiNFeatures.NUpdateFeatureAssignment(models.ReqCreateNFeatureAssignment{
		ID:         req.Data.ID,
		FeatureId:  req.Data.FeatureId,
		MerchantId: req.Data.MerchantId,
		CreatedAt:  dbTime,
		UpdatedAt:  dbTime,
		CreatedBy:  "sys",
		UpdatedBy:  "sys",
	})
	if err != nil {
		log.Println(logErr+"NUpdateFeatureAssignment", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", response)
	return ctx.JSON(http.StatusOK, result)
}

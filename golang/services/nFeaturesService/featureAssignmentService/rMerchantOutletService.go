package featuresassignmentservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (svc nFeaturesAssignmentServices) GetFeaturesAssignmentService(ctx echo.Context) error {
	var (
		svcName  = "GetFeaturesAssignmentService"
		logErr   = "Error " + svcName
		req      models.ReqGetListNFeatureAssignment
		response models.RespGetListNFeatureAssignment
		// dbTime   = time.Now().Format(time.RFC3339)
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println(logErr+"BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.ApiNFeatures.NReadSingleFeatureAssignment(models.ReqCreateNFeatureAssignment{
		ID:         req.Data.ID,
		FeatureId:  req.Data.FeatureId,
		MerchantId: req.Data.MerchantId,
	})
	if err != nil {
		log.Println(logErr+"NReadSingleFeatures", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Not Found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	response.ID = resp.ID
	response.FeatureId = resp.FeatureId
	response.MerchantId = resp.MerchantId
	response.CreatedAt = resp.CreatedAt
	response.CreatedBy = resp.CreatedBy
	response.UpdatedAt = resp.UpdatedAt
	response.UpdatedBy = resp.UpdatedBy
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", response)
	return ctx.JSON(http.StatusOK, result)
}

func (svc nFeaturesAssignmentServices) GetListFeaturesAssignmentService(ctx echo.Context) error {
	var (
		svcName = "GetListFeaturesAssignmentService"
		logErr  = "Error " + svcName
		req     models.ReqGetListNFeatureAssignment
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println(logErr+"BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resP, err := svc.services.ApiNFeatures.NReadFeatureAssignment(req)
	if err != nil {
		log.Println(logErr+"NReadFeatureAssignment", err.Error())
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

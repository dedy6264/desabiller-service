package featuresassignmentservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/services"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type nFeaturesAssignmentServices struct {
	services services.UsecaseService
}

func NewApiNFeaturesAssignmentServices(services services.UsecaseService) nFeaturesAssignmentServices {
	return nFeaturesAssignmentServices{
		services: services,
	}
}
func (svc nFeaturesAssignmentServices) CreateFeaturesAssignmentService(ctx echo.Context) error {
	var (
		svcName  = "CreateFeaturesAssignmentService"
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
	if req.Data.FeatureId == 0 {
		log.Println(logErr + " Feature ID cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Feature ID cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Data.MerchantId == 0 {
		log.Println(logErr + " Merchant ID cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Merchant ID cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	idx, _ := strconv.Atoi((strconv.Itoa(req.Data.FeatureId) + strconv.Itoa(req.Data.MerchantId)))
	id, err := svc.services.ApiNFeatures.NCreateFeatureAssignment(models.ReqCreateNFeatureAssignment{
		ID:         0,
		FeatureId:  req.Data.FeatureId,
		MerchantId: req.Data.MerchantId,
		Index:      idx,
		CreatedAt:  dbTime,
		UpdatedAt:  dbTime,
		CreatedBy:  "sys",
		UpdatedBy:  "sys",
	})
	if err != nil {
		log.Println(logErr+"NCreateFeatureAssignment", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	response.ID = id
	response.FeatureId = req.Data.FeatureId
	response.MerchantId = req.Data.MerchantId
	response.CreatedAt = dbTime
	response.CreatedBy = "sys"
	response.UpdatedAt = dbTime
	response.UpdatedBy = "sys"
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", response)
	return ctx.JSON(http.StatusOK, result)
}

package helperservices

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/services"
	"desabiller/utils"
	"net/http"

	"github.com/labstack/echo"
)

type helperService struct {
	service services.UsecaseService
}

func NewApiHelperService(service services.UsecaseService) helperService {
	return helperService{service: service}
}
func (svc helperService) GetOperatorService(ctx echo.Context) error {
	var (
		svcName = "[HELPER]GetOperatorService"
		respSvc models.ResponseList
	)
	req := new(models.GetPrefix)
	//binding *req
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_VALIDATION_FAILED[0], configs.RC_VALIDATION_FAILED[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	} //tes validation
	if len(req.SubscriberId) < 5 {
		utils.Log(" Subcriber count char", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_INVALID_PARAM[0], "karakter kurang dari 5 digit", "karakter kurang dari 5 digit", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.SubscriberId = utils.NumberFixer(req.SubscriberId)
	if req.SubscriberId == "" {
		utils.Log(" NumberFixer", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_INVALID_PARAM[0], "format number error", "format number error", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.service.HelperRepo.GetProductReferenceById(req.SubscriberId)
	if err != nil {
		utils.Log(" GetProductReferenceById", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			configs.RC_SYSTEM_ERROR[1], nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respSvc.Data = resp
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.RC_SUCCESS[0],
		configs.RC_SUCCESS[1],
		configs.RC_SUCCESS[1], respSvc)
	return ctx.JSON(http.StatusOK, result)
}

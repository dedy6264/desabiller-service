package helperservices

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/services"
	"desabiller/utils"
	"log"
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
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	} //tes validation
	if len(req.SubscriberId) < 5 {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "karakter kurang dari 5 digit", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.SubscriberId = utils.NumberFixer(req.SubscriberId)
	if req.SubscriberId == "" {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "format number error", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.service.HelperRepo.GetProductReferenceById(req.SubscriberId)
	if err != nil {
		log.Println("Err ", svcName, " GetProductReferenceById ", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "operator :: not found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respSvc.Data = resp
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "SUCCESS", respSvc)
	return ctx.JSON(http.StatusOK, result)
}

package hierarchyservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func (svc HierarcyService) UpdateMerchantOutlet(ctx echo.Context) error {
	var (
		svcName = "UpdateMerchantOutlet"
	)
	req := new(models.ReqGetListMerchantOutlet)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.MerchantOutletName == "" {
		log.Println("Error " + svcName + " // Validated :: MerchantOutletName is null")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.MerchantOutletName = strings.ToUpper(req.MerchantOutletName)
	// resMerch, status := svc.service.ApiHierarchy.GetListMerchant(*req)
	// if len(resMerch) == 0 {
	// 	log.Println("Error " + svcName + " // GetListMerchant :: ")
	// 	respGlobal.StatusCode = configs.DB_NOT_FOUND
	// 	respGlobal.Message = "Not Found"
	// 	respGlobal.Success = false
	// 	respGlobal.ResponseDatetime = dbTime
	// 	return ctx.JSON(http.StatusOK, result)
	// }
	resUpdt, status := svc.service.ApiHierarchy.UpdateMerchantOutlet(*req)
	if !status {
		log.Println("Error " + svcName + " // UpdateMerchant :: ")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success ", resUpdt)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) DropMerchantOutlet(ctx echo.Context) error {
	var (
		svcName = "DropMerchantOutlet"
	)
	req := new(models.ReqGetListMerchantOutlet)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.ID == 0 {
		log.Println("Error " + svcName + " // Validated :: id is null")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	status = svc.service.ApiHierarchy.DropMerchantOutlet(*req)
	if !status {
		log.Println("Error " + svcName + " // DropMerchant :: ")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) GetListMerchantOutlet(ctx echo.Context) error {
	var (
		svcName = "GetListMerchantOutlet"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetListMerchantOutlet)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	req.MerchantOutletName = strings.ToUpper(req.MerchantOutletName)
	req.MerchantName = strings.ToUpper(req.MerchantName)

	count, status := svc.service.ApiHierarchy.GetListMerchantOutletCount(*req)
	if count == 0 {
		log.Println("Error " + svcName + " // GetListMerchantOutletCount :: ")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resMerch, status := svc.service.ApiHierarchy.GetListMerchantOutlet(*req)
	if !status {
		log.Println("Error " + svcName + " // GetListMerchantOutlet :: ")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respSvc.Data = resMerch
	respSvc.TotalData = count
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", respSvc)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) AddMerchantOutlet(ctx echo.Context) error {
	var (
		svcName = "AddMerchantOutlet"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetListMerchantOutlet)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	req.MerchantOutletName = strings.ToUpper(req.MerchantOutletName)
	if req.ClientId == 0 {
		log.Println("Error " + svcName + " :: clientId invalid")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Client id invalid", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.MerchantId == 0 {
		log.Println("Error " + svcName + " :: merchantId invalid")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Merchant id invalid", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.MerchantOutletName == "" {
		log.Println("Error " + svcName + " :: merchantOutletName invalid")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Merchant Outlet name id invalid", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resCLi, _ := svc.service.ApiHierarchy.GetListClientCount(models.ReqGetListClient{
		ID:         req.ClientId,
		ClientName: "",
		Limit:      0,
		Offset:     0,
		OrderBy:    "",
		StartDate:  "",
		EndDate:    "",
		Username:   "",
	})
	if resCLi != 1 {
		log.Println("Error " + svcName + " // GetListClientCount :: client Not Found")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Not Found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resMer, _ := svc.service.ApiHierarchy.GetListMerchantCount(models.ReqGetListMerchant{
		ID:           req.MerchantId,
		MerchantName: "",
		ClientId:     0,
		ClientName:   "",
		Limit:        0,
		Offset:       0,
		OrderBy:      "",
		StartDate:    "",
		EndDate:      "",
		Username:     "",
	})
	if resMer != 1 {
		log.Println("Error " + svcName + " // GetListMerchantCount :: Merchant Not Found")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Not Found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	//cek existensi merchant name
	resMerch, status := svc.service.ApiHierarchy.GetListMerchantOutlet(models.ReqGetListMerchantOutlet{
		ID:                 0,
		MerchantOutletName: req.MerchantOutletName,
		MerchantId:         req.MerchantId,
		ClientId:           req.ClientId,
		Limit:              0,
		Offset:             0,
		OrderBy:            "",
		StartDate:          "",
		EndDate:            "",
		Username:           "",
	})
	if status {
		if len(resMerch) > 0 {
			for _, data := range resMerch {
				if data.MerchantOutletName == req.MerchantOutletName {
					log.Println("Error " + svcName + " // GetListMerchantOutlet :: rejected merchant name")
					result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "rejected", nil)
					return ctx.JSON(http.StatusOK, result)
				}
			}
		}
	}

	//simpan
	status = svc.service.ApiHierarchy.AddMerchantOutlet(models.ReqGetListMerchantOutlet{
		ID:                 0,
		MerchantOutletName: req.MerchantOutletName,
		MerchantId:         req.MerchantId,
		ClientId:           req.ClientId,
		Limit:              0,
		Offset:             0,
		OrderBy:            "",
		StartDate:          "",
		EndDate:            "",
		Username:           "",
	})
	if !status {
		log.Println("Error " + svcName + " // AddMerchantOutlet :: failed add merchant")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respSvc.Data = req
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", respSvc)
	return ctx.JSON(http.StatusOK, result)
}

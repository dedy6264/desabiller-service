package hierarchyservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

func (svc HierarcyService) UpdateOutletDevice(ctx echo.Context) error {
	var (
		svcName    = "UpdateOutletDevice"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListOutletDevice)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.DeviceSn == "" {
		log.Println("Error " + svcName + " :: Nickname invalid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Nickname id invalid"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.DeviceType == "" {
		log.Println("Error " + svcName + " :: OutletPassword invalid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "OutletPassword id invalid"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}

	resUpdt, status := svc.service.ApiHierarchy.UpdateOutletDevice(*req)
	if !status {
		log.Println("Error " + svcName + " // UpdateDevice :: ")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Message = "Success"
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	respGlobal.Result = resUpdt
	return ctx.JSON(http.StatusOK, respGlobal)

}

func (svc HierarcyService) DropOutletDevice(ctx echo.Context) error {
	var (
		svcName    = "DropOutletDevice"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListOutletDevice)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.ID == 0 {
		log.Println("Error " + svcName + " // Validated :: id is null")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	status = svc.service.ApiHierarchy.DropOutletDevice(*req)
	if !status {
		log.Println("Error " + svcName + " // DropUserOutlet :: ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "Failed"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Message = "Success"
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	return ctx.JSON(http.StatusOK, respGlobal)
}
func (svc HierarcyService) GetListOutletDevice(ctx echo.Context) error {
	var (
		svcName    = "GetListOutletDevice"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
		respSvc    models.ResponseList
	)
	req := new(models.ReqGetListOutletDevice)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	req.DeviceSn = strings.ToUpper(req.DeviceSn)
	req.DeviceType = strings.ToUpper(req.DeviceType)
	req.MerchantOutletName = strings.ToUpper(req.MerchantOutletName)
	req.MerchantName = strings.ToUpper(req.MerchantName)

	count, status := svc.service.ApiHierarchy.GetListOutletDeviceCount(*req)
	if count == 0 {
		log.Println("Error " + svcName + " // GetListOutletDeviceCount :: ")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Message = "Not Found"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	resMerch, status := svc.service.ApiHierarchy.GetListOutletDevice(*req)
	if !status {
		log.Println("Error " + svcName + " // GetListOutletDevice :: ")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Message = "Not Found"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	respSvc.Data = resMerch
	respSvc.TotalData = count
	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Message = "Success"
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	respGlobal.Result = respSvc
	return ctx.JSON(http.StatusOK, respGlobal)
}

func (svc HierarcyService) AddOutletDevice(ctx echo.Context) error {
	var (
		svcName    = "AddOutletDevice"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
		respSvc    models.ResponseList
	)
	req := new(models.ReqGetListOutletDevice)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	// req.MerchantName = strings.ToUpper(req.MerchantName)
	// req.MerchantOutletName = strings.ToUpper(req.MerchantOutletName)
	if req.ClientId == 0 {
		log.Println("Error " + svcName + " :: clientId invalid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Client id invalid"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.MerchantId == 0 {
		log.Println("Error " + svcName + " :: merchantId invalid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Merchant id invalid"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.MerchantOutletId == 0 {
		log.Println("Error " + svcName + " :: merchantOutletId invalid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Merchant Outlet id invalid"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.DeviceSn == "" {
		log.Println("Error " + svcName + " :: SN invalid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Nickname id invalid"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.DeviceType == "" {
		log.Println("Error " + svcName + " :: Type invalid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "OutletPassword id invalid"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
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
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Message = "Not Found"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
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
	fmt.Println(resMer)

	if resMer != 1 {
		log.Println("Error " + svcName + " // GetListMerchantCount :: Merchant Not Found")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Message = "Not Found"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	resMeOu, _ := svc.service.ApiHierarchy.GetListMerchantOutletCount(models.ReqGetListMerchantOutlet{
		ID:                 req.MerchantOutletId,
		MerchantOutletName: "",
		MerchantId:         0,
		MerchantName:       "",
		ClientId:           0,
		Limit:              0,
		Offset:             0,
		OrderBy:            "",
		StartDate:          "",
		EndDate:            "",
		Username:           "",
	})
	if resMeOu != 1 {
		log.Println("Error " + svcName + " // GetListMerchantOutletCount :: Merchant Outlet Not Found")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Message = "Not Found"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	//cek existensi merchant name
	resMerch, status := svc.service.ApiHierarchy.GetListOutletDevice(*req)
	if status {
		if len(resMerch) > 0 {
			for _, data := range resMerch {
				if data.DeviceSn == req.DeviceSn {
					log.Println("Error " + svcName + " // GetListOutletDevice :: rejected username")
					respGlobal.StatusCode = configs.DB_NOT_FOUND
					respGlobal.Message = "rejected"
					respGlobal.Success = false
					respGlobal.ResponseDatetime = dbTime
					return ctx.JSON(http.StatusOK, respGlobal)
				}
			}
		}
	}

	//simpan
	status = svc.service.ApiHierarchy.AddOutletDevice(*req)
	if !status {
		log.Println("Error " + svcName + " // AddOutletDevice :: failed add device")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	respSvc.Data = req
	respGlobal.Result = respSvc
	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	return ctx.JSON(http.StatusOK, respGlobal)
}

package hierarchyservice

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

func (svc HierarcyService) UpdateMerchant(ctx echo.Context) error {
	var (
		svcName    = "UpdateMerchant"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListMerchant)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	req.MerchantName = strings.ToUpper(req.MerchantName)
	// resMerch, status := svc.service.ApiHierarchy.GetListMerchant(*req)
	// if len(resMerch) == 0 {
	// 	log.Println("Error " + svcName + " // GetListMerchant :: ")
	// 	respGlobal.StatusCode = configs.DB_NOT_FOUND
	// 	respGlobal.Message = "Not Found"
	// 	respGlobal.Success = false
	// 	respGlobal.ResponseDatetime = dbTime
	// 	return ctx.JSON(http.StatusOK, respGlobal)
	// }
	resUpdt, status := svc.service.ApiHierarchy.UpdateMerchant(*req)
	if !status {
		log.Println("Error " + svcName + " // UpdateMerchant :: ")
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
func (svc HierarcyService) DropMerchant(ctx echo.Context) error {
	var (
		svcName    = "DropMerchant"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListMerchant)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	status = svc.service.ApiHierarchy.DropMerchant(*req)
	if !status {
		log.Println("Error " + svcName + " // DropMerchant :: ")
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
func (svc HierarcyService) GetListMerchant(ctx echo.Context) error {
	var (
		svcName    = "GetListMerchant"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
		respSvc    models.ResponseList
	)
	req := new(models.ReqGetListMerchant)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	req.MerchantName = strings.ToUpper(req.MerchantName)
	req.ClientName = strings.ToUpper(req.ClientName)

	count, status := svc.service.ApiHierarchy.GetListMerchantCount(*req)
	if count == 0 {
		log.Println("Error " + svcName + " // GetListMerchantCount :: ")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Message = "Not Found"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	resMerch, status := svc.service.ApiHierarchy.GetListMerchant(*req)
	if !status {
		log.Println("Error " + svcName + " // GetListMerchant :: ")
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
func (svc HierarcyService) AddMerchant(ctx echo.Context) error {
	var (
		svcName    = "AddMerchant"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
		respSvc    models.ResponseList
	)
	req := new(models.ReqGetListMerchant)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	req.MerchantName = strings.ToUpper(req.MerchantName)
	if req.ClientId == 0 {
		log.Println("Error " + svcName + " :: clientId invalid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Client id invalid"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.MerchantName == "" {
		log.Println("Error " + svcName + " :: merchantName invalid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Merchant  name id invalid"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	// cek client id
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
	//cek existensi merchant name
	resMerch, status := svc.service.ApiHierarchy.GetListMerchant(models.ReqGetListMerchant{
		ID:           0,
		MerchantName: req.MerchantName,
		ClientId:     req.ClientId,
		Limit:        0,
		Offset:       0,
		OrderBy:      "",
		StartDate:    "",
		EndDate:      "",
		Username:     "",
	})
	if status {
		if len(resMerch) > 0 {
			for _, data := range resMerch {
				if data.MerchantName == req.MerchantName {
					log.Println("Error " + svcName + " // GetListMerchant :: rejected merchant name")
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
	status = svc.service.ApiHierarchy.AddMerchant(models.ReqGetListMerchant{
		ID:           0,
		MerchantName: req.MerchantName,
		ClientId:     req.ClientId,
		Limit:        0,
		Offset:       0,
		OrderBy:      "",
		StartDate:    "",
		EndDate:      "",
		Username:     "",
	})
	if !status {
		log.Println("Error " + svcName + " // AddMerchant :: failed add merchant")
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

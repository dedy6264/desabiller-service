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

func (svc HierarcyService) UpdateUserOutlet(ctx echo.Context) error {
	var (
		svcName    = "UpdateUserOutlet"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListUserOutlet)
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

	if req.OutletPassword != "" {
		req.OutletPassword, err = helpers.PswEnc(req.OutletPassword)
		fmt.Println(":::::", req.OutletPassword)
		if err != nil {
			log.Println("Error " + svcName + " :: Encrypt password faild")
			respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
			respGlobal.Message = "Failed"
			respGlobal.Success = false
			respGlobal.ResponseDatetime = dbTime
			return ctx.JSON(http.StatusOK, respGlobal)
		}
	}

	resUpdt, status := svc.service.ApiHierarchy.UpdateUserOutlet(*req)
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

func (svc HierarcyService) DropUserOutlet(ctx echo.Context) error {
	var (
		svcName    = "DropUserOutlet"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListUserOutlet)
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
	status = svc.service.ApiHierarchy.DropUserOutlet(*req)
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
func (svc HierarcyService) GetListUserOutlet(ctx echo.Context) error {
	var (
		svcName    = "GetListUserOutlet"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
		respSvc    models.ResponseList
	)
	req := new(models.ReqGetListUserOutlet)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	req.MerchantOutletName = strings.ToUpper(req.MerchantOutletName)
	req.MerchantName = strings.ToUpper(req.MerchantName)

	count, status := svc.service.ApiHierarchy.GetListUserOutletCount(*req)
	if count == 0 {
		log.Println("Error " + svcName + " // GetListUserOutletCount :: ")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Message = "Not Found"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	resMerch, status := svc.service.ApiHierarchy.GetListUserOutlet(*req)
	if !status {
		log.Println("Error " + svcName + " // GetListUserOutlet :: ")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Message = "Not Found"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	for i, _ := range resMerch {
		resMerch[i].OutletPassword = ""
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

func (svc HierarcyService) AddUserOutlet(ctx echo.Context) error {
	var (
		svcName    = "AddUserOutlet"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
		respSvc    models.ResponseList
	)
	req := new(models.ReqGetListUserOutlet)
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
	if req.Nickname == "" {
		log.Println("Error " + svcName + " :: Nickname invalid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Nickname id invalid"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.OutletPassword == "" {
		log.Println("Error " + svcName + " :: OutletPassword invalid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "OutletPassword id invalid"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.OutletUsername == "" {
		log.Println("Error " + svcName + " :: OutletUsername invalid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "OutletUsername id invalid"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	req.OutletPassword, err = helpers.PswEnc(req.OutletPassword)
	if err != nil {
		log.Println("Error " + svcName + " :: Encrypt password faild")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	fmt.Println(":::PSWRD::", req.OutletPassword)
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
	resMerch, status := svc.service.ApiHierarchy.GetListUserOutlet(models.ReqGetListUserOutlet{
		ID:             0,
		OutletUsername: req.OutletUsername,
		MerchantId:     req.MerchantId,
		ClientId:       req.ClientId,
		Limit:          0,
		Offset:         0,
		OrderBy:        "",
		StartDate:      "",
		EndDate:        "",
		Username:       "",
	})
	if status {
		if len(resMerch) > 0 {
			for _, data := range resMerch {
				if data.OutletUsername == req.OutletUsername {
					log.Println("Error " + svcName + " // GetListUserOutlet :: rejected username")
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
	status = svc.service.ApiHierarchy.AddUserOutlet(*req)
	if !status {
		log.Println("Error " + svcName + " // AddUserOutlet :: failed add merchant")
		respGlobal.StatusCode = configs.DB_NOT_FOUND
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	req.OutletPassword = ""
	respSvc.Data = req
	respGlobal.Result = respSvc
	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	return ctx.JSON(http.StatusOK, respGlobal)
}

package productservice

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

func (svc ProductService) GetListProductBillerProvider(ctx echo.Context) error {
	var (
		svcName    = "GetListProductBillerProvider"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListProductBillerProvider)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	resProTy, _ := svc.service.ApiProduct.GetListProductBillerProvider(*req)
	if len(resProTy) == 0 {
		log.Println("Error " + svcName + " // not found :: ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}

	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Message = "SUCCESS"
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	respGlobal.Result = resProTy
	return ctx.JSON(http.StatusOK, respGlobal)
}

func (svc ProductService) AddProductBillerProvider(ctx echo.Context) error {
	var (
		svcName    = "AddProductBillerProvider"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListProductBillerProvider)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	req.ProductProviderName = strings.ToUpper(req.ProductProviderName)
	req.ProductProviderCode = strings.ToUpper(req.ProductProviderCode)

	if req.ProductCategoryId == 0 || req.ProductTypeId == 0 {
		log.Println("Error " + svcName + " // id :: not valid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.ProductProviderName == "" || req.ProductProviderCode == "" {
		log.Println("Error " + svcName + " // updateable :: must true")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	resProTy, _ := svc.service.ApiProduct.GetListProductBillerProvider(models.ReqGetListProductBillerProvider{
		ProductProviderCode: req.ProductProviderCode,
	})
	fmt.Println(":::", resProTy)
	if len(resProTy) > 0 {
		log.Println("Error " + svcName + " // not found :: ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	_, status = svc.service.ApiProduct.AddProductBillerProvider(*req)
	if !status {
		log.Println("Error " + svcName + " // Failed :: ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Message = "SUCCESS"
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	respGlobal.Result = resProTy
	return ctx.JSON(http.StatusOK, respGlobal)
}

func (svc ProductService) UpdateProductBillerProvider(ctx echo.Context) error {
	var (
		svcName    = "UpdateProductBillerProvider"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListProductBillerProvider)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.ProductProviderName == "" {
		log.Println("Error " + svcName + " // product prov name :: not valid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.ProductProviderCode == "" {
		log.Println("Error " + svcName + " // product prov code :: not valid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.ProductCategoryId == 0 {
		log.Println("Error " + svcName + " // category product :: not valid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.ProductTypeId == 0 {
		log.Println("Error " + svcName + " // Type Product :: not valid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}

	req.ProductProviderName = strings.ToUpper(req.ProductProviderName)
	req.ProductProviderCode = strings.ToUpper(req.ProductProviderCode)

	resProTy, _ := svc.service.ApiProduct.GetListProductBillerProvider(models.ReqGetListProductBillerProvider{
		ID:                         req.ID,
		ProductProviderName:        "",
		ProductProviderCode:        "",
		ProductProviderPrice:       0,
		ProductProviderAdminFee:    0,
		ProductProviderMerchantFee: 0,
		IsOpen:                     false,
		ProductTypeId:              0,
		ProductCategoryId:          0,
		Limit:                      0,
		Offset:                     0,
		OrderBy:                    "",
		StartDate:                  "",
		EndDate:                    "",
		Username:                   "",
	})
	fmt.Println(":::", resProTy)
	if len(resProTy) == 0 {
		log.Println("Error " + svcName + " // not found :: ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}

	_, status = svc.service.ApiProduct.UpdateProductBillerProvider(*req)
	if !status {
		log.Println("Error " + svcName + " // Failed :: ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Message = "SUCCESS"
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	respGlobal.Result = req
	return ctx.JSON(http.StatusOK, respGlobal)
}

func (svc ProductService) DropProductBillerProvider(ctx echo.Context) error {
	var (
		svcName    = "DropProductBillerProvider"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListProductBillerProvider)
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
	resProTy, _ := svc.service.ApiProduct.GetListProductBillerProvider(models.ReqGetListProductBillerProvider{
		ID:                         req.ID,
		ProductProviderName:        "",
		ProductProviderCode:        "",
		ProductProviderPrice:       0,
		ProductProviderAdminFee:    0,
		ProductProviderMerchantFee: 0,
		IsOpen:                     false,
		ProductTypeId:              0,
		ProductCategoryId:          0,
		Limit:                      0,
		Offset:                     0,
		OrderBy:                    "",
		StartDate:                  "",
		EndDate:                    "",
		Username:                   "",
	})
	if len(resProTy) == 0 {
		log.Println("Error " + svcName + " // not found :: ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	for _, data := range resProTy {
		if data.IsOpen {
			log.Println("Error " + svcName + " // product is Open :: ")
			respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
			respGlobal.Message = "Failed"
			respGlobal.Success = false
			respGlobal.ResponseDatetime = dbTime
			return ctx.JSON(http.StatusOK, respGlobal)
		}
	}
	status = svc.service.ApiProduct.DropProductBillerProvider(*req)
	if !status {
		log.Println("Error " + svcName + " // DropProductBillerProvider :: ")
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

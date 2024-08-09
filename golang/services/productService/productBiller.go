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

func (svc ProductService) GetListProductBiller(ctx echo.Context) error {
	var (
		svcName    = "GetListProductBiller"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListProductBiller)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	resProTy, _ := svc.service.ApiProduct.GetListProductBiller(*req)
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

func (svc ProductService) AddProductBiller(ctx echo.Context) error {
	var (
		svcName    = "AddProductBiller"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListProductBiller)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	req.ProductName = strings.ToUpper(req.ProductName)
	req.ProductCode = strings.ToUpper(req.ProductCode)

	if req.ProductCategoryId == 0 || req.ProductTypeId == 0 {
		log.Println("Error " + svcName + " // id :: not valid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.ProductName == "" || req.ProductCode == "" {
		log.Println("Error " + svcName + " // updateable :: must true")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	resProTy, _ := svc.service.ApiProduct.GetListProductBiller(models.ReqGetListProductBiller{
		ProductCode: req.ProductCode,
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
	_, status = svc.service.ApiProduct.AddProductBiller(*req)
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

func (svc ProductService) UpdateProductBiller(ctx echo.Context) error {
	var (
		svcName    = "UpdateProductBiller"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListProductBiller)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.ProductName == "" {
		log.Println("Error " + svcName + " // product prov name :: not valid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.ProductCode == "" {
		log.Println("Error " + svcName + " // product prov code :: not valid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.ProductProviderId == 0 {
		log.Println("Error " + svcName + " // provider product :: not valid")
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

	req.ProductName = strings.ToUpper(req.ProductName)
	req.ProductCode = strings.ToUpper(req.ProductCode)

	resProTy, _ := svc.service.ApiProduct.GetListProductBiller(models.ReqGetListProductBiller{
		ID:                req.ID,
		ProductName:       "",
		ProductCode:       "",
		ProductProviderId: 0,
		IsOpen:            false,
		ProductTypeId:     0,
		ProductCategoryId: 0,
		Limit:             0,
		Offset:            0,
		OrderBy:           "",
		StartDate:         "",
		EndDate:           "",
		Username:          "",
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

	_, status = svc.service.ApiProduct.UpdateProductBiller(*req)
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

func (svc ProductService) DropProductBiller(ctx echo.Context) error {
	var (
		svcName    = "DropProductBiller"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
	)
	req := new(models.ReqGetListProductBiller)
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
	resProTy, _ := svc.service.ApiProduct.GetListProductBiller(models.ReqGetListProductBiller{
		ID:                req.ID,
		ProductName:       "",
		ProductCode:       "",
		ProductProviderId: 0,
		IsOpen:            false,
		ProductTypeId:     0,
		ProductCategoryId: 0,
		Limit:             0,
		Offset:            0,
		OrderBy:           "",
		StartDate:         "",
		EndDate:           "",
		Username:          "",
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
	status = svc.service.ApiProduct.DropProductBiller(*req)
	if !status {
		log.Println("Error " + svcName + " // DropProductBiller :: ")
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

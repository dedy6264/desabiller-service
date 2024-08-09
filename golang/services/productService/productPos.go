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

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
)

func (svc ProductService) GetListProductPos(ctx echo.Context) error {
	var (
		svcName    = "GetListProductPos"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
		resSvc     models.ResponseList
		// snDevice   string
		// cID float64
		// oID float64
		mID float64
		req models.ReqGetListProductPos
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	a := ctx.Get("user").(*jwt.Token)
	fmt.Println("::CEK JWT:: ", a)
	claim := a.Claims.(jwt.MapClaims)
	fmt.Println("::CEK CLAIM:: ", claim)
	if a != nil {
		claim := a.Claims.(jwt.MapClaims)
		if claim["userDashboardId"] == nil {
			// cID = claim["clientId"].(float64)
			// oID = claim["outletId"].(float64)
			mID = claim["merchantId"].(float64)
			req.MerchantId = int(mID)
			// req.MerchantOutletId = int(oID)
			// req.ClientId = int(cID)
		} else {
			res, err := svc.service.ApiNUserDashboard.NReadSingleUserDashboard(models.ReqCreateNUserDashboard{
				ID: int(claim["userDashboardId"].(float64)),
			})
			if err != nil {
				log.Println("FAILLED Get data user", err.Error())
				result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED Get data user "+err.Error(), nil)
				return ctx.JSON(http.StatusNotFound, result)
			}
			if res.ClientId != (-1) {
				// if res.ClientId != 0 {
				// 	req.ClientId = res.ClientId
				// }
				if res.MerchantId != 0 {
					req.MerchantId = res.MerchantId
				}
				// if res.MerchantOutletId != 0 {
				// 	req.MerchantOutletId = res.MerchantOutletId
				// }
			}
		}
	}
	if req.Draw == 0 {
		req.Draw = 1
	}
	req.Offset = (req.Limit * req.Draw) - req.Limit
	resProTy, status := svc.service.ApiProduct.GetListProductPos(req)
	// fmt.Println("::::::::>>>>", resProTy, "::", status)

	if !status {
		log.Println("Error " + svcName + " // not found :: ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}

	resSvc.TotalData = 0
	resSvc.TotalRow = 0
	resSvc.Data = resProTy
	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Message = "SUCCESS"
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	respGlobal.Result = resSvc
	return ctx.JSON(http.StatusOK, respGlobal)
}

func (svc ProductService) AddProductPos(ctx echo.Context) error {
	var (
		svcName    = "AddProductPos"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
		// snDevice   string
		// uID int
		// oID int
		mID float64
		req models.ReqGetListProductPos
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.ProductPrice == 0 {
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "Product Price min 1"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	a := ctx.Get("user").(*jwt.Token)
	if a != nil {
		claim := a.Claims.(jwt.MapClaims)
		if claim["userDashboardId"] == nil {
			// cID = claim["clientId"].(float64)
			// oID = claim["outletId"].(float64)
			mID = claim["merchantId"].(float64)
			req.MerchantId = int(mID)
			// req.MerchantOutletId = int(oID)
			// req.ClientId = int(cID)
		} else {
			res, err := svc.service.ApiNUserDashboard.NReadSingleUserDashboard(models.ReqCreateNUserDashboard{
				ID: int(claim["userDashboardId"].(float64)),
			})
			if err != nil {
				log.Println("FAILLED Get data user", err.Error())
				result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED Get data user "+err.Error(), nil)
				return ctx.JSON(http.StatusNotFound, result)
			}
			if res.ClientId != (-1) {
				// if res.ClientId != 0 {
				// 	req.ClientId = res.ClientId
				// }
				if res.MerchantId != 0 {
					req.MerchantId = res.MerchantId
				}
				// if res.MerchantOutletId != 0 {
				// 	req.MerchantOutletId = res.MerchantOutletId
				// }
			}
		}
	}
	req.ProductName = strings.ToUpper(req.ProductName)
	req.ProductCode = strings.ToUpper(req.ProductCode)

	if req.ProductCategoryId == 0 || req.ProductTypeId == 0 {
		log.Println("Error " + svcName + " // id :: not valid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Result = "PRODUCT CATEGORY OR TYPE ID IS NOT VALID"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.ProductName == "" || req.ProductCode == "" {
		log.Println("Error " + svcName + " // Product name or code :: cannot be empty")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Result = " PRODUCT NAME OR CODE IS NOT VALID"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.MerchantId == 0 {
		log.Println("Error " + svcName + " // Merchant :: invalid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Result = " MERCHANT NAME OR ID IS NOT VALID"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	//harus cek merchant exist
	resProTy, status := svc.service.ApiProduct.GetListProductPos(models.ReqGetListProductPos{
		ProductCode: req.ProductCode,
		MerchantId:  req.MerchantId,
	})
	fmt.Println(":::", resProTy)
	if len(resProTy) > 0 {
		log.Println("Error " + svcName + " // not found :: ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Result = " PRODUCT IS EXIST"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	_, status = svc.service.ApiProduct.AddProductPos(req)
	if !status {
		log.Println("Error " + svcName + " // Failed :: ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Result = " ADD PRODUCT FAILED"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	respGlobal.StatusCode = configs.SUCCESS_CODE
	respGlobal.Message = "SUCCESS"
	respGlobal.Success = true
	respGlobal.ResponseDatetime = dbTime
	respGlobal.Result = resProTy
	return ctx.JSON(http.StatusOK, respGlobal)
}

func (svc ProductService) UpdateProductPos(ctx echo.Context) error {
	var (
		svcName    = "UpdateProductPos"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
		mID        float64
		req        models.ReqGetListProductPos
	)
	status, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}

	a := ctx.Get("user").(*jwt.Token)
	if a != nil {
		claim := a.Claims.(jwt.MapClaims)
		if claim["userDashboardId"] == nil {
			// cID = claim["clientId"].(float64)
			// oID = claim["outletId"].(float64)
			mID = claim["merchantId"].(float64)
			req.MerchantId = int(mID)
			// req.MerchantOutletId = int(oID)
			// req.ClientId = int(cID)
		} else {
			res, err := svc.service.ApiNUserDashboard.NReadSingleUserDashboard(models.ReqCreateNUserDashboard{
				ID: int(claim["userDashboardId"].(float64)),
			})
			if err != nil {
				log.Println("FAILLED Get data user", err.Error())
				result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED Get data user "+err.Error(), nil)
				return ctx.JSON(http.StatusNotFound, result)
			}
			if res.ClientId != (-1) {
				// if res.ClientId != 0 {
				// 	req.ClientId = res.ClientId
				// }
				if res.MerchantId != 0 {
					req.MerchantId = res.MerchantId
				}
				// if res.MerchantOutletId != 0 {
				// 	req.MerchantOutletId = res.MerchantOutletId
				// }
			}
		}
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
	if req.ProductPrice == 0 {
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		respGlobal.Message = "Product Price min 1"
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	req.ProductName = strings.ToUpper(req.ProductName)
	req.ProductCode = strings.ToUpper(req.ProductCode)

	resProTy, _ := svc.service.ApiProduct.GetListProductPos(models.ReqGetListProductPos{
		ID:                   req.ID,
		ProductName:          "",
		ProductCode:          "",
		ProductPriceProvider: 0,
		MerchantId:           0,
		MerchantName:         "",
		ProductPrice:         0,
		IsOpen:               false,
		ProductTypeId:        0,
		ProductCategoryId:    0,
		Limit:                0,
		Offset:               0,
		OrderBy:              "",
		StartDate:            "",
		EndDate:              "",
		Username:             "",
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

	_, status = svc.service.ApiProduct.UpdateProductPos(req)
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

func (svc ProductService) DropProductPos(ctx echo.Context) error {
	var (
		svcName    = "DropProductPos"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
		mID        float64
		req        models.ReqGetListProductPos
	)
	status, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	a := ctx.Get("user").(*jwt.Token)
	if a != nil {
		claim := a.Claims.(jwt.MapClaims)
		if claim["userDashboardId"] == nil {
			// cID = claim["clientId"].(float64)
			// oID = claim["outletId"].(float64)
			mID = claim["merchantId"].(float64)
			req.MerchantId = int(mID)
			// req.MerchantOutletId = int(oID)
			// req.ClientId = int(cID)
		} else {
			res, err := svc.service.ApiNUserDashboard.NReadSingleUserDashboard(models.ReqCreateNUserDashboard{
				ID: int(claim["userDashboardId"].(float64)),
			})
			if err != nil {
				log.Println("FAILLED Get data user", err.Error())
				result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED Get data user "+err.Error(), nil)
				return ctx.JSON(http.StatusNotFound, result)
			}
			if res.ClientId != (-1) {
				// if res.ClientId != 0 {
				// 	req.ClientId = res.ClientId
				// }
				if res.MerchantId != 0 {
					req.MerchantId = res.MerchantId
				}
				// if res.MerchantOutletId != 0 {
				// 	req.MerchantOutletId = res.MerchantOutletId
				// }
			}
		}
	}
	if req.ID == 0 {
		log.Println("Error " + svcName + " // Validated :: id is null")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	// resProTy, _ := svc.service.ApiProduct.GetListProductPos(models.ReqGetListProductPos{
	// 	ID:                   req.ID,
	// 	ProductName:          "",
	// 	ProductCode:          "",
	// 	ProductPriceProvider: 0,
	// 	MerchantId:           0,
	// 	MerchantName:         "",
	// 	ProductPrice:         0,
	// 	IsOpen:               false,
	// 	ProductTypeId:        0,
	// 	ProductCategoryId:    0,
	// 	Limit:                0,
	// 	Offset:               0,
	// 	OrderBy:              "",
	// 	StartDate:            "",
	// 	EndDate:              "",
	// 	Username:             "",
	// })
	// if len(resProTy) == 0 {
	// 	log.Println("Error " + svcName + " // not found :: ")
	// 	respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
	// 	respGlobal.Message = "Failed"
	// 	respGlobal.Success = false
	// 	respGlobal.ResponseDatetime = dbTime
	// 	return ctx.JSON(http.StatusOK, respGlobal)
	// }
	// for _, data := range resProTy {
	// 	if data.IsOpen {
	// 		log.Println("Error " + svcName + " // product is Open :: ")
	// 		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
	// 		respGlobal.Message = "Failed"
	// 		respGlobal.Success = false
	// 		respGlobal.ResponseDatetime = dbTime
	// 		return ctx.JSON(http.StatusOK, respGlobal)
	// 	}
	// }
	status = svc.service.ApiProduct.DropProductPos(req)
	if !status {
		log.Println("Error " + svcName + " // DropProductPos :: ")
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

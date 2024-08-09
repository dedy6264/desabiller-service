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

func (svc ProductService) GetListProductCategory(ctx echo.Context) error {
	var (
		svcName    = "GetListProductCategory"
		respGlobal models.Response
		resSvc     models.ResponseList
		dbTime     = time.Now().Format(time.RFC3339)
		// snDevice   string
		// uID        int
		oID float64
		mID float64
		cID float64
		req models.ReqGetListProductCategory
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	fmt.Println(":::::", ctx.Get("user").(*jwt.Token).Claims)
	a := ctx.Get("user").(*jwt.Token)
	if a != nil {
		claim := a.Claims.(jwt.MapClaims)
		if claim["userDashboardId"] == nil {
			// snDevice = claim["snDevice"].(string)
			// uID = claim["userId"].(int)
			cID = claim["clientId"].(float64)
			oID = claim["outletId"].(float64)
			mID = claim["merchantId"].(float64)
			req.MerchantId = int(mID)
			req.MerchantOutletId = int(oID)
			req.ClientId = int(cID)

			// claims := ctx.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
			// merchantOutletID := claims["merchantOutletID"].(float64)
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
				if res.ClientId != 0 {
					req.ClientId = res.ClientId
				}
				if res.MerchantId != 0 {
					req.MerchantId = res.MerchantId
				}
				if res.MerchantOutletId != 0 {
					req.MerchantOutletId = res.MerchantOutletId
				}
			}
		}
	}
	if req.Draw == 0 {
		req.Draw = 1
	}
	req.Offset = (req.Limit * req.Draw) - req.Limit
	resProTy, _ := svc.service.ApiProduct.GetListProductCategory(req)
	if len(resProTy) == 0 {
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
func (svc ProductService) AddProductCategory(ctx echo.Context) error {
	var (
		svcName    = "AddProductCategory"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
		oID        float64
		mID        float64
		cID        float64
		req        models.ReqGetListProductCategory
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	a := ctx.Get("user").(*jwt.Token)
	if a != nil {
		claim := a.Claims.(jwt.MapClaims)
		if claim["userDashboardId"] == nil {
			// snDevice = claim["snDevice"].(string)
			// uID = claim["userId"].(int)
			cID = claim["clientId"].(float64)
			oID = claim["outletId"].(float64)
			mID = claim["merchantId"].(float64)
			req.MerchantId = int(mID)
			req.MerchantOutletId = int(oID)
			req.ClientId = int(cID)
			// claims := ctx.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
			// merchantOutletID := claims["merchantOutletID"].(float64)
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
				if res.ClientId != 0 {
					req.ClientId = res.ClientId
				}
				if res.MerchantId != 0 {
					req.MerchantId = res.MerchantId
				}
				if res.MerchantOutletId != 0 {
					req.MerchantOutletId = res.MerchantOutletId
				}
			}
		}
	}
	req.ProductCategoryName = strings.ToUpper(req.ProductCategoryName)
	switch req.Updateable {
	case true:
		if req.MerchantOutletId == 0 {
			if req.MerchantId == 0 {
				if req.ClientId == 0 {
					log.Println("Error " + svcName + " // id :: not valid")
					respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
					respGlobal.Message = "Failed"
					respGlobal.Success = false
					respGlobal.ResponseDatetime = dbTime
					return ctx.JSON(http.StatusOK, respGlobal)
				}
			}
		}
	default:
		log.Println("Error " + svcName + " // updateable :: must true")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	//harus cek merchant exist
	resProTy, _ := svc.service.ApiProduct.GetListProductCategory(req)
	fmt.Println(":::", resProTy)
	if len(resProTy) != 0 {
		log.Println("Error " + svcName + " // not found :: ")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	_, status := svc.service.ApiProduct.AddProductCategory(req)
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
func (svc ProductService) UpdateProductCategory(ctx echo.Context) error {
	var (
		svcName    = "UpdateProductCategory"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
		oID        float64
		mID        float64
		cID        float64
		req        models.ReqGetListProductCategory
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
			cID = claim["clientId"].(float64)
			oID = claim["outletId"].(float64)
			mID = claim["merchantId"].(float64)
			req.MerchantId = int(mID)
			req.MerchantOutletId = int(oID)
			req.ClientId = int(cID)
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
				if res.ClientId != 0 {
					req.ClientId = res.ClientId
				}
				if res.MerchantId != 0 {
					req.MerchantId = res.MerchantId
				}
				if res.MerchantOutletId != 0 {
					req.MerchantOutletId = res.MerchantOutletId
				}
			}
		}
	}
	if req.ProductCategoryName == "" {
		log.Println("Error " + svcName + " // category name :: not valid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}
	if req.ProductCategoryCode == "" {
		log.Println("Error " + svcName + " // category code :: not valid")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}

	req.ProductCategoryName = strings.ToUpper(req.ProductCategoryName)
	switch req.Updateable {
	case true:
		if req.MerchantId == 0 || req.ClientId == 0 {
			log.Println("Error " + svcName + " // id :: not valid")
			respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
			respGlobal.Message = "Failed"
			respGlobal.Success = false
			respGlobal.ResponseDatetime = dbTime
			return ctx.JSON(http.StatusOK, respGlobal)
		}
	default:
		log.Println("Error " + svcName + " // updateable :: must true")
		respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
		respGlobal.Message = "Failed"
		respGlobal.Success = false
		respGlobal.ResponseDatetime = dbTime
		return ctx.JSON(http.StatusOK, respGlobal)
	}

	resProTy, _ := svc.service.ApiProduct.GetListProductCategory(models.ReqGetListProductCategory{
		ID:                  req.ID,
		ProductCategoryName: "",
		ProductCategoryCode: "",
		MerchantId:          req.MerchantId,
		MerchantName:        "",
		Updateable:          req.Updateable,
		ClientName:          "",
		ClientId:            req.ClientId,
		Limit:               0,
		Offset:              0,
		OrderBy:             "",
		StartDate:           "",
		EndDate:             "",
		Username:            "",
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

	_, status = svc.service.ApiProduct.UpdateProductCategory(req)
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
func (svc ProductService) DropProductCategory(ctx echo.Context) error {
	var (
		svcName    = "DropProductCategory"
		respGlobal models.Response
		dbTime     = time.Now().Format(time.RFC3339)
		oID        float64
		mID        float64
		cID        float64
		req        models.ReqGetListProductCategory
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	a := ctx.Get("user").(*jwt.Token)
	if a != nil {
		claim := a.Claims.(jwt.MapClaims)
		if claim["userDashboardId"] == nil {
			cID = claim["clientId"].(float64)
			oID = claim["outletId"].(float64)
			mID = claim["merchantId"].(float64)
			req.MerchantId = int(mID)
			req.MerchantOutletId = int(oID)
			req.ClientId = int(cID)
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
				if res.ClientId != 0 {
					req.ClientId = res.ClientId
				}
				if res.MerchantId != 0 {
					req.MerchantId = res.MerchantId
				}
				if res.MerchantOutletId != 0 {
					req.MerchantOutletId = res.MerchantOutletId
				}
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
	resProTy, _ := svc.service.ApiProduct.GetListProductCategory(models.ReqGetListProductCategory{
		ID:                  req.ID,
		ProductCategoryName: "",
		ProductCategoryCode: "",
		MerchantId:          req.MerchantId,
		MerchantName:        "",
		Updateable:          req.Updateable,
		ClientName:          "",
		ClientId:            req.ClientId,
		Limit:               0,
		Offset:              0,
		OrderBy:             "",
		StartDate:           "",
		EndDate:             "",
		Username:            "",
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
		if !data.Updateable {
			log.Println("Error " + svcName + " // not found :: ")
			respGlobal.StatusCode = configs.VALIDATE_ERROR_CODE
			respGlobal.Message = "Failed"
			respGlobal.Success = false
			respGlobal.ResponseDatetime = dbTime
			return ctx.JSON(http.StatusOK, respGlobal)
		}
	}
	status := svc.service.ApiProduct.DropProductCategory(req)
	if !status {
		log.Println("Error " + svcName + " // DropProductCategory :: ")
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

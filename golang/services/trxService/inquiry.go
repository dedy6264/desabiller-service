package trxservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (svc trxService) InquiryBiller(ctx echo.Context) error {
	var (
		svcName = "InquiryBiller"
		// respSvc models.ResponseList
		respOutlet models.RespGetMerchantOutlet
	)
	//get product
	//check price-provider price
	// check merchant fee-provider merchant fee
	//hit inq to partner
	//record to db
	//assing to response
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMPTRX)
	dbTimeTrx := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	req := new(models.ReqGetTrx)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ProductId == 0 {
		log.Println("Err ", svcName, "product id cannot be null")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "product id cannot be null", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.CustomerId == "" {
		log.Println("Err ", svcName, "customer id cannot be null")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "customer id cannot be null", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ReferenceNumber, err = svc.services.ApiTrx.GenerateNo("DB-"+dbTime, "", 7)
	if err != nil {
		log.Println("Err ", svcName, "GenerateNo", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respProduct, err := svc.services.ApiProduct.GetProduct(models.ReqGetProduct{
		ID: req.ProductId,
	})
	if err != nil {
		log.Println("Err ", svcName, "GetProduct", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if respProduct.ProductTypeId == 2 { //PREPAID
		if respProduct.ProductPrice < respProduct.ProductProviderPrice {
			log.Println("Err ", svcName, "product price invalid", respProduct.ProductPrice, respProduct.ProductProviderPrice)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "product price invalid", nil)
			return ctx.JSON(http.StatusOK, result)
		}
	} else {
		if respProduct.ProductMerchantFee > respProduct.ProductProviderMerchantFee {
			log.Println("Err ", svcName, "product merchant fee invalid", respProduct.ProductMerchantFee, respProduct.ProductProviderMerchantFee)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "product merchant fee invalid", nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}

	if req.ClientId == 0 {
		data := helpers.TokenJWTDecode(ctx)
		respOutlet, err = svc.services.ApiHierarchy.GetMerchantOutlet(models.ReqGetMerchantOutlet{
			ID: data.MerchantOutletId,
		})
		if err != nil {
			log.Println("Err ", svcName, "GetMerchantOutlet", err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}
	recordInq := models.ReqGetTrx{
		ProductClanId:              respProduct.ProductClanId,
		ProductClanName:            respProduct.ProductClanName,
		ProductCategoryId:          respProduct.ProductCategoryId,
		ProductCategoryName:        respProduct.ProductCategoryName,
		ProductTypeId:              respProduct.ProductTypeId,
		ProductTypeName:            respProduct.ProductTypeName,
		ProductId:                  respProduct.ID,
		ProductName:                respProduct.ProductName,
		ProductCode:                respProduct.ProductCode,
		ProductPrice:               respProduct.ProductPrice,
		ProductAdminFee:            respProduct.ProductAdminFee,
		ProductMerchantFee:         respProduct.ProductMerchantFee,
		ProductProviderId:          respProduct.ProviderId,
		ProductProviderName:        respProduct.ProviderName,
		ProductProviderCode:        respProduct.ProductCode,
		ProductProviderPrice:       respProduct.ProductProviderPrice,
		ProductProviderAdminFee:    respProduct.ProductProviderAdminFee,
		ProductProviderMerchantFee: respProduct.ProductProviderMerchantFee,
		StatusCode:                 configs.INQUIRY_SUCCESS_CODE,
		StatusMessage:              "INQUIRY " + configs.SUCCESS_MSG,
		StatusDesc:                 "INQUIRY " + configs.SUCCESS_MSG,
		ReferenceNumber:            req.ReferenceNumber,
		ProviderStatusCode:         "",
		ProviderStatusMessage:      "INQUIRY " + configs.SUCCESS_MSG,
		ProviderStatusDesc:         "INQUIRY " + configs.SUCCESS_MSG,
		ProviderReferenceNumber:    "-",
		ClientId:                   respOutlet.ClientId,
		ClientName:                 respOutlet.ClientName,
		GroupId:                    respOutlet.GroupId,
		GroupName:                  respOutlet.GroupName,
		MerchantId:                 respOutlet.MerchantId,
		MerchantName:               respOutlet.MerchantName,
		MerchantOutletId:           respOutlet.ID,
		MerchantOutletName:         respOutlet.MerchantOutletName,
		MerchantOutletUsername:     respOutlet.MerchantOutletUsername,
		CustomerId:                 req.CustomerId,
		OtherMsg:                   "-",
		Filter: models.FilterReq{
			CreatedAt: dbTimeTrx,
		},
	}
	// byte, status, er := utils.WorkerPostWithBearer())
	err = svc.services.ApiTrx.InsertTrx(recordInq, nil)
	if err != nil {
		log.Println("Err ", svcName, "InsertTrx", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.ApiTrx.InsertTrxStatus(models.ReqGetTrxStatus{
		ReferenceNumber:         recordInq.ReferenceNumber,
		ProviderReferenceNumber: recordInq.ProviderReferenceNumber,
		StatusCode:              recordInq.StatusCode,
		StatusMessage:           recordInq.StatusMessage,
	})
	if err != nil {
		log.Println("Err ", svcName, "InsertTrxStatus", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, configs.SUCCESS_MSG, recordInq)
	return ctx.JSON(http.StatusOK, result)
}

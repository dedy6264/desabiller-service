package trxservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	helperIakservice "desabiller/services/helperIakService"
	"encoding/json"
	"fmt"
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
		url        string
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
	req.ReferenceNumber, err = svc.services.ApiTrx.GenerateNo("DB-"+dbTime, "", 7)
	if err != nil {
		log.Println("Err ", svcName, "GenerateNo", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respProduct, err := svc.services.ApiProduct.GetProduct(models.ReqGetProduct{
		// ID: req.ProductId,
		ProductCode: req.ProductCode,
	})
	if err != nil {
		log.Println("Err ", svcName, "GetProduct", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
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
		ProviderId:                 respProduct.ProviderId,
		ProviderName:               respProduct.ProviderName,
		ProductProviderId:          respProduct.ProductProviderId,
		ProductProviderName:        respProduct.ProductProviderName,
		ProductProviderCode:        respProduct.ProductProviderCode,
		ProductProviderPrice:       respProduct.ProductProviderPrice,
		ProductProviderAdminFee:    respProduct.ProductProviderAdminFee,
		ProductProviderMerchantFee: respProduct.ProductProviderMerchantFee,
		StatusCode:                 configs.INQUIRY_SUCCESS_CODE,
		StatusMessage:              "INQUIRY " + configs.SUCCESS_MSG,
		StatusDesc:                 "INQUIRY " + configs.SUCCESS_MSG,
		ReferenceNumber:            req.ReferenceNumber,
		ProviderStatusCode:         "-",
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
	if respProduct.ProductTypeId == 2 { //PREPAID
		// if configs.AppEnv == "DEV" {
		// 	url = configs.IakDevUrlPrepaid + configs.ENDPOINT_IAK_PREPAID
		// }
		// if configs.AppEnv == "PROD" {
		// 	url = configs.IakProdUrlPrepaid + configs.ENDPOINT_IAK_PREPAID
		// }
		if respProduct.ProductPrice < respProduct.ProductProviderPrice {
			log.Println("Err ", svcName, "product price invalid", respProduct.ProductPrice, respProduct.ProductProviderPrice)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "product price invalid", nil)
			return ctx.JSON(http.StatusOK, result)
		}
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
		}, nil)
		if err != nil {
			log.Println("Err ", svcName, "InsertTrxStatus", err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		}
	} else { //postpaid
		if configs.AppEnv == "DEV" {
			url = configs.IakDevUrlPostpaid + configs.ENDPOINT_IAK_POSTPAID
		}
		if configs.AppEnv == "PROD" {
			url = configs.IakProdUrlPostpaid + configs.ENDPOINT_IAK_POSTPAID
		}
		if respProduct.ProductMerchantFee > respProduct.ProductProviderMerchantFee {
			log.Println("Err ", svcName, "product merchant fee invalid", respProduct.ProductMerchantFee, respProduct.ProductProviderMerchantFee)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "product merchant fee invalid", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		//inquiry ke partner
		respWorker, err := helperIakservice.IakPLNPostpaidWorkerInquiry(models.ReqInqIak{
			ProductCode: respProduct.ProductProviderCode,
			CustomerId:  req.CustomerId,
			RefId:       req.ReferenceNumber,
			Url:         url,
		})
		if err != nil {
			log.Println("Err ", svcName, "IakPLNPrepaidWorkerInquiry", err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Trx failed", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		byte, _ := json.Marshal(respWorker.BillInfo)
		statusCode := helpers.ErrorCodeGateway(respWorker.InquiryStatus, "INQ")
		recordInq.StatusCode = statusCode
		recordInq.StatusMessage = "INQUIRY " + respWorker.InquiryStatusDesc
		recordInq.StatusDesc = respWorker.InquiryStatusDesc
		recordInq.ReferenceNumber = req.ReferenceNumber
		recordInq.ProviderStatusCode = respWorker.InquiryStatusDetail
		recordInq.ProviderStatusMessage = respWorker.InquiryStatusDescDetail
		recordInq.ProviderStatusDesc = respWorker.InquiryStatusDescDetail
		recordInq.ProviderReferenceNumber = respWorker.TrxProviderReferenceNumber
		recordInq.OtherMsg = string(byte)
		recordInq.Filter = models.FilterReq{
			CreatedAt: dbTime,
		}
		fmt.Println("=", string(byte))
		fmt.Println("=", respWorker.BillInfo)

		// byte, status, er := utils.WorkerPostWithBearer())
		err = svc.services.ApiTrx.InsertTrx(recordInq, nil)
		if err != nil {
			log.Println("Err ", svcName, "UpdateTrx", err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		}
		err = svc.services.ApiTrx.InsertTrxStatus(models.ReqGetTrxStatus{
			ReferenceNumber:         recordInq.ReferenceNumber,
			ProviderReferenceNumber: recordInq.ProviderReferenceNumber,
			StatusCode:              recordInq.StatusCode,
			StatusMessage:           recordInq.StatusMessage,
		}, nil)
		if err != nil {
			log.Println("Err ", svcName, "InsertTrxStatus", err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}

	// byte, status, er := utils.WorkerPostWithBearer())
	respInquiry := models.RespInquiry{
		StatusCode:             recordInq.StatusCode,
		StatusMessage:          recordInq.StatusMessage,
		StatusDesc:             recordInq.StatusDesc,
		ReferenceNumber:        recordInq.ReferenceNumber,
		CreatedAt:              recordInq.Filter.CreatedAt,
		CustomerId:             recordInq.CustomerId,
		BillInfo:               recordInq.OtherMsg,
		ProductId:              recordInq.ProductId,
		ProductName:            recordInq.ProductName,
		ProductCode:            recordInq.ProductCode,
		ProductPrice:           recordInq.ProductPrice,
		ProductAdminFee:        recordInq.ProductAdminFee,
		ProductMerchantFee:     recordInq.ProductMerchantFee,
		MerchantOutletId:       recordInq.MerchantOutletId,
		MerchantOutletName:     recordInq.MerchantOutletName,
		MerchantOutletUsername: recordInq.MerchantOutletUsername,
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, recordInq.StatusCode, recordInq.StatusMessage, respInquiry)
	return ctx.JSON(http.StatusOK, result)
}

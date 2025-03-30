package trxservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

func (svc trxService) InquiryBiller(ctx echo.Context) error {
	var (
		svcName = "InquiryBiller"
		// respSvc models.ResponseList
		respOutlet                                                                                                           models.RespGetMerchantOutlet
		referenceNumber, url, statusCode, statusMessage, statusDesc, statusCodeDetail, statusMessageDetail, statusDescDetail string
		billInfo                                                                                                             map[string]interface{}
		respWorker                                                                                                           models.ResponseWorkerInquiry
		result                                                                                                               models.Response

		productPrice, productAdminFee, productMerchantFee                         float64
		productProviderPrice, productProviderAdminFee, productProviderMerchantFee float64
	)
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMPTRX)
	dbTimeTrx := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	req := new(models.ReqInquiry)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.INQUIRY_FAILED_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ProductCode == "" {
		log.Println("Err ", svcName, "product code cannot be null")
		result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.INQUIRY_FAILED_CODE, "product code cannot be null", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	// if req.AdditionalField.SubscriberNumber == "" {
	// 	log.Println("Err ", svcName, "customer id cannot be null")
	// 	result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.INQUIRY_FAILED_CODE, "customer id cannot be null", nil)
	// 	return ctx.JSON(http.StatusOK, result)
	// }
	data := helpers.TokenJWTDecode(ctx)
	respOutlet, err = svc.services.RepoHierarchy.GetMerchantOutlet(models.ReqGetMerchantOutlet{
		ID: data.MerchantOutletId,
	})
	if err != nil {
		log.Println("Err ", svcName, "GetMerchantOutlet", err)
		result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.INQUIRY_FAILED_CODE, "Not Found", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	respProduct, err := svc.services.RepoProduct.GetProduct(models.ReqGetProduct{
		// ID: req.ProductId,
		ProductCode: req.ProductCode,
	})
	if err != nil {
		log.Println("Err Product Invalid ", svcName, "GetProduct", err)
		result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.INQUIRY_FAILED_CODE, "Product Invalid", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	referenceNumber, err = svc.services.RepoTrx.GenerateNo("DB-"+dbTime, "", 7)
	if err != nil {
		log.Println("Err ", svcName, "GenerateNo", err)
		result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.INQUIRY_FAILED_CODE, "Not Found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	statusCode = configs.INQUIRY_SUCCESS_CODE
	statusDesc = "INQUIRY " + configs.SUCCESS_MSG
	statusMessage = "INQUIRY " + configs.SUCCESS_MSG
	statusCodeDetail = "-"
	statusMessageDetail = "INQUIRY " + configs.SUCCESS_MSG
	statusDescDetail = "INQUIRY " + configs.SUCCESS_MSG
	recordInq := models.ReqGetTrx{
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
		StatusCode:                 statusCode,
		StatusMessage:              statusMessage,
		StatusDesc:                 statusDesc,
		ReferenceNumber:            referenceNumber,
		ProviderStatusCode:         statusCodeDetail,
		ProviderStatusMessage:      statusMessageDetail,
		ProviderStatusDesc:         statusDescDetail,
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
		CustomerId:                 req.AdditionalField.SubscriberNumber,
		TotalTrxAmount:             respProduct.ProductPrice,
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
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.INQUIRY_FAILED_CODE, "invalid product price ", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		err = svc.services.RepoTrx.InsertTrx(recordInq, nil)
		if err != nil {
			log.Println("Err ", svcName, "InsertTrx", err)
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.INQUIRY_FAILED_CODE, "failed", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		err = svc.services.RepoTrx.InsertTrxStatus(models.ReqGetTrxStatus{
			ReferenceNumber:         recordInq.ReferenceNumber,
			ProviderReferenceNumber: recordInq.ProviderReferenceNumber,
			StatusCode:              recordInq.StatusCode,
			StatusMessage:           recordInq.StatusMessage,
		}, nil)
		if err != nil {
			log.Println("Err ", svcName, "InsertTrxStatus", err)
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.INQUIRY_FAILED_CODE, err.Error(), nil)
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
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.INQUIRY_FAILED_CODE, "product merchant fee invalid", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		//inquiry ke partner
		respWorker, err = svc.InqProviderSwitcher(models.ProviderInqRequest{
			ProviderName:         respProduct.ProviderName,
			ProviderId:           respProduct.ProviderId,
			ProductCode:          respProduct.ProductProviderCode,
			SubscriberNumber:     req.AdditionalField.SubscriberNumber,
			SubscriberName:       req.AdditionalField.SubscriberName,
			ReferenceNumber:      referenceNumber,
			Url:                  url,
			Periode:              req.AdditionalField.Periode,
			ProductReferenceCode: respProduct.ProductReferenceCode,
		})
		if err != nil {
			log.Println("Err ", svcName, err)
			if strings.Contains(err.Error(), "BPJSKSValidate") {
				result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.INQUIRY_FAILED_CODE, err.Error(), nil)
			} else {
				result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.INQUIRY_FAILED_CODE, "Trx failed", nil)
			}
			return ctx.JSON(http.StatusOK, result)
		}
		productProviderPrice = respWorker.TrxAmount
		productProviderAdminFee = respWorker.AdminFee
		productProviderMerchantFee = respProduct.ProductProviderMerchantFee

		productAdminFee = respWorker.AdminFee
		productMerchantFee = respProduct.ProductMerchantFee
		productPrice = respWorker.TotalTrxAmount - respWorker.AdminFee

		billInfo = respWorker.BillInfo
		byte, _ := json.Marshal(billInfo)
		statusCode = helpers.ErrorCodeGateway(respWorker.InquiryStatus, "INQ")
		recordInq.StatusCode = statusCode
		recordInq.StatusMessage = "INQUIRY " + respWorker.InquiryStatusDesc
		recordInq.StatusDesc = respWorker.InquiryStatusDesc
		recordInq.ReferenceNumber = referenceNumber
		recordInq.ProviderStatusCode = respWorker.InquiryStatusDetail
		recordInq.ProviderStatusMessage = respWorker.InquiryStatusDescDetail
		recordInq.ProviderStatusDesc = respWorker.InquiryStatusDescDetail
		recordInq.ProviderReferenceNumber = respWorker.TrxProviderReferenceNumber
		recordInq.OtherMsg = string(byte)
		recordInq.TotalTrxAmount = respWorker.TotalTrxAmount
		recordInq.Filter = models.FilterReq{
			CreatedAt: dbTime,
		}
		recordInq.ProductPrice = productPrice
		recordInq.ProductAdminFee = productAdminFee
		recordInq.ProductMerchantFee = productMerchantFee
		recordInq.ProductProviderPrice = productProviderPrice
		recordInq.ProductProviderAdminFee = productProviderAdminFee
		recordInq.ProductProviderMerchantFee = productProviderMerchantFee
		recordInq.ProductReferenceCode = respProduct.ProductReferenceCode

		err = svc.services.RepoTrx.InsertTrx(recordInq, nil)
		if err != nil {
			log.Println("Err ", svcName, "UpdateTrx", err)
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.INQUIRY_FAILED_CODE, err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		}
		err = svc.services.RepoTrx.InsertTrxStatus(models.ReqGetTrxStatus{
			ReferenceNumber:         recordInq.ReferenceNumber,
			ProviderReferenceNumber: recordInq.ProviderReferenceNumber,
			StatusCode:              recordInq.StatusCode,
			StatusMessage:           recordInq.StatusMessage,
		}, nil)
		if err != nil {
			log.Println("Err ", svcName, "InsertTrxStatus", err)
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.INQUIRY_FAILED_CODE, err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}

	// byte, status, er := utils.WorkerPostWithBearer())
	respInquiry := models.RespInquiry{
		StatusMessage:          recordInq.ProviderStatusMessage,
		CreatedAt:              recordInq.Filter.CreatedAt,
		MerchantOutletName:     recordInq.MerchantOutletName,
		MerchantOutletUsername: recordInq.MerchantOutletUsername,
		ReferenceNumber:        recordInq.ReferenceNumber,
		ProductName:            recordInq.ProductName,
		ProductCode:            recordInq.ProductCode,
		SubscriberNumber:       recordInq.CustomerId,
		ProductPrice:           recordInq.ProductPrice,
		ProductAdminFee:        recordInq.ProductAdminFee,
		ProductMerchantFee:     recordInq.ProductMerchantFee,
		TotalTrxAmount:         recordInq.TotalTrxAmount,
		BillInfo:               billInfo,
	}
	result = helpers.ResponseJSON(configs.TRUE_VALUE, recordInq.StatusCode, recordInq.StatusMessage, respInquiry)
	return ctx.JSON(http.StatusOK, result)
}

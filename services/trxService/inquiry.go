package trxservice

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func (svc trxService) InquiryBiller(ctx echo.Context) error {
	var (
		svcName = "InquiryBiller"
		data    models.DataToken
		// respSvc models.ResponseList
		respUserApp models.UserApp
		referenceNumber, url,
		statusCode, statusMsg, statusDesc,
		statusCodeDetail, statusMsgDetail, statusDescDetail string
		billInfo    map[string]interface{}
		respWorker  models.ProviderResponse
		result      models.Response
		datainquiry models.ReqGetTransaction
		productId   int
		productPrice, productAdminFee, productMerchantFee,
		productProviderPrice, productProviderAdminFee, productProviderMerchantFee, transactionTotalAmount float64
	)
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMPTRX)
	dbTimeTrx := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	req := new(models.ReqInquiry)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log("", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	{ //validasi #1
		if req.ProductCode == "" {
			utils.Log("ProductCode cannot be null", svcName, nil)
			result := helpers.ResponseJSON(configs.FALSE_VALUE,
				configs.VALIDATE_ERROR_CODE,
				"Product Code cannot be empty",
				"Product Code cannot be empty",
				nil)
			return ctx.JSON(http.StatusOK, result)
		}
		if req.AdditionalField.SubscriberNumber == "" {
			utils.Log("SubscriberNumber cannot be null", svcName, nil)
			result := helpers.ResponseJSON(configs.FALSE_VALUE,
				configs.VALIDATE_ERROR_CODE,
				"SubscriberNumber cannot be empty",
				"SubscriberNumber cannot be empty",
				nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}
	if req.ProductCode == "BPJSKS" && req.AdditionalField.Periode == 0 {
		if req.ProductCode == "" {
			utils.Log("Period canot be empty", svcName, nil)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Period cannot be empty",
				"Period cannot be empty", nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}
	//get user app data
	if ctx.Get("user") != nil {
		data = helpers.TokenJWTDecode(ctx)
	}
	respUserApp, err = svc.services.RepoHierarchy.GetUserApp(models.ReqGetUserApp{
		Filter: models.UserApp{
			ID: int64(data.UserAppId),
		},
	})
	if err != nil {
		utils.Log("GetUserApp", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	respProduct, err := svc.services.RepoProduct.GetProduct(models.ReqGetProduct{
		Filter: models.Product{
			ProductCode: req.ProductCode,
		},
	})
	if err != nil {
		utils.Log("GetProduct", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	productId = respProduct.ID
	productPrice = respProduct.ProductPrice
	productAdminFee = respProduct.ProductAdminFee
	productMerchantFee = respProduct.ProductMerchantFee
	productProviderPrice = respProduct.ProductProviderPrice
	productProviderAdminFee = respProduct.ProductProviderAdminFee
	productProviderMerchantFee = respProduct.ProductProviderMerchantFee
	transactionTotalAmount = productPrice + productAdminFee
	referenceNumber, err = svc.services.RepoTrx.GenerateNo("VDB-"+dbTime, "", 7)
	if err != nil {
		utils.Log("GenerateNo", svcName, err)
		result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", "Failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	_, err = svc.services.RepoTrx.GetTrx(models.ReqGetTransaction{
		Filter: models.Transaction{
			ReferenceNumber: referenceNumber,
		},
	})
	if err != nil && err != sql.ErrNoRows {
		// if err == sql.ErrNoRows {
		// 	utils.Log("GetTrx", svcName, err)
		// 	result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_INVALID_TRANSACTION[0], configs.RC_INVALID_TRANSACTION[1], "Invalid Noreff", nil)
		// 	return ctx.JSON(http.StatusOK, result)
		// }
		utils.Log("GetTrx", svcName, err)
		result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED[0], configs.RC_FAILED[1], "Invalid Noreff", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	switch configs.AppEnv {
	case "DEV":
		url = configs.DevUrl + configs.ENDPOINT_PROVIDER_INQUIRY
	case "PROD":
		url = configs.ProdUrl + configs.ENDPOINT_PROVIDER_INQUIRY
	default:
		url = configs.LocalUrl + configs.ENDPOINT_PROVIDER_INQUIRY
	}
	if respProduct.ProductTypeID == 2 { //PREPAID

		if respProduct.ProductPrice < respProduct.ProductProviderPrice {
			utils.Log("Invalid Product Price : "+strconv.Itoa(int(respProduct.ProductPrice))+" "+strconv.Itoa(int(respProduct.ProductProviderPrice)), svcName, nil)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Period cannot be empty",
				"Period cannot be empty", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		// if respProduct.ProductReferenceID == 10 {
		reqprovider := models.ReqInquiryProvider{
			ProductCode:             respProduct.ProductProviderCode,
			ReferenceNumber:         "",
			ReferenceNumberMerchant: referenceNumber,
			CustomerID:              req.AdditionalField.SubscriberNumber,
			Periode:                 strconv.Itoa(req.AdditionalField.Periode),
			Amount:                  0,
		}
		respByte, _, err := utils.WorkerPostWithBearer(url, configs.TOKEN, reqprovider, "json")

		if err != nil {
			utils.Log("WorkerPostWithBearer", svcName, err)
			statusCode = configs.RC_INQUIRY_FAILED[0]
			statusMsg = "INQUIRY"
			statusDesc = configs.RC_INQUIRY_FAILED[1]
			statusMsgDetail = err.Error()
			statusDescDetail = err.Error()
		} else {
			err = json.Unmarshal(respByte, &respWorker)
			if err != nil {
				utils.Log("Unmarshal", svcName, err)
				result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_INQUIRY_FAILED[0], configs.RC_INQUIRY_FAILED[1], "Failed", nil)
				return ctx.JSON(http.StatusOK, result)
			}
			//converter
			code, msg, _ := helpers.ResponseConverter(respWorker.ResponseCode, respWorker.ResponseMessage, true)
			statusCode = code
			statusMsg = "INQUIRY"
			statusDesc = msg
			statusCodeDetail = respWorker.ResponseCode
			statusMsgDetail = respWorker.ResponseMessage
			statusDescDetail = respWorker.ResponseMessage
		}
		transactionTotalAmount = respProduct.ProductPrice + respProduct.ProductAdminFee
		datainquiry = models.ReqGetTransaction{
			Filter: models.Transaction{
				ProductProviderName:        respProduct.ProductProviderName,
				ProductProviderCode:        respProduct.ProductProviderCode,
				ProductProviderPrice:       productProviderPrice,
				ProductProviderAdminFee:    productProviderAdminFee,
				ProductProviderMerchantFee: productProviderMerchantFee,
				ProductID:                  int64(productId),
				ProductName:                respProduct.ProductName,
				ProductCode:                respProduct.ProductCode,
				ProductPrice:               productPrice,
				ProductAdminFee:            productAdminFee,
				ProductMerchantFee:         productMerchantFee,
				ProductCategoryID:          int64(respProduct.ProductCategoryID),
				ProductCategoryName:        respProduct.ProductCategoryName,
				ProductTypeID:              int64(respProduct.ProductTypeID),
				ProductTypeName:            respProduct.ProductTypeName,
				ReferenceNumber:            referenceNumber,
				ProviderReferenceNumber:    respWorker.Result.ReferenceNumber,
				StatusCode:                 statusCode,
				StatusMessage:              statusMsg,
				StatusDesc:                 statusDesc,
				StatusCodeDetail:           statusCodeDetail,
				StatusMessageDetail:        statusMsgDetail,
				StatusDescDetail:           statusDescDetail,
				ProductReferenceID:         int64(respProduct.ProductReferenceID),
				ProductReferenceCode:       respProduct.ProductReferenceCode,
				CustomerID:                 req.AdditionalField.SubscriberNumber,
				TransactionTotalAmount:     transactionTotalAmount,
				SavingAccountID:            respUserApp.AccountID,
				SavingAccountNumber:        respUserApp.AccountNumber,
				UserAppID:                  respUserApp.ID,
				Username:                   respUserApp.Username,
				CreatedAt:                  dbTimeTrx,
				CreatedBy:                  respUserApp.Username,
				UpdatedAt:                  dbTimeTrx,
				UpdatedBy:                  respUserApp.Username,
			},
		}
	} else { //postpaid
		if respProduct.ProductMerchantFee > respProduct.ProductProviderMerchantFee {
			utils.Log("Invalid Product Price : "+strconv.Itoa(int(respProduct.ProductPrice))+" "+strconv.Itoa(int(respProduct.ProductProviderPrice)), svcName, err)
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", "failed", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		//inquiry ke partner
		reqprovider := models.ReqInquiryProvider{
			ProductCode:             respProduct.ProductProviderCode,
			ReferenceNumber:         "",
			ReferenceNumberMerchant: referenceNumber,
			CustomerID:              req.AdditionalField.SubscriberNumber,
			Periode:                 strconv.Itoa(req.AdditionalField.Periode),
			Amount:                  0,
		}
		respByte, _, err := utils.WorkerPostWithBearer(url, configs.TOKEN, reqprovider, "json")
		if err != nil {
			utils.Log("WorkerPostWithBearer", svcName, err)
			statusCode = configs.RC_INQUIRY_FAILED[0]
			statusMsg = "INQUIRY"
			statusDesc = configs.RC_INQUIRY_FAILED[1]
			statusMsgDetail = err.Error()
			statusDescDetail = err.Error()
		} else {
			err = json.Unmarshal(respByte, &respWorker)
			if err != nil {
				utils.Log("Unmarshal", svcName, err)
				result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_INQUIRY_FAILED[0], configs.RC_INQUIRY_FAILED[1], "Failed", nil)
				return ctx.JSON(http.StatusOK, result)
			}
			//converter
			code, msg, _ := helpers.ResponseConverter(respWorker.ResponseCode, respWorker.ResponseMessage, true)
			statusCode = code
			statusMsg = "INQUIRY"
			statusDesc = msg
			statusCodeDetail = respWorker.ResponseCode
			statusMsgDetail = respWorker.ResponseMessage
			statusDescDetail = respWorker.ResponseMessage
		}
		datainquiry = models.ReqGetTransaction{
			Filter: models.Transaction{
				ProductProviderName:        respProduct.ProductProviderName,
				ProductProviderCode:        respProduct.ProductProviderCode,
				ProductProviderPrice:       productProviderPrice,
				ProductProviderAdminFee:    productProviderAdminFee,
				ProductProviderMerchantFee: productProviderMerchantFee,
				ProductID:                  int64(productId),
				ProductName:                respProduct.ProductName,
				ProductCode:                respProduct.ProductCode,
				ProductPrice:               productPrice,
				ProductAdminFee:            productAdminFee,
				ProductMerchantFee:         productMerchantFee,
				ProductCategoryID:          int64(respProduct.ProductCategoryID),
				ProductCategoryName:        respProduct.ProductCategoryName,
				ProductTypeID:              int64(respProduct.ProductTypeID),
				ProductTypeName:            respProduct.ProductTypeName,
				ReferenceNumber:            referenceNumber,
				ProviderReferenceNumber:    respWorker.Result.ReferenceNumber,
				StatusCode:                 statusCode,
				StatusMessage:              statusMsg,
				StatusDesc:                 statusDesc,
				StatusCodeDetail:           statusCodeDetail,
				StatusMessageDetail:        statusMsgDetail,
				StatusDescDetail:           statusDescDetail,
				ProductReferenceID:         int64(respProduct.ProductReferenceID),
				ProductReferenceCode:       respProduct.ProductReferenceCode,
				CustomerID:                 req.AdditionalField.SubscriberNumber,
				TransactionTotalAmount:     transactionTotalAmount,
				SavingAccountID:            respUserApp.AccountID,
				SavingAccountNumber:        respUserApp.AccountNumber,
				UserAppID:                  respUserApp.ID,
				Username:                   respUserApp.Username,
				CreatedAt:                  dbTimeTrx,
				CreatedBy:                  respUserApp.Username,
				UpdatedAt:                  dbTimeTrx,
				UpdatedBy:                  respUserApp.Username,
			},
		}
	}
	if respWorker.Result.ProductPrice > respProduct.ProductPrice {
		statusCode = configs.RC_PRODUCT_DISRUPTION[0]
		statusDesc = configs.RC_PRODUCT_DISRUPTION[1]
		datainquiry.Filter.StatusCode = statusCode
		datainquiry.Filter.StatusMessage = statusMsg
		datainquiry.Filter.StatusDesc = statusDesc
		datainquiry.Filter.ProductProviderPrice = float64(respWorker.Result.ProductPrice)
	}
	err = svc.services.RepoTrx.InsertTrx(datainquiry, nil)
	if err != nil {
		utils.Log("InsertTrx", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	// byte, status, er := utils.WorkerPostWithBearer())
	respInquiry := models.RespInquiry{
		// StatusMessage:          dataInquiry.ProviderStatusMessage,
		CreatedAt: dbTimeTrx,
		// MerchantOutletName:     dataInquiry.MerchantOutletName,
		// MerchantOutletUsername: dataInquiry.MerchantOutletUsername,
		ReferenceNumber: referenceNumber,
		ProductName:     respProduct.ProductName,
		// ProductCode:            dataInquiry.ProductCode,
		SubscriberNumber:       req.AdditionalField.SubscriberNumber,
		ProductPrice:           productPrice,
		ProductAdminFee:        productAdminFee,
		ProductMerchantFee:     productMerchantFee,
		TransactionTotalAmount: transactionTotalAmount,
		BillInfo:               billInfo,
	}
	result = helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMsg, statusDesc, respInquiry)
	return ctx.JSON(http.StatusOK, result)
}

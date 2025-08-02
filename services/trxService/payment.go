package trxservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (svc trxService) PaymentBiller(ctx echo.Context) error {
	var (
		svcName = "PaymentBiller"
		statusCode,
		statusMsg,
		statusDesc,
		// referenceNumber,
		// statusMsg,
		statusCodeDetail,
		statusMsgDetail,
		statusDescDetail,
		// providerReferenceNumber,
		url string
		// billInfo     map[string]interface{}
		respProvider models.ProviderResponse
		// respSvc models.ResponseList
		// respOutlet models.RespGetMerchantOutlet
		result models.Response
		// transactionTotalAmount float64
		// billDescByte           []byte
	)
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMPTRX)
	req := new(models.ReqPaymentTrx)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log("", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ReferenceNumber == "" {
		utils.Log("ReferenceNumber id cannot be null", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, configs.VALIDATE_ERROR_CODE, "ReferenceNumber id cannot be null", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	// if req.AccountNumber == "" {
	// 	log.Println("Err ", svcName, "Account Number cannot be null")
	// 	result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.FAILED_CODE, "Account Number id cannot be null", nil)
	// 	return ctx.JSON(http.StatusOK, result)
	// }
	// if req.AccountPIN == "" {
	// 	log.Println("Err ", svcName, "PIN cannot be null")
	// 	result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.FAILED_CODE, "PIN id cannot be null", nil)
	// 	return ctx.JSON(http.StatusOK, result)
	// }
	respInqTrx, err := svc.services.RepoTrx.GetTrx(models.ReqGetTransaction{
		Filter: models.Transaction{
			ReferenceNumber: req.ReferenceNumber,
		},
	})
	if err != nil {
		utils.Log("GetTrx", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], "Not found", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	if respInqTrx.StatusCode != configs.RC_INQUIRY_SUCCESS[0] {
		utils.Log("Invalid Transaction Number", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_INVALID_TRANSACTION[0], configs.RC_INVALID_TRANSACTION[1], "Not found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	//----->> trigger kredit
	err = svc.TriggerKredit(respInqTrx.TransactionTotalAmount, respInqTrx.SavingAccountNumber, req.AccountPIN, respInqTrx.ReferenceNumber, configs.TRX_CODE_PAYMENT)
	if err != nil {
		switch err.Error() {
		case "UNSETPIN":
			utils.Log("TriggerKredit", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_UNSET_PIN[0], configs.RC_FAILED_UNSET_PIN[1], err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		case "WRONGPIN":
			utils.Log("TriggerKredit", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_WRONG_PIN[0], configs.RC_FAILED_WRONG_PIN[1], err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		case "BALANCE_NOT_ENOUGH":
			utils.Log("TriggerKredit", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_EXCESSIVE_BALANCE[0], configs.RC_EXCESSIVE_BALANCE[1], err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)

		}
		utils.Log("TriggerKredit", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED[0], configs.RC_FAILED[1], "Not found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	// billDescPay := BillInfoBPJS{}
	if configs.AppEnv == "LOCAL" {
		url = configs.LocalUrl + configs.ENDPOINT_PROVIDER_PAYMENT
	}
	if configs.AppEnv == "DEV" {
		url = configs.DevUrl + configs.ENDPOINT_PROVIDER_PAYMENT
	}
	if configs.AppEnv == "PROD" {
		url = configs.ProdUrl + configs.ENDPOINT_PROVIDER_PAYMENT
	}
	reqprovider := models.ReqInquiryProvider{
		ReferenceNumber:         respInqTrx.ProviderReferenceNumber,
		ReferenceNumberMerchant: respInqTrx.ReferenceNumber,
	}
	respByte, _, err := utils.WorkerPostWithBearer(url, configs.TOKEN, reqprovider, "json")

	if err != nil {
		utils.Log("WorkerPostWithBearer", svcName, err)
		statusCode = configs.RC_PENDING[0]
		statusMsg = "PAYMENT"
		statusDesc = configs.RC_PENDING[1]
		statusMsgDetail = err.Error()
		statusDescDetail = err.Error()
	} else {
		err = json.Unmarshal(respByte, &respProvider)
		if err != nil {
			utils.Log("Unmarshal", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_INQUIRY_FAILED[0], configs.RC_INQUIRY_FAILED[1], "Failed", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		//converter
		code, msg, _ := helpers.ResponseConverter(respProvider.ResponseCode, respProvider.ResponseMessage, false)
		statusCode = code
		statusMsg = "PAYMENT"
		statusDesc = msg
		statusCodeDetail = respProvider.ResponseCode
		statusMsgDetail = respProvider.ResponseMessage
		statusDescDetail = respProvider.ResponseMessage
	}
	// transactionTotalAmount = respInqTrx.ProductPrice + respInqTrx.ProductAdminFee
	// fmt.Println(
	// 	statusCode,
	// 	statusMsg,
	// 	statusDesc,
	// 	referenceNumber,
	// 	statusMsg,
	// 	statusCodeDetail,
	// 	statusMsgDetail,
	// 	statusDescDetail,
	// 	providerReferenceNumber,
	// 	transactionTotalAmount,
	// )
	switch respInqTrx.ProductTypeID {
	case 1: //Postpaid
		// switch respInqTrx.ProductReferenceCode {
		// case "BPJSKS":
		// 	billDescInq := models.BillInfoBPJS{}
		// 	otherMsgInq := helpers.JsonDescape(respInqTrx.OtherMsg)
		// 	_ = json.Unmarshal([]byte(otherMsgInq), &billDescInq)
		// 	if len(respProvider.BillInfo) != 0 {
		// 		snVal, _ := respProvider.BillInfo["sn"].(string)
		// 		billDescInq.Sn = snVal
		// 	}
		// 	billDescByte, _ = json.Marshal(billDescInq)
		// }
	case 2: //PREPAID
		// billDescInq := models.BillDescPLN{}
		// otherMsgInq := helpers.JsonDescape(respInqTrx.OtherMsg)
		// _ = json.Unmarshal([]byte(otherMsgInq), &billDescInq)
		// if len(respProvider.BillInfo) != 0 {
		// 	// snVal, _ := respProvider.BillInfo["sn"].(string)
		// 	// billDescInq.Sn = snVal
		// }
		// billDescByte, _ = json.Marshal(billDescInq)
	}
	updateTrx := models.Transaction{
		ID:                         int64(respInqTrx.Id),
		ProductProviderName:        respInqTrx.ProductProviderName,
		ProductProviderCode:        respInqTrx.ProductProviderCode,
		ProductProviderPrice:       respInqTrx.ProductProviderPrice,
		ProductProviderAdminFee:    respInqTrx.ProductProviderAdminFee,
		ProductProviderMerchantFee: respInqTrx.ProductProviderMerchantFee,
		ProductID:                  respInqTrx.ProductID,
		ProductName:                respInqTrx.ProductName,
		ProductCode:                respInqTrx.ProductCode,
		ProductPrice:               respInqTrx.ProductPrice,
		ProductAdminFee:            respInqTrx.ProductAdminFee,
		ProductMerchantFee:         respInqTrx.ProductMerchantFee,
		ProductCategoryID:          respInqTrx.ProductCategoryID,
		ProductCategoryName:        respInqTrx.ProductCategoryName,
		ProductTypeID:              respInqTrx.ProductTypeID,
		ProductTypeName:            respInqTrx.ProductTypeName,
		ReferenceNumber:            respInqTrx.ReferenceNumber,
		ProviderReferenceNumber:    respProvider.Result.ReferenceNumber,
		StatusCode:                 statusCode,
		StatusMessage:              statusMsg,
		StatusDesc:                 statusDesc,
		StatusCodeDetail:           statusCodeDetail,
		StatusMessageDetail:        statusMsgDetail,
		StatusDescDetail:           statusDescDetail,
		ProductReferenceID:         respInqTrx.ProductReferenceID,
		ProductReferenceCode:       respInqTrx.ProductReferenceCode,
		CustomerID:                 respInqTrx.CustomerID,
		OtherReff:                  respInqTrx.OtherReff,
		OtherCustomerInfo:          respInqTrx.OtherCustomerInfo,
		SavingAccountName:          respInqTrx.SavingAccountName,
		SavingAccountID:            respInqTrx.SavingAccountID,
		SavingAccountNumber:        respInqTrx.SavingAccountNumber,
		TransactionTotalAmount:     respInqTrx.TransactionTotalAmount,
		UserAppID:                  respInqTrx.UserAppID,
		Username:                   respInqTrx.Username,
		CreatedAt:                  respInqTrx.CreatedAt,
		CreatedBy:                  respInqTrx.CreatedBy,
		UpdatedAt:                  dbTime,
		UpdatedBy:                  respInqTrx.UpdatedBy,
	}

	err = svc.services.RepoTrx.UpdateTrx(models.ReqGetTransaction{
		Filter: updateTrx,
	}, nil)
	if err != nil {
		utils.Log("UpdateTrx", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], "Failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	responsePayment := models.RespPayment{
		ReferenceNumber:     respInqTrx.ReferenceNumber,
		ProductName:         respInqTrx.ProductName,
		ProductCategoryId:   int(respInqTrx.ProductCategoryID),
		ProductCategoryName: respInqTrx.ProductCategoryName,
		SubscriberNumber:    respInqTrx.CustomerID,
		ProductPrice:        respInqTrx.ProductPrice,
		ProductAdminFee:     respInqTrx.ProductAdminFee,
		ProductMerchantFee:  respInqTrx.ProductMerchantFee,
		TotalTrxAmount:      respInqTrx.TransactionTotalAmount,
		BillInfo:            respProvider.Result.BillInfo,
	}
	if statusCode != configs.RC_PENDING[0] && statusCode != configs.RC_SUCCESS[0] {
		svc.TriggerDebet(respInqTrx.TransactionTotalAmount, respInqTrx.SavingAccountNumber, req.AccountPIN, respInqTrx.ReferenceNumber, configs.TRX_CODE_REVERSAL)
		result = helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMsg, statusDesc, responsePayment)
		return ctx.JSON(http.StatusOK, result)
	}
	result = helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMsg, statusDesc, responsePayment)
	return ctx.JSON(http.StatusOK, result)
}

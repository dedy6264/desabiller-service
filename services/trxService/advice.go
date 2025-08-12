package trxservice

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/utils"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (svc trxService) Advice(ctx echo.Context) error {
	var (
		svcName = "[IAK]Advice"
		url, statusCode,
		statusMsg, statusDesc, statusMsgDetail,
		statusDescDetail, statusCodeDetail string
		respProvider models.ProviderResponse
		billdesc     map[string]interface{}
	)
	req := new(models.ReqAviceTrx)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log("", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respTrx, err := svc.services.RepoTrx.GetTrx(models.ReqGetTransaction{
		Filter: models.Transaction{
			ReferenceNumber: req.ReferenceNumber,
		},
	})
	if err != nil {
		utils.Log("GetTrx", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], "Not found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	statusCode = respTrx.StatusCode
	statusMsg = respTrx.StatusMessage
	statusDesc = respTrx.StatusDesc
	if respTrx.OtherReff != "" {
		err = json.Unmarshal([]byte(respTrx.OtherReff), &billdesc)
		if err != nil {
			utils.Log("Unmarshal", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED[0], configs.RC_FAILED[1], "Not found", nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}
	if respTrx.StatusCode != configs.RC_PENDING[0] {
		respPayment := models.RespPayment{
			ReferenceNumber:  respTrx.ReferenceNumber,
			CreatedAt:        respTrx.CreatedAt,
			SubscriberNumber: respTrx.CustomerID,
			BillInfo:         billdesc,
			ProductName:      respTrx.ProductName,
			// ProductCode:            respTrx.ProductCode,
			ProductPrice:       respTrx.ProductPrice,
			ProductAdminFee:    respTrx.ProductAdminFee,
			ProductMerchantFee: respTrx.ProductMerchantFee,
			// MerchantOutletName:     resp.MerchantOutletName,
			// MerchantOutletUsername: resp.MerchantOutletUsername,
		}
		result := helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMsg, statusDesc, respPayment)
		return ctx.JSON(http.StatusOK, result)
	}
	if configs.AppEnv == "DEV" {
		url = configs.DevUrl + "/api/v1/advice"
	}
	if configs.AppEnv == "PROD" {
		url = configs.ProdUrl + "/api/v1/advice"
	}
	if configs.AppEnv == "LOCAL" {
		url = configs.LocalUrl + "/api/v1/advice"
	}
	reqprovider := models.ReqPaymentProvider{
		ReferenceNumber:         respTrx.ProviderReferenceNumber,
		ReferenceNumberMerchant: respTrx.ReferenceNumber,
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
	if statusCode == configs.RC_PENDING[0] {
		respPayment := models.RespPayment{
			ReferenceNumber:  respTrx.ReferenceNumber,
			CreatedAt:        respTrx.CreatedAt,
			SubscriberNumber: respTrx.CustomerID,
			BillInfo:         billdesc,
			ProductName:      respTrx.ProductName,
			// ProductCode:            respTrx.ProductCode,
			ProductPrice:       respTrx.ProductPrice,
			ProductAdminFee:    respTrx.ProductAdminFee,
			ProductMerchantFee: respTrx.ProductMerchantFee,
			// MerchantOutletName:     resp.MerchantOutletName,
			// MerchantOutletUsername: resp.MerchantOutletUsername,
		}
		result := helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMsg, statusDesc, respPayment)
		return ctx.JSON(http.StatusOK, result)
	}
	if statusCode == configs.RC_FAILED[0] {
		err = svc.TriggerDebet(respTrx.TransactionTotalAmount, respTrx.SavingAccountNumber, "", respTrx.ReferenceNumber, configs.TRX_CODE_REVERSAL)
		if err != nil {
			log.Println("TriggerDebet", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED[0], configs.RC_FAILED[1], "Failed to reverse transaction", nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}
	billInfo, _ := json.Marshal(respProvider.Result.BillInfo)
	err = UpdateAndInsertStatusTrx(respTrx, models.ResponseWorkerPayment{
		PaymentStatus:              statusCode,
		PaymentStatusDesc:          statusDesc,
		PaymentStatusMsg:           statusMsg,
		PaymentStatusDetail:        statusCodeDetail,
		PaymentStatusDescDetail:    statusDescDetail,
		PaymentStatusMsgDetail:     statusMsgDetail,
		TrxReferenceNumber:         respTrx.ReferenceNumber,
		TrxProviderReferenceNumber: reqprovider.ReferenceNumber,
		TransactionTotalAmount:     respTrx.TransactionTotalAmount,
		AdminFee:                   respTrx.ProductAdminFee,
		BillInfo: map[string]interface{}{
			"billDesc": string(billInfo),
		},
	}, svc)
	if err != nil {
		log.Println("Err UpdateAndInsertStatusTrx", svcName, err)
		result := helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMsg, statusDesc, nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respPayment := models.RespPayment{
		ReferenceNumber:  respTrx.ReferenceNumber,
		CreatedAt:        respTrx.CreatedAt,
		SubscriberNumber: respTrx.CustomerID,
		BillInfo:         string(billInfo),
		ProductName:      respTrx.ProductName,
		// ProductCode:            respTrx.ProductCode,
		ProductPrice:       respTrx.ProductPrice,
		ProductAdminFee:    respTrx.ProductAdminFee,
		ProductMerchantFee: respTrx.ProductMerchantFee,
		// MerchantOutletName:     resp.MerchantOutletName,
		// MerchantOutletUsername: resp.MerchantOutletUsername,
	}
	// switch resp.ProductTypeId {
	// case 1: //postpaid

	// 	var billInfo map[string]interface{}
	// 	err = json.Unmarshal([]byte(resp.OtherMsg), &billInfo)
	// 	if err != nil {
	// 		log.Println("Err ", svcName, "Unmarshal", err)
	// 		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.FAILED_CODE, configs.FAILED_MSG, err.Error(), nil)
	// 		return ctx.JSON(http.StatusOK, result)
	// 	}
	// 	respProvider.TrxReferenceNumber = resp.ReferenceNumber
	// 	respProvider.TrxProviderReferenceNumber = resp.ProviderReferenceNumber
	// 	respProvider.TransactionTotalAmount = resp.TransactionTotalAmount
	// 	respProvider.BillInfo = billInfo
	// 	statusCode = helpers.ErrorCodeGateway(respProvider.PaymentStatus, "PAY")
	// 	if statusCode == configs.PENDING_CODE {
	// 		result := helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMessage, statusDesc, respPayment)
	// 		return ctx.JSON(http.StatusOK, result)
	// 	}
	// 	err = UpdateAndInsertStatusTrx(resp, respProvider, svc)
	// 	if err != nil {
	// 		log.Println("Err UpdateAndInsertStatusTrx", svcName, err)
	// 		result := helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMessage, statusDesc, respPayment)
	// 		return ctx.JSON(http.StatusOK, result)
	// 	}

	// 	byte, _ := json.Marshal(respProvider.BillInfo)
	// 	statusMessage = "PAYMENT " + respProvider.PaymentStatusDesc
	// 	respPayment.BillInfo = string(byte)
	// 	result := helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMessage, statusDesc, respPayment)
	// 	return ctx.JSON(http.StatusOK, result)
	// case 2: //PREPAID
	// 	respProvider, err := svc.CheckStatusProviderSwitcher(models.ProviderInqRequest{
	// 		ReferenceNumber:      resp.ReferenceNumber,
	// 		Url:                  url,
	// 		ProviderName:         resp.ProviderName,
	// 		ProductReferenceCode: resp.ProductReferenceCode,
	// 		ProductReferenceId:   resp.ProductReferenceId,
	// 	})
	// 	// respProvider, err := iakworkerservice.IakPrepaidWorkerCheckStatus(models.ReqInqIak{
	// 	// 	RefId: resp.ReferenceNumber,
	// 	// 	Url:   url,
	// 	// })
	// 	if err != nil {
	// 		log.Println("Err ", svcName, "IakPrepaidiakworkerservice", err)
	// 		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", "Trx failed", nil)
	// 		return ctx.JSON(http.StatusOK, result)
	// 	}
	// 	respProvider.TrxReferenceNumber = resp.ReferenceNumber
	// 	respProvider.TrxProviderReferenceNumber = resp.ProviderReferenceNumber
	// 	respProvider.TransactionTotalAmount = resp.TransactionTotalAmount
	// 	err = UpdateAndInsertStatusTrx(resp, respProvider, svc)
	// 	if err != nil {
	// 		log.Println("Err UpdateAndInsertStatusTrx", svcName, err)
	// 		result := helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMessage, statusDesc, respPayment)
	// 		return ctx.JSON(http.StatusOK, result)
	// 	}
	// 	statusCode = helpers.ErrorCodeGateway(respProvider.PaymentStatus, "PAY")
	// 	if statusCode == configs.PENDING_CODE {
	// 		result := helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMessage, statusDesc, respPayment)
	// 		return ctx.JSON(http.StatusOK, result)
	// 	}
	// 	byte, _ := json.Marshal(respProvider.BillInfo)
	// 	statusMessage = "PAYMENT " + respProvider.PaymentStatusDesc
	// 	respPayment.BillInfo = string(byte)

	// 	result := helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMessage, statusDesc, respPayment)
	// 	return ctx.JSON(http.StatusOK, result)
	// default:
	// 	log.Println("Err ", svcName, "Unknown Product Type ID")
	// 	result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", "Unknown Product Type ID", nil)
	// 	return ctx.JSON(http.StatusOK, result)
	// }
	// if resp.ProductTypeId == 1 {
	// 	if configs.AppEnv == "DEV" {
	// 		url = configs.DevUrl + "/api/v1/bill/check"
	// 	}
	// 	if configs.AppEnv == "PROD" {
	// 		url = configs.IakProdUrlPostpaid + "/api/v1/bill/check"
	// 	}
	// 	respProvider, err := svc.CheckStatusProviderSwitcher(models.ProviderInqRequest{
	// 		ReferenceNumber:      resp.ReferenceNumber,
	// 		Url:                  url,
	// 		ProviderName:         resp.ProviderName,
	// 		ProductReferenceCode: resp.ProductReferenceCode,
	// 	})
	// 	if err != nil {
	// 		log.Println("Err ", svcName, "CheckStatusProviderSwitcher", err)
	// 		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.FAILED_CODE, configs.FAILED_MSG, "Trx failed", nil)
	// 		return ctx.JSON(http.StatusOK, result)
	// 	}
	// 	var billInfo map[string]interface{}
	// 	err = json.Unmarshal([]byte(resp.OtherMsg), &billInfo)
	// 	if err != nil {
	// 		log.Println("Err ", svcName, "Unmarshal", err)
	// 		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.FAILED_CODE, configs.FAILED_MSG, err.Error(), nil)
	// 		return ctx.JSON(http.StatusOK, result)
	// 	}
	// 	respProvider.TrxReferenceNumber = resp.ReferenceNumber
	// 	respProvider.TrxProviderReferenceNumber = resp.ProviderReferenceNumber
	// 	respProvider.TransactionTotalAmount = resp.TransactionTotalAmount
	// 	respProvider.BillInfo = billInfo
	// 	statusCode = helpers.ErrorCodeGateway(respProvider.PaymentStatus, "PAY")
	// 	if statusCode == configs.PENDING_CODE {
	// 		result := helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMessage, statusDesc, respPayment)
	// 		return ctx.JSON(http.StatusOK, result)
	// 	}
	// 	err = UpdateAndInsertStatusTrx(resp, respProvider, svc)
	// 	if err != nil {
	// 		log.Println("Err UpdateAndInsertStatusTrx", svcName, err)
	// 		result := helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMessage, statusDesc, respPayment)
	// 		return ctx.JSON(http.StatusOK, result)
	// 	}

	// 	byte, _ := json.Marshal(respProvider.BillInfo)
	// 	statusMessage = "PAYMENT " + respProvider.PaymentStatusDesc
	// 	respPayment.BillInfo = string(byte)
	// 	result := helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMessage, statusDesc, respPayment)
	// 	return ctx.JSON(http.StatusOK, result)
	// }
	// if resp.ProductTypeId == 2 {
	// 	if configs.AppEnv == "DEV" {
	// 		url = configs.IakDevUrlPrepaid + "/api/check-status"
	// 	}
	// 	if configs.AppEnv == "PROD" {
	// 		url = configs.IakProdUrlPrepaid + "/api/check-status"
	// 	}
	// 	respProvider, err := iakworkerservice.IakPrepaidWorkerCheckStatus(models.ReqInqIak{
	// 		RefId: resp.ReferenceNumber,
	// 		Url:   url,
	// 	})
	// 	if err != nil {
	// 		log.Println("Err ", svcName, "IakPrepaidiakworkerservice", err)
	// 		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", "Trx failed", nil)
	// 		return ctx.JSON(http.StatusOK, result)
	// 	}
	// 	respProvider.TrxReferenceNumber = resp.ReferenceNumber
	// 	respProvider.TrxProviderReferenceNumber = resp.ProviderReferenceNumber
	// 	respProvider.TransactionTotalAmount = resp.TransactionTotalAmount
	// 	err = UpdateAndInsertStatusTrx(resp, respProvider, svc)
	// 	if err != nil {
	// 		log.Println("Err UpdateAndInsertStatusTrx", svcName, err)
	// 		result := helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMessage, statusDesc, respPayment)
	// 		return ctx.JSON(http.StatusOK, result)
	// 	}
	// 	statusCode = helpers.ErrorCodeGateway(respProvider.PaymentStatus, "PAY")
	// 	if statusCode == configs.PENDING_CODE {
	// 		result := helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMessage, statusDesc, respPayment)
	// 		return ctx.JSON(http.StatusOK, result)
	// 	}
	// 	byte, _ := json.Marshal(respProvider.BillInfo)
	// 	statusMessage = "PAYMENT " + respProvider.PaymentStatusDesc
	// 	respPayment.BillInfo = string(byte)

	result := helpers.ResponseJSON(configs.TRUE_VALUE, statusCode, statusMsg, statusDesc, respPayment)
	return ctx.JSON(http.StatusOK, result)
	// }
	// return nil
}
func UpdateAndInsertStatusTrx(dataPayment models.RespGetTrx, dataAdvice models.ResponseWorkerPayment, svc trxService) error {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMPTRX)
	statusCode := helpers.ErrorCodeGateway(dataAdvice.PaymentStatus, "PAY")
	if statusCode == configs.PENDING_CODE {
		return errors.New("trx was pending")
	}
	updatePayment := models.ReqGetTransaction{
		Filter: models.Transaction{
			ProductReferenceID:         dataPayment.ProductReferenceID,
			ProductReferenceCode:       dataPayment.ProductReferenceCode,
			ProductCategoryID:          dataPayment.ProductCategoryID,
			ProductCategoryName:        dataPayment.ProductCategoryName,
			ProductTypeID:              dataPayment.ProductTypeID,
			ProductTypeName:            dataPayment.ProductTypeName,
			ProductID:                  dataPayment.ProductID,
			ProductName:                dataPayment.ProductName,
			ProductCode:                dataPayment.ProductCode,
			ProductPrice:               dataPayment.ProductPrice,
			ProductAdminFee:            dataPayment.ProductAdminFee,
			ProductMerchantFee:         dataPayment.ProductMerchantFee,
			ProductProviderName:        dataPayment.ProductProviderName,
			ProductProviderCode:        dataPayment.ProductProviderCode,
			ProductProviderPrice:       dataPayment.ProductProviderPrice,
			ProductProviderAdminFee:    dataPayment.ProductProviderAdminFee,
			ProductProviderMerchantFee: dataPayment.ProductProviderMerchantFee,
			StatusCode:                 dataAdvice.PaymentStatus,
			StatusMessage:              dataAdvice.PaymentStatusMsg,
			StatusDesc:                 dataAdvice.PaymentStatusDesc,
			ReferenceNumber:            dataPayment.ReferenceNumber,
			StatusCodeDetail:           dataAdvice.PaymentStatusDetail,
			StatusDescDetail:           dataAdvice.PaymentStatusDescDetail,
			StatusMessageDetail:        dataAdvice.PaymentStatusMsgDetail,
			ProviderReferenceNumber:    dataAdvice.TrxProviderReferenceNumber,
			OtherReff:                  dataPayment.OtherReff,
			OtherCustomerInfo:          dataPayment.OtherCustomerInfo,
			SavingAccountName:          dataPayment.SavingAccountName,
			SavingAccountID:            dataPayment.SavingAccountID,
			SavingAccountNumber:        dataPayment.SavingAccountNumber,
			TransactionTotalAmount:     dataPayment.TransactionTotalAmount,
			UserAppID:                  dataPayment.UserAppID,
			Username:                   dataPayment.Username,
			CreatedAt:                  dataPayment.CreatedAt,
			CreatedBy:                  dataPayment.CreatedBy,
			UpdatedAt:                  dbTime,
			UpdatedBy:                  "sys",
			ID:                         int64(dataPayment.Id),
			CustomerID:                 dataPayment.CustomerID,
		},
	}
	if statusCode == configs.SUCCESS_CODE {
		byte, _ := json.Marshal(dataAdvice.BillInfo)
		updatePayment.Filter.OtherReff = string(byte)
	}
	err := helpers.DBTransaction(svc.services.RepoDB, func(Tx *sql.Tx) error {
		err := svc.services.RepoTrx.UpdateTrx(updatePayment, Tx)
		if err != nil {
			// log.Println("Err ", svcName, "UpdateTrx", err)
			// result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_VALIDATION_FAILED[0], configs.RC_VALIDATION_FAILED[1], err.Error(), nil)
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

package trxservice

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	helperservice "desabiller/services/helperService"
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
		statusMessage,
		statusDesc,
		providerStatusCode,
		providerStatusMessage,
		providerStatusDesc string
		// respSvc models.ResponseList
	)
	req := new(models.ReqAviceTrx)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.ApiTrx.GetTrx(models.ReqGetTrx{
		ReferenceNumber: req.ReferenceNumber,
	})
	if err != nil {
		log.Println("Err ", svcName, "GetTrx", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "transaction not found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	statusCode = resp.StatusCode
	statusMessage = resp.StatusMessage
	statusDesc = resp.StatusDesc
	providerStatusCode = resp.ProviderStatusCode
	providerStatusMessage = resp.ProviderStatusMessage
	providerStatusDesc = resp.ProviderStatusDesc
	respPayment := models.RespPayment{
		Id:                      resp.Id,
		StatusCode:              statusCode,
		StatusMessage:           statusMessage,
		StatusDesc:              statusDesc,
		ReferenceNumber:         resp.ReferenceNumber,
		ProviderStatusCode:      providerStatusCode,
		ProviderStatusMessage:   providerStatusMessage,
		ProviderStatusDesc:      providerStatusDesc,
		ProviderReferenceNumber: resp.ProviderReferenceNumber,
		CreatedAt:               resp.CreatedAt,
		UpdatedAt:               resp.UpdatedAt,
		CustomerId:              resp.CustomerId,
		OtherMsg:                resp.OtherMsg,
		ProductId:               resp.ProductId,
		ProductName:             resp.ProductName,
		ProductCode:             resp.ProductCode,
		ProductPrice:            resp.ProductPrice,
		ProductAdminFee:         resp.ProductAdminFee,
		ProductMerchantFee:      resp.ProductMerchantFee,
		ClientId:                resp.ClientId,
		ClientName:              resp.ClientName,
		GroupId:                 resp.GroupId,
		GroupName:               resp.GroupName,
		MerchantId:              resp.MerchantId,
		MerchantName:            resp.MerchantName,
		MerchantOutletId:        resp.MerchantOutletId,
		MerchantOutletName:      resp.MerchantOutletName,
		MerchantOutletUsername:  resp.MerchantOutletUsername,
	}
	if resp.StatusCode != configs.PENDING_CODE {
		respPayment.StatusCode = statusCode
		respPayment.StatusMessage = statusMessage
		respPayment.StatusDesc = statusDesc
		respPayment.ProviderStatusCode = providerStatusCode
		respPayment.ProviderStatusMessage = providerStatusMessage
		respPayment.ProviderStatusDesc = providerStatusDesc
		result := helpers.ResponseJSON(configs.TRUE_VALUE, resp.StatusCode, resp.StatusMessage, respPayment)
		return ctx.JSON(http.StatusOK, result)
	}
	if resp.ProductTypeId == 1 {
		if configs.AppEnv == "DEV" {
			url = configs.IakDevUrlPostpaid + "/api/v1/bill/check"
		}
		if configs.AppEnv == "PROD" {
			url = configs.IakProdUrlPostpaid + "/api/v1/bill/check"
		}
	}
	if resp.ProductTypeId == 2 {
		if configs.AppEnv == "DEV" {
			url = configs.IakDevUrlPrepaid + "/api/check-status"
		}
		if configs.AppEnv == "PROD" {
			url = configs.IakProdUrlPrepaid + "/api/check-status"
		}
		respProvider, err := helperservice.IakPrepaidHelperService(models.ReqPaymentPrepaidIak{
			// CustomerId:  resp.CustomerId,
			// ProductCode: resp.ProductProviderCode,
			RefId:    resp.ReferenceNumber,
			Username: configs.IakUsername,
			Sign:     helpers.SignIakEncrypt(""),
		}, url)
		if err != nil {
			log.Println("Err ", svcName, "IakPrepaidHelperService", err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Trx failed", nil)
			return ctx.JSON(http.StatusOK, result)
		}

		err = UpdateAndInsertStatusTrx(resp, respProvider, svc)
		if err != nil {
			log.Println("Err UpdateAndInsertStatusTrx", svcName, err)
			respPayment.StatusCode = statusCode
			respPayment.StatusMessage = statusMessage
			respPayment.StatusDesc = statusDesc
			respPayment.ProviderStatusCode = providerStatusCode
			respPayment.ProviderStatusMessage = providerStatusMessage
			respPayment.ProviderStatusDesc = providerStatusDesc
			result := helpers.ResponseJSON(configs.TRUE_VALUE, resp.StatusCode, resp.StatusMessage, respPayment)
			return ctx.JSON(http.StatusOK, result)
		}
		statusCode = helpers.ErrorCodeGateway(respProvider.PaymentStatus)
		if statusCode == configs.PENDING_CODE {
			result := helpers.ResponseJSON(configs.TRUE_VALUE, resp.StatusCode, resp.StatusMessage, respPayment)
			return ctx.JSON(http.StatusOK, result)
		}
		respPayment.StatusCode = statusCode
		respPayment.StatusMessage = "PAYMENT " + respProvider.PaymentStatusDesc
		respPayment.StatusDesc = respProvider.PaymentStatusDesc
		respPayment.ProviderStatusCode = respProvider.PaymentStatusDetail
		respPayment.ProviderStatusMessage = respProvider.PaymentStatusDescDetail
		respPayment.ProviderStatusDesc = respProvider.PaymentStatusDescDetail
		result := helpers.ResponseJSON(configs.TRUE_VALUE, respPayment.StatusCode, respPayment.StatusMessage, respPayment)
		return ctx.JSON(http.StatusOK, result)
	}
	return nil
}
func UpdateAndInsertStatusTrx(dataPayment models.RespGetTrx, dataAdvice models.ResponseWorkerPayment, svc trxService) error {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMPTRX)
	statusCode := helpers.ErrorCodeGateway(dataAdvice.PaymentStatus)
	if statusCode == configs.PENDING_CODE {
		return errors.New("trx was pending")
	}
	updatePayment := models.ReqGetTrx{
		ProductClanId:              dataPayment.ProductClanId,
		ProductClanName:            dataPayment.ProductClanName,
		ProductCategoryId:          dataPayment.ProductCategoryId,
		ProductCategoryName:        dataPayment.ProductCategoryName,
		ProductTypeId:              dataPayment.ProductTypeId,
		ProductTypeName:            dataPayment.ProductTypeName,
		ProductId:                  dataPayment.ProductId,
		ProductName:                dataPayment.ProductName,
		ProductCode:                dataPayment.ProductCode,
		ProductPrice:               dataPayment.ProductPrice,
		ProductAdminFee:            dataPayment.ProductAdminFee,
		ProductMerchantFee:         dataPayment.ProductMerchantFee,
		ProductProviderId:          dataPayment.ProductProviderId,
		ProductProviderName:        dataPayment.ProductProviderName,
		ProductProviderCode:        dataPayment.ProductCode,
		ProductProviderPrice:       dataPayment.ProductProviderPrice,
		ProductProviderAdminFee:    dataPayment.ProductProviderAdminFee,
		ProductProviderMerchantFee: dataPayment.ProductProviderMerchantFee,
		StatusCode:                 statusCode,
		StatusMessage:              "PAYMENT " + dataAdvice.PaymentStatusDesc,
		StatusDesc:                 dataAdvice.PaymentStatusDesc,
		ReferenceNumber:            dataPayment.ReferenceNumber,
		ProviderStatusCode:         dataAdvice.PaymentStatusDetail,
		ProviderStatusMessage:      dataAdvice.PaymentStatusDescDetail,
		ProviderStatusDesc:         dataAdvice.PaymentStatusDescDetail,
		ProviderReferenceNumber:    dataAdvice.TrxProviderReferenceNumber,
		ClientId:                   dataPayment.ClientId,
		ClientName:                 dataPayment.ClientName,
		GroupId:                    dataPayment.GroupId,
		GroupName:                  dataPayment.GroupName,
		MerchantId:                 dataPayment.MerchantId,
		MerchantName:               dataPayment.MerchantName,
		MerchantOutletId:           dataPayment.MerchantOutletId,
		MerchantOutletName:         dataPayment.MerchantOutletName,
		MerchantOutletUsername:     dataPayment.MerchantOutletUsername,
		CustomerId:                 dataPayment.CustomerId,
		OtherMsg:                   dataPayment.OtherMsg,
		Filter: models.FilterReq{
			UpdatedAt: dbTime,
			UpdatedBy: "sys",
		},
	}
	err := helpers.DBTransaction(svc.services.RepoDB, func(Tx *sql.Tx) error {
		err := svc.services.ApiTrx.UpdateTrx(updatePayment, Tx)
		if err != nil {
			// log.Println("Err ", svcName, "UpdateTrx", err)
			// result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
			return err
		}
		err = svc.services.ApiTrx.InsertTrxStatus(models.ReqGetTrxStatus{
			ReferenceNumber:         updatePayment.ReferenceNumber,
			ProviderReferenceNumber: updatePayment.ProviderReferenceNumber,
			StatusCode:              updatePayment.StatusCode,
			StatusMessage:           updatePayment.StatusMessage,
		}, Tx)
		if err != nil {
			// log.Println("Err ", svcName, "InsertTrxStatus", err)
			// result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

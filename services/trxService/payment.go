package trxservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	iakworkerservice "desabiller/services/IAKWorkerService"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (svc trxService) PaymentBiller(ctx echo.Context) error {
	var (
		svcName = "PaymentBiller"
		statusCode,
		statusMessage,
		statusDesc,
		referenceNumber,
		providerStatusCode,
		providerStatusMessage,
		providerStatusDesc,
		providerReferenceNumber,
		url string
		// billInfo     map[string]interface{}
		respProvider models.ResponseWorkerPayment
		// respSvc models.ResponseList
		// respOutlet models.RespGetMerchantOutlet
	)
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMPTRX)
	req := new(models.ReqPaymentTrx)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.FAILED_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ReferenceNumber == "" {
		log.Println("Err ", svcName, "ReferenceNumber id cannot be null")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.FAILED_CODE, "ReferenceNumber id cannot be null", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respInqTrx, err := svc.services.ApiTrx.GetTrx(models.ReqGetTrx{
		ReferenceNumber: req.ReferenceNumber,
	})
	if err != nil {
		log.Println("Err ", svcName, "GetTrx", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.FAILED_CODE, "transaction not found", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	if respInqTrx.StatusCode != configs.INQUIRY_SUCCESS_CODE {
		log.Println("Err ", svcName, "Transaction invalid")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.FAILED_CODE, "Transaction invalid", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	fmt.Println("SINI 1")
	billDescInq := models.BillInfoBPJS{}
	// billDescPay := BillInfoBPJS{}
	switch respInqTrx.ProductTypeId {
	case 1:
		if configs.AppEnv == "DEV" {
			url = configs.IakDevUrlPostpaid + configs.ENDPOINT_IAK_POSTPAID
		}
		if configs.AppEnv == "PROD" {
			url = configs.IakProdUrlPostpaid + configs.ENDPOINT_IAK_POSTPAID
		}
		fmt.Println("SINI 2")

		respProvider, err = svc.PayProviderSwitcher(models.ProviderPayRequest{
			ProviderReferenceNumber: respInqTrx.ProviderReferenceNumber,
			Url:                     url,
			ProductClan:             respInqTrx.ProductClanName,
			ProviderName:            respInqTrx.ProviderName,
		})
		if err != nil {
			log.Println("Err ", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.PENDING_CODE, configs.PENDING_MSG, nil)
			return ctx.JSON(http.StatusOK, result)
		}
		switch respInqTrx.ProductClanName {
		case "BPJSKS":
			otherMsgInq := helpers.JsonDescape(respInqTrx.OtherMsg)
			_ = json.Unmarshal([]byte(otherMsgInq), &billDescInq)
			if len(respProvider.BillInfo) != 0 {
				snVal, _ := respProvider.BillInfo["sn"].(string)
				billDescInq.Sn = snVal
			}
		}
	case 2:
		if configs.AppEnv == "DEV" {
			url = configs.IakDevUrlPrepaid + configs.ENDPOINT_IAK_PREPAID
		}
		if configs.AppEnv == "PROD" {
			url = configs.IakProdUrlPrepaid + configs.ENDPOINT_IAK_PREPAID
		}
		if respInqTrx.ProviderId == 1 { //IAK
			respProvider, err = iakworkerservice.IakPulsaWorkerPayment(models.ReqInqIak{
				CustomerId:  respInqTrx.CustomerId,
				ProductCode: respInqTrx.ProductProviderCode,
				RefId:       respInqTrx.ReferenceNumber,
				Url:         url,
			})
			if err != nil {
				log.Println("Err ", svcName, "IakPulsaWorkerPayment", err)
				result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.PENDING_CODE, configs.PENDING_MSG, nil)
				return ctx.JSON(http.StatusOK, result)
			}
		}
	}

	byte, _ := json.Marshal(billDescInq)
	updatePayment := models.ReqGetTrx{
		ProductClanId:              respInqTrx.ProductClanId,
		ProductClanName:            respInqTrx.ProductClanName,
		ProductCategoryId:          respInqTrx.ProductCategoryId,
		ProductCategoryName:        respInqTrx.ProductCategoryName,
		ProductTypeId:              respInqTrx.ProductTypeId,
		ProductTypeName:            respInqTrx.ProductTypeName,
		ProductId:                  respInqTrx.ProductId,
		ProductName:                respInqTrx.ProductName,
		ProductCode:                respInqTrx.ProductCode,
		ProductPrice:               respInqTrx.ProductPrice,
		ProductAdminFee:            respInqTrx.ProductAdminFee,
		ProductMerchantFee:         respInqTrx.ProductMerchantFee,
		ProductProviderId:          respInqTrx.ProductProviderId,
		ProductProviderName:        respInqTrx.ProductProviderName,
		ProductProviderCode:        respInqTrx.ProductCode,
		ProductProviderPrice:       respInqTrx.ProductProviderPrice,
		ProductProviderAdminFee:    respInqTrx.ProductProviderAdminFee,
		ProductProviderMerchantFee: respInqTrx.ProductProviderMerchantFee,
		StatusCode:                 statusCode,              //statusCode,
		StatusMessage:              statusMessage,           //"PAYMENT " + respProvider.PaymentStatusDesc,
		StatusDesc:                 statusDesc,              //respProvider.PaymentStatusDesc,
		ReferenceNumber:            referenceNumber,         //respInqTrx.ReferenceNumber,
		ProviderStatusCode:         providerStatusCode,      //respProvider.PaymentStatusDetail,
		ProviderStatusMessage:      providerStatusMessage,   //respProvider.PaymentStatusDescDetail,
		ProviderStatusDesc:         providerStatusDesc,      //respProvider.PaymentStatusDescDetail,
		ProviderReferenceNumber:    providerReferenceNumber, //respProvider.TrxProviderReferenceNumber,
		ClientId:                   respInqTrx.ClientId,
		ClientName:                 respInqTrx.ClientName,
		GroupId:                    respInqTrx.GroupId,
		GroupName:                  respInqTrx.GroupName,
		MerchantId:                 respInqTrx.MerchantId,
		MerchantName:               respInqTrx.MerchantName,
		MerchantOutletId:           respInqTrx.MerchantOutletId,
		MerchantOutletName:         respInqTrx.MerchantOutletName,
		MerchantOutletUsername:     respInqTrx.MerchantOutletUsername,
		CustomerId:                 respInqTrx.CustomerId,
		OtherMsg:                   respInqTrx.OtherMsg,
		Filter: models.FilterReq{
			CreatedAt: dbTime,
		},
	}
	statusCode = helpers.ErrorCodeGateway(respProvider.PaymentStatus, "PAY")
	updatePayment.StatusCode = statusCode
	updatePayment.StatusMessage = "PAYMENT " + respProvider.PaymentStatusDesc
	updatePayment.StatusDesc = respProvider.PaymentStatusDesc
	updatePayment.ReferenceNumber = respInqTrx.ReferenceNumber
	updatePayment.ProviderStatusCode = respProvider.PaymentStatusDetail
	updatePayment.ProviderStatusMessage = respProvider.PaymentStatusDescDetail
	updatePayment.ProviderStatusDesc = respProvider.PaymentStatusDescDetail
	updatePayment.ProviderReferenceNumber = respProvider.TrxProviderReferenceNumber
	updatePayment.OtherMsg = string(byte)
	updatePayment.Filter = models.FilterReq{
		CreatedAt: dbTime,
	}
	updatePayment.TotalTrxAmount = respProvider.TotalTrxAmount
	if configs.TrxPaymentPending == "YES" && statusCode == configs.SUCCESS_CODE {
		updatePayment.StatusCode = configs.PENDING_CODE
		updatePayment.StatusMessage = "PAYMENT PENDING"
	}
	err = svc.services.ApiTrx.UpdateTrx(updatePayment, nil)
	if err != nil {
		log.Println("Err ", svcName, "UpdateTrx", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.PENDING_CODE, configs.PENDING_MSG, nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.ApiTrx.InsertTrxStatus(models.ReqGetTrxStatus{
		ReferenceNumber:         updatePayment.ReferenceNumber,
		ProviderReferenceNumber: updatePayment.ProviderReferenceNumber,
		StatusCode:              updatePayment.StatusCode,
		StatusMessage:           updatePayment.StatusMessage,
	}, nil)
	if err != nil {
		log.Println("Err ", svcName, "InsertTrxStatus", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.PENDING_CODE, configs.PENDING_MSG, nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if updatePayment.StatusCode == configs.FAILED_CODE {
		result := helpers.ResponseJSON(configs.TRUE_VALUE, updatePayment.StatusCode, updatePayment.StatusMessage, nil)
		return ctx.JSON(http.StatusOK, result)
	}
	responsePayment := models.RespPayment{
		ReferenceNumber:        updatePayment.ReferenceNumber,
		CreatedAt:              respInqTrx.CreatedAt,
		SubscriberNumber:       updatePayment.CustomerId,
		BillInfo:               billDescInq,
		ProductName:            updatePayment.ProductName,
		ProductCode:            updatePayment.ProductCode,
		ProductPrice:           updatePayment.ProductPrice,
		ProductAdminFee:        updatePayment.ProductAdminFee,
		MerchantOutletName:     updatePayment.MerchantOutletName,
		MerchantOutletUsername: updatePayment.MerchantOutletUsername,
		TotalTrxAmount:         updatePayment.ProductPrice + updatePayment.ProductAdminFee,
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, updatePayment.StatusCode, updatePayment.StatusMessage, responsePayment)
	return ctx.JSON(http.StatusOK, result)
}

package iakworkerservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/utils"
	"encoding/json"
	"log"
	"strconv"
)

func IakBPJSWorkerPayment(req models.ReqInqIak) (respWorker models.ResponseWorkerPayment, err error) {

	var (
		helperName   = "[IAK][WKR]IakBPJSWorkerPayment"
		respProvider models.RespPaymentBPJSIak
		statusCode,
		statusMsg, statusDesc,
		statusCodeDetail,
		statusMsgDetail string
		// paymentDetail    models.PaymentDetails
		respUndefined  models.RespWorkerUndefined
		respUndefinedI models.RespWorkerUndefinedI
	)
	providerRequest := models.ReqPaymentPostpaidIak{
		Commands: "pay-pasca",
		Username: configs.IakUsername,
		TrID:     req.RefId,
		Sign:     helpers.SignIakEncrypt(req.RefId),
	}
	respByte, _, err := utils.WorkerPostWithBearer(req.Url, "", providerRequest, "json")
	if err != nil {
		log.Println("Err ", helperName, err)
		return respWorker, err
	}
	//bind struct response
	err = json.Unmarshal(respByte, &respProvider)
	if err != nil {
		log.Println("Err ", helperName, err)
		return respWorker, err
	}
	if respProvider.Data.RefID == "" {
		err = json.Unmarshal(respByte, &respUndefined)
		if err != nil {
			log.Println("Err ", helperName, err)
			return respWorker, err
		}
		if respUndefined.ResponseCode == "" {
			err = json.Unmarshal(respByte, &respUndefinedI)
			if err != nil {
				log.Println("Err ", helperName, err)
				return respWorker, err
			}
			respProvider.Data.ResponseCode = respUndefinedI.Data.ResponseCode
			respProvider.Data.Message = respUndefinedI.Data.Message
		} else {
			respProvider.Data.ResponseCode = respUndefined.ResponseCode
			respProvider.Data.Message = respUndefined.Message
		}
	}
	statusCodeDetail = respProvider.Data.ResponseCode
	statusMsgDetail = respProvider.Data.Message
	statusCode, statusMsg, statusDesc = helpers.IakResponseConverter(respProvider.Data.ResponseCode, respProvider.Data.Message)
	respWorker.BillInfo = map[string]interface{}{
		"sn": respProvider.Data.Noref,
	}
	respWorker.PaymentStatus = statusCode
	respWorker.PaymentStatusMsg = statusMsg
	respWorker.PaymentStatusDesc = statusDesc
	respWorker.PaymentStatusDetail = statusCodeDetail
	respWorker.PaymentStatusDescDetail = statusMsgDetail
	// respWorker.TotalTrxAmount, _ = strconv.ParseFloat(strconv.Itoa(respProvider.Data.Price), 64)
	respWorker.TrxReferenceNumber = providerRequest.TrID
	respWorker.TrxProviderReferenceNumber = strconv.Itoa(respProvider.Data.TrID)

	return respWorker, nil
}

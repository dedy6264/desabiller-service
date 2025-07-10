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

func IakPrepaidWorkerCheckStatus(req models.ReqInqIak) (respWorker models.ResponseWorkerPayment, err error) {
	var (
		helperName       = "[IAK][WKR]IakPrepaidWorkerCheckStatus"
		respProvider     models.RespCheckStatusPrepaidIak
		statusCode       string
		statusMsg        string
		statusCodeDetail string
		statusMsgDetail  string
		// paymentDetail    models.PaymentDetails
		respUndefined  models.RespWorkerUndefined
		respUndefinedI models.RespWorkerUndefinedI
	)
	providerRequest := models.ReqPaymentPrepaidIak{
		CustomerId:  req.CustomerId,
		ProductCode: req.ProductCode,
		RefId:       req.RefId,
		Username:    configs.IakUsername,
		Sign:        helpers.SignIakEncrypt(req.RefId),
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
			respProvider.Data.Rc = respUndefinedI.Data.ResponseCode
			respProvider.Data.Message = respUndefinedI.Data.Message
		} else {
			respProvider.Data.Rc = respUndefined.ResponseCode
			respProvider.Data.Message = respUndefined.Message
		}
	}
	statusCodeDetail = respProvider.Data.Rc
	statusMsgDetail = respProvider.Data.Message
	if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"201", "39", "05", "02"}); ok {
		statusCode = configs.WORKER_PENDING_CODE
		statusMsg = "PENDING"
	}
	if respProvider.Data.Rc != "00" {
		switch respProvider.Data.Rc {
		case "07":
			statusCode = configs.WORKER_FAILED_CODE
			statusMsg = "FAILED"
		default:
			statusCode = configs.WORKER_PENDING_CODE
			statusMsg = "PENDING"
		}
	}
	// respWorker.PaymentDetail = paymentDetail
	respWorker.PaymentStatus = statusCode
	respWorker.PaymentStatusDesc = statusMsg
	respWorker.PaymentStatusDetail = statusCodeDetail
	respWorker.PaymentStatusDescDetail = statusMsgDetail
	respWorker.TotalTrxAmount, _ = strconv.ParseFloat(strconv.Itoa(respProvider.Data.Price), 64)
	respWorker.TrxProviderReferenceNumber = strconv.Itoa(respProvider.Data.TrID)
	respWorker.BillInfo = map[string]interface{}{
		"sn": respProvider.Data.Sn,
	}

	return respWorker, nil
}

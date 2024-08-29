package helperIakservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/utils"
	"encoding/json"
	"log"
	"strconv"
)

func IakPulsaWorkerPayment(req models.ReqInqIak) (respWorker models.ResponseWorkerPayment, err error) {

	var (
		helperName       = "[IAK][WKR]IakPulsaWorkerPayment"
		respProvider     models.RespPaymentPrepaidIak
		statusCode       string
		statusMsg        string
		statusCodeDetail string
		statusMsgDetail  string
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
	statusCodeDetail = respProvider.Data.Rc
	statusMsgDetail = respProvider.Data.Message
	if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"201", "39"}); ok {
		statusCode = configs.WORKER_PENDING_CODE
		statusMsg = "PENDING"
	}
	if respProvider.Data.Rc != "00" {
		if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"06", "07", "13", "18", "20", "21", "132", "106"}); ok {
			statusCode = configs.WORKER_FAILED_CODE
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"203", "205", "107"}); ok {
			statusCode = configs.WORKER_INVALID_PARAM
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"102"}); ok {
			statusCode = configs.WORKER_CREDENTIAL_ERROR
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"14", "16", "19", "131", "141", "142", "206"}); ok {
			statusCode = configs.WORKER_VALIDATION_ERROR
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"404", "12", "204", "17", "110", "202", "207", "121", "117", "10"}); ok {
			statusCode = configs.WORKER_SYSTEM_ERROR
			statusMsg = "FAILED"
		}
	} else {
		statusCode = configs.WORKER_SUCCESS_CODE
		statusMsg = "SUCCESS"
		// billdesc := models.BillDescPulsa{
		// 	CustomerId: respProvider.Data.CustomerID,
		// 	Sn:         respProvider.Data.Sn,
		// }
		// byte, _ := json.Marshal(billdesc)
		respWorker.BillInfo = map[string]interface{}{
			"customerId": respProvider.Data.CustomerID,
			"sn":         respProvider.Data.Sn,
		}
		// respWorker.BillInfo = string(byte)
	}
	respWorker.PaymentStatus = statusCode
	respWorker.PaymentStatusDesc = statusMsg
	respWorker.PaymentStatusDetail = statusCodeDetail
	respWorker.PaymentStatusDescDetail = statusMsgDetail
	respWorker.TotalTrxAmount, _ = strconv.ParseFloat(strconv.Itoa(respProvider.Data.Price), 64)
	respWorker.TrxReferenceNumber = providerRequest.RefId
	respWorker.TrxProviderReferenceNumber = strconv.Itoa(respProvider.Data.TrID)

	return respWorker, nil
}

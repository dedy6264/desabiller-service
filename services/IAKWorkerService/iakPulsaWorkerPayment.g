package iakworkerservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/utils"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func IakPulsaWorkerPayment(req models.ReqInqIak) (respWorker models.ResponseWorkerPayment, err error) {

	var (
		helperName   = "[IAK][WKR]IakPulsaWorkerPayment"
		respProvider models.RespPaymentPrepaidIak
		statusCode, statusDesc,
		statusMsg,
		statusCodeDetail,
		statusMsgDetail string
		respUndefined   models.RespWorkerUndefined
		respUndefinedI  models.RespWorkerUndefinedI
		respUndefinedII models.RespWorkerUndefinedII
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
			if respUndefinedI.Data.ResponseCode == "" {
				err = json.Unmarshal(respByte, &respUndefinedII)
				if err != nil {
					log.Println("Err ", helperName, err)
					return respWorker, err
				}
				respProvider.Data.Rc = respUndefinedII.Data.Rc
				respProvider.Data.Message = respUndefinedII.Data.Message
			} else {
				respProvider.Data.Rc = respUndefinedI.Data.ResponseCode
				respProvider.Data.Message = respUndefinedI.Data.Message
			}
		} else {
			respProvider.Data.Rc = respUndefined.ResponseCode
			respProvider.Data.Message = respUndefined.Message
		}
	}
	byte, _ := json.Marshal(respProvider)

	statusCodeDetail = respProvider.Data.Rc
	fmt.Println("===", respProvider, string(byte), "=====", statusCodeDetail)
	statusMsgDetail = respProvider.Data.Message
	// if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"201", "39", "05", "02"}); ok {
	// 	statusCode = configs.WORKER_PENDING_CODE
	// 	statusMsg = "PENDING"
	// 	statusDesc = configs.BILLER_DISRUPTION_MSG

	// }
	// if respProvider.Data.Rc != "00" {
	// 	if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"06", "07", "13", "18", "20", "21", "132", "106", "30", "33", "37", "38", "91", "92", "105"}); ok {
	// 		statusCode = configs.WORKER_FAILED_CODE
	// 		statusMsg = "FAILED"
	// 		statusDesc = configs.BILLER_DISRUPTION_MSG

	// 	}
	// 	if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"203", "205", "107", "93"}); ok {
	// 		statusCode = configs.WORKER_INVALID_PARAM
	// 		statusMsg = "FAILED"
	// 		statusDesc = configs.BILLER_DISRUPTION_MSG

	// 	}
	// 	if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"102"}); ok {
	// 		statusCode = configs.WORKER_CREDENTIAL_ERROR
	// 		statusMsg = "FAILED"
	// 		statusDesc = configs.BILLER_DISRUPTION_MSG

	// 	}
	// 	if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"14", "16", "19", "131", "141", "142", "206", "01", "03", "04", "08", "11", "31", "32", "34", "35", "36", "40", "41", "42", "100", "101", "103"}); ok {
	// 		statusCode = configs.WORKER_VALIDATION_ERROR
	// 		statusMsg = "FAILED"
	// 		statusDesc = configs.BILLER_DISRUPTION_MSG

	// 	}
	// 	if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"404", "12", "204", "17", "110", "202", "207", "121", "117", "10", "94", "108", "109"}); ok {
	// 		statusCode = configs.WORKER_SYSTEM_ERROR
	// 		statusMsg = "FAILED"
	// 		statusDesc = configs.BILLER_DISRUPTION_MSG

	// 	}
	// } else {
	// 	statusCode = configs.WORKER_SUCCESS_CODE
	// 	statusMsg = "SUCCESS"
	// 	// billdesc := models.BillDescPulsa{
	// 	// 	CustomerId: respProvider.Data.CustomerID,
	// 	// 	Sn:         respProvider.Data.Sn,
	// 	// }
	// 	// byte, _ := json.Marshal(billdesc)
	// 	respWorker.BillInfo = map[string]interface{}{
	// 		"customerId": respProvider.Data.CustomerID,
	// 		"sn":         respProvider.Data.Sn,
	// 	}
	// 	// respWorker.BillInfo = string(byte)
	// }
	if respProvider.Data.Rc != "00" {
		if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"07", "08", "09", "91", "92", "94", "102", "103", "105", "109", "110", "117", "02", "37", "38", "04", "202"}); ok { //failed/Biller Disruption
			statusCode = configs.BILLER_DISRUPTION
			statusDesc = configs.BILLER_DISRUPTION_MSG
			statusMsg = configs.FAILED_MSG
		} else if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"01", "34"}); ok { //failed/paid bill
			statusCode = configs.WORKER_FAILED_CODE
			statusDesc = configs.BIIL_PAID_MSG
			statusMsg = configs.FAILED_MSG
		} else if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"201"}); ok { //failed/Undefined
			statusCode = configs.WORKER_UNDEFINED_ERROR
			statusDesc = configs.UNDEFINED_MSG
			statusMsg = configs.FAILED_MSG
		} else if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"10", "15", "18", "30", "33", "32", "16", "108", "14", "31", "35", "36", "40", "41", "76", "77"}); ok { //failed/BILL UNREADY
			statusCode = configs.WORKER_BILL_NOTFOUND_CODE
			statusDesc = configs.BILL_NOTFOUND_MSG
			statusMsg = configs.FAILED_MSG
		} else if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"39"}); ok { //failed/expired
			statusCode = configs.WORKER_PENDING_CODE
			statusMsg = configs.PENDING_MSG
			statusDesc = statusMsg
		} else {
			statusCode = configs.WORKER_PENDING_CODE
			statusMsg = configs.PENDING_MSG
			statusDesc = statusMsg
		}
	} else {
		statusCode = configs.WORKER_SUCCESS_CODE
		statusMsg = configs.SUCCESS_MSG
		statusDesc = statusMsg
		respWorker.BillInfo = map[string]interface{}{
			"customerId": respProvider.Data.CustomerID,
			"sn":         respProvider.Data.Sn,
		}
	}
	fmt.Println("++++", statusCode, "||", statusMsg)
	respWorker.PaymentStatus = statusCode
	respWorker.PaymentStatusDesc = statusDesc
	respWorker.PaymentStatusMsg = statusMsg
	respWorker.PaymentStatusDetail = statusCodeDetail
	respWorker.PaymentStatusDescDetail = statusMsgDetail
	respWorker.TotalTrxAmount, _ = strconv.ParseFloat(strconv.Itoa(respProvider.Data.Price), 64)
	respWorker.TrxReferenceNumber = providerRequest.RefId
	respWorker.TrxProviderReferenceNumber = strconv.Itoa(respProvider.Data.TrID)

	return respWorker, nil
}

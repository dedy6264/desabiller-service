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

func IakPLNPostpaidWorkerPayment(req models.ReqInqIak) (respWorker models.ResponseWorkerPayment, err error) {

	var (
		helperName       = "[IAK][WKR]IakPLNPostpaidWorkerPayment"
		respProvider     models.RespPaymentPLNPostpaidIak
		statusCode       string
		statusMsg        string
		statusCodeDetail string
		statusMsgDetail  string
		paymentDetail    models.PaymentDetails
		respUndefined    models.RespWorkerUndefined
		respUndefinedI   models.RespWorkerUndefinedI
	)
	providerRequest := models.ReqPaymentPostpaidIak{
		Commands: req.Commands,
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
	if ok, _ := helpers.InArray(respProvider.Data.ResponseCode, []string{"201", "39", "05", "02"}); ok {
		statusCode = configs.WORKER_PENDING_CODE
		statusMsg = "PENDING"
	}
	if respProvider.Data.ResponseCode != "00" {
		if ok, _ := helpers.InArray(respProvider.Data.ResponseCode, []string{"06", "07", "13", "18", "20", "21", "132", "106", "09", "30", "33", "37", "38", "91", "92", "105"}); ok {
			statusCode = configs.WORKER_FAILED_CODE
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(respProvider.Data.ResponseCode, []string{"203", "205", "107", "93"}); ok {
			statusCode = configs.WORKER_INVALID_PARAM
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(respProvider.Data.ResponseCode, []string{"102"}); ok {
			statusCode = configs.WORKER_CREDENTIAL_ERROR
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(respProvider.Data.ResponseCode, []string{"14", "16", "19", "131", "141", "142", "206", "01", "03", "04", "08", "11", "31", "32", "34", "35", "36", "40", "41", "42", "100", "101", "103"}); ok {
			statusCode = configs.WORKER_VALIDATION_ERROR
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(respProvider.Data.ResponseCode, []string{"404", "12", "204", "17", "110", "202", "207", "121", "117", "10", "94", "108", "109"}); ok {
			statusCode = configs.WORKER_SYSTEM_ERROR
			statusMsg = "FAILED"
		}
	} else {
		var (
			detail  models.DetailBillDescPLN
			details []models.DetailBillDescPLN
		)
		paymentDetail = models.PaymentDetails{
			Price:    float64(respProvider.Data.Nominal),
			AdminFee: float64(respProvider.Data.Admin),
		}
		// tarif, _ := strconv.ParseFloat(respProvider.Data.Desc.Tarif, 64)
		lemTag, _ := strconv.Atoi(respProvider.Data.Desc.LembarTagihan)
		statusCode = configs.WORKER_SUCCESS_CODE
		statusMsg = "SUCCESS"
		if len(respProvider.Data.Desc.Tagihan.Detail) != 0 {
			for _, data := range respProvider.Data.Desc.Tagihan.Detail {
				admin, _ := strconv.ParseFloat(data.Admin, 64)
				denda, _ := strconv.ParseFloat(data.Denda, 64)
				tagihan, _ := strconv.ParseFloat(data.NilaiTagihan, 64)
				detail = models.DetailBillDescPLN{
					Periode:    data.Periode,
					Admin:      admin,
					Denda:      denda,
					Tagihan:    tagihan,
					MeterAwal:  data.MeterAwal,
					MeterAkhir: data.MeterAkhir,
				}
				details = append(details, detail)
			}
		}
		billdesc := models.BillDescPLN{
			CustomerId:    strconv.Itoa(respProvider.Data.TrID),
			Tarif:         respProvider.Data.Desc.Tarif,
			Daya:          strconv.Itoa(respProvider.Data.Desc.Daya),
			LembarTagihan: lemTag,
			Detail:        details,
		}
		// byte, _ := json.Marshal(billdesc)
		respWorker.BillInfo = map[string]interface{}{
			"billDesc": billdesc,
		}
		// respWorker.BillInfo = string(byte)
	}
	respWorker.PaymentDetail = paymentDetail
	respWorker.PaymentStatus = statusCode
	respWorker.PaymentStatusDesc = statusMsg
	respWorker.PaymentStatusDetail = statusCodeDetail
	respWorker.PaymentStatusDescDetail = statusMsgDetail
	respWorker.TotalTrxAmount, _ = strconv.ParseFloat(strconv.Itoa(respProvider.Data.Price), 64)
	respWorker.TrxReferenceNumber = providerRequest.TrID
	respWorker.TrxProviderReferenceNumber = strconv.Itoa(respProvider.Data.TrID)

	return respWorker, nil
}

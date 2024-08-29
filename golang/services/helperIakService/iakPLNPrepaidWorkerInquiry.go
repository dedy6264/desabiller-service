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

func IakPLNPostpaidWorkerInquiry(req models.ReqInqIak) (respWorker models.ResponseWorkerInquiry, err error) {

	var (
		helperName       = "[IAK][WKR]IakPLNPostpaidWorkerInquiry"
		respProvider     models.RespInquiryPLNPostpaidIak
		statusCode       string
		statusMsg        string
		statusCodeDetail string
		statusMsgDetail  string
		inquiryDetail    models.InquiryDetail
	)
	providerRequest := models.ReqInquiryPostpaidIak{
		Commands: "inq-pasca",
		Hp:       req.CustomerId,
		Code:     req.ProductCode,
		RefId:    req.RefId,
		Username: configs.IakUsername,
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
	statusCodeDetail = respProvider.Data.ResponseCode
	statusMsgDetail = respProvider.Data.Message
	if ok, _ := helpers.InArray(respProvider.Data.ResponseCode, []string{"201", "39"}); ok {
		statusCode = configs.WORKER_PENDING_CODE
		statusMsg = "PENDING"
	}
	if respProvider.Data.ResponseCode != "00" {
		if ok, _ := helpers.InArray(respProvider.Data.ResponseCode, []string{"06", "07", "13", "18", "20", "21", "132", "106"}); ok {
			statusCode = configs.WORKER_FAILED_CODE
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(respProvider.Data.ResponseCode, []string{"203", "205", "107"}); ok {
			statusCode = configs.WORKER_INVALID_PARAM
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(respProvider.Data.ResponseCode, []string{"102"}); ok {
			statusCode = configs.WORKER_CREDENTIAL_ERROR
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(respProvider.Data.ResponseCode, []string{"14", "16", "19", "131", "141", "142", "206"}); ok {
			statusCode = configs.WORKER_VALIDATION_ERROR
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(respProvider.Data.ResponseCode, []string{"404", "12", "204", "17", "110", "202", "207", "121", "117", "10"}); ok {
			statusCode = configs.WORKER_SYSTEM_ERROR
			statusMsg = "FAILED"
		}
	} else {
		var (
			detail  models.DetailBillDescPLN
			details []models.DetailBillDescPLN
		)
		inquiryDetail = models.InquiryDetail{
			Price:    float64(respProvider.Data.Nominal),
			AdminFee: float64(respProvider.Data.Admin),
		}
		tarif, _ := strconv.ParseFloat(respProvider.Data.Desc.Tarif, 64)
		lemTag, _ := strconv.Atoi(respProvider.Data.Desc.LembarTagihan)
		statusCode = configs.WORKER_SUCCESS_CODE
		statusMsg = "SUCCESS"
		if len(respProvider.Data.Desc.Tagihan.Detail) != 0 {
			for _, data := range respProvider.Data.Desc.Tagihan.Detail {
				admin, _ := strconv.ParseFloat(data.Admin, 64)
				denda, _ := strconv.ParseFloat(data.Denda, 64)
				tagihan, _ := strconv.ParseFloat(data.NilaiTagihan, 64)
				detail = models.DetailBillDescPLN{
					Periode: data.Periode,
					Admin:   admin,
					Denda:   denda,
					Tagihan: tagihan,
				}
				details = append(details, detail)
			}
		}
		billdesc := models.BillDescPLN{
			CustomerId:    strconv.Itoa(respProvider.Data.TrID),
			Tarif:         tarif,
			Daya:          strconv.Itoa(respProvider.Data.Desc.Daya),
			LembarTagihan: lemTag,
			Detail:        details,
		}
		byte, _ := json.Marshal(billdesc)
		respWorker.BillInfo = map[string]interface{}{
			"billDesc": string(byte),
		}
		// respWorker.BillInfo = string(byte)
	}
	respWorker.InquiryDetail = inquiryDetail
	respWorker.InquiryStatus = statusCode
	respWorker.InquiryStatusDesc = statusMsg
	respWorker.InquiryStatusDetail = statusCodeDetail
	respWorker.InquiryStatusDescDetail = statusMsgDetail
	respWorker.TotalTrxAmount, _ = strconv.ParseFloat(strconv.Itoa(respProvider.Data.Price), 64)
	respWorker.TrxReferenceNumber = providerRequest.RefId
	respWorker.TrxProviderReferenceNumber = strconv.Itoa(respProvider.Data.TrID)

	return respWorker, nil
}

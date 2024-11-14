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

func IakPLNPostpaidWorkerPayment(req models.ReqInqIak) (respWorker models.ResponseWorkerPayment, err error) {

	var (
		helperName       = "[IAK][WKR]IakPLNPostpaidWorkerPayment"
		respProvider     models.RespPaymentPLNPostpaidIak
		statusCode       string
		statusMsg        string
		statusCodeDetail string
		statusMsgDetail  string
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
	statusCode, statusMsg = helpers.IakPayResponseConverter(respProvider.Data.ResponseCode)
	if statusCode == configs.WORKER_SUCCESS_CODE || statusCode == configs.WORKER_PENDING_CODE {
		var (
			detail  models.DetailBillDescPLN
			details []models.DetailBillDescPLN
		)
		// paymentDetail = models.PaymentDetails{
		// 	Price:    float64(respProvider.Data.Nominal),
		// 	AdminFee: float64(respProvider.Data.Admin),
		// }
		// tarif, _ := strconv.ParseFloat(respProvider.Data.Desc.Tarif, 64)
		lemTag, _ := strconv.Atoi(respProvider.Data.Desc.LembarTagihan)
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
	// respWorker.PaymentDetail = paymentDetail
	respWorker.PaymentStatus = statusCode
	respWorker.PaymentStatusDesc = statusMsg
	respWorker.PaymentStatusDetail = statusCodeDetail
	respWorker.PaymentStatusDescDetail = statusMsgDetail
	respWorker.TotalTrxAmount, _ = strconv.ParseFloat(strconv.Itoa(respProvider.Data.Price), 64)
	respWorker.TrxReferenceNumber = providerRequest.TrID
	respWorker.TrxProviderReferenceNumber = strconv.Itoa(respProvider.Data.TrID)

	return respWorker, nil
}

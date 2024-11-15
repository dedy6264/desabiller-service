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

func IakPostpaidWorkerCheckStatus(req models.ReqInqIak) (respWorker models.ResponseWorkerPayment, err error) {

	var (
		helperName       = "[IAK][WKR]IakPostpaidWorkerCheckStatus"
		respProvider     models.RespCheckStatusPostpaidIak
		statusCode       string
		statusMsg        string
		statusCodeDetail string
		statusMsgDetail  string
		// paymentDetail    models.PaymentDetails
		respUndefined  models.RespWorkerUndefined
		respUndefinedI models.RespWorkerUndefinedI
	)
	providerRequest := models.ReqCheckStatusPostpaidIak{
		Commands: req.Commands,
		Username: configs.IakUsername,
		RefId:    req.RefId,
		Sign:     helpers.SignIakEncrypt("cs"),
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
	if respProvider.Data.RefID == "" { // pasti response error
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
	if respProvider.Data.ResponseCode == "00" {
		statusCode = configs.WORKER_SUCCESS_CODE
		statusMsg = "SUCCESS"
	} else if respProvider.Data.ResponseCode == "07" {
		statusCode = configs.WORKER_FAILED_CODE
		statusMsg = "FAILED"
	} else {
		statusCode = configs.WORKER_PENDING_CODE
		statusMsg = "PENDING"
	}
	fmt.Println(statusCode, statusMsg)
	if statusCode == configs.WORKER_SUCCESS_CODE || statusCode == configs.WORKER_PENDING_CODE {
		if req.ProductClan == "" {
			respWorker.PaymentStatus = configs.WORKER_PENDING_CODE
			respWorker.PaymentStatusDesc = "PENDING"
			respWorker.PaymentStatusDescDetail = "PENDING"
			log.Println("Err ", helperName, "Invalid Product Clan", err)
			return
		}
		switch req.ProductClan {
		case "BPJSKS":
			fmt.Println("udah bener")
			if respProvider.Data.Desc == "" {
				var respProvider models.RespPaymentBPJSIak
				err := json.Unmarshal(respByte, &respProvider)
				if err != nil {
					log.Println("Err ", helperName, err)
					return respWorker, err
				}
				respWorker.BillInfo = map[string]interface{}{
					"sn": respProvider.Data.Noref,
				}
				respWorker.TotalTrxAmount, _ = strconv.ParseFloat(strconv.Itoa(respProvider.Data.Price), 64)
				respWorker.TrxProviderReferenceNumber = strconv.Itoa(respProvider.Data.TrID)
			}
		case "PLN POST":
			if respProvider.Data.Desc == "" {
				var respProvider models.RespPaymentPLNPostpaidIak
				err := json.Unmarshal(respByte, &respProvider)
				if err != nil {
					log.Println("Err ", helperName, err)
					return respWorker, err
				}
				var (
					detail  models.DetailBillDescPLN
					details []models.DetailBillDescPLN
				)
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
							Tarif:      respProvider.Data.Desc.Tarif,
							Daya:       strconv.Itoa(respProvider.Data.Desc.Daya),
						}
						details = append(details, detail)
					}
				}
				billdesc := models.BillDescPLN{
					SubscriberNumber: strconv.Itoa(respProvider.Data.TrID),
					SubscriberName:   respProvider.Data.TrName,
					LembarTagihan:    lemTag,
					Detail:           details,
				}
				// byte, _ := json.Marshal(billdesc)
				respWorker.BillInfo = map[string]interface{}{
					"billDesc": billdesc,
				}
				respWorker.TotalTrxAmount, _ = strconv.ParseFloat(strconv.Itoa(respProvider.Data.Price), 64)
				respWorker.TrxProviderReferenceNumber = strconv.Itoa(respProvider.Data.TrID)
			}
		default:
			fmt.Println("nyasar kene")
			respWorker.PaymentStatus = configs.WORKER_PENDING_CODE
			respWorker.PaymentStatusDesc = "PENDING"
			return respWorker, nil
		}
	}
	// respWorker.PaymentDetail = paymentDetail
	respWorker.PaymentStatus = statusCode
	respWorker.PaymentStatusDesc = statusMsg
	respWorker.PaymentStatusDetail = statusCodeDetail
	respWorker.PaymentStatusDescDetail = statusMsgDetail

	return respWorker, nil
}

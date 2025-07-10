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

func IakBPJSWorkerInquiry(req models.ReqInqIak) (respWorker models.ResponseWorkerInquiry, err error) {

	var (
		helperName   = "[IAK][WKR]IakBPJSWorkerInquiry"
		respProvider models.RespInquiryBPJSIak
		statusCode,
		statusMsg, statusDesc,
		statusCodeDetail,
		statusMsgDetail string
		respUndefined  models.RespWorkerUndefined
		respUndefinedI models.RespWorkerUndefinedI
	)
	// if req.Month == "" || req.Month == "0" {
	// 	err = errors.New("BPJSKSValidate :: Month/Period can be null")
	// 	log.Println("Err ", helperName, err)
	// 	return respWorker, err
	// }
	providerRequest := models.ReqInquiryPostpaidIak{
		Commands: "inq-pasca",
		Hp:       req.CustomerId,
		Code:     req.ProductCode,
		RefId:    req.RefId,
		Username: configs.IakUsername,
		Sign:     helpers.SignIakEncrypt(req.RefId),
		Month:    req.Month,
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
	if statusCode == configs.WORKER_SUCCESS_CODE || statusCode == configs.WORKER_PENDING_CODE {
		var (
			detail  models.DetailBillDescBPJS
			details []models.DetailBillDescBPJS
		)
		admin, _ := strconv.ParseFloat(strconv.Itoa(respProvider.Data.Admin), 64)
		tagihan, _ := strconv.ParseFloat(strconv.Itoa(respProvider.Data.Nominal), 64)
		detail = models.DetailBillDescBPJS{
			Periode:    respProvider.Data.Period,
			Admin:      admin,
			Denda:      0,
			Tagihan:    tagihan, //tagihan/productPrice
			JmlPeserta: respProvider.Data.Desc.JumlahPeserta,
		}
		details = append(details, detail)
		billdesc := models.BillDescBPJS{
			CustomerId:   respProvider.Data.Hp,
			CustomerName: respProvider.Data.TrName,
			Detail:       details,
		}
		// byte, _ := json.Marshal(billdesc)
		respWorker.BillInfo = map[string]interface{}{
			"billDesc": billdesc,
		}
		respWorker.TotalTrxAmount, _ = strconv.ParseFloat(strconv.Itoa(respProvider.Data.Price), 64)   //totaltrx
		respWorker.AdminFee, _ = strconv.ParseFloat(strconv.Itoa(respProvider.Data.Admin), 64)         //adminFee
		respWorker.TrxAmount, _ = strconv.ParseFloat(strconv.Itoa(respProvider.Data.SellingPrice), 64) //productProviderPrice
	}
	respWorker.SubscriberNumber = respProvider.Data.Hp
	respWorker.SubscriberName = respProvider.Data.TrName
	respWorker.InquiryStatus = statusCode
	respWorker.InquiryStatusMsg = statusMsg
	respWorker.InquiryStatusDesc = statusDesc
	respWorker.InquiryStatusDetail = statusCodeDetail
	respWorker.InquiryStatusDescDetail = statusMsgDetail

	respWorker.TrxReferenceNumber = providerRequest.RefId
	respWorker.TrxProviderReferenceNumber = strconv.Itoa(respProvider.Data.TrID)

	return respWorker, nil
}

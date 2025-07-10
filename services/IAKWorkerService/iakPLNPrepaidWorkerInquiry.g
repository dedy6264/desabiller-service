package iakworkerservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/utils"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func IakPLNPrepaidWorkerInquiry(req models.ReqInqIak) (respWorker models.ResponseWorkerInquiry, err error) {

	var (
		helperName    = "[IAK][WKR]IakPLNPrepaidWorkerInquiry"
		respProvider  models.RespInquiryPLNPrepaidIak
		respUndefined models.RespWorkerUndefinedII
		statusCode,
		statusMsg, statusDesc,
		statusCodeDetail,
		statusMsgDetail string
	)
	providerRequest := models.ReqInquiryPLNPrepaidIak{
		Commands: "inquiry_pln",
		Hp:       req.CustomerId,
		Username: configs.IakUsername,
		Sign:     helpers.SignIakEncrypt(req.CustomerId),
	}

	respByte, _, err := utils.WorkerPostWithBearer(req.Url, "", providerRequest, "json")
	if err != nil {
		log.Println("Err ", helperName, err)
		return respWorker, err
	}
	//bind struct response
	err = json.Unmarshal(respByte, &respProvider)
	if err != nil {
		err = json.Unmarshal(respByte, &respUndefined)
		if err != nil {
			log.Println("Err ", helperName, err)
			return respWorker, err
		}
		respProvider.Data.Rc = respUndefined.Data.Rc
		respProvider.Data.Message = respUndefined.Data.Message
	}

	statusCodeDetail = respProvider.Data.Rc
	statusMsgDetail = respProvider.Data.Message
	fmt.Println(statusCode, statusMsg, statusDesc)
	statusCode, statusMsg, statusDesc = helpers.IakResponseConverter(respProvider.Data.Rc, respProvider.Data.Message)
	if statusCode == configs.WORKER_SUCCESS_CODE {
		var (
			detail models.DetailBillDescPLN
		)
		segmentPower := strings.Split(strings.ReplaceAll(respProvider.Data.SegmentPower, `\/`, ``), " ")
		if len(segmentPower) > 1 {
			detail = models.DetailBillDescPLN{
				Tarif: segmentPower[0],
				Daya:  strings.ReplaceAll(segmentPower[1], `/`, ``),
			}
		}

		billdesc := models.BillDescPLN{
			MeterNo:          respProvider.Data.MeterNo,
			SubscriberNumber: respProvider.Data.SubscriberId,
			SubscriberName:   respProvider.Data.Name,
			Detail:           append([]models.DetailBillDescPLN{}, detail),
		}
		// byte, _ := json.Marshal(billdesc)
		respWorker.BillInfo = map[string]interface{}{
			"billDesc": billdesc,
		}
	}

	respWorker.InquiryStatus = statusCode
	respWorker.InquiryStatusMsg = statusMsg
	respWorker.InquiryStatusDesc = statusDesc
	respWorker.InquiryStatusDetail = statusCodeDetail
	respWorker.InquiryStatusDescDetail = statusMsgDetail

	return respWorker, nil
}

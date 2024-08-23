package helperservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/utils"
	"encoding/json"
)

func IakHelperServicePayment(providerRequest models.ReqPaymentIak) (respByte []byte, err error) {
	var (
		helperName       = "[IAK]IakHelperServicePayment"
		respProvider     models.RespPaymentIak
		statusCode       string
		statusMsg        string
		statusCodeDetail string
		statusMsgDetail  string
	)
	sign := helpers.SignIakEncrypt("")
	providerRequest.Sign = sign
	respByte, _, err = utils.WorkerPostWithBearer(configs.IakDevUrl, "", providerRequest, "json")
	if err != nil {
		return //
	}
	//bind struct response
	err = json.Unmarshal(respByte, &respProvider)
	if err != nil {
		return //
	}
	statusCodeDetail = respProvider.Data.Rc
	statusMsgDetail = respProvider.Data.Message
	if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"201", "39"}); ok {
		statusCode = configs.WORKER_PENDING_CODE
		statusMsg = "PENDING"
	}
	if respProvider.Data.Rc != "00" {
		if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"06", "07", "13"}); ok {
			statusCode = configs.WORKER_FAILED_CODE
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"14", "16", "18", "19", "20", "21", "131", "132", "141", "142", "203", "206"}); ok {
			statusCode = configs.WORKER_FAILED_CODE
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(respProvider.Data.Rc, []string{"102", "12", "17", "110", "106", "107", "204", "205", "202", "207", "121", "117", "10"}); ok {
			statusCode = configs.WORKER_FAILED_CODE
			statusMsg = "FAILED"
		}
	} else {
		statusCode = configs.WORKER_SUCCESS_CODE
		statusMsg = "SUCCESS"
	}
	return
}

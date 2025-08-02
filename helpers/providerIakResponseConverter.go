package helpers

import "desabiller/configs"

func ResponseConverter(respCode, respDesc string, inq bool) (statusCode, statusMsg, statusDesc string) {
	statusDesc = respDesc //desc adalah message detail atau message dari provider
	switch respCode {
	case "04": //inq sukses
		statusCode = configs.RC_INQUIRY_SUCCESS[0]
		statusMsg = configs.RC_INQUIRY_SUCCESS[1]
	case "00": //pay sukses
		statusCode = configs.RC_SUCCESS[0]
		statusMsg = configs.RC_SUCCESS[1]
	default:
		if inq {
			if ok, _ := InArray(respCode, []string{"18"}); ok {
				statusCode = configs.RC_INQUIRY_ALREADY_PAID[0]
				statusMsg = configs.RC_INQUIRY_ALREADY_PAID[1]
			} else {
				statusCode = configs.RC_INQUIRY_FAILED[0]
				statusMsg = configs.RC_INQUIRY_FAILED[1]
				statusDesc = statusMsg
			}
		} else {
			if ok, _ := InArray(respCode, []string{"03", "11", "05", "39", "08", "09", "10", "401", "12", "13", "14", "15", "16", "19", "20", "21", "22", "24", "25", "26", "27", "28", "29", "30", "36", "37", "38", "40", "96"}); ok {
				statusCode = configs.RC_FAILED[0]
				statusMsg = "PAYMENT " + configs.RC_FAILED[1]
			} else if ok, _ := InArray(respCode, []string{"02", "06", "07"}); ok {
				statusCode = configs.RC_PENDING[0]
				statusMsg = configs.RC_PENDING[1]
			} else if ok, _ := InArray(respCode, []string{"17"}); ok {
				statusCode = configs.RC_ALREADY_PAID[0]
				statusMsg = configs.RC_ALREADY_PAID[1]
			} else {
				statusCode = configs.RC_PENDING[0]
				statusMsg = configs.RC_PENDING[1]
				statusDesc = statusMsg
			}
		}
		// if ok, _ := InArray(respCode, []string{"07", "08", "09", "91", "92", "94", "102", "103", "105", "109", "110", "117", "02", "37", "38", "04", "12", "13", "17", "110", "121", "131", "132", "202", "203", "205", "206", "207"}); ok { //failed/Biller Disruption
		// 	statusCode = configs.BILLER_DISRUPTION
		// 	statusDesc = configs.BILLER_DISRUPTION_MSG
		// 	statusMsg = configs.FAILED_MSG
		// } else if ok, _ := InArray(respCode, []string{"01", "34", "11", "19", "20", "107", "143"}); ok { //failed/paid bill
		// 	statusCode = configs.WORKER_FAILED_CODE
		// 	statusDesc = configs.BIIL_PAID_MSG
		// 	statusMsg = configs.FAILED_MSG
		// } else if ok, _ := InArray(respCode, []string{"204"}); ok { //failed/paid bill
		// 	statusCode = configs.BILLER_DISRUPTION
		// 	statusDesc = configs.BILLER_DISRUPTION_MSG
		// 	statusMsg = statusDesc
		// } else if ok, _ := InArray(respCode, []string{"201"}); ok { //wrong auth
		// 	statusCode = configs.WORKER_UNDEFINED_ERROR
		// 	statusDesc = configs.UNDEFINED_MSG
		// 	statusMsg = configs.FAILED_MSG
		// } else if ok, _ := InArray(respCode, []string{"10", "15", "18", "30", "33", "32", "16", "108", "14", "31", "35", "36", "40", "41", "76", "77", "06", "21", "106", "141", "142"}); ok { //failed/BILL UNREADY
		// 	statusCode = configs.WORKER_BILL_NOTFOUND_CODE
		// 	statusDesc = configs.BILL_NOTFOUND_MSG
		// 	statusMsg = configs.FAILED_MSG
		// } else
	}
	return
}

package helpers

import "desabiller/configs"

func IakResponseConverter(respCode, respDesc string) (statusCode, statusMsg, statusDesc string) {
	statusDesc = respDesc //desc adalah message detail atau message dari provider
	if respCode != "00" {
		if ok, _ := InArray(respCode, []string{"07", "08", "09", "91", "92", "94", "102", "103", "105", "109", "110", "117", "02", "37", "38", "04", "12", "13", "17", "110", "121", "131", "132", "202", "203", "204", "205", "206", "207"}); ok { //failed/Biller Disruption
			statusCode = configs.BILLER_DISRUPTION
			statusDesc = configs.BILLER_DISRUPTION_MSG
			statusMsg = configs.FAILED_MSG
		} else if ok, _ := InArray(respCode, []string{"01", "34", "11", "19", "20", "107", "143"}); ok { //failed/paid bill
			statusCode = configs.WORKER_FAILED_CODE
			statusDesc = configs.BIIL_PAID_MSG
			statusMsg = configs.FAILED_MSG
		} else if ok, _ := InArray(respCode, []string{"201"}); ok { //failed/Undefined
			statusCode = configs.WORKER_UNDEFINED_ERROR
			statusDesc = configs.UNDEFINED_MSG
			statusMsg = configs.FAILED_MSG
		} else if ok, _ := InArray(respCode, []string{"10", "15", "18", "30", "33", "32", "16", "108", "14", "31", "35", "36", "40", "41", "76", "77", "06", "21", "106", "141", "142"}); ok { //failed/BILL UNREADY
			statusCode = configs.WORKER_BILL_NOTFOUND_CODE
			statusDesc = configs.BILL_NOTFOUND_MSG
			statusMsg = configs.FAILED_MSG
		} else if ok, _ := InArray(respCode, []string{"39"}); ok { //failed/expired
			statusCode = configs.WORKER_PENDING_CODE
			// statusDesc = configs.BIIL_PAID_MSG/
			statusMsg = configs.PENDING_MSG
		} else {
			statusCode = configs.WORKER_PENDING_CODE
			statusMsg = configs.PENDING_MSG
			statusDesc = statusMsg
		}
	} else {
		statusCode = configs.WORKER_SUCCESS_CODE
		statusMsg = configs.SUCCESS_MSG
	}
	return
}
func IakInqResponseConverter(rc string) (statusCode, statusMsg, statusDesc string) {
	if rc != "00" {
		if ok, _ := InArray(rc, []string{"201", "39", "05", "02"}); ok {
			statusCode = configs.WORKER_UNDEFINED_ERROR
			statusDesc = "Biller Disruption"
			statusMsg = configs.FAILED_MSG
		}
		if ok, _ := InArray(rc, []string{"06", "07", "13", "18", "20", "21", "132", "106", "09", "30", "33", "37", "38", "91", "92", "105"}); ok {
			statusCode = configs.WORKER_FAILED_CODE
			statusDesc = "Biller Disruption"
			statusMsg = configs.FAILED_MSG
		}
		if ok, _ := InArray(rc, []string{"203", "205", "107", "93"}); ok {
			statusCode = configs.WORKER_INVALID_PARAM
			statusDesc = "Invalid Param"
			statusMsg = configs.FAILED_MSG
		}
		if ok, _ := InArray(rc, []string{"102"}); ok {
			statusCode = configs.WORKER_CREDENTIAL_ERROR
			statusDesc = "Credential Error"
			statusMsg = configs.FAILED_MSG
		}
		if ok, _ := InArray(rc, []string{"14", "16", "19", "131", "141", "142", "206", "01", "03", "04", "08", "11", "31", "32", "34", "35", "36", "40", "41", "42", "100", "101", "103"}); ok {
			statusCode = configs.WORKER_VALIDATION_ERROR
			statusDesc = "Invalid Param"
			statusMsg = configs.FAILED_MSG
		}
		if ok, _ := InArray(rc, []string{"404", "12", "204", "17", "110", "202", "207", "121", "117", "10", "94", "108", "109"}); ok {
			statusCode = configs.WORKER_SYSTEM_ERROR
			statusDesc = "Biller Disruption"
			statusMsg = configs.FAILED_MSG
		}
	} else {
		statusCode = configs.WORKER_SUCCESS_CODE
		statusDesc = configs.SUCCESS_MSG
		statusMsg = configs.SUCCESS_MSG
	}
	return
}
func IakPayResponseConverter(rc string) (statusCode, statusMsg, statusDesc string) {
	if ok, _ := InArray(rc, []string{"201", "39", "05", "02"}); ok {
		statusCode = configs.WORKER_PENDING_CODE
		statusDesc = configs.PENDING_MSG
		statusMsg = configs.PENDING_MSG
	}
	if rc != "00" {
		if ok, _ := InArray(rc, []string{"06", "07", "13", "18", "20", "21", "132", "106", "09", "30", "33", "37", "38", "91", "92", "105"}); ok {
			statusCode = configs.WORKER_FAILED_CODE
			statusDesc = "Biller Disruption"
			statusMsg = configs.FAILED_MSG
		}
		if ok, _ := InArray(rc, []string{"203", "205", "107", "93"}); ok {
			statusCode = configs.WORKER_INVALID_PARAM
			statusDesc = "Invalid Param"
			statusMsg = configs.FAILED_MSG
		}
		if ok, _ := InArray(rc, []string{"102"}); ok {
			statusCode = configs.WORKER_CREDENTIAL_ERROR
			statusDesc = "Credential Error"
			statusMsg = configs.FAILED_MSG
		}
		if ok, _ := InArray(rc, []string{"14", "16", "19", "131", "141", "142", "206", "01", "03", "04", "08", "11", "31", "32", "34", "35", "36", "40", "41", "42", "100", "101", "103"}); ok {
			statusCode = configs.WORKER_VALIDATION_ERROR
			statusDesc = "Invalid Param"
			statusMsg = configs.FAILED_MSG
		}
		if ok, _ := InArray(rc, []string{"404", "12", "204", "17", "110", "202", "207", "121", "117", "10", "94", "108", "109"}); ok {
			statusCode = configs.WORKER_SYSTEM_ERROR
			statusDesc = "Biller Disruption"
			statusMsg = configs.FAILED_MSG
		}
	} else {
		statusCode = configs.WORKER_SUCCESS_CODE
		statusDesc = configs.SUCCESS_MSG
		statusMsg = configs.SUCCESS_MSG
	}
	return
}

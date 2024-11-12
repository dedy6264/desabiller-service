package helpers

import "desabiller/configs"

func IakInqResponseConverter(rc string) (statusCode, statusMsg string) {
	// if ok, _ := InArray(rc, []string{"201", "39", "05", "02"}); ok {
	// 	statusCode = configs.WORKER_PENDING_CODE
	// 	statusMsg = "PENDING"
	// }
	if rc != "00" {
		if ok, _ := InArray(rc, []string{"06", "07", "13", "18", "20", "21", "132", "106", "09", "30", "33", "37", "38", "91", "92", "105"}); ok {
			statusCode = configs.WORKER_FAILED_CODE
			statusMsg = "FAILED"
		}
		if ok, _ := InArray(rc, []string{"203", "205", "107", "93"}); ok {
			statusCode = configs.WORKER_INVALID_PARAM
			statusMsg = "FAILED"
		}
		if ok, _ := InArray(rc, []string{"102"}); ok {
			statusCode = configs.WORKER_CREDENTIAL_ERROR
			statusMsg = "FAILED"
		}
		if ok, _ := InArray(rc, []string{"14", "16", "19", "131", "141", "142", "206", "01", "03", "04", "08", "11", "31", "32", "34", "35", "36", "40", "41", "42", "100", "101", "103"}); ok {
			statusCode = configs.WORKER_VALIDATION_ERROR
			statusMsg = "FAILED"
		}
		if ok, _ := InArray(rc, []string{"404", "12", "204", "17", "110", "202", "207", "121", "117", "10", "94", "108", "109"}); ok {
			statusCode = configs.WORKER_SYSTEM_ERROR
			statusMsg = "FAILED"
		}
	} else {
		statusCode = configs.WORKER_SUCCESS_CODE
		statusMsg = "SUCCESS"
	}
	return
}
func IakPayResponseConverter(rc string) (statusCode, statusMsg string) {
	if ok, _ := InArray(rc, []string{"201", "39", "05", "02"}); ok {
		statusCode = configs.WORKER_PENDING_CODE
		statusMsg = "PENDING"
	}
	if rc != "00" {
		if ok, _ := InArray(rc, []string{"06", "07", "13", "18", "20", "21", "132", "106", "09", "30", "33", "37", "38", "91", "92", "105"}); ok {
			statusCode = configs.WORKER_FAILED_CODE
			statusMsg = "FAILED"
		}
		if ok, _ := InArray(rc, []string{"203", "205", "107", "93"}); ok {
			statusCode = configs.WORKER_INVALID_PARAM
			statusMsg = "FAILED"
		}
		if ok, _ := InArray(rc, []string{"102"}); ok {
			statusCode = configs.WORKER_CREDENTIAL_ERROR
			statusMsg = "FAILED"
		}
		if ok, _ := InArray(rc, []string{"14", "16", "19", "131", "141", "142", "206", "01", "03", "04", "08", "11", "31", "32", "34", "35", "36", "40", "41", "42", "100", "101", "103"}); ok {
			statusCode = configs.WORKER_VALIDATION_ERROR
			statusMsg = "FAILED"
		}
		if ok, _ := InArray(rc, []string{"404", "12", "204", "17", "110", "202", "207", "121", "117", "10", "94", "108", "109"}); ok {
			statusCode = configs.WORKER_SYSTEM_ERROR
			statusMsg = "FAILED"
		}
	} else {
		statusCode = configs.WORKER_SUCCESS_CODE
		statusMsg = "SUCCESS"
	}
	return
}

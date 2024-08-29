package helpers

import (
	"desabiller/configs"
)

func ErrorCodeGateway(errCode string, state string) string {
	var (
		success string
		pending string
		failed  string
	)
	if state == "INQ" {
		success = configs.INQUIRY_SUCCESS_CODE
		pending = configs.INQUIRY_PENDING_CODE
		failed = configs.INQUIRY_FAILED_CODE
	}
	if state == "PAY" {
		success = configs.SUCCESS_CODE
		pending = configs.PENDING_CODE
		failed = configs.FAILED_CODE
	}

	if errCode != configs.WORKER_SUCCESS_CODE {
		if errCode == configs.WORKER_PENDING_CODE {
			return pending
		} else if ok, _ := InArray(errCode, []string{configs.WORKER_FAILED_CODE, configs.WORKER_INVALID_PARAM, configs.WORKER_CREDENTIAL_ERROR, configs.WORKER_VALIDATION_ERROR, configs.WORKER_SYSTEM_ERROR}); ok {
			return failed
		}
	}
	return success
}

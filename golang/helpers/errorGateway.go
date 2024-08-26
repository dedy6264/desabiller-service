package helpers

import (
	"desabiller/configs"
)

func ErrorCodeGateway(errCode string) string {
	if errCode != configs.WORKER_SUCCESS_CODE {
		if errCode == configs.WORKER_PENDING_CODE {
			return configs.PENDING_CODE
		} else if ok, _ := InArray(errCode, []string{configs.WORKER_FAILED_CODE, configs.WORKER_INVALID_PARAM, configs.WORKER_CREDENTIAL_ERROR, configs.WORKER_VALIDATION_ERROR, configs.WORKER_SYSTEM_ERROR}); ok {
			return configs.FAILED_CODE
		}
	}
	return configs.SUCCESS_CODE
}

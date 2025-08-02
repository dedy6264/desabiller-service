package hierarchyservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/utils"
	"net/http"

	"github.com/labstack/echo"
)

func (svc HierarcyService) GetUser(ctx echo.Context) error {
	var (
		svcName = "GetUser"
		respSvc models.ResponseList
	)
	data := helpers.TokenJWTDecode(ctx)
	resUserApp, err := svc.service.RepoHierarchy.GetUserApp(models.ReqGetUserApp{
		Filter: models.UserApp{
			ID: int64(data.UserAppId),
		},
	})
	if err != nil {
		utils.Log("GetUserApp", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	resUserApp.Password = ""
	respSvc.Data = resUserApp
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1], respSvc)
	return ctx.JSON(http.StatusOK, result)
}

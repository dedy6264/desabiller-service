package trxservice

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (svc trxService) HistoryTrxBillerReports(ctx echo.Context) error {
	var (
		svcName = "HistoryTrxBillerReports"
		respSvc models.ResponseList
		data    models.DataToken
	)
	req := new(models.ReqGetTransaction)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if ctx.Get("user") != nil {
		data = helpers.TokenJWTDecode(ctx)
	}
	// respUserApp, err := svc.services.RepoHierarchy.GetUserApp(models.ReqGetUserApp{
	// 	Filter: models.UserApp{
	// 		ID: int64(data.UserAppId),
	// 	},
	// })
	// if err != nil {
	// 	utils.Log("GetUserApp", svcName, err)
	// 	result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], "Data :: empty", nil)
	// 	return ctx.JSON(http.StatusOK, result)
	// }
	req.Filter.UserAppID = int64(data.UserAppId)
	count, err := svc.services.RepoTrx.GetTrxCount(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetTrxCount", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.FAILED_CODE, "Failed", "transaction not found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respInqTrx, err := svc.services.RepoTrx.GetPaymentTrxs(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetPaymentTrxs", err)
		if err == sql.ErrNoRows {
			result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], "transaction not found", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.FAILED_CODE, "Failed", "transaction not found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if len(respInqTrx) != 0 {
		respSvc.Data = respInqTrx
	} else {
		respSvc.Data = []models.RespGetTrx{}
	}
	respSvc.RecordsTotal = count
	respSvc.RecordsFiltered = len(respInqTrx)
	respSvc.Draw = 1
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, configs.SUCCESS_MSG, configs.SUCCESS_MSG, respSvc)
	return ctx.JSON(http.StatusOK, result)
}
func (svc trxService) TrxBillerReports(ctx echo.Context) error {
	var (
		svcName = "TrxBillerReports"
		respSvc models.ResponseList
		data    models.DataToken
	)
	req := new(models.ReqGetTransaction)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if ctx.Get("user") != nil {
		data = helpers.TokenJWTDecode(ctx)
		req.Filter.UserAppID = int64(data.UserAppId)
	}
	count, err := svc.services.RepoTrx.GetTrxCount(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetTrxCount", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.FAILED_CODE, "Failed", "transaction not found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respInqTrx, err := svc.services.RepoTrx.GetTrxs(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetTrx", err)
		if err == sql.ErrNoRows {
			result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], "transaction not found", respInqTrx)
			return ctx.JSON(http.StatusOK, result)
		}
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.FAILED_CODE, "Failed", "transaction not found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if len(respInqTrx) != 0 {
		respSvc.Data = respInqTrx
	} else {
		respSvc.Data = []models.RespGetTrx{}
	}
	respSvc.RecordsTotal = count
	respSvc.RecordsFiltered = len(respInqTrx)
	respSvc.Draw = 1
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, configs.SUCCESS_MSG, configs.SUCCESS_MSG, respSvc)
	return ctx.JSON(http.StatusOK, result)
}
func (svc trxService) TrxBillerReport(ctx echo.Context) error {
	var (
		svcName = "TrxBillerReport"
		respSvc models.ResponseList
		data    models.DataToken
	)
	req := new(models.ReqGetTransaction)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if ctx.Get("user") != nil {
		data = helpers.TokenJWTDecode(ctx)
		req.Filter.UserAppID = int64(data.UserAppId)
	}
	respInqTrx, err := svc.services.RepoTrx.GetTrx(*req)
	if err != nil {
		log.Println("Err ", svcName, "GetTrx", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.FAILED_CODE, "Failed", "transaction not found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	// if len(respInqTrx) != 0 {
	respSvc.Data = respInqTrx
	// } else {
	// 	respSvc.Data = []models.RespGetTrx{}
	// }
	respSvc.RecordsTotal = 1
	respSvc.RecordsFiltered = 1
	respSvc.Draw = 1
	// responsePayment := models.RespPayment{
	// 	ReferenceNumber:        respInqTrx.ReferenceNumber,
	// 	CreatedAt:              respInqTrx.CreatedAt,
	// 	SubscriberNumber:       respInqTrx.CustomerId,
	// 	// BillInfo:               billDescInq,
	// 	ProductName:            respInqTrx.ProductName,
	// 	ProductCode:            respInqTrx.ProductCode,
	// 	ProductCategoryId:      respInqTrx.ProductCategoryId,
	// 	ProductCategoryName:    respInqTrx.ProductCategoryName,
	// 	ProductPrice:           respInqTrx.ProductPrice,
	// 	ProductAdminFee:        respInqTrx.ProductAdminFee,
	// 	MerchantOutletName:     respInqTrx.MerchantOutletName,
	// 	MerchantOutletUsername: respInqTrx.MerchantOutletUsername,
	// 	TransactionTotalAmount:         respInqTrx.ProductPrice + respInqTrx.ProductAdminFee,
	// }
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, configs.SUCCESS_MSG, configs.SUCCESS_MSG, respSvc)
	return ctx.JSON(http.StatusOK, result)
}

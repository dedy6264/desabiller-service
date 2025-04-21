package trxservice

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (svc trxService) IAKCallback(ctx echo.Context) error {
	var (
		svcName = "[IAK]IAKCallback"
		statusCode,
		statusMsg,
		statusCodeDetail,
		statusMsgDetail string
		respByte []byte
		// respSvc models.ResponseList
	)
	req := new(models.IakCallback)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.RepoTrx.GetTrx(models.ReqGetTrx{
		ReferenceNumber: req.Data.RefID,
	})
	if err != nil {
		log.Println("Err ", svcName, "GetTrx", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", "transaction not found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	statusCodeDetail = req.Data.Rc
	statusMsgDetail = req.Data.Message
	if ok, _ := helpers.InArray(req.Data.Rc, []string{"201", "39"}); ok {
		statusCode = configs.WORKER_PENDING_CODE
		statusMsg = "PENDING"
	}
	if req.Data.Rc != "00" {
		if ok, _ := helpers.InArray(req.Data.Rc, []string{"06", "07", "13", "18", "20", "21", "132", "106"}); ok {
			statusCode = configs.WORKER_FAILED_CODE
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(req.Data.Rc, []string{"203", "205", "107"}); ok {
			statusCode = configs.WORKER_INVALID_PARAM
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(req.Data.Rc, []string{"102"}); ok {
			statusCode = configs.WORKER_CREDENTIAL_ERROR
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(req.Data.Rc, []string{"14", "16", "19", "131", "141", "142", "206"}); ok {
			statusCode = configs.WORKER_VALIDATION_ERROR
			statusMsg = "FAILED"
		}
		if ok, _ := helpers.InArray(req.Data.Rc, []string{"404", "12", "204", "17", "110", "202", "207", "121", "117", "10"}); ok {
			statusCode = configs.WORKER_SYSTEM_ERROR
			statusMsg = "FAILED"
		}
	} else {
		statusCode = configs.WORKER_SUCCESS_CODE
		statusMsg = "SUCCESS"

		// billdesc := models.BillDescPulsa{
		// 	CustomerId: req.Data.CustomerID,
		// 	Sn:         req.Data.Sn,
		// }
		billInfo := map[string]interface{}{
			"customerId": req.Data.CustomerID,
			"sn":         req.Data.Sn,
		}
		respByte, _ = json.Marshal(billInfo)
	}
	statusCode = helpers.ErrorCodeGateway(statusCode, "PAY")
	updateTrx := models.ReqGetTrx{
		Id:                         resp.Id,
		ProductReferenceId:         resp.ProductReferenceId,
		ProductReferenceCode:       resp.ProductReferenceCode,
		ProductCategoryId:          resp.ProductCategoryId,
		ProductCategoryName:        resp.ProductCategoryName,
		ProductTypeId:              resp.ProductTypeId,
		ProductTypeName:            resp.ProductTypeName,
		ProductId:                  resp.ProductId,
		ProductName:                resp.ProductName,
		ProductCode:                resp.ProductCode,
		ProductPrice:               resp.ProductPrice,
		ProductAdminFee:            resp.ProductAdminFee,
		ProductMerchantFee:         resp.ProductMerchantFee,
		ProductProviderId:          resp.ProductProviderId,
		ProductProviderName:        resp.ProductProviderName,
		ProductProviderCode:        resp.ProductProviderCode,
		ProductProviderPrice:       resp.ProductProviderPrice,
		ProductProviderAdminFee:    resp.ProductProviderAdminFee,
		ProductProviderMerchantFee: resp.ProductProviderMerchantFee,

		StatusCode:              statusCode,
		StatusMessage:           "PAYMENT " + statusMsg,
		StatusDesc:              statusMsg,
		ReferenceNumber:         resp.ReferenceNumber,
		ProviderStatusCode:      statusCodeDetail,
		ProviderStatusMessage:   statusMsgDetail,
		ProviderStatusDesc:      statusMsgDetail,
		ProviderReferenceNumber: resp.ProviderReferenceNumber,

		ClientId:               resp.ClientId,
		ClientName:             resp.ClientName,
		GroupId:                resp.GroupId,
		GroupName:              resp.GroupName,
		MerchantId:             resp.MerchantId,
		MerchantName:           resp.MerchantName,
		MerchantOutletId:       resp.MerchantOutletId,
		MerchantOutletName:     resp.MerchantOutletName,
		MerchantOutletUsername: resp.MerchantOutletUsername,
		CustomerId:             resp.CustomerId,
		OtherMsg:               string(respByte),
		Filter:                 models.FilterReq{},
	}
	//cek trx, jika pending, update
	if resp.StatusCode == configs.PENDING_CODE {
		err = helpers.DBTransaction(svc.services.RepoDB, func(Tx *sql.Tx) error {
			err = svc.services.RepoTrx.UpdateTrx(updateTrx, Tx)
			if err != nil {
				return err
			}
			err = svc.services.RepoTrx.InsertTrxStatus(models.ReqGetTrxStatus{
				ReferenceNumber:         updateTrx.ReferenceNumber,
				ProviderReferenceNumber: updateTrx.ProviderReferenceNumber,
				StatusCode:              updateTrx.StatusCode,
				StatusMessage:           updateTrx.StatusMessage,
			}, Tx)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return ctx.JSON(http.StatusOK, false)
		}
	}
	return ctx.JSON(http.StatusOK, true)
}

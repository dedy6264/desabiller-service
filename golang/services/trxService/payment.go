package trxservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (svc trxService) PaymentBiller(ctx echo.Context) error {
	var (
		svcName = "PaymentBiller"
		// respSvc models.ResponseList
		// respOutlet models.RespGetMerchantOutlet
	)
	//get product
	//check price-provider price
	// check merchant fee-provider merchant fee
	//hit inq to partner
	//record to db
	//assing to response
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMPTRX)
	req := new(models.ReqPaymentTrx)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ReferenceNumber == "" {
		log.Println("Err ", svcName, "ReferenceNumber id cannot be null")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "ReferenceNumber id cannot be null", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ProviderReferenceNumber == "" {
		log.Println("Err ", svcName, "ProviderReferenceNumber id cannot be null")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "ProviderReferenceNumber id cannot be null", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	//payment method validation
	// if err != nil {
	// 	log.Println("Err ", svcName, "GetProduct", err)
	// 	result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
	// 	return ctx.JSON(http.StatusOK, result)
	// }
	respInqTrx, err := svc.services.ApiTrx.GetTrx(models.ReqGetTrx{
		ReferenceNumber: req.ReferenceNumber,
	})
	if err != nil {
		log.Println("Err ", svcName, "GetTrx", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}

	if respInqTrx.StatusCode != configs.INQUIRY_SUCCESS_CODE {
		log.Println("Err ", svcName, "Transaction invalid")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Transaction invalid", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	updatePayment := models.ReqGetTrx{
		ProductClanId:              respInqTrx.ProductClanId,
		ProductClanName:            respInqTrx.ProductClanName,
		ProductCategoryId:          respInqTrx.ProductCategoryId,
		ProductCategoryName:        respInqTrx.ProductCategoryName,
		ProductTypeId:              respInqTrx.ProductTypeId,
		ProductTypeName:            respInqTrx.ProductTypeName,
		ProductId:                  respInqTrx.ProductId,
		ProductName:                respInqTrx.ProductName,
		ProductCode:                respInqTrx.ProductCode,
		ProductPrice:               respInqTrx.ProductPrice,
		ProductAdminFee:            respInqTrx.ProductAdminFee,
		ProductMerchantFee:         respInqTrx.ProductMerchantFee,
		ProductProviderId:          respInqTrx.ProductProviderId,
		ProductProviderName:        respInqTrx.ProductProviderName,
		ProductProviderCode:        respInqTrx.ProductCode,
		ProductProviderPrice:       respInqTrx.ProductProviderPrice,
		ProductProviderAdminFee:    respInqTrx.ProductProviderAdminFee,
		ProductProviderMerchantFee: respInqTrx.ProductProviderMerchantFee,
		StatusCode:                 configs.SUCCESS_CODE,
		StatusMessage:              "PAYMENT " + configs.SUCCESS_MSG,
		StatusDesc:                 "PAYMENT " + configs.SUCCESS_MSG,
		ReferenceNumber:            respInqTrx.ReferenceNumber,
		ProviderStatusCode:         req.ProviderReferenceNumber,
		ProviderStatusMessage:      "PAYMENT " + configs.SUCCESS_MSG,
		ProviderStatusDesc:         "PAYMENT " + configs.SUCCESS_MSG,
		ProviderReferenceNumber:    "-",
		ClientId:                   respInqTrx.ClientId,
		ClientName:                 respInqTrx.ClientName,
		GroupId:                    respInqTrx.GroupId,
		GroupName:                  respInqTrx.GroupName,
		MerchantId:                 respInqTrx.MerchantId,
		MerchantName:               respInqTrx.MerchantName,
		MerchantOutletId:           respInqTrx.MerchantOutletId,
		MerchantOutletName:         respInqTrx.MerchantOutletName,
		MerchantOutletUsername:     respInqTrx.MerchantOutletUsername,
		CustomerId:                 respInqTrx.CustomerId,
		OtherMsg:                   respInqTrx.OtherMsg,
		Filter: models.FilterReq{
			CreatedAt: dbTime,
		},
	}
	// byte, status, er := utils.WorkerPostWithBearer())
	err = svc.services.ApiTrx.UpdateTrx(updatePayment, ``)
	if err != nil {
		log.Println("Err ", svcName, "UpdateTrx", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.ApiTrx.InsertTrxStatus(models.ReqGetTrxStatus{
		ReferenceNumber:         updatePayment.ReferenceNumber,
		ProviderReferenceNumber: updatePayment.ProviderReferenceNumber,
		StatusCode:              updatePayment.StatusCode,
		StatusMessage:           updatePayment.StatusMessage,
	})
	if err != nil {
		log.Println("Err ", svcName, "InsertTrxStatus", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, configs.SUCCESS_MSG, updatePayment)
	return ctx.JSON(http.StatusOK, result)
}

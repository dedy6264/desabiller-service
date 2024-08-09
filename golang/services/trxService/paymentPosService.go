package trxservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
)

func (svc TrxServices) PaymentPos(ctx echo.Context) error {
	var (
		svcName = "PaymentPos"
		t       = time.Now()
		dbTime  = t.Local().Format(configs.LAYOUT_TIMESTAMP)
		// dbTrx         = t.Local().Format(configs.LAYOUT_TIMESTAMPTRX)
		logErr            = "Err :: " + svcName
		tableNm           = "trx_poses"
		nickname          string
		paymentMethodId   int
		paymentMethodName string
	)
	a := ctx.Get("user").(*jwt.Token)
	claim := a.Claims.(jwt.MapClaims)
	snDevice := claim["snDevice"].(string)
	req := new(models.ReqPaymentTrx)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println(logErr+" FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	//cek payment method
	respay, _ := svc.service.ApiPayment.GetListPaymentMethod(models.ReqGetListPaymentMethod{
		Id: req.PaymentMethodId,
	})
	if len(respay) == 0 {
		log.Println(logErr + " Not Found")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Not Found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	for _, datapay := range respay {
		paymentMethodId = datapay.Id
		paymentMethodName = datapay.PaymentMethodName
	}
	//create no trx
	resGet, status := svc.service.ApiTrx.GetTrxPos(models.ReqTrx{
		TrxNumber: req.TrxNumber,
	}, tableNm)

	if !status {
		log.Println(logErr + " Not Found")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Not Found", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if resGet.OutletDeviceSn != snDevice {
		log.Println(logErr + " Device not authorized")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Device not authorized", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if resGet.StatusCode == configs.SUCCESS_CODE {
		log.Println(logErr + " Trx has payment")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Trx has payment", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	status = svc.service.ApiTrx.UpdateTrxPos(models.ReqUpdateTrx{
		Id:                   resGet.Id,
		TrxNumber:            resGet.TrxNumber,
		TrxNumberPartner:     "",
		PaymentAt:            dbTime,
		UpdatedBy:            nickname,
		UpdatedAt:            dbTime,
		StatusCode:           configs.SUCCESS_CODE,
		StatusMessage:        "SUCCESS",
		StatusDesc:           "PAYMENT SUCCESS",
		StatusCodePartner:    resGet.StatusCodePartner,
		StatusMessagePartner: resGet.StatusMessagePartner,
		StatusDescPartner:    resGet.StatusDescPartner,
		BillInfo:             resGet.BillInfo,
		PaymentMethodId:      paymentMethodId,
		PaymentMethodName:    paymentMethodName,
	}, tableNm)
	if !status {
		log.Println(logErr + " Inq Failed")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Inq Failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resPayment := models.ReqInsertTrx{
		StatusCode:           configs.SUCCESS_CODE,
		StatusMessage:        "SUCCESS",
		StatusDesc:           "PAYMENT SUCCESS",
		TrxNumber:            resGet.TrxNumber,
		TrxNumberPartner:     resGet.TrxNumberPartner,
		PaymentAt:            resGet.PaymentAt,
		CreatedBy:            resGet.CreatedBy,
		CreatedAt:            resGet.CreatedAt,
		UpdatedBy:            resGet.UpdatedBy,
		UpdatedAt:            resGet.UpdatedAt,
		StatusCodePartner:    resGet.StatusCodePartner,
		StatusMessagePartner: resGet.StatusMessagePartner,
		StatusDescPartner:    resGet.StatusDescPartner,
		SegmentId:            resGet.SegmentId,
		ProductTypeId:        resGet.ProductTypeId,
		ProductTypeName:      resGet.ProductTypeName,
		ProductCategoryId:    resGet.ProductCategoryId,
		ProductCategoryName:  resGet.ProductCategoryName,
		ProductId:            resGet.ProductId,
		ProductCode:          resGet.ProductCode,
		ProductName:          resGet.ProductName,
		ProductPrice:         resGet.ProductPrice,
		ProductAdminFee:      resGet.ProductAdminFee,
		ProductMerchantFee:   resGet.ProductMerchantFee,
		SubTotal:             resGet.SubTotal,
		GrandTotal:           resGet.GrandTotal,
		CustomerId:           resGet.CustomerId,
		BillInfo:             resGet.BillInfo,
		ClientId:             resGet.ClientId,
		ClientName:           resGet.ClientName,
		MerchantId:           resGet.MerchantId,
		MerchantName:         resGet.MerchantName,
		MerchantOutletId:     resGet.MerchantOutletId,
		MerchantOutletName:   resGet.MerchantOutletName,
		UserOutletId:         resGet.UserOutletId,
		UserOutletName:       resGet.UserOutletName,
		OutletDeviceId:       resGet.OutletDeviceId,
		OutletDeviceType:     resGet.OutletDeviceType,
		OutletDeviceSn:       resGet.OutletDeviceSn,
		PaymentMethodId:      paymentMethodId,
		PaymentMethodName:    paymentMethodName,
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", resPayment)
	return ctx.JSON(http.StatusOK, result)
}

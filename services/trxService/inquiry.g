package trxservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func (svc trxService) InquiryBiller(ctx echo.Context) error {
	var (
		svcName = "InquiryBiller"
		// respSvc models.ResponseList
		respOutlet                                                                                      models.RespGetMerchantOutlet
		referenceNumber, url, statusCode, statusCodeProvider, statusMessageProvider, statusDescProvider string
		billInfo                                                                                        map[string]interface{}
		respWorker                                                                                      models.ResponseWorkerInquiry
		result                                                                                          models.Response

		productPrice, productAdminFee, productMerchantFee                         float64
		productProviderPrice, productProviderAdminFee, productProviderMerchantFee float64
	)
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMPTRX)
	dbTimeTrx := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	req := new(models.ReqInquiry)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		helpers.ErrLogger(svcName, "BindValidate", err)
		result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	{ //valitasi #1
		if req.ProductCode == "" {
			helpers.ErrLogger(svcName, "Product Code cannot be null", nil)
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", "Product Code cannot be null", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		if req.AdditionalField.SubscriberNumber == "" {
			helpers.ErrLogger(svcName, "Customer Id cannot be null", nil)
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", "Customer Id cannot be null", nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}
	// if req.AdditionalField.SubscriberNumber == "" {
	// 	log.Println("Err ", svcName, "customer id cannot be null")
	// 	result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "customer id cannot be null", nil)
	// 	return ctx.JSON(http.StatusOK, result)
	// }
	data := helpers.TokenJWTDecode(ctx)
	respOutlet, err = svc.services.RepoHierarchy.GetMerchantOutlet(models.ReqGetMerchantOutlet{
		ID: data.MerchantOutletId,
	})
	if err != nil {
		helpers.ErrLogger(svcName, "GetMerchantOutlet", err)
		result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", "Merchant Not Found", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	respProduct, err := svc.services.RepoProduct.GetProduct(models.ReqGetProduct{
		// ID: req.ProductId,
		ProductCode: req.ProductCode,
	})
	if err != nil {
		helpers.ErrLogger(svcName, "GetProduct", err)
		result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", "Invalid Product", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if respProduct.ProductCode == "BPJSKS" && req.AdditionalField.Periode == 0 {
		if req.ProductCode == "" {
			helpers.ErrLogger(svcName, "Periode cannot be null", nil)
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", "Periode cannot be null", nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}
	referenceNumber, err = svc.services.RepoTrx.GenerateNo("DB-"+dbTime, "", 7)
	if err != nil {
		helpers.ErrLogger(svcName, "GenerateNo", err)
		result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", "Failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	// statusCode = configs.INQUIRY_SUCCESS_CODE
	// statusDesc = "INQUIRY " + configs.SUCCESS_MSG
	// statusMessage = "INQUIRY " + configs.SUCCESS_MSG
	// statusCodeProvider = "-"
	// statusMessageProvider = "INQUIRY " + configs.SUCCESS_MSG
	// statusDescProvider = "INQUIRY " + configs.SUCCESS_MSG
	fmt.Println("::::::::", respProduct.ProductReferenceId)
	dataInquiry := models.ReqGetTrx{
		ProductReferenceId:         respProduct.ProductReferenceId,
		ProductCategoryId:          respProduct.ProductCategoryId,
		ProductCategoryName:        respProduct.ProductCategoryName,
		ProductTypeId:              respProduct.ProductTypeId,
		ProductTypeName:            respProduct.ProductTypeName,
		ProductId:                  respProduct.ID,
		ProductName:                respProduct.ProductName,
		ProductCode:                respProduct.ProductCode,
		ProductPrice:               respProduct.ProductPrice,
		ProductAdminFee:            respProduct.ProductAdminFee,
		ProductMerchantFee:         respProduct.ProductMerchantFee,
		ProviderId:                 respProduct.ProviderId,
		ProviderName:               respProduct.ProviderName,
		ProductProviderId:          respProduct.ProductProviderId,
		ProductProviderName:        respProduct.ProductProviderName,
		ProductProviderCode:        respProduct.ProductProviderCode,
		ProductProviderPrice:       respProduct.ProductProviderPrice,
		ProductProviderAdminFee:    respProduct.ProductProviderAdminFee,
		ProductProviderMerchantFee: respProduct.ProductProviderMerchantFee,
		StatusCode:                 configs.INQUIRY_SUCCESS_CODE,
		StatusMessage:              "INQUIRY",
		StatusDesc:                 "INQUIRY" + configs.SUCCESS_MSG,
		ReferenceNumber:            referenceNumber,
		ProviderStatusCode:         statusCodeProvider,
		ProviderStatusMessage:      statusMessageProvider,
		ProviderStatusDesc:         statusDescProvider,
		ProviderReferenceNumber:    "-",
		ClientId:                   respOutlet.ClientId,
		ClientName:                 respOutlet.ClientName,
		GroupId:                    respOutlet.GroupId,
		GroupName:                  respOutlet.GroupName,
		MerchantId:                 respOutlet.MerchantId,
		MerchantName:               respOutlet.MerchantName,
		MerchantOutletId:           respOutlet.ID,
		MerchantOutletName:         respOutlet.MerchantOutletName,
		MerchantOutletUsername:     respOutlet.MerchantOutletUsername,
		CustomerId:                 req.AdditionalField.SubscriberNumber,
		TotalTrxAmount:             respProduct.ProductPrice,
		OtherMsg:                   "-",
		Filter: models.FilterReq{
			CreatedAt: dbTimeTrx,
		},
	}
	if respProduct.ProductTypeId == 2 { //PREPAID
		if configs.AppEnv == "DEV" {
			url = configs.IakDevUrlPrepaid + configs.ENDPOINT_IAK_PREPAID
		}
		if configs.AppEnv == "PROD" {
			url = configs.IakProdUrlPrepaid + configs.ENDPOINT_IAK_PREPAID
		}
		if respProduct.ProductPrice < respProduct.ProductProviderPrice {
			helpers.ErrLogger(svcName, "Invalid Product Price : "+strconv.Itoa(int(respProduct.ProductPrice))+" "+strconv.Itoa(int(respProduct.ProductProviderPrice)), err)
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", "Failed ", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		if respProduct.ProductReferenceId == 10 {
			url = configs.IakDevUrlPrepaid + configs.ENDPOINT_IAK_INQ_PLN_PREPAID
			respWorker, err = svc.InqProviderSwitcher(models.ProviderInqRequest{
				ProviderName:         respProduct.ProviderName,
				ProviderId:           respProduct.ProviderId,
				ProductCode:          respProduct.ProductProviderCode,
				SubscriberNumber:     req.AdditionalField.SubscriberNumber,
				SubscriberName:       req.AdditionalField.SubscriberName,
				ReferenceNumber:      referenceNumber,
				Url:                  url,
				Periode:              req.AdditionalField.Periode,
				ProductReferenceCode: respProduct.ProductReferenceCode,
				ProductReferenceId:   respProduct.ProductReferenceId,
			})
			if err != nil {
				helpers.ErrLogger(svcName, "", err)
				// result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.BILLER_DISRUPTION, "Failed", configs.BILLER_DISRUPTION_MSG, nil)
				// return ctx.JSON(http.StatusOK, result)
				dataInquiry.StatusCode = configs.BILLER_DISRUPTION
				dataInquiry.StatusMessage = "INQUIRY " + respWorker.InquiryStatusMsg
				dataInquiry.StatusDesc = respWorker.InquiryStatusDesc
				dataInquiry.ProviderStatusCode = respWorker.InquiryStatusDetail
				dataInquiry.ProviderStatusMessage = respWorker.InquiryStatusDescDetail
				dataInquiry.ProviderStatusDesc = respWorker.InquiryStatusDescDetail
				dataInquiry.ProviderReferenceNumber = respWorker.TrxProviderReferenceNumber
			} else {
				statusCode = helpers.ErrorCodeGateway(respWorker.InquiryStatus, "INQ")
				dataInquiry.StatusCode = statusCode
				dataInquiry.StatusMessage = "INQUIRY " + respWorker.InquiryStatusMsg
				dataInquiry.StatusDesc = respWorker.InquiryStatusDesc
				dataInquiry.ProviderStatusCode = respWorker.InquiryStatusDetail
				dataInquiry.ProviderStatusMessage = respWorker.InquiryStatusDescDetail
				dataInquiry.ProviderStatusDesc = respWorker.InquiryStatusDescDetail
				dataInquiry.ProviderReferenceNumber = respWorker.TrxProviderReferenceNumber
			}
			billInfo = respWorker.BillInfo
			byte, _ := json.Marshal(billInfo)
			dataInquiry.ReferenceNumber = referenceNumber
			dataInquiry.OtherMsg = string(byte)
			// dataInquiry.TotalTrxAmount = respWorker.TotalTrxAmount
			dataInquiry.Filter = models.FilterReq{
				CreatedAt: dbTime,
			}
			dataInquiry.ProductReferenceCode = respProduct.ProductReferenceCode
		}

		err = svc.services.RepoTrx.InsertTrx(dataInquiry, nil)
		if err != nil {
			helpers.ErrLogger(svcName, "InsertTrx", err)
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", "failed", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		err = svc.services.RepoTrx.InsertTrxStatus(models.ReqGetTrxStatus{
			ReferenceNumber:         dataInquiry.ReferenceNumber,
			ProviderReferenceNumber: dataInquiry.ProviderReferenceNumber,
			StatusCode:              dataInquiry.StatusCode,
			StatusMessage:           dataInquiry.StatusMessage,
		}, nil)
		if err != nil {
			helpers.ErrLogger(svcName, "InsertTrxStatus", err)
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", "failed", nil)
			return ctx.JSON(http.StatusOK, result)
		}
	} else { //postpaid
		if configs.AppEnv == "DEV" {
			url = configs.IakDevUrlPostpaid + configs.ENDPOINT_IAK_POSTPAID
		}
		if configs.AppEnv == "PROD" {
			url = configs.IakProdUrlPostpaid + configs.ENDPOINT_IAK_POSTPAID
		}
		if respProduct.ProductMerchantFee > respProduct.ProductProviderMerchantFee {
			helpers.ErrLogger(svcName, "Invalid Product Price : "+strconv.Itoa(int(respProduct.ProductPrice))+" "+strconv.Itoa(int(respProduct.ProductProviderPrice)), err)
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", "failed", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		//inquiry ke partner
		respWorker, err = svc.InqProviderSwitcher(models.ProviderInqRequest{
			ProviderName:         respProduct.ProviderName,
			ProviderId:           respProduct.ProviderId,
			ProductCode:          respProduct.ProductProviderCode,
			SubscriberNumber:     req.AdditionalField.SubscriberNumber,
			SubscriberName:       req.AdditionalField.SubscriberName,
			ReferenceNumber:      referenceNumber,
			Url:                  url,
			Periode:              req.AdditionalField.Periode,
			ProductReferenceCode: respProduct.ProductReferenceCode,
			ProductReferenceId:   respProduct.ProductReferenceId,
		})
		if err != nil {
			helpers.ErrLogger(svcName, "", err)

			// if strings.Contains(err.Error(), "BPJSKSValidate") {
			// 	result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, err.Error(), nil)
			// } else {
			// result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.INQUIRY_FAILED_CODE, "Failed", "Trx failed", nil)
			// // }
			// return ctx.JSON(http.StatusOK, result)
			dataInquiry.StatusCode = configs.BILLER_DISRUPTION
			dataInquiry.StatusMessage = "INQUIRY " + respWorker.InquiryStatusMsg
			dataInquiry.StatusDesc = respWorker.InquiryStatusDesc
			dataInquiry.ProviderStatusCode = respWorker.InquiryStatusDetail
			dataInquiry.ProviderStatusMessage = respWorker.InquiryStatusDescDetail
			dataInquiry.ProviderStatusDesc = respWorker.InquiryStatusDescDetail
			dataInquiry.ProviderReferenceNumber = respWorker.TrxProviderReferenceNumber
		} else {

			productProviderPrice = respWorker.TrxAmount
			productProviderAdminFee = respWorker.AdminFee
			productProviderMerchantFee = respProduct.ProductProviderMerchantFee

			productAdminFee = respWorker.AdminFee
			productMerchantFee = respProduct.ProductMerchantFee
			productPrice = respWorker.TotalTrxAmount - respWorker.AdminFee
			statusCode = helpers.ErrorCodeGateway(respWorker.InquiryStatus, "INQ")
			dataInquiry.StatusCode = statusCode
			dataInquiry.StatusMessage = "INQUIRY " + respWorker.InquiryStatusMsg
			dataInquiry.StatusDesc = respWorker.InquiryStatusDesc
			dataInquiry.ProviderStatusCode = respWorker.InquiryStatusDetail
			dataInquiry.ProviderStatusMessage = respWorker.InquiryStatusDescDetail
			dataInquiry.ProviderStatusDesc = respWorker.InquiryStatusDescDetail
			dataInquiry.ProviderReferenceNumber = respWorker.TrxProviderReferenceNumber
		}

		billInfo = respWorker.BillInfo
		byte, _ := json.Marshal(billInfo)
		dataInquiry.ReferenceNumber = referenceNumber
		dataInquiry.OtherMsg = string(byte)
		dataInquiry.TotalTrxAmount = respWorker.TotalTrxAmount
		dataInquiry.Filter = models.FilterReq{
			CreatedAt: dbTime,
		}
		dataInquiry.ProductPrice = productPrice
		dataInquiry.ProductAdminFee = productAdminFee
		dataInquiry.ProductMerchantFee = productMerchantFee
		dataInquiry.ProductProviderPrice = productProviderPrice
		dataInquiry.ProductProviderAdminFee = productProviderAdminFee
		dataInquiry.ProductProviderMerchantFee = productProviderMerchantFee
		dataInquiry.ProductReferenceCode = respProduct.ProductReferenceCode

		err = svc.services.RepoTrx.InsertTrx(dataInquiry, nil)
		if err != nil {
			helpers.ErrLogger(svcName, "InsertTrx", err)
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", "failed", nil)
			return ctx.JSON(http.StatusOK, result)
		}
		err = svc.services.RepoTrx.InsertTrxStatus(models.ReqGetTrxStatus{
			ReferenceNumber:         dataInquiry.ReferenceNumber,
			ProviderReferenceNumber: dataInquiry.ProviderReferenceNumber,
			StatusCode:              dataInquiry.StatusCode,
			StatusMessage:           dataInquiry.StatusMessage,
		}, nil)
		if err != nil {
			helpers.ErrLogger(svcName, "InsertTrxStatus", err)
			result = helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Failed", "failed", nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}

	// byte, status, er := utils.WorkerPostWithBearer())
	respInquiry := models.RespInquiry{
		StatusMessage:          dataInquiry.ProviderStatusMessage,
		CreatedAt:              dataInquiry.Filter.CreatedAt,
		MerchantOutletName:     dataInquiry.MerchantOutletName,
		MerchantOutletUsername: dataInquiry.MerchantOutletUsername,
		ReferenceNumber:        dataInquiry.ReferenceNumber,
		ProductName:            dataInquiry.ProductName,
		ProductCode:            dataInquiry.ProductCode,
		SubscriberNumber:       dataInquiry.CustomerId,
		ProductPrice:           dataInquiry.ProductPrice,
		ProductAdminFee:        dataInquiry.ProductAdminFee,
		ProductMerchantFee:     dataInquiry.ProductMerchantFee,
		TotalTrxAmount:         dataInquiry.TotalTrxAmount,
		BillInfo:               billInfo,
	}
	result = helpers.ResponseJSON(configs.TRUE_VALUE, dataInquiry.StatusCode, dataInquiry.StatusMessage, dataInquiry.StatusDesc, respInquiry)
	return ctx.JSON(http.StatusOK, result)
}

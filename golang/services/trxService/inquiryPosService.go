package trxservice

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
)

func (svc TrxServices) InquiryPos(ctx echo.Context) error {
	var (
		svcName    = "InquiryPos"
		t          = time.Now()
		dbTime     = t.Local().Format(configs.LAYOUT_TIMESTAMP)
		dbTrx      = t.Local().Format(configs.LAYOUT_TIMESTAMPTRX)
		logErr     = "Err :: " + svcName
		outletNm   string
		merchantNm string
		clientNm   string
		// cID           int
		tableNm       = "trx_poses"
		nickname      string
		productType   = 1
		productTypeId int
		// productTypeName     string
		productCategoryId int
		// productCategoryName string
		productId    int
		productCode  string
		productName  string
		productPrice float64
	)
	a := ctx.Get("user").(*jwt.Token)
	claim := a.Claims.(jwt.MapClaims)
	snDevice := claim["snDevice"].(string)
	uID := claim["userId"].(float64)
	oID := claim["outletId"].(float64)
	mID := claim["merchantId"].(float64)
	cID := claim["clientId"].(float64)

	// claims := ctx.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
	// merchantOutletID := claims["merchantOutletID"].(float64)

	req := new(models.ReqInqTrx)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", logErr+err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	// if req.PaymentMethodId == 0 {
	// 	log.Println("FAILLED BINDING Validate Payment Method", logErr)
	// 	result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING Validate Payment Method", nil)
	// 	return ctx.JSON(http.StatusOK, result)
	// }

	//validasi payment method/ok
	//validasi device/ok
	//validasi user outlet
	//validasi produk
	//get no trx
	//kalkulasi totalan
	resDvc, _ := svc.service.ApiHierarchy.GetListOutletDevice(models.ReqGetListOutletDevice{
		DeviceSn:         snDevice,
		MerchantOutletId: int(oID),
		MerchantId:       int(mID),
	})
	if len(resDvc) != 1 {
		log.Println(logErr + "GetListOutletDevice")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "GetListOutletDevice", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	// if resDvc[0].DeviceSn != snDevice {
	// 	log.Println(logErr + "GetListOutletDevice")
	// 	result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "GetListOutletDevice", nil)
	// 	return ctx.JSON(http.StatusOK, result)
	// }
	resUs, _ := svc.service.ApiHierarchy.GetListUserOutlet(models.ReqGetListUserOutlet{
		ID:               int(uID),
		MerchantOutletId: int(oID),
		MerchantId:       int(mID),
	})
	fmt.Println("RES count::", resUs)

	if len(resUs) != 1 {
		log.Println(logErr + "GetListUserOutletCount")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "GetListUserOutletCount", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	for _, dataUs := range resUs {
		nickname = dataUs.Nickname
	}
	// var idProduct []int
	var subtotal float64
	var reqInsertDetails []models.ReqInsertTrxDetails

	//VALIDASI HARGA N PRODUK
	//create no trx
	// resGet, status := svc.service.ApiTrx.GetTrxList(models.ReqTrx{
	// 	TrxNumber: "",
	// }, tableNm)
	// if len(resGet) > 0 {
	// 	log.Println(logErr + "GetTrxList")
	// 	result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "GetTrxList", nil)
	// 	return ctx.JSON(http.StatusOK, result)
	// }
	// generate no trx		: typeproduct<>ddmmyyyy<>padding(3)no urut
	dataType := "M" + strconv.Itoa(productType) + dbTrx
	code, err := svc.service.ApiNoTrx.GenerateNo(dataType, "", 7)
	if err != nil {
		log.Println(logErr + "Generate Failed")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Generate Failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	for _, data := range req.ProductDetails {
		subtotal = subtotal + (data.ProductPrice * float64(data.Qty))
	}
	helpers.DBTransaction(svc.service.RepoDB, func(tx *sql.Tx) error {
		var (
			errMsg error
		)
		idTrx, status := svc.service.ApiTrx.InsertTrxPos(models.ReqInsertTrx{
			TrxNumber:            dataType + code,
			TrxNumberPartner:     "",
			CreatedBy:            nickname,
			CreatedAt:            dbTime,
			UpdatedBy:            nickname,
			UpdatedAt:            dbTime,
			StatusCode:           "66",
			StatusMessage:        "SUCCESS",
			StatusDesc:           "INQUIRY SUCCESS",
			StatusCodePartner:    "",
			StatusMessagePartner: "",
			StatusDescPartner:    "",
			SegmentId:            0,
			ProductTypeId:        productTypeId,
			// ProductTypeName:      "",
			// ProductCategoryId:    productCategoryId,
			// ProductCategoryName:  "",
			// ProductId:            productId,
			// ProductCode:          productCode,
			// ProductName:          productName,
			// ProductPrice:         productPrice,
			ProductAdminFee:    0,
			ProductMerchantFee: 0,
			// Quantity:             req.Qty,
			SubTotal:   subtotal,
			GrandTotal: subtotal,
			// CustomerId:           req.CustomerId,
			// BillInfo:           "",
			ClientId:           int(cID),
			ClientName:         clientNm,
			MerchantId:         int(mID),
			MerchantName:       merchantNm,
			MerchantOutletId:   int(oID),
			MerchantOutletName: merchantNm,
			UserOutletId:       int(uID),
			UserOutletName:     outletNm,
			OutletDeviceId:     int(oID),
			OutletDeviceType:   "",
			OutletDeviceSn:     snDevice,
			PaymentMethodId:    0,
			PaymentMethodName:  "",
		}, tableNm, tx)
		if !status {
			log.Println(logErr + "Inq Failed")
			// result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Inq Failed", nil)
			errMsg = errors.New("InsertTrxPos Failed")
			return errMsg
		}
		for _, data := range req.ProductDetails {
			// idProduct = append(idProduct, data.ProductId)
			// subtotal = subtotal + (data.ProductPrice * float64(data.Qty))
			// subtotal, _ = strconv.ParseFloat(strconv.Itoa(req.Qty*int(productPrice)), 64)
			reqDetails := models.ReqInsertTrxDetails{
				ID:            idTrx,
				ProductTypeId: req.ProductTypeId,
				ProductId:     data.ProductId,
				ProductCode:   data.ProductCode,
				ProductName:   data.ProductName,
				ProductPrice:  data.ProductPrice,
				Quantity:      data.Qty,
				CustomerId:    "",
				BillInfo:      "",
			}
			reqInsertDetails = append(reqInsertDetails, reqDetails)
		}
		status = svc.service.ApiTrx.InsertTrxDetails(reqInsertDetails, tx)
		fmt.Println(status)
		if !status {
			log.Println(logErr + "Inq Failed")
			// result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Inq Failed", nil)
			errMsg = errors.New("InsertTrxDetails Failed")
			return errMsg
		}
		return nil
	})
	resInq := models.ReqInsertTrx{
		TrxNumber:            dataType + code,
		TrxNumberPartner:     "",
		PaymentAt:            "",
		CreatedBy:            nickname,
		CreatedAt:            dbTime,
		UpdatedBy:            nickname,
		UpdatedAt:            dbTime,
		StatusCode:           "66",
		StatusMessage:        "SUCCESS",
		StatusDesc:           "INQUIRY SUCCESS",
		StatusCodePartner:    "",
		StatusMessagePartner: "",
		StatusDescPartner:    "",
		SegmentId:            0,
		ProductTypeId:        productTypeId,
		ProductTypeName:      "",
		ProductCategoryId:    productCategoryId,
		ProductCategoryName:  "",
		ProductId:            productId,
		ProductCode:          productCode,
		ProductName:          productName,
		ProductPrice:         productPrice,
		ProductAdminFee:      0,
		ProductMerchantFee:   0,
		// Quantity:             req.Qty,
		SubTotal:   subtotal,
		GrandTotal: subtotal,
		// CustomerId:           req.CustomerId,
		BillInfo:           "",
		ClientId:           int(cID),
		ClientName:         clientNm,
		MerchantId:         0,
		MerchantName:       merchantNm,
		MerchantOutletId:   int(oID),
		MerchantOutletName: merchantNm,
		UserOutletId:       int(uID),
		UserOutletName:     outletNm,
		OutletDeviceId:     int(oID),
		OutletDeviceType:   "",
		OutletDeviceSn:     snDevice,
		PaymentMethodId:    0,
		PaymentMethodName:  "",
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", resInq)
	return ctx.JSON(http.StatusOK, result)
	//cek product type
	//cek produk category
	//cek produk

}

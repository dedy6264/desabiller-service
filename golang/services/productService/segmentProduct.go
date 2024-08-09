package productservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func (svc ProductService) DropSegmentProduct(ctx echo.Context) error {
	var (
		svcName = "DropSegmentProduct"
		result  models.Response
		// resultSvc models.ResponseList
		formatLogError = "Erorr " + svcName + " ::"
		dbTime         = time.Now().Format(time.RFC3339)
		// dbTime         = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqListSegmentProduct)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	resPro, _ := svc.service.ApiProduct.GetListSegmentProduct(models.ReqListSegmentProduct{
		ID: req.ID,
	})
	if len(resPro) > 0 {
		for _, data := range resPro {
			if data.IsOpen {
				log.Println(formatLogError + " DropSegmentProduct Product is Active")
				result = models.Response{
					StatusCode:       configs.VALIDATE_ERROR_CODE,
					Success:          false,
					ResponseDatetime: dbTime,
					Result:           "",
					Message:          "Failed",
				}
				return ctx.JSON(http.StatusOK, result)
			}
		}

	}
	status := svc.service.ApiProduct.DropSegmentProduct(*req)
	if !status {
		log.Println(formatLogError + " DropSegmentProduct Failed")
		result = models.Response{
			StatusCode:       configs.DB_NOT_FOUND,
			Success:          false,
			ResponseDatetime: dbTime,
			Result:           "",
			Message:          "Failed",
		}
		return ctx.JSON(http.StatusOK, result)
	}
	result = models.Response{
		StatusCode:       configs.SUCCESS_CODE,
		Success:          true,
		ResponseDatetime: dbTime,
		Result:           "",
		Message:          "Success",
	}
	return ctx.JSON(http.StatusOK, result)
}
func (svc ProductService) UpdateSegmentProduct(ctx echo.Context) error {
	var (
		svcName = "UpdateSegmentProduct"
		result  models.Response
		// resultSvc models.ResponseList
		formatLogError = "Erorr " + svcName + " ::"
		dbTime         = time.Now().Format(time.RFC3339)
		// dbTime         = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqListSegmentProduct)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.ID == 0 || req.ProductBillerId == 0 || req.ProductBillerProviderId == 0 || req.SegmentId == 0 {
		log.Println(formatLogError + " ID is null")
		result = models.Response{
			StatusCode:       configs.VALIDATE_ERROR_CODE,
			Success:          false,
			ResponseDatetime: dbTime,
			Result:           "",
			Message:          "Failed",
		}
		return ctx.JSON(http.StatusOK, result)
	}
	req.SegmentProductPrefix = strconv.Itoa(req.SegmentId) + strconv.Itoa(req.ProductBillerId)
	// if req.SegmentProductName == "" {
	// 	log.Println(formatLogError + " segmentProduct Name is null")
	// 	result = models.Response{
	// 		StatusCode:       configs.VALIDATE_ERROR_CODE,
	// 		Success:          false,
	// 		ResponseDatetime: dbTime,
	// 		Result:           "",
	// 		Message:          "Failed",
	// 	}
	// 	return ctx.JSON(http.StatusOK, result)
	// }
	// resGet, status := svc.service.ApiProduct.GetListSegmentProduct(models.ReqListSegmentProduct{
	// 	ID: req.ID,
	// })
	// if len(resGet) != 0 {
	// 	for _, data := range resGet {
	// 		if data.SegmentProductName == req.SegmentProductName {
	// 			log.Println(formatLogError + " segmentProduct Name is exist")
	// 			result = models.Response{
	// 				StatusCode:       configs.VALIDATE_ERROR_CODE,
	// 				Success:          false,
	// 				ResponseDatetime: dbTime,
	// 				Result:           "",
	// 				Message:          "Failed",
	// 			}
	// 			return ctx.JSON(http.StatusOK, result)
	// 		}
	// 	}
	// } else {
	// 	log.Println(formatLogError + " segmentProduct not exist")
	// 	result = models.Response{
	// 		StatusCode:       configs.VALIDATE_ERROR_CODE,
	// 		Success:          false,
	// 		ResponseDatetime: dbTime,
	// 		Result:           "",
	// 		Message:          "Failed",
	// 	}
	// 	return ctx.JSON(http.StatusOK, result)
	// }
	_, status := svc.service.ApiProduct.UpdateSegmentProduct(*req)
	if !status {
		log.Println(formatLogError + " UpdateSegmentProduct Failed")
		result = models.Response{
			StatusCode:       configs.DB_NOT_FOUND,
			Success:          false,
			ResponseDatetime: dbTime,
			Result:           "",
			Message:          "Failed",
		}
		return ctx.JSON(http.StatusOK, result)
	}
	result = models.Response{
		StatusCode:       configs.SUCCESS_CODE,
		Success:          true,
		ResponseDatetime: dbTime,
		Result:           "",
		Message:          "Success",
	}
	return ctx.JSON(http.StatusOK, result)
}
func (svc ProductService) GetListSegmentProduct(ctx echo.Context) error {
	var (
		svcName = "GetListSegmentProduct"
		result  models.Response
		// resultSvc models.ResponseList
		formatLogError = "Erorr " + svcName + " ::"
		dbTime         = time.Now().Format(time.RFC3339)
		// dbTime         = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqListSegmentProduct)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	resSeg, _ := svc.service.ApiProduct.GetListSegmentProduct(*req)
	if len(resSeg) == 0 {
		log.Println(formatLogError + " GetListSegmentProduct Not Found")
		result = models.Response{
			StatusCode:       configs.DB_NOT_FOUND,
			Success:          false,
			ResponseDatetime: dbTime,
			Result:           "",
			Message:          "Not Found",
		}
		return ctx.JSON(http.StatusOK, result)
	}
	result = models.Response{
		StatusCode:       configs.SUCCESS_CODE,
		Success:          true,
		ResponseDatetime: dbTime,
		Result:           resSeg,
		Message:          "Success",
	}
	return ctx.JSON(http.StatusOK, result)
}
func (svc ProductService) AddSegmentProduct(ctx echo.Context) error {
	var (
		svcName = "AddSegmentProduct"
		result  models.Response
		// resultSvc models.ResponseList
		formatLogError = "Erorr " + svcName + " ::"
		dbTime         = time.Now().Format(time.RFC3339)
		// dbTime         = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqListSegmentProduct)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	req.SegmentProductPrefix = strconv.Itoa(req.SegmentId) + strconv.Itoa(req.ProductBillerId)
	if req.ProductBillerId == 0 || req.ProductBillerProviderId == 0 || req.SegmentId == 0 {
		log.Println(formatLogError + " ID is null")
		result = models.Response{
			StatusCode:       configs.VALIDATE_ERROR_CODE,
			Success:          false,
			ResponseDatetime: dbTime,
			Result:           "",
			Message:          "Failed",
		}
		return ctx.JSON(http.StatusOK, result)
	}
	resSeg, status := svc.service.ApiProduct.GetListSegmentProduct(models.ReqListSegmentProduct{
		SegmentProductPrefix: req.SegmentProductPrefix,
	})
	if len(resSeg) > 0 {
		log.Println(formatLogError + "GetListSegmentProduct product is exist")
		result = models.Response{
			StatusCode:       configs.VALIDATE_ERROR_CODE,
			Success:          false,
			ResponseDatetime: dbTime,
			Result:           "",
			Message:          "Failed",
		}
		return ctx.JSON(http.StatusOK, result)
	}
	_, status = svc.service.ApiProduct.AddSegmentProduct(*req)
	if !status {
		log.Println(formatLogError + "AddSegmentProduct Failled add new segmentProduct")
		result = models.Response{
			StatusCode:       configs.VALIDATE_ERROR_CODE,
			Success:          false,
			ResponseDatetime: dbTime,
			Result:           "",
			Message:          "Failed",
		}
		return ctx.JSON(http.StatusOK, result)
	}
	result = models.Response{
		StatusCode:       configs.SUCCESS_CODE,
		Success:          true,
		ResponseDatetime: dbTime,
		Result:           req,
		Message:          "Success",
	}
	return ctx.JSON(http.StatusOK, result)
}

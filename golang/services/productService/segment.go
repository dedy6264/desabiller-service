package productservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (svc ProductService) DropSegment(ctx echo.Context) error {
	var (
		svcName = "DropSegment"
		result  models.Response
		// resultSvc models.ResponseList
		formatLogError = "Erorr " + svcName + " ::"
		dbTime         = time.Now().Format(time.RFC3339)
		// dbTime         = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqListSegment)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	status := svc.service.ApiProduct.DropSegment(*req)
	if !status {
		log.Println(formatLogError + " DropSegment Failed")
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
func (svc ProductService) UpdateSegment(ctx echo.Context) error {
	var (
		svcName = "UpdateSegment"
		result  models.Response
		// resultSvc models.ResponseList
		formatLogError = "Erorr " + svcName + " ::"
		dbTime         = time.Now().Format(time.RFC3339)
		// dbTime         = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqListSegment)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.ID == 0 {
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
	if req.SegmentName == "" {
		log.Println(formatLogError + " segment Name is null")
		result = models.Response{
			StatusCode:       configs.VALIDATE_ERROR_CODE,
			Success:          false,
			ResponseDatetime: dbTime,
			Result:           "",
			Message:          "Failed",
		}
		return ctx.JSON(http.StatusOK, result)
	}
	resGet, status := svc.service.ApiProduct.GetListSegment(models.ReqListSegment{
		ID: req.ID,
	})
	if len(resGet) != 0 {
		for _, data := range resGet {
			if data.SegmentName == req.SegmentName {
				log.Println(formatLogError + " segment Name is exist")
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
	} else {
		log.Println(formatLogError + " segment not exist")
		result = models.Response{
			StatusCode:       configs.VALIDATE_ERROR_CODE,
			Success:          false,
			ResponseDatetime: dbTime,
			Result:           "",
			Message:          "Failed",
		}
		return ctx.JSON(http.StatusOK, result)
	}
	_, status = svc.service.ApiProduct.UpdateSegment(*req)
	if !status {
		log.Println(formatLogError + " UpdateSegment Failed")
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
func (svc ProductService) GetListSegment(ctx echo.Context) error {
	var (
		svcName = "GetListSegment"
		result  models.Response
		// resultSvc models.ResponseList
		formatLogError = "Erorr " + svcName + " ::"
		dbTime         = time.Now().Format(time.RFC3339)
		// dbTime         = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqListSegment)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	resSeg, _ := svc.service.ApiProduct.GetListSegment(*req)
	if len(resSeg) == 0 {
		log.Println(formatLogError + " GetListSegment Not Found")
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
func (svc ProductService) AddSegment(ctx echo.Context) error {
	var (
		svcName = "AddSegment"
		result  models.Response
		// resultSvc models.ResponseList
		formatLogError = "Erorr " + svcName + " ::"
		dbTime         = time.Now().Format(time.RFC3339)
		// dbTime         = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqListSegment)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("FAILLED BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED BINDING"+err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.SegmentName == "" {
		log.Println(formatLogError + " Segment Name is null")
		result = models.Response{
			StatusCode:       configs.VALIDATE_ERROR_CODE,
			Success:          false,
			ResponseDatetime: dbTime,
			Result:           "",
			Message:          "Failed",
		}
		return ctx.JSON(http.StatusOK, result)
	}
	_, status := svc.service.ApiProduct.AddSegment(*req)
	if !status {
		log.Println(formatLogError + "AddSegment Failled add new segment")
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

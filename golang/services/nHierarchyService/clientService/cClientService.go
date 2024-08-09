package clientservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/services"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

type nHierarchyClientServices struct {
	services services.UsecaseService
}

func NewApiNHierarchyClientServices(services services.UsecaseService) nHierarchyClientServices {
	return nHierarchyClientServices{
		services: services,
	}
}
func (svc nHierarchyClientServices) CreateClientService(ctx echo.Context) error {
	var (
		svcName  = "CreateClientService"
		logErr   = "Error " + svcName
		req      models.ReqGetListClient
		response models.ResGetNClient
		dbTime   = time.Now().Format(time.RFC3339)
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println(logErr+"BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ClientName == "" {
		log.Println(logErr + " Client Name cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Client Name cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.ClientName = strings.ToUpper(req.ClientName)
	id, err := svc.services.ApiNHierarchy.NCreateClient(models.ReqGetListNClient{
		ClientName: req.ClientName,
	})
	if err != nil {
		log.Println(logErr+"NCreateClient", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_ERROR, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	response.ID = id
	response.ClientName = req.ClientName
	response.CreatedAt = dbTime
	response.CreatedBy = "sys"
	response.UpdatedAt = dbTime
	response.UpdatedBy = "sys"
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", response)
	return ctx.JSON(http.StatusOK, result)
}

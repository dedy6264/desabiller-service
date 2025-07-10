package administrationservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/utils"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (svc AdministrationService) Login(ctx echo.Context) error {
	var (
		svcName   = "Login"
		resultSvc models.RespLogin
		isSetPin  string
	)
	req := new(models.ReqLogin)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_VALIDATION_FAILED[0], configs.RC_VALIDATION_FAILED[1], err.Error(), nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	if req.Username == "" {
		utils.Log("username cannot emptyEnc ", svcName, nil)
		result := helpers.ResponseJSON(false, configs.RC_INVALID_PARAM[0], "Failed", "username cannot empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Password == "" {
		utils.Log(" Password cannot emptyEnc", svcName, nil)
		result := helpers.ResponseJSON(false, configs.RC_INVALID_PARAM[0], "Failed", "password cannot empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Password, err = helpers.PswEnc(req.Password)
	if err != nil {
		utils.Log(" wrong username or password", svcName, err)
		result := helpers.ResponseJSON(false, configs.RC_SYSTEM_ERROR[0], configs.RC_SYSTEM_ERROR[1], configs.RC_SYSTEM_ERROR[1], nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respGet, err := svc.service.RepoHierarchy.GetUserApp(models.ReqGetUserApp{
		Filter: models.UserApp{
			Username: req.Username,
		},
	})
	if err != nil {
		utils.Log(" GetUserApp", svcName, err)
		result := helpers.ResponseJSON(false, configs.RC_FAILED_WRONG_PWD_USRNAME[0], configs.RC_FAILED_WRONG_PWD_USRNAME[1], configs.RC_FAILED_WRONG_PWD_USRNAME[1], nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if respGet.Password != req.Password {
		log.Println("Err ", svcName, "Password cannot emptyEnc")
		utils.Log(" wrong username or password", svcName, err)
		result := helpers.ResponseJSON(false, configs.RC_FAILED_WRONG_PWD_USRNAME[0], configs.RC_FAILED_WRONG_PWD_USRNAME[1], configs.RC_FAILED_WRONG_PWD_USRNAME[1], nil)
		return ctx.JSON(http.StatusOK, result)
	}
	token, err := helpers.TokenJwtGenerate(int(respGet.ID))
	if err != nil {
		utils.Log(" TokenJwtGenerate", svcName, err)
		result := helpers.ResponseJSON(false, configs.RC_SYSTEM_ERROR[0], configs.RC_SYSTEM_ERROR[1], configs.RC_SYSTEM_ERROR[1], nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respAccount, err := svc.service.SavingRepo.GetAccount(models.ReqGetAccountSaving{
		Filter: models.Account{
			ID: respGet.AccountID,
		},
	})
	if err != nil {
		utils.Log(" GetAccount", svcName, err)
		result := helpers.ResponseJSON(false, configs.RC_SYSTEM_ERROR[0], configs.RC_SYSTEM_ERROR[1], configs.RC_SYSTEM_ERROR[1], nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if respAccount.AccountPin == "" {
		isSetPin = "N"
	}
	resultSvc = models.RespLogin{
		Data: models.Data{
			ID:             int(respGet.ID),
			Username:       respGet.Username,
			Password:       "",
			Name:           respGet.Name,
			IdentityType:   respGet.IdentityType,
			IdentityNumber: respGet.IdentityNumber,
			Phone:          respGet.Phone,
			Email:          respGet.Email,
			Gender:         respGet.Gender,
			Province:       respGet.Province,
			City:           respGet.City,
			Address:        respGet.Address,
			AccountID:      respGet.AccountID,
			Status:         respGet.Status,
			// AccountNumber:  ,
			IsSetPin: isSetPin,
		},
		Token: token,
	}
	result := helpers.ResponseJSON(false, configs.RC_SUCCESS[0],
		configs.RC_SUCCESS[1],
		configs.RC_SUCCESS[1], resultSvc)
	return ctx.JSON(http.StatusOK, result)
}

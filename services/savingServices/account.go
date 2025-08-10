package savingservices

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/utils"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (svc savingServices) AddAccount(ctx echo.Context) error {
	var (
		svcName   = "AddAccount"
		t         = time.Now()
		dbTime    = t.Local().Format(configs.LAYOUT_TIMESTAMP)
		resAddAcc models.RespGetAccount
	)
	req := new(models.ReqGetAccountSaving)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILED", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.CifID == 0 {
		utils.Log(" ", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"CIF is empty", "CIF is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.AccountNumber == "" {
		utils.Log(" ", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Account Number is empty", "Account Number is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.AccountPin == "" {
		utils.Log(" ", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Pin is empty", "Pin is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.SavingSegmentID == 0 {
		utils.Log(" ", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Segment is empty", "Segment is empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.AccountPin != "" {
		req.Filter.AccountPin, err = helpers.PassEncrypt(req.Filter.AccountPin)
		if err != nil {
			utils.Log(" PassEncrypt", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE,
				configs.RC_SYSTEM_ERROR[0],
				configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}
	req.Filter.CreatedAt = dbTime
	req.Filter.UpdatedAt = dbTime
	req.Filter.CreatedBy = "sys"
	req.Filter.UpdatedBy = "sys"
	req.Filter.AccountPin, err = helpers.PassEncrypt(req.Filter.AccountPin)
	if err != nil {
		utils.Log(" PassEncrypt", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resAddAcc, err = svc.services.SavingRepo.AddAccount(*req, nil)
	if err != nil {
		utils.Log(" AddAccount", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1],
		resAddAcc)
	return ctx.JSON(http.StatusOK, result)
}
func (svc savingServices) GetAccounts(ctx echo.Context) error {
	var (
		svcName = "GetAccount"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetAccountSaving)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	count, err := svc.services.SavingRepo.GetAccountCount(*req)
	if err != nil {
		utils.Log(" GetAccountCount", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0],
			configs.RC_FAILED_DB_NOT_FOUND[1], "", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	resp, err := svc.services.SavingRepo.GetAccounts(*req)
	if err != nil {
		utils.Log(" GetAccounts", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0],
			configs.RC_FAILED_DB_NOT_FOUND[1], "", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respSvc.Data = resp
	respSvc.RecordsTotal = count
	respSvc.RecordsFiltered = count
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1],
		respSvc)
	return ctx.JSON(http.StatusOK, result)
}
func (svc savingServices) DropAccount(ctx echo.Context) error {
	var (
		svcName = "DropAccount"
	)
	req := new(models.ReqGetAccountSaving)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.services.SavingRepo.DropAccount(int(req.Filter.ID), nil)
	if err != nil {
		utils.Log(" DropAccount", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1],
		nil)
	return ctx.JSON(http.StatusOK, result)
}

// func (svc savingServices) SetPin(ctx echo.Context) error {
// 	var (
// 		svcName = "SetPin"
// 		t       = time.Now()
// 		pin     string
// 		dbTime  = t.Local().Format(configs.LAYOUT_TIMESTAMP)
// 	)
// 	req := new(models.ReqSetPin)
// 	_, err := helpers.BindValidate(req, ctx)
// 	if err != nil {
// 		utils.Log(" ", svcName, err)
// 		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
// 		return ctx.JSON(http.StatusOK, result)
// 	}
// 	if req.Pin == "" {
// 		utils.Log(" Pin cannot be null", svcName, nil)
// 		result := helpers.ResponseJSON(configs.FALSE_VALUE,
// 			configs.RC_SYSTEM_ERROR[0],
// 			configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
// 		return ctx.JSON(http.StatusOK, result)
// 	}
// 	pin, err = helpers.PassEncrypt(req.Pin)
// 	if err != nil {
// 		utils.Log(" PassEncrypt", svcName, err)
// 		result := helpers.ResponseJSON(configs.FALSE_VALUE,
// 			configs.RC_SYSTEM_ERROR[0],
// 			configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
// 		return ctx.JSON(http.StatusOK, result)
// 	}
// 	//get user apps
// 	respCif, err := svc.services.SavingRepo.GetCif(models.ReqGetCIF{
// 		Filter: models.CIF{
// 			CifPhone: req.Phone,
// 		},
// 	})
// 	if err != nil {
// 		utils.Log(" PassEncrypt", svcName, err)
// 		result := helpers.ResponseJSON(configs.FALSE_VALUE,
// 			configs.RC_SYSTEM_ERROR[0],
// 			configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
// 		return ctx.JSON(http.StatusOK, result)
// 	}
// 	respAcc, err := svc.services.SavingRepo.GetAccount(models.ReqGetAccountSaving{
// 		Filter: models.Account{
// 			CifID: respCif.ID,
// 		},
// 	})
// 	if err != nil {
// 		utils.Log(" PassEncrypt", svcName, err)
// 		result := helpers.ResponseJSON(configs.FALSE_VALUE,
// 			configs.RC_SYSTEM_ERROR[0],
// 			configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
// 		return ctx.JSON(http.StatusOK, result)
// 	}

//		err = svc.services.SavingRepo.UpdateAccount(models.ReqGetAccountSaving{
//			Filter: models.Account{
//				ID:              int64(respAcc.ID),
//				CifID:           respCif.ID,
//				AccountNumber:   respAcc.AccountNumber,
//				AccountPin:      pin,
//				Balance:         respAcc.Balance,
//				SavingSegmentID: int64(respAcc.SavingSegmentID),
//				UpdatedAt:       dbTime,
//				UpdatedBy:       "sys",
//			},
//		}, nil)
//		if err != nil {
//			utils.Log(" UpdateAccount", svcName, err)
//			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
//				configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
//			return ctx.JSON(http.StatusOK, result)
//		}
//		result := helpers.ResponseJSON(configs.TRUE_VALUE,
//			configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1],
//			nil)
//		return ctx.JSON(http.StatusOK, result)
//	}
func (svc savingServices) UpdateAccount(ctx echo.Context) error {
	var (
		svcName = "UpdateAccount"
		t       = time.Now()
		pin     string
		dbTime  = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqGetAccountSaving)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log(" ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.AccountPin != "" {
		pin, err = helpers.PassEncrypt(req.Filter.AccountPin)
		if err != nil {
			utils.Log(" PassEncrypt", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE,
				configs.RC_SYSTEM_ERROR[0],
				configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}
	req.Filter.UpdatedAt = dbTime
	req.Filter.UpdatedBy = "sys"
	err = svc.services.SavingRepo.UpdateAccount(models.ReqGetAccountSaving{
		Filter: models.Account{
			ID:              req.Filter.ID,
			CifID:           req.Filter.CifID,
			AccountNumber:   req.Filter.AccountNumber,
			AccountPin:      pin,
			Balance:         req.Filter.Balance,
			SavingSegmentID: req.Filter.SavingSegmentID,
			UpdatedAt:       req.Filter.UpdatedAt,
			UpdatedBy:       req.Filter.UpdatedBy,
		},
	}, nil)
	if err != nil {
		utils.Log(" UpdateAccount", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE,
		configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1],
		nil)
	return ctx.JSON(http.StatusOK, result)
}

package hierarchyservice

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
)

func (svc HierarcyService) GetUserApps(ctx echo.Context) error {
	var (
		svcName = "GetUserApps"
		respSvc models.ResponseList
	)
	req := new(models.ReqGetUserApp)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log("", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}

	req.Filter.Name = strings.ToUpper(req.Filter.Name)

	count, err := svc.service.RepoHierarchy.GetUserAppCount(*req)
	if err != nil {
		utils.Log("GetUserAppCount", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	resUserApp, err := svc.service.RepoHierarchy.GetUserApps(*req)
	if err != nil {
		utils.Log("GetUserApps", svcName, err)
		if err == sql.ErrNoRows {
			result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1], respSvc)
			return ctx.JSON(http.StatusOK, result)
		}
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], "Data :: empty", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	for i := range resUserApp {
		resUserApp[i].Password = ""
	}
	respSvc.RecordsTotal = count
	respSvc.RecordsFiltered = count
	respSvc.Data = resUserApp
	respSvc.Draw = req.Draw
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1], respSvc)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) AddUserApp(ctx echo.Context) error {
	var (
		svcName = "AddUserApp"
		t       = time.Now()
		dbTime  = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqGetUserApp)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log("", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.Name == "" {
		utils.Log("", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Name cannot be empty",
			"Name cannot be empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.IdentityNumber == "" {
		utils.Log("", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"ID Number cannot be empty",
			"ID Number cannot be empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.IdentityType == "" {
		utils.Log("", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"ID Type cannot be empty",
			"ID Type cannot be empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.Phone == "" {
		utils.Log("", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Phone Number cannot be empty",
			"Phone Number cannot be empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.Username == "" {
		utils.Log("", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Username cannot be empty",
			"Username cannot be empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Filter.Password == "" {
		utils.Log("", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"Password cannot be empty",
			"Password cannot be empty",
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Filter.Name = strings.ToUpper(req.Filter.Name)
	req.Filter.CreatedAt = dbTime
	req.Filter.UpdatedAt = dbTime
	req.Filter.CreatedBy = "sys"
	req.Filter.UpdatedBy = "sys"
	req.Filter.Password, err = helpers.PswEnc(req.Filter.Password)
	if err != nil {
		utils.Log("Generate Password ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			configs.RC_SYSTEM_ERROR[1],
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	_, err = svc.service.RepoHierarchy.GetUserApp(*req)
	if err != nil {
		if err == sql.ErrNoRows {
			err = helpers.DBTransaction(svc.service.RepoDB, func(Tx *sql.Tx) error {
				var (
					respCif     models.CIF
					respAccount models.RespGetAccount
					idUserApp   int
				)
				//add cifs
				respCif, err = svc.service.SavingRepo.GetCif(models.ReqGetCIF{
					Filter: models.CIF{
						CifNoID: req.Filter.IdentityNumber,
					},
				})
				// a, _ := json.Marshal(respCif)
				// fmt.Println(":::", string(a))
				if err != nil { //ga ketemu cif nya, bikin ulang semua
					if err == sql.ErrNoRows {
						fmt.Println("ga ketemu cif nya, bikin ulang semua")
						respCif, err = svc.service.SavingRepo.AddCif(models.ReqGetCIF{
							Filter: models.CIF{
								CifName:    req.Filter.Name,
								CifNoID:    req.Filter.IdentityNumber,
								CifTypeID:  req.Filter.IdentityType,
								CifIDIndex: req.Filter.IdentityNumber + req.Filter.IdentityType,
								CifPhone:   req.Filter.Phone,
								CifEmail:   req.Filter.Email,
								CifAddress: req.Filter.Address,
								CreatedAt:  req.Filter.CreatedAt,
								UpdatedAt:  req.Filter.CreatedAt,
								CreatedBy:  req.Filter.CreatedBy,
								UpdatedBy:  req.Filter.UpdatedBy,
							},
						}, Tx)
						if err != nil {
							utils.Log("AddCif ", svcName, err)
							return err
						}
						respAccount, err = svc.service.SavingRepo.AddAccount(models.ReqGetAccountSaving{
							Filter: models.Account{
								CifID:           respCif.ID,
								AccountNumber:   "19940812" + strconv.Itoa(int(respCif.ID)),
								AccountPin:      "",
								Balance:         0,
								SavingSegmentID: 2, //default 2
								CreatedAt:       req.Filter.CreatedAt,
								UpdatedAt:       req.Filter.CreatedAt,
								CreatedBy:       req.Filter.CreatedBy,
								UpdatedBy:       req.Filter.UpdatedBy,
							},
						}, Tx)
						if err != nil {
							utils.Log("AddAccount ", svcName, err)
							return err
						}
						a, _ := json.Marshal(respAccount)
						fmt.Println("respAccount|||| ", string(a))
						req.Filter.AccountID = int64(respAccount.ID)
						idUserApp, err = svc.service.RepoHierarchy.AddUserApp(models.ReqGetUserApp{
							Filter: models.UserApp{
								AccountID:      int64(respAccount.ID),
								Username:       req.Filter.Username,
								Password:       req.Filter.Password,
								Name:           req.Filter.Name,
								IdentityType:   req.Filter.IdentityType,
								IdentityNumber: req.Filter.IdentityNumber,
								Phone:          req.Filter.Phone,
								Email:          req.Filter.Email,
								Gender:         req.Filter.Gender,
								Province:       req.Filter.Province,
								City:           req.Filter.City,
								Address:        req.Filter.Address,
								Status:         req.Filter.Status,
								CreatedAt:      req.Filter.CreatedAt,
								UpdatedAt:      req.Filter.CreatedAt,
								CreatedBy:      req.Filter.CreatedBy,
								UpdatedBy:      req.Filter.UpdatedBy,
							},
						}, Tx)
						if err != nil {
							utils.Log("AddUserApp", svcName, err)
							return err
						}
					}
				} else {
					fmt.Println("ketemu cif nya")
					respAccount, err = svc.service.SavingRepo.GetAccount(models.ReqGetAccountSaving{
						Filter: models.Account{
							CifID: respCif.ID,
						},
					})
					if err != nil {
						utils.Log("GetAccount", svcName, err)
						return err
					}
					d, _ := json.Marshal(respAccount)
					fmt.Println("respAccount", string(d))
					//jika cif ada langsung buat user apps
					idUserApp, err = svc.service.RepoHierarchy.AddUserApp(models.ReqGetUserApp{
						Filter: models.UserApp{
							AccountID:      int64(respAccount.ID),
							Username:       req.Filter.Username,
							Password:       req.Filter.Password,
							Name:           req.Filter.Name,
							IdentityType:   req.Filter.IdentityType,
							IdentityNumber: req.Filter.IdentityNumber,
							Phone:          req.Filter.Phone,
							Email:          req.Filter.Email,
							Gender:         req.Filter.Gender,
							Province:       req.Filter.Province,
							City:           req.Filter.City,
							Address:        req.Filter.Address,
							Status:         req.Filter.Status,
							CreatedAt:      req.Filter.CreatedAt,
							UpdatedAt:      req.Filter.CreatedAt,
							CreatedBy:      req.Filter.CreatedBy,
							UpdatedBy:      req.Filter.UpdatedBy,
						},
					}, Tx)
					if err != nil {
						utils.Log("AddUserApp", svcName, err)
						return err
					}
				}
				d, _ := json.Marshal(respAccount)
				fmt.Println("respAccount", string(d))
				//generate otp
				op := helpers.Otp()
				err = svc.service.RepoHierarchy.AddOtp(models.ReqGetOtp{
					Filter: models.Otp{
						UserAppID:       int64(idUserApp),
						Username:        req.Filter.Username,
						Otp:             op,
						ExpiredDuration: 30,
						Phone:           req.Filter.Phone,
						CreatedBy:       req.Filter.UpdatedBy,
						UpdatedBy:       req.Filter.UpdatedBy,
						CreatedAt:       req.Filter.CreatedAt,
						UpdatedAt:       dbTime,
					},
				}, Tx)

				if err != nil {
					utils.Log("GetUserApp", svcName, err)
					return err
				}
				reFonte := models.ReqFonnte{
					Target: req.Filter.Phone,
					Message: `VILLER, 
					jangan berikan kode ini kepada siapapun : ` + op,
				}
				respByte, _, err := utils.WorkerPost("https://api.fonnte.com/send", "LyfkJ2o1LA8wER8RiMBe", reFonte, "json")
				if err != nil {
					utils.Log("WorkerPostWithBearer", svcName, err)
					return err
				}
				fmt.Println("resp fonnte: ", string(respByte))
				return nil
			})
			if err != nil {
				utils.Log("DBTransaction", svcName, err)
				result := helpers.ResponseJSON(configs.FALSE_VALUE,
					configs.RC_FAILED_USER_EXISTING[0],
					configs.RC_FAILED_USER_EXISTING[1],
					configs.RC_FAILED_USER_EXISTING[1],
					nil)
				return ctx.JSON(http.StatusOK, result)
			}
		} else {
			log.Println("Err ", svcName, "GetUserApp", err)
			utils.Log("GetUserApp ", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE,
				configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1],
				"failed",
				nil)
			return ctx.JSON(http.StatusOK, result)
		}
	} else {
		utils.Log("GetUserApp ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.RC_FAILED_USER_EXISTING[0],
			configs.RC_FAILED_USER_EXISTING[1],
			configs.RC_FAILED_USER_EXISTING[1],
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Filter.Password = ""
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1],
		req.Filter)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) DropUserApp(ctx echo.Context) error {
	var (
		svcName = "DropUserApp"
	)
	req := new(models.ReqGetUserApp)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.service.RepoHierarchy.DropUserApp(int(req.Filter.ID), nil)
	if err != nil {
		log.Println("Err ", svcName, "DropUserApp", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			configs.RC_SYSTEM_ERROR[1],
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1], nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) UpdateUserApp(ctx echo.Context) error {
	var (
		svcName = "UpdateUserApp"
		// respSvc    models.ResponseList
		t      = time.Now()
		dbTime = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.ReqGetUserApp)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		log.Println("Err ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	req.Filter.Name = strings.ToUpper(req.Filter.Name)
	req.Filter.UpdatedAt = dbTime
	req.Filter.UpdatedBy = "sys"
	req.Filter.Password, err = helpers.PswEnc(req.Filter.Password)
	if err != nil {
		utils.Log("Generate Password ", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			configs.RC_SYSTEM_ERROR[1],
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	err = svc.service.RepoHierarchy.UpdateUserApp(*req, nil)
	if err != nil {
		log.Println("Err ", svcName, "UpdateUserApp", err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.RC_SYSTEM_ERROR[0],
			configs.RC_SYSTEM_ERROR[1],
			configs.RC_SYSTEM_ERROR[1],
			nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1], req.Filter)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) VerificationOTP(ctx echo.Context) error {
	var (
		svcName = "VerificationOTP"
		t       = time.Now()
		dbTime  = t.Local().Format(time.RFC3339)
		// dbTime  = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	req := new(models.Otp)
	_, err := helpers.BindValidate(req, ctx)
	if err != nil {
		utils.Log("", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Otp == "" {
		utils.Log("OTP cannot be null", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.Phone == "" {
		utils.Log("Phone cannot be null", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respOtp, err := svc.service.RepoHierarchy.GetOtp(models.ReqGetOtp{
		Filter: models.Otp{
			Otp:   req.Otp,
			Phone: req.Phone,
		},
	})
	if err != nil {
		utils.Log("GetOtp", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}

	layout := time.RFC3339
	createdAt, err := time.Parse(layout, respOtp.CreatedAt)
	if err != nil {
		utils.Log("Gagal parse waktu", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	fmt.Println(":", createdAt)
	// Hitung batas waktu valid
	batasWaktu := createdAt.Add(time.Duration(respOtp.ExpiredDuration) * time.Minute)

	// Ambil waktu saat ini
	now := time.Now()
	// fmt.Println(now.After(batasWaktu))
	if now.After(batasWaktu) {
		utils.Log("Waktu sudah kadaluarsa", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], "", nil)
		return ctx.JSON(http.StatusOK, result)
	} else {
		utils.Log("Masih dalam durasi", svcName, nil)
		respUserApp, err := svc.service.RepoHierarchy.GetUserApp(models.ReqGetUserApp{
			Filter: models.UserApp{
				ID: respOtp.UserAppID,
			},
		})
		if err != nil {
			utils.Log("GetUserApp", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		}
		//update verifikasi
		err = svc.service.RepoHierarchy.UpdateUserApp(models.ReqGetUserApp{
			Filter: models.UserApp{
				ID:              respUserApp.ID,
				Username:        respUserApp.Username,
				Password:        respUserApp.Password,
				Name:            respUserApp.Name,
				IdentityType:    respUserApp.IdentityType,
				IdentityNumber:  respUserApp.IdentityNumber,
				Phone:           respUserApp.Phone,
				Email:           respUserApp.Email,
				Gender:          respUserApp.Gender,
				Province:        respUserApp.Province,
				City:            respUserApp.City,
				Address:         respUserApp.Address,
				AccountID:       respUserApp.AccountID,
				Status:          "isverified",
				CreatedBy:       respUserApp.Username,
				UpdatedBy:       respUserApp.Username,
				CreatedAt:       respUserApp.CreatedAt,
				UpdatedAt:       dbTime,
				AccountNumber:   respUserApp.AccountNumber,
				Balance:         respUserApp.Balance,
				SavingSegmentID: respUserApp.SavingSegmentID,
			},
		}, nil)
		if err != nil {
			utils.Log("GetUserApp", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}
	// resp:=map[string]string{
	// 	"cifId":,
	// }
	result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], "", nil)
	return ctx.JSON(http.StatusOK, result)
}

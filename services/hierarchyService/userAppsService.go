package hierarchyservice

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"desabiller/utils"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

// pendaftaran
// 1. cek cif
// 2. exist-> create user app
// 3. not exist->create cif->create userapp
// 4. generate otp -> insert otp&cif id (done)
//
// 5. verification otp,(cif_id,otp)-> get otps by cif_id & otp
// 6. update userapp isferivied->delete otps(done)
//
// 7. resend otp, (cif_id)->get otps by cif_id
// 8. generate otp
// 9. update otp(done)
//
// 10. aktivasi saving,(cif_id, pin)->get cif by cif_id
// 11. exist, get saving by cif_id
// 12. exist, tolak
// 13. not exist-> create saving

func (svc HierarcyService) AddUserApp(ctx echo.Context) error {
	var (
		svcName     = "AddUserApp"
		t           = time.Now()
		dbTime      = t.Local().Format(time.RFC3339)
		respUserApp models.UserApp
		idUserApp   int
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
	req.Filter.IdentityType = strings.ToUpper(req.Filter.IdentityType)
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
	//cek cif by nik
	respCif, err := svc.service.SavingRepo.GetCif(models.ReqGetCIF{
		Filter: models.CIF{
			CifNoID: req.Filter.IdentityNumber,
		},
	})
	if err != nil {
		if err == sql.ErrNoRows {
			//start dbtransaction
			err = helpers.DBTransaction(svc.service.RepoDB, func(Tx *sql.Tx) error {
				//create cif
				respCif, err = svc.service.SavingRepo.AddCif(models.ReqGetCIF{
					Filter: models.CIF{
						CifName:    req.Filter.Name,
						CifNoID:    req.Filter.IdentityNumber,
						CifTypeID:  req.Filter.IdentityType,
						CifIDIndex: req.Filter.IdentityNumber + req.Filter.IdentityType,
						CifEmail:   req.Filter.Email,
						CifAddress: req.Filter.Address,
						CreatedBy:  req.Filter.Username,
						UpdatedBy:  req.Filter.Username,
						CreatedAt:  dbTime,
						UpdatedAt:  dbTime,
					},
				}, Tx)
				fmt.Println(":::", respCif.ID)
				if err != nil {
					utils.Log("AddCif ", svcName, err)
					return err
				}
				// create userapp
				respUserApp, err = svc.service.RepoHierarchy.AddUserApp(models.ReqGetUserApp{
					Filter: models.UserApp{
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
						CifID:          respCif.ID,
						Status:         "N",
						CreatedBy:      req.Filter.Username,
						UpdatedBy:      req.Filter.Username,
						CreatedAt:      dbTime,
						UpdatedAt:      dbTime,
					},
				}, Tx)
				if err != nil {
					utils.Log("AddUserApp ", svcName, err)
					return err
				}
				//create otp
				otp := helpers.Otp()
				_, err := svc.service.RepoHierarchy.AddOtp(models.ReqGetOtp{
					Filter: models.Otp{
						CifID:           respCif.ID,
						Username:        req.Filter.Username,
						Otp:             otp,
						ExpiredDuration: 30,
						Phone:           req.Filter.Phone,
						CreatedBy:       req.Filter.Username,
						UpdatedBy:       req.Filter.Username,
						CreatedAt:       dbTime,
						UpdatedAt:       dbTime,
					},
				}, Tx)
				if err != nil {
					utils.Log("AddOtp ", svcName, err)
					return err
				}
				//send otp
				reFonte := models.ReqFonnte{
					Target:  req.Filter.Phone,
					Message: "OTP : " + otp,
				}
				respByte, _, err := utils.WorkerPost("https://api.fonnte.com/send", "LyfkJ2o1LA8wER8RiMBe", reFonte, "json")
				if err != nil {
					utils.Log("WorkerPost", svcName, err)
					return err
				}
				fmt.Println("resp fonnte: ", string(respByte))
				return nil
			})
			if err != nil {
				utils.Log("DBTransaction", svcName, err)
				result := helpers.ResponseJSON(configs.FALSE_VALUE,
					configs.RC_FAILED_DB_NOT_FOUND[0],
					configs.RC_FAILED_DB_NOT_FOUND[1],
					configs.RC_FAILED_DB_NOT_FOUND[1],
					nil)
				return ctx.JSON(http.StatusOK, result)
			}
		} else {
			if err != nil {
				utils.Log("GetCif", svcName, err)
				result := helpers.ResponseJSON(configs.FALSE_VALUE,
					configs.RC_FAILED_DB_NOT_FOUND[0],
					configs.RC_FAILED_DB_NOT_FOUND[1],
					"Failed",
					nil)
				return ctx.JSON(http.StatusOK, result)
			}
		}
	} else {
		//if exist, cek user app
		respUserApp, err = svc.service.RepoHierarchy.GetUserApp(models.ReqGetUserApp{
			Filter: models.UserApp{
				CifID: respCif.ID,
			},
		})
		if err != nil {
			if err == sql.ErrNoRows {
				//create user app
				err = helpers.DBTransaction(svc.service.RepoDB, func(Tx *sql.Tx) error {
					// create userapp
					respUserApp, err = svc.service.RepoHierarchy.AddUserApp(models.ReqGetUserApp{
						Filter: models.UserApp{
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
							CifID:          respCif.ID,
							Status:         "N",
							CreatedBy:      req.Filter.Username,
							UpdatedBy:      req.Filter.Username,
							CreatedAt:      dbTime,
							UpdatedAt:      dbTime,
						},
					}, Tx)
					if err != nil {
						utils.Log("AddUserApp ", svcName, err)
						return err
					}
					//create otp
					otp := helpers.Otp()
					_, err := svc.service.RepoHierarchy.AddOtp(models.ReqGetOtp{
						Filter: models.Otp{
							CifID:           respCif.ID,
							Username:        req.Filter.Username,
							Otp:             otp,
							ExpiredDuration: 30,
							Phone:           req.Filter.Phone,
							CreatedBy:       req.Filter.Username,
							UpdatedBy:       req.Filter.Username,
							CreatedAt:       dbTime,
							UpdatedAt:       dbTime,
						},
					}, Tx)
					if err != nil {
						utils.Log("AddOtp ", svcName, err)
						return err
					}
					//send otp
					reFonte := models.ReqFonnte{
						Target:  req.Filter.Phone,
						Message: "OTP : " + otp,
					}
					respByte, _, err := utils.WorkerPost("https://api.fonnte.com/send", "LyfkJ2o1LA8wER8RiMBe", reFonte, "json")
					if err != nil {
						utils.Log("WorkerPost", svcName, err)
						return err
					}
					fmt.Println("resp fonnte: ", string(respByte))
					return nil
				})
				if err != nil {
					utils.Log("DBTransaction", svcName, err)
					result := helpers.ResponseJSON(configs.FALSE_VALUE,
						configs.RC_FAILED_DB_NOT_FOUND[0],
						configs.RC_FAILED_DB_NOT_FOUND[1],
						configs.RC_FAILED_DB_NOT_FOUND[1],
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
			//user app is exist, tolak
			utils.Log("GetUserApp ", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE,
				configs.RC_FAILED_USER_EXISTING[0],
				configs.RC_FAILED_USER_EXISTING[1],
				configs.RC_FAILED_USER_EXISTING[1],
				nil)
			return ctx.JSON(http.StatusOK, result)
		}
	}
	respUserApp = models.UserApp{
		ID:             int64(idUserApp),
		Username:       respUserApp.Username,
		Password:       "",
		Name:           respUserApp.Name,
		IdentityType:   respUserApp.IdentityType,
		IdentityNumber: respUserApp.IdentityNumber,
		Phone:          respUserApp.Phone,
		Email:          respUserApp.Email,
		Gender:         respUserApp.Gender,
		Province:       respUserApp.Province,
		City:           respUserApp.City,
		Address:        respUserApp.Address,
		Status:         respUserApp.Status,
		CreatedBy:      respUserApp.CreatedBy,
		UpdatedBy:      respUserApp.UpdatedBy,
		CreatedAt:      respUserApp.CreatedAt,
		UpdatedAt:      respUserApp.UpdatedAt,
		CifID:          respCif.ID,
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], configs.RC_SUCCESS[1],
		respUserApp)
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
	if req.CifID == 0 {
		utils.Log("Cif Id cannot be null", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respOtp, err := svc.service.RepoHierarchy.GetOtp(models.ReqGetOtp{
		Filter: models.Otp{
			Otp:   req.Otp,
			CifID: req.CifID,
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
		err = svc.service.RepoHierarchy.DropOtp(int(respOtp.ID), nil)
		if err != nil {
			utils.Log("DropOtp", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		}
		utils.Log("Waktu sudah kadaluarsa", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], "", nil)
		return ctx.JSON(http.StatusOK, result)
	} else {
		utils.Log("Masih dalam durasi", svcName, nil)
		//get user apps
		respUserApp, err := svc.service.RepoHierarchy.GetUserApp(models.ReqGetUserApp{
			Filter: models.UserApp{
				CifID: respOtp.CifID,
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
				ID:             respUserApp.ID,
				Username:       respUserApp.Username,
				Password:       respUserApp.Password,
				Name:           respUserApp.Name,
				IdentityType:   respUserApp.IdentityType,
				IdentityNumber: respUserApp.IdentityNumber,
				Phone:          respUserApp.Phone,
				Email:          respUserApp.Email,
				Gender:         respUserApp.Gender,
				Province:       respUserApp.Province,
				City:           respUserApp.City,
				Address:        respUserApp.Address,
				Status:         "Y",
				CreatedBy:      respUserApp.Username,
				UpdatedBy:      respUserApp.Username,
				CreatedAt:      respUserApp.CreatedAt,
				UpdatedAt:      dbTime,
				CifID:          respOtp.CifID,
			},
		}, nil)
		if err != nil {
			utils.Log("GetUserApp", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		}
		err = svc.service.RepoHierarchy.DropOtp(int(respOtp.ID), nil)
		if err != nil {
			utils.Log("DropOtp", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		}
		//send otp
		reFonte := models.ReqFonnte{
			Target:  respUserApp.Phone,
			Message: "Selamat " + respUserApp.Name + ", akun Viller anda telah terverifikasi, silahkan lanjutkan untuk set pin untuk bisa melakukan transaksi",
		}
		respByte, _, err := utils.WorkerPost("https://api.fonnte.com/send", "LyfkJ2o1LA8wER8RiMBe", reFonte, "json")
		if err != nil {
			utils.Log("WorkerPost", svcName, err)
			result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], err.Error(), nil)
			return ctx.JSON(http.StatusOK, result)
		}
		fmt.Println("resp fonnte: ", string(respByte))
	}
	result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], "", nil)
	return ctx.JSON(http.StatusOK, result)
}
func (svc HierarcyService) ResendOtp(ctx echo.Context) error {
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
	if req.CifID == 0 {
		utils.Log("Cif Id cannot be null", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	respOtp, err := svc.service.RepoHierarchy.GetOtp(models.ReqGetOtp{
		Filter: models.Otp{
			CifID: req.CifID,
		},
	})
	if err != nil {
		utils.Log("GetOtp", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	//get userapp->phone
	respUserApp, err := svc.service.RepoHierarchy.GetUserApp(models.ReqGetUserApp{
		Filter: models.UserApp{
			CifID: respOtp.CifID,
		},
	})
	if err != nil {
		utils.Log("GetUserApp", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_FAILED_DB_NOT_FOUND[0], configs.RC_FAILED_DB_NOT_FOUND[1], err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	//generate
	otp := helpers.Otp()
	//update otp
	err = helpers.DBTransaction(svc.service.RepoDB, func(Tx *sql.Tx) error {
		err = svc.service.RepoHierarchy.UpdateOtp(models.ReqGetOtp{
			Filter: models.Otp{
				ID:              respOtp.ID,
				CifID:           respOtp.CifID,
				Username:        respUserApp.Username,
				Otp:             otp,
				ExpiredDuration: 30,
				Phone:           respUserApp.Phone,
				UpdatedBy:       respUserApp.Username,
				UpdatedAt:       dbTime,
			},
		}, Tx)
		if err != nil {
			utils.Log("UpdateOtp", svcName, err)
			return err
		}
		reFonte := models.ReqFonnte{
			Target:  respUserApp.Phone,
			Message: "OTP : " + otp,
		}
		respByte, _, err := utils.WorkerPost("https://api.fonnte.com/send", "LyfkJ2o1LA8wER8RiMBe", reFonte, "json")
		if err != nil {
			utils.Log("WorkerPost", svcName, err)
			return err
		}
		fmt.Println("resp fonnte: ", string(respByte))
		return nil
	})
	if err != nil {
		utils.Log("DBTransaction", svcName, err)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.RC_FAILED_DB_NOT_FOUND[0],
			configs.RC_FAILED_DB_NOT_FOUND[1],
			configs.RC_FAILED_DB_NOT_FOUND[1],
			nil)
		return ctx.JSON(http.StatusOK, result)
	}

	result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.RC_SUCCESS[0], configs.RC_SUCCESS[1], "", respUserApp)
	return ctx.JSON(http.StatusOK, result)
}
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
	if req.Filter.ID == 0 {
		utils.Log("", svcName, nil)
		result := helpers.ResponseJSON(configs.FALSE_VALUE,
			configs.VALIDATE_ERROR_CODE,
			"ID cannot be empty",
			"ID cannot be empty",
			nil)
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

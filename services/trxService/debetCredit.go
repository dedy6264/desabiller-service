package trxservice

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"encoding/json"
	"errors"
	"log"
	"time"
)

func (svc trxService) TriggerDebet(
	amount float64,
	accountNumber, pin, referenceNumber, trxCode string,
) (err error) {
	var (
		subSvcName = "TriggerKredit"
		t          = time.Now()
		dbTimeTrx  = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	resp, err := svc.services.SavingRepo.GetAccount(models.ReqGetAccountSaving{
		Filter: models.Account{AccountNumber: accountNumber},
	})
	if err != nil {
		log.Println(subSvcName+" ", err)
		return err
	}
	// if resp.Balance > amount {
	resp.Balance += amount
	// }
	{
		err = helpers.DBTransaction(svc.services.RepoDB, func(Tx *sql.Tx) error {
			err = svc.services.SavingRepo.UpdateAccount(models.ReqGetAccountSaving{
				Filter: models.Account{
					ID:              int64(resp.ID),
					CifID:           int64(resp.CifID),
					AccountNumber:   resp.AccountNumber,
					Balance:         resp.Balance,
					SavingSegmentID: int64(resp.SavingSegmentID),
					AccountPin:      resp.AccountPin,
				},
			}, Tx)
			if err != nil {
				log.Println(subSvcName+" FAILED ", err)
				return err
			}
			_, err = svc.services.SavingRepo.AddSavingTransaction(models.ReqGetSavingTransaction{
				Filter: models.SavingTransaction{
					ReferenceNumber:        referenceNumber,
					ReferenceNumberPartner: referenceNumber,
					DcType:                 configs.TRX_TYPE_DEBET,
					TransactionAmount:      amount,
					TransactionCode:        trxCode,
					AccountID:              resp.ID,
					AccountNumber:          resp.AccountNumber,
					LastBalance:            resp.Balance - amount,
					CreatedAt:              dbTimeTrx,
					CreatedBy:              "sys",
					UpdatedAt:              dbTimeTrx,
					UpdatedBy:              "sys",
				},
			}, Tx)
			if err != nil {
				log.Println(subSvcName+" FAILED ", err)
				return err
			}
			return nil
		},
		)
		if err != nil {
			log.Println(subSvcName+" FAILED ", err)
			return err
		}
	}
	return nil
}

// type ReqTriggerKredit struct {
// 	amount
// 	accountNumber
// 	accountPin
// 	referenceNumber
// }

var ()

func (svc trxService) TriggerKredit(
	amount float64,
	accountNumber, pin, referenceNumber, trxCode string,
) (err error) {
	var (
		subSvcName = "TriggerKredit"
		t          = time.Now()
		dbTimeTrx  = t.Local().Format(configs.LAYOUT_TIMESTAMP)
	)
	resp, err := svc.services.SavingRepo.GetAccount(models.ReqGetAccountSaving{
		Filter: models.Account{AccountNumber: accountNumber},
	})
	if err != nil {
		log.Println(subSvcName+" ", err)
		return err
	}
	a, _ := json.Marshal(resp)
	log.Println(pin+" Account ", string(a))
	if resp.AccountPin == "" { //set pin dlu
		log.Println(subSvcName+" Pin Set before ", err)
		return errors.New("UNSETPIN")
	}
	// pin, err = helpers.PassEncrypt(resp.AccountPin)
	// if err != nil {
	// 	log.Println(subSvcName+" encrypt ", err)
	// 	return err
	// }
	err = helpers.PassCheck(resp.AccountPin, pin)
	if err != nil {
		log.Println(subSvcName+" WRONG PIN ", err)
		return errors.New("WRONGPIN")
	}
	//blm bener pengurangannya

	if resp.Balance > amount {
		resp.Balance -= amount
	} else {
		log.Println(subSvcName+" Balance not enough ", err)
		return errors.New("BALANCE_NOT_ENOUGH")
	}
	//--->db transaction untuk update amount dan insert saving transaction
	{
		err = helpers.DBTransaction(svc.services.RepoDB, func(Tx *sql.Tx) error {
			err = svc.services.SavingRepo.UpdateAccount(models.ReqGetAccountSaving{
				Filter: models.Account{
					ID:              int64(resp.ID),
					CifID:           int64(resp.CifID),
					AccountNumber:   resp.AccountNumber,
					Balance:         resp.Balance,
					SavingSegmentID: int64(resp.SavingSegmentID),
					AccountPin:      resp.AccountPin,
				},
			}, Tx)
			if err != nil {
				log.Println(subSvcName+" FAILED ", err)
				return err
			}
			_, err = svc.services.SavingRepo.AddSavingTransaction(models.ReqGetSavingTransaction{
				Filter: models.SavingTransaction{
					ReferenceNumber:        referenceNumber,
					ReferenceNumberPartner: referenceNumber,
					DcType:                 configs.TRX_TYPE_CREDIT,
					TransactionAmount:      amount,
					TransactionCode:        trxCode,
					AccountID:              resp.ID,
					AccountNumber:          resp.AccountNumber,
					LastBalance:            resp.Balance + amount,
					CreatedAt:              dbTimeTrx,
					CreatedBy:              "sys",
					UpdatedAt:              dbTimeTrx,
					UpdatedBy:              "sys",
				},
			}, Tx)
			if err != nil {
				log.Println(subSvcName+" FAILED ", err)
				return err
			}
			return nil
		},
		)
		if err != nil {
			log.Println(subSvcName+" FAILED ", err)
			return err
		}
	}
	return nil
}

// Transaction Code
// Pembelian 100
// Pembayaran 150
//

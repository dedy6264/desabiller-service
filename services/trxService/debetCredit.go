package trxservice

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
)

func (svc trxService) TriggerDebet(
	amount float64,
	accountNumber, pin, referenceNumber, trxCode string,
) (err error) {
	var (
		subSvcName = "TriggerKredit"
	)
	resp, err := svc.services.SavingRepo.GetAccount(models.ReqGetAccount{
		AccountNumber: accountNumber,
	})
	if err != nil {
		log.Println(subSvcName+" ", err)
		return err
	}
	if resp.Balance > amount {
		resp.Balance += amount
	}
	{
		err = helpers.DBTransaction(svc.services.RepoDB, func(Tx *sql.Tx) error {
			err = svc.services.SavingRepo.UpdateAccount(models.ReqGetAccount{
				ID:              resp.ID,
				CifID:           resp.CifID,
				AccountNumber:   resp.AccountNumber,
				Balance:         resp.Balance,
				SavingSegmentID: resp.SavingSegmentID,
				AccountPin:      resp.AccountPin,
			}, Tx)
			if err != nil {
				log.Println(subSvcName+" FAILED ", err)
				return err
			}
			_, err = svc.services.SavingRepo.AddSavingTransaction(models.ReqGetSavingTransaction{
				ReferenceNumber:       referenceNumber,
				SavingReferenceNumber: referenceNumber,
				DcType:                configs.TRX_TYPE_DEBET,
				TransactionAmount:     amount,
				TransactionCode:       trxCode,
				AccountID:             resp.ID,
				AccountNumber:         resp.AccountNumber,
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
	)
	resp, err := svc.services.SavingRepo.GetAccount(models.ReqGetAccount{
		AccountNumber: accountNumber,
	})
	if err != nil {
		log.Println(subSvcName+" ", err)
		return err
	}
	if resp.AccountPin == "" { //set pin dlu
		log.Println(subSvcName+" Pin Set before ", err)
		return err
	}
	// pin, err = helpers.PassEncrypt(resp.AccountPin)
	// if err != nil {
	// 	log.Println(subSvcName+" encrypt ", err)
	// 	return err
	// }
	err = helpers.PassCheck(resp.AccountPin, pin)
	if err != nil {
		log.Println(subSvcName+" WRONG PIN ", err)
		return err
	}
	//blm bener pengurangannya

	if resp.Balance > amount {
		resp.Balance -= amount
	}
	//--->db transaction untuk update amount dan insert saving transaction
	{
		err = helpers.DBTransaction(svc.services.RepoDB, func(Tx *sql.Tx) error {
			err = svc.services.SavingRepo.UpdateAccount(models.ReqGetAccount{
				ID:              resp.ID,
				CifID:           resp.CifID,
				AccountNumber:   resp.AccountNumber,
				Balance:         resp.Balance,
				SavingSegmentID: resp.SavingSegmentID,
				AccountPin:      resp.AccountPin,
			}, Tx)
			if err != nil {
				log.Println(subSvcName+" FAILED ", err)
				return err
			}
			_, err = svc.services.SavingRepo.AddSavingTransaction(models.ReqGetSavingTransaction{
				ReferenceNumber:       referenceNumber,
				SavingReferenceNumber: referenceNumber,
				DcType:                configs.TRX_TYPE_CREDIT,
				TransactionAmount:     amount,
				TransactionCode:       trxCode,
				AccountID:             resp.ID,
				AccountNumber:         resp.AccountNumber,
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

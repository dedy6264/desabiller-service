package savingrepo

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/models"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (ctx savingRepository) GetSavingTransactionCount(req models.ReqGetSavingTransaction) (result int, err error) {
	query := `select count(id) from saving_transactions where true `
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.ReferenceNumber != "" {
		query += ` and reference_number='` + req.ReferenceNumber + `'`
	}
	if req.SavingReferenceNumber != "" {
		query += ` and saving_reference_number='` + req.SavingReferenceNumber + `'`
	}
	if req.DcType != "" {
		query += ` and dc_type='` + req.DcType + `'`
	}
	if req.TransactionCode != "" {
		query += ` and transaction_code='` + req.TransactionCode + `'`
	}
	if req.AccountNumber != "" {
		query += ` and account_number='` + req.AccountNumber + `'`
	}
	if req.AccountID != 0 {
		query += ` and account_id = ` + strconv.Itoa(req.AccountID)
	}
	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		log.Println("GetCount :: ", err.Error())
		return 0, err
	}
	return result, nil
}
func (ctx savingRepository) DropSavingTransaction(id int, tx *sql.Tx) (err error) {
	query := `delete from saving_transactions where id = $1`
	if tx != nil {
		_, err = tx.Exec(query, id)
	} else {
		_, err = ctx.repo.Db.Exec(query, id)
	}
	if err != nil {
		log.Println("DropSavingTransaction :: ", err.Error())
		return err
	}
	return nil
}
func (ctx savingRepository) UpdateSavingTransaction(req models.ReqGetSavingTransaction, tx *sql.Tx) (err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update saving_transactions set 
				reference_number=$1,
				saving_reference_number=$2,
				dc_type=$3,
				transaction_amount=$4,
				transaction_code=$5,
				account_id=$6,
				account_number=$7,
				last_balance=$8,
				updated_at = $9,
				updated_by =$10
				where id = $11
				`
	if tx != nil {
		_, err = tx.Exec(query, req.ReferenceNumber, req.SavingReferenceNumber, req.DcType,
			req.TransactionAmount, req.TransactionCode, req.AccountID, req.AccountNumber, dbTime, "sys", req.ID)
	} else {
		_, err = ctx.repo.Db.Exec(query, req.ReferenceNumber, req.SavingReferenceNumber, req.DcType,
			req.TransactionAmount, req.TransactionCode, req.AccountID, req.AccountNumber, dbTime, "sys", req.ID)
	}
	if err != nil {
		log.Println("UpdateSavingTransaction :: ", err.Error())
		return err
	}
	return nil
}
func (ctx savingRepository) AddSavingTransaction(req models.ReqGetSavingTransaction, tx *sql.Tx) (result models.RespGetSavingTransaction, err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into saving_transactions (reference_number,
	saving_reference_number,
	dc_type,
	transaction_amount,
	transaction_code,
	account_id,
	account_number,
	last_balance,
	created_at,
	updated_at,
	 created_by,
	  updated_by) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) returning id,reference_number,
	saving_reference_number,
	dc_type,
	transaction_amount,
	transaction_code,
	account_id,
	account_number,
	last_balance,created_at,updated_at, created_by,  updated_by`
	fmt.Println(query, dbTime)
	if tx != nil {
		err = tx.QueryRow(query, req.ReferenceNumber, req.SavingReferenceNumber, req.DcType,
			req.TransactionAmount, req.TransactionCode, req.AccountID, req.AccountNumber, req.LastBalance,
			dbTime,
			dbTime, "sys", "sys").Scan(
			&result.ID,
			&result.ReferenceNumber, &result.SavingReferenceNumber, &result.DcType,
			&result.TransactionAmount, &result.TransactionCode, &result.AccountID, &result.AccountNumber, &result.LastBalance,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	} else {
		err = ctx.repo.Db.QueryRow(query, req.ReferenceNumber, req.SavingReferenceNumber, req.DcType,
			req.TransactionAmount, req.TransactionCode, req.AccountID, req.AccountNumber, req.LastBalance,
			dbTime,
			dbTime, "sys", "sys").Scan(
			&result.ID,
			&result.ReferenceNumber, &result.SavingReferenceNumber, &result.DcType,
			&result.TransactionAmount, &result.TransactionCode, &result.AccountID, &result.AccountNumber, &result.LastBalance,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	}

	if err != nil {
		log.Println("AddSavingTransaction :: ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetSavingTransaction(req models.ReqGetSavingTransaction) (result models.RespGetSavingTransaction, err error) {
	query := `select id,
reference_number,
	saving_reference_number,
	dc_type,
	transaction_amount,
	transaction_code,
	account_id,
	account_number,
	last_balance,created_at, created_by, updated_at, updated_by from saving_transactions where true`
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.ReferenceNumber != "" {
		query += ` and reference_number='` + req.ReferenceNumber + `'`
	}
	if req.SavingReferenceNumber != "" {
		query += ` and saving_reference_number='` + req.SavingReferenceNumber + `'`
	}
	if req.DcType != "" {
		query += ` and dc_type='` + req.DcType + `'`
	}
	if req.TransactionCode != "" {
		query += ` and transaction_code='` + req.TransactionCode + `'`
	}
	if req.AccountNumber != "" {
		query += ` and account_number='` + req.AccountNumber + `'`
	}
	if req.AccountID != 0 {
		query += ` and account_id = ` + strconv.Itoa(req.AccountID)
	}
	err = ctx.repo.Db.QueryRow(query).Scan(&result.ID,
		&result.ReferenceNumber, &result.SavingReferenceNumber, &result.DcType,
		&result.TransactionAmount, &result.TransactionCode, &result.AccountID, &result.AccountNumber,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy)
	if err != nil {
		log.Println("GetSavingTransaction :: ", err.Error())
		return result, err
	}

	return result, nil
}
func (ctx savingRepository) GetSavingTransactions(req models.ReqGetSavingTransaction) (result []models.RespGetSavingTransaction, err error) {
	query := `select id,
reference_number,
	saving_reference_number,
	dc_type,
	transaction_amount,
	transaction_code,
	account_id,
	account_number,
	last_balance,created_at, created_by, updated_at, updated_by from saving_transactions where true `
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.ReferenceNumber != "" {
		query += ` and reference_number='` + req.ReferenceNumber + `'`
	}
	if req.SavingReferenceNumber != "" {
		query += ` and saving_reference_number='` + req.SavingReferenceNumber + `'`
	}
	if req.DcType != "" {
		query += ` and dc_type='` + req.DcType + `'`
	}
	if req.TransactionCode != "" {
		query += ` and transaction_code='` + req.TransactionCode + `'`
	}
	if req.AccountNumber != "" {
		query += ` and account_number='` + req.AccountNumber + `'`
	}
	if req.AccountID != 0 {
		query += ` and account_id = ` + strconv.Itoa(req.AccountID)
	}
	if req.Filter.Length != 0 {
		query += ` limit  ` + strconv.Itoa(req.Filter.Length) + `  offset  ` + strconv.Itoa(req.Filter.Start)
	} else {
		if req.Filter.OrderBy != "" {
			query += `  order by '` + req.Filter.OrderBy + `' asc`
		} else {
			query += `  order by id asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("GetSavingTransactions :: ", err.Error())
		return result, err
	}
	defer rows.Close()
	result, err = SavingTransactionDataRow(rows)
	if err != nil {
		log.Println("GetSavingTransactions :: ", err.Error())
		return result, err
	}
	if len(result) == 0 {
		return result, sql.ErrNoRows
	}
	return result, nil

}
func SavingTransactionDataRow(rows *sql.Rows) (result []models.RespGetSavingTransaction, err error) {
	for rows.Next() {
		var val models.RespGetSavingTransaction
		err := rows.Scan(
			&val.ID,
			&val.ReferenceNumber, &val.SavingReferenceNumber, &val.DcType,
			&val.TransactionAmount, &val.TransactionCode, &val.AccountID, &val.AccountNumber,
			&val.CreatedAt,
			&val.CreatedBy,
			&val.UpdatedAt,
			&val.UpdatedBy,
		)
		if err != nil {
			log.Println("SavingTransactionDataRow :: ", err.Error())
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}

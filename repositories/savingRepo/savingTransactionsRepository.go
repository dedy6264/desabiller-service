package savingrepo

import (
	"database/sql"
	"desabiller/models"
	"desabiller/utils"
	"strconv"
)

func (ctx savingRepository) GetSavingTransactionCount(req models.ReqGetSavingTransaction) (result int, err error) {
	query := `select count(id) from saving_transactions where true `

	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
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
		return err
	}
	return nil
}
func (ctx savingRepository) UpdateSavingTransaction(req models.ReqGetSavingTransaction, tx *sql.Tx) (err error) {
	query := `update saving_transactions set 
				reference_number=?,
				reference_number_partner=?,
				dc_type=?,
				transaction_amount=?,
				transaction_code=?,
				account_id=?,
				account_number=?,
				last_balance=?,
				updated_at = ?,
				updated_by =?
				where id = ?
				`
	query = utils.QuerySupport(query)
	if tx != nil {
		_, err = tx.Exec(query, req.Filter.ReferenceNumber, req.Filter.ReferenceNumberPartner, req.Filter.DcType,
			req.Filter.TransactionAmount, req.Filter.TransactionCode, req.Filter.AccountID, req.Filter.AccountNumber, req.Filter.UpdatedAt,
			req.Filter.UpdatedBy, req.Filter.ID)
	} else {
		_, err = ctx.repo.Db.Exec(query, req.Filter.ReferenceNumber, req.Filter.ReferenceNumberPartner, req.Filter.DcType,
			req.Filter.TransactionAmount, req.Filter.TransactionCode, req.Filter.AccountID, req.Filter.AccountNumber, req.Filter.UpdatedAt,
			req.Filter.UpdatedBy, req.Filter.ID)
	}
	if err != nil {
		return err
	}
	return nil
}
func (ctx savingRepository) AddSavingTransaction(req models.ReqGetSavingTransaction, tx *sql.Tx) (result models.SavingTransaction, err error) {
	query := `insert into saving_transactions (reference_number,
	reference_number_partner,
	dc_type,
	transaction_amount,
	transaction_code,
	account_id,
	account_number,
	last_balance,
	created_at,
	updated_at,
	 created_by,
	  updated_by) values (?,?,?,?,?,?,?,?,?,?,?,?) returning id,reference_number,
	reference_number_partner,
	dc_type,
	transaction_amount,
	transaction_code,
	account_id,
	account_number,
	last_balance,created_at,updated_at, created_by,  updated_by`
	query = utils.QuerySupport(query)
	if tx != nil {
		err = tx.QueryRow(query, req.Filter.ReferenceNumber, req.Filter.ReferenceNumberPartner, req.Filter.DcType,
			req.Filter.TransactionAmount, req.Filter.TransactionCode, req.Filter.AccountID, req.Filter.AccountNumber, req.Filter.LastBalance,
			req.Filter.CreatedAt,
			req.Filter.UpdatedAt, req.Filter.CreatedBy, req.Filter.UpdatedBy).Scan(
			&result.ID,
			&result.ReferenceNumber, &result.ReferenceNumberPartner, &result.DcType,
			&result.TransactionAmount, &result.TransactionCode, &result.AccountID, &result.AccountNumber, &result.LastBalance,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	} else {
		err = ctx.repo.Db.QueryRow(query, req.Filter.ReferenceNumber, req.Filter.ReferenceNumberPartner, req.Filter.DcType,
			req.Filter.TransactionAmount, req.Filter.TransactionCode, req.Filter.AccountID, req.Filter.AccountNumber, req.Filter.LastBalance,
			req.Filter.CreatedAt,
			req.Filter.UpdatedAt, req.Filter.CreatedBy, req.Filter.UpdatedBy).Scan(
			&result.ID,
			&result.ReferenceNumber, &result.ReferenceNumberPartner, &result.DcType,
			&result.TransactionAmount, &result.TransactionCode, &result.AccountID, &result.AccountNumber, &result.LastBalance,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	}

	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetSavingTransaction(req models.ReqGetSavingTransaction) (result models.SavingTransaction, err error) {
	query := `select id,
	reference_number,
	reference_number_partner,
	dc_type,
	transaction_amount,
	transaction_code,
	account_id,
	account_number,
	last_balance,created_at, created_by, updated_at, updated_by from saving_transactions where true`
	if req.Filter.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.Filter.ID)
	}
	if req.Filter.ReferenceNumber != "" {
		query += ` and reference_number='` + req.Filter.ReferenceNumber + `'`
	}
	if req.Filter.ReferenceNumberPartner != "" {
		query += ` and reference_number_partner='` + req.Filter.ReferenceNumberPartner + `'`
	}
	if req.Filter.DcType != "" {
		query += ` and dc_type='` + req.Filter.DcType + `'`
	}
	if req.Filter.TransactionCode != "" {
		query += ` and transaction_code='` + req.Filter.TransactionCode + `'`
	}
	if req.Filter.AccountNumber != "" {
		query += ` and account_number='` + req.Filter.AccountNumber + `'`
	}
	if req.Filter.AccountID != 0 {
		query += ` and account_id = ` + strconv.Itoa(req.Filter.AccountID)
	}
	err = ctx.repo.Db.QueryRow(query).Scan(&result.ID,
		&result.ReferenceNumber, &result.ReferenceNumber, &result.ReferenceNumberPartner, &result.DcType,
		&result.TransactionAmount, &result.TransactionCode, &result.AccountID, &result.AccountNumber,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy)
	if err != nil {
		return result, err
	}

	return result, nil
}
func (ctx savingRepository) GetSavingTransactions(req models.ReqGetSavingTransaction) (result []models.SavingTransaction, err error) {
	query := `select id,
	reference_number,
	reference_number_partner,
	dc_type,
	transaction_amount,
	transaction_code,
	account_id,
	account_number,
	last_balance,created_at, created_by, updated_at, updated_by from saving_transactions where true `
	if req.Filter.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.Filter.ID)
	}
	if req.Filter.ReferenceNumber != "" {
		query += ` and reference_number='` + req.Filter.ReferenceNumber + `'`
	}
	if req.Filter.ReferenceNumberPartner != "" {
		query += ` and reference_number_partner='` + req.Filter.ReferenceNumberPartner + `'`
	}
	if req.Filter.DcType != "" {
		query += ` and dc_type='` + req.Filter.DcType + `'`
	}
	if req.Filter.TransactionCode != "" {
		query += ` and transaction_code='` + req.Filter.TransactionCode + `'`
	}
	if req.Filter.AccountNumber != "" {
		query += ` and account_number='` + req.Filter.AccountNumber + `'`
	}
	if req.Filter.AccountID != 0 {
		query += ` and account_id = ` + strconv.Itoa(req.Filter.AccountID)
	}
	if req.Lenght != 0 {
		query += ` limit  ` + strconv.Itoa(int(req.Lenght)) + `  offset  ` + strconv.Itoa(int(req.Start))
	} else {
		if req.Order != "" {
			query += `  order by '` + req.Order + `' asc`
		} else {
			query += `  order by id asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	result, err = SavingTransactionDataRow(rows)
	if err != nil {
		return result, err
	}
	if len(result) == 0 {
		return result, sql.ErrNoRows
	}
	return result, nil

}
func SavingTransactionDataRow(rows *sql.Rows) (result []models.SavingTransaction, err error) {
	for rows.Next() {
		var val models.SavingTransaction
		err := rows.Scan(
			&val.ID,
			&val.ReferenceNumber, &val.ReferenceNumber, &val.ReferenceNumberPartner, &val.DcType,
			&val.TransactionAmount, &val.TransactionCode, &val.AccountID, &val.AccountNumber,
			&val.CreatedAt,
			&val.CreatedBy,
			&val.UpdatedAt,
			&val.UpdatedBy,
		)
		if err != nil {
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}

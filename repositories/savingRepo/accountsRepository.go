package savingrepo

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/models"
	"desabiller/utils"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func (ctx savingRepository) GetAccountCount(req models.ReqGetAccountSaving) (result int, err error) {
	query := `select count(id) from accounts where true `

	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func (ctx savingRepository) DropAccount(id int, tx *sql.Tx) (err error) {
	query := `delete from accounts where id = $1`
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
func (ctx savingRepository) UpdateAccount(req models.ReqGetAccountSaving, tx *sql.Tx) (err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	dd, _ := json.Marshal(req.Filter)
	fmt.Println("UpdateAccount dbTime: ", string(dd))
	if req.Filter.AccountPin == "" {
		query := `update accounts set 
					cif_id=?,
					account_number=?,
					balance=?,
					saving_segment_id=?,
					updated_at = ?,
					updated_by =?
					where id = ?
					`
		query = utils.QuerySupport(query)
		fmt.Println("UpdateAccount query: ", query)

		if tx != nil {
			_, err = tx.Exec(query,
				req.Filter.CifID,
				req.Filter.AccountNumber,
				req.Filter.Balance,
				req.Filter.SavingSegmentID,
				dbTime, "sys", req.Filter.ID)
		} else {
			_, err = ctx.repo.Db.Exec(query, req.Filter.CifID,
				req.Filter.AccountNumber,
				req.Filter.Balance,
				req.Filter.SavingSegmentID, dbTime, "sys", req.Filter.ID)
		}
	} else {
		query := `update accounts set 
		cif_id=?,
		account_number=?,
		account_pin=?,
		balance=?,
		saving_segment_id=?,
		updated_at = ?,
		updated_by =?
		where id = ?
		`
		query = utils.QuerySupport(query)
		fmt.Println("UpdateAccount query: ", query)
		if tx != nil {
			_, err = tx.Exec(query, req.Filter.CifID,
				req.Filter.AccountNumber,
				req.Filter.AccountPin,
				req.Filter.Balance,
				req.Filter.SavingSegmentID, dbTime, "sys", req.Filter.ID)
		} else {
			_, err = ctx.repo.Db.Exec(query, req.Filter.CifID,
				req.Filter.AccountNumber,
				req.Filter.AccountPin,
				req.Filter.Balance,
				req.Filter.SavingSegmentID, dbTime, "sys", req.Filter.ID)
		}
	}
	if err != nil {
		return err
	}
	return nil
}
func (ctx savingRepository) AddAccount(req models.ReqGetAccountSaving, tx *sql.Tx) (result models.RespGetAccount, err error) {
	query := `insert into accounts (cif_id,
account_number,
account_pin,
balance,
saving_segment_id,created_at,updated_at, created_by,  updated_by) values (?,?,?,?,?,?,?,?,?) returning id,cif_id,account_number,balance,saving_segment_id,created_at,updated_at, created_by,  updated_by`
	query = utils.QuerySupport(query)

	if tx != nil {
		err = tx.QueryRow(query, req.Filter.CifID,
			req.Filter.AccountNumber,
			req.Filter.AccountPin,
			req.Filter.Balance,
			req.Filter.SavingSegmentID,
			req.Filter.CreatedAt,
			req.Filter.UpdatedAt, req.Filter.CreatedBy, req.Filter.UpdatedBy).Scan(
			&result.ID,
			&result.CifID,
			&result.AccountNumber,
			&result.Balance,
			&result.SavingSegmentID,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	} else {
		err = ctx.repo.Db.QueryRow(query, req.Filter.CifID,
			req.Filter.AccountNumber,
			req.Filter.AccountPin,
			req.Filter.Balance,
			req.Filter.SavingSegmentID,
			req.Filter.CreatedAt,
			req.Filter.UpdatedAt, req.Filter.CreatedBy, req.Filter.UpdatedBy).Scan(
			&result.ID,
			&result.CifID,
			&result.AccountNumber,
			&result.Balance,
			&result.SavingSegmentID,
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
func (ctx savingRepository) GetAccount(req models.ReqGetAccountSaving) (result models.RespGetAccount, err error) {
	query := `select id,
	cif_id,account_number,balance,saving_segment_id,COALESCE(account_pin,'') as account_pin,created_at, created_by, updated_at, updated_by from accounts where true`

	if req.Filter.ID != 0 {
		query += ` and id = ` + strconv.Itoa(int(req.Filter.ID))
	}
	if req.Filter.CifID != 0 {
		query += ` and cif_id =` + strconv.Itoa(int(req.Filter.CifID))
	}
	if req.Filter.SavingSegmentID != 0 {
		query += ` and saving_segment_id =` + strconv.Itoa(int(req.Filter.SavingSegmentID))
	}
	if req.Filter.AccountNumber != "" {
		query += ` and account_number='` + req.Filter.AccountNumber + `'`
	}
	err = ctx.repo.Db.QueryRow(query).Scan(&result.ID,
		&result.CifID,
		&result.AccountNumber,
		&result.Balance,
		&result.SavingSegmentID,
		&result.AccountPin,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetAccounts(req models.ReqGetAccountSaving) (result []models.RespGetAccount, err error) {
	query := `select 
	a.id,
	a.cif_id,
	b.cif_name,
	b.cif_email,
	a.account_number,
	a.balance,
	a.saving_segment_id,
	a.created_at,
	a.created_by,
	a.updated_at,
	a.updated_by from accounts as a
	join cifs as b on b.id = a.cif_id where true `
	if req.Filter.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(int(req.Filter.ID))
	}
	if req.Filter.CifID != 0 {
		query += ` and a.cif_id =` + strconv.Itoa(int(req.Filter.CifID))
	}
	if req.Filter.SavingSegmentID != 0 {
		query += ` and a.saving_segment_id =` + strconv.Itoa(int(req.Filter.SavingSegmentID))
	}
	if req.Filter.AccountNumber != "" {
		query += ` and a.account_number='` + req.Filter.AccountNumber + `'`
	}
	if req.Lenght != 0 {
		query += ` limit  ` + strconv.Itoa(int(req.Lenght)) + `  offset  ` + strconv.Itoa(int(req.Start))
	} else {
		if req.Order != "" {
			query += `  order by '` + req.Order + `' asc`
		} else {
			query += `  order by a.id asc`
		}
	}
	fmt.Println("GetAccounts query: ", query)
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	result, err = AccountDataRow(rows)
	if err != nil {
		return result, err
	}
	if len(result) == 0 {
		return result, sql.ErrNoRows
	}
	return result, nil

}
func AccountDataRow(rows *sql.Rows) (result []models.RespGetAccount, err error) {
	for rows.Next() {
		var val models.RespGetAccount
		err := rows.Scan(
			&val.ID,
			&val.CifID,
			&val.CifName,
			&val.CifEmail,
			&val.AccountNumber,
			&val.Balance,
			&val.SavingSegmentID,
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

// func (ctx savingRepository) GetUserAppComplete(idUserApp int) {
// 	query := `select
// 	a.id,
// 	a.username,
// 	a.password,
// 	a.name,
// 	a.identity_type,
// 	a.identity_number,
// 	a.phone,
// 	a.email,
// 	a.gender,
// 	a.province,
// 	a.city,
// 	a.address,
// 	a.account_id,
// 	a.status, a.created_at, a.created_by, a.updated_at, a.updated_by,
// 	b.account_number, b.balance, b.saving_segment_id
// 	from user_apps as a
// 	join accounts as b on b.id=a.account_id where true `

// }

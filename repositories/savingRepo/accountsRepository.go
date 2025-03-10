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

func (ctx savingRepository) GetAccountCount(req models.ReqGetAccount) (result int, err error) {
	query := `select count(id) from accounts where true `
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.CifID != 0 {
		query += ` and cif_id =` + strconv.Itoa(req.CifID)
	}
	if req.SavingSegmentID != 0 {
		query += ` and saving_segment_id =` + strconv.Itoa(req.SavingSegmentID)
	}
	if req.AccountNumber != "" {
		query += ` and account_number='` + req.AccountNumber + `'`
	}

	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		log.Println("GetCount :: ", err.Error())
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
		log.Println("DropAccount :: ", err.Error())
		return err
	}
	return nil
}
func (ctx savingRepository) UpdateAccount(req models.ReqGetAccount, tx *sql.Tx) (err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	if req.AccountPin == "" {
		query := `update accounts set 
					cif_id=$1,
					account_number=$2,
					balance=$3,
					saving_segment_id=$4,
					updated_at = $5,
					updated_by =$6
					where id = $7
					`
		if tx != nil {
			_, err = tx.Exec(query, req.CifID, req.AccountNumber, req.Balance,
				req.SavingSegmentID, dbTime, "sys", req.ID)
		} else {
			_, err = ctx.repo.Db.Exec(query, req.CifID, req.AccountNumber, req.Balance,
				req.SavingSegmentID, dbTime, "sys", req.ID)
		}
	} else {
		query := `update accounts set 
		cif_id=$1,
		account_number=$2,
		balance=$3,
		saving_segment_id=$4,
		account_pin=$5,
		updated_at = $6,
		updated_by =$7
		where id = $8
		`
		if tx != nil {
			_, err = tx.Exec(query, req.CifID, req.AccountNumber, req.Balance,
				req.SavingSegmentID, req.AccountPin, dbTime, "sys", req.ID)
		} else {
			_, err = ctx.repo.Db.Exec(query, req.CifID, req.AccountNumber, req.Balance,
				req.SavingSegmentID, req.AccountPin, dbTime, "sys", req.ID)
		}
	}
	if err != nil {
		log.Println("UpdateAccount :: ", err.Error())
		return err
	}
	return nil
}
func (ctx savingRepository) AddAccount(req models.ReqGetAccount, tx *sql.Tx) (result models.RespGetAccount, err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into accounts (cif_id,account_number,balance,saving_segment_id,account_pin,created_at,updated_at, created_by,  updated_by) values ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id,cif_id,account_number,balance,saving_segment_id,created_at,updated_at, created_by,  updated_by`
	fmt.Println(query, dbTime)
	if tx != nil {
		err = tx.QueryRow(query, req.CifID, req.AccountNumber, req.Balance,
			req.SavingSegmentID, req.AccountPin,
			dbTime,
			dbTime, "sys", "sys").Scan(
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
		err = ctx.repo.Db.QueryRow(query, req.CifID, req.AccountNumber, req.Balance,
			req.SavingSegmentID,
			dbTime,
			dbTime, "sys", "sys").Scan(
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
		log.Println("AddAccount :: ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetAccount(req models.ReqGetAccount) (result models.RespGetAccount, err error) {
	query := `select id,
cif_id,account_number,balance,saving_segment_id,COALESCE(account_pin,'') as account_pin,created_at, created_by, updated_at, updated_by from accounts where true`
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.CifID != 0 {
		query += ` and cif_id =` + strconv.Itoa(req.CifID)
	}
	if req.SavingSegmentID != 0 {
		query += ` and saving_segment_id =` + strconv.Itoa(req.SavingSegmentID)
	}
	if req.AccountNumber != "" {
		query += ` and account_number='` + req.AccountNumber + `'`
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
		log.Println("GetAccount :: ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetAccounts(req models.ReqGetAccount) (result []models.RespGetAccount, err error) {
	query := `select id,
cif_id,account_number,balance,saving_segment_id,created_at, created_by, updated_at, updated_by from accounts where true `
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.CifID != 0 {
		query += ` and cif_id =` + strconv.Itoa(req.CifID)
	}
	if req.SavingSegmentID != 0 {
		query += ` and saving_segment_id =` + strconv.Itoa(req.SavingSegmentID)
	}
	if req.AccountNumber != "" {
		query += ` and account_number='` + req.AccountNumber + `'`
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
		log.Println("GetAccounts :: ", err.Error())
		return result, err
	}
	defer rows.Close()
	result, err = AccountDataRow(rows)
	if err != nil {
		log.Println("GetAccounts :: ", err.Error())
		return result, err
	}
	return result, nil

}
func AccountDataRow(rows *sql.Rows) (result []models.RespGetAccount, err error) {
	for rows.Next() {
		var val models.RespGetAccount
		err := rows.Scan(
			&val.ID,
			&val.CifID,
			&val.AccountNumber,
			&val.Balance,
			&val.SavingSegmentID,
			&val.CreatedAt,
			&val.CreatedBy,
			&val.UpdatedAt,
			&val.UpdatedBy,
		)
		if err != nil {
			log.Println("AccountDataRow :: ", err.Error())
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}

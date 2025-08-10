package hierarchyrepo

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

func (ctx hierarchy) AddOtp(req models.ReqGetOtp, tx *sql.Tx) (result models.Otp, err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into otps (otp,cif_id,username,phone,expired_duration,created_at,created_by,updated_at,updated_by) values (?,?,?,?,?,?,?,?,?) returning 
	id,
	otp,
	cif_id,
	username,
	phone,
	expired_duration,
	created_at,
	created_by,
	updated_at,
	updated_by
				`
	query = utils.QuerySupport(query)
	fmt.Println(query)
	if tx != nil {
		err = tx.QueryRow(query, req.Filter.Otp, req.Filter.CifID, req.Filter.Username, req.Filter.Phone, req.Filter.ExpiredDuration, dbTime, "sys", dbTime, "sys").Scan(
			&result.ID,
			&result.CifID,
			&result.Username,
			&result.Otp,
			&result.ExpiredDuration,
			&result.Phone,
			&result.CreatedBy,
			&result.UpdatedBy,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	} else {
		err = ctx.repo.Db.QueryRow(query, req.Filter.Otp, req.Filter.CifID, req.Filter.Username, req.Filter.Phone, req.Filter.ExpiredDuration, dbTime, "sys", dbTime, "sys").Scan(
			&result.ID,
			&result.CifID,
			&result.Username,
			&result.Otp,
			&result.ExpiredDuration,
			&result.Phone,
			&result.CreatedBy,
			&result.UpdatedBy,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	}
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx hierarchy) DropOtp(id int, tx *sql.Tx) (err error) {
	query := `delete from otps where id = $1`
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
func (ctx hierarchy) UpdateOtp(req models.ReqGetOtp, tx *sql.Tx) (err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update otps set 
				otp=?,
				cif_id=?,
				username=?,
				phone=?,
				expired_duration=?,
				updated_at = ?,
				updated_by =?
				where id = ? 
				`
	query = utils.QuerySupport(query)
	a, _ := json.Marshal(req.Filter)
	fmt.Println("UpdateUserApp query: ", string(a))
	if tx != nil {
		_, err = tx.Exec(query,
			req.Filter.Otp,
			req.Filter.CifID,
			req.Filter.Username,
			req.Filter.Phone,
			req.Filter.ExpiredDuration,
			dbTime,
			"sys",
			req.Filter.ID)
	} else {
		_, err = ctx.repo.Db.Exec(query, req.Filter.Otp, req.Filter.CifID, req.Filter.Username, req.Filter.Phone, req.Filter.ExpiredDuration, dbTime, "sys", req.Filter.ID)
	}
	if err != nil {
		return err
	}
	return nil
}
func (ctx hierarchy) GetOtp(req models.ReqGetOtp) (result models.Otp, err error) {
	query := `select 
	id, otp,cif_id,username,phone,expired_duration, created_at, created_by, updated_at, updated_by
	from otps where true `
	if req.Filter.ID != 0 {
		query += ` and id =` + strconv.Itoa(int(req.Filter.ID))
	}
	if req.Filter.Otp != "" {
		query += ` and otp ='` + req.Filter.Otp + `'`
	}
	if req.Filter.Phone != "" {
		query += ` and phone ='` + req.Filter.Phone + `'`
	}
	fmt.Println("GetUserApp query: ", query)
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.Otp,
		&result.CifID,
		&result.Username,
		&result.Phone,
		&result.ExpiredDuration,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy,
	)
	if err != nil {
		return result, err
	}
	return result, nil
}

package hierarchyrepo

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/models"
	"desabiller/utils"
	"fmt"
	"strconv"
	"time"
)

const userapp = `id, username,
password,
name,
identity_type,
identity_number,
phone,
email,
gender,
province,
city,
address,
account_id,
status, created_at, created_by, updated_at, updated_by`

func (ctx hierarchy) AddUserApp(req models.ReqGetUserApp, tx *sql.Tx) (err error) {

	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into user_apps (username,password,name,identity_type,identity_number,phone,email,gender,province,city,address,account_id,status,created_at,created_by,updated_at,updated_by) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) returning id
				`
	query = utils.QuerySupport(query)
	if tx != nil {
		_, err = tx.Exec(query, req.Filter.Username, req.Filter.Password, req.Filter.Name, req.Filter.IdentityType, req.Filter.IdentityNumber, req.Filter.Phone, req.Filter.Email, req.Filter.Gender, req.Filter.Province, req.Filter.City, req.Filter.Address, req.Filter.AccountID, req.Filter.Status, dbTime, "sys", dbTime, "sys")
	} else {
		_, err = ctx.repo.Db.Exec(query, req.Filter.Username, req.Filter.Password, req.Filter.Name, req.Filter.IdentityType, req.Filter.IdentityNumber, req.Filter.Phone, req.Filter.Email, req.Filter.Gender, req.Filter.Province, req.Filter.City, req.Filter.Address, req.Filter.AccountID, req.Filter.Status, dbTime, "sys", dbTime, "sys")
	}
	if err != nil {
		return err
	}
	return nil
}
func (ctx hierarchy) DropUserApp(id int, tx *sql.Tx) (err error) {
	query := `delete from user_apps where id = $1`
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
func (ctx hierarchy) UpdateUserApp(req models.ReqGetUserApp, tx *sql.Tx) (err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update user_apps set 
				username=?,
				password=?,
				name=?,
				identity_type=?,
				identity_number=?,
				phone=?,
				email=?,
				gender=?,
				province=?,
				city=?,
				address=?,
				account_id=?,
				status=?,
				updated_at = ?,
				updated_by =?
				where id = ? 
				`
	query = utils.QuerySupport(query)
	if tx != nil {
		_, err = tx.Exec(query, req.Filter.Username, req.Filter.Password, req.Filter.Name, req.Filter.IdentityType, req.Filter.IdentityNumber, req.Filter.Phone, req.Filter.Email, req.Filter.Gender, req.Filter.Province, req.Filter.City, req.Filter.Address, req.Filter.AccountID, req.Filter.Status, dbTime, "sys", req.Filter.ID)
	} else {
		_, err = ctx.repo.Db.Exec(query, req.Filter.Username, req.Filter.Password, req.Filter.Name, req.Filter.IdentityType, req.Filter.IdentityNumber, req.Filter.Phone, req.Filter.Email, req.Filter.Gender, req.Filter.Province, req.Filter.City, req.Filter.Address, req.Filter.AccountID, req.Filter.Status, dbTime, "sys", req.Filter.ID)
	}
	if err != nil {
		return err
	}
	return nil
}
func (ctx hierarchy) GetUserAppCount(req models.ReqGetUserApp) (result int, err error) {
	query := `select count(id) from user_apps where true `

	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func (ctx hierarchy) GetUserApps(req models.ReqGetUserApp) (result []models.UserApp, err error) {
	query := `select ` + userapp + ` from user_apps where true `
	if req.Filter.ID != 0 {
		query += ` and id =` + strconv.Itoa(int(req.Filter.ID))
	}
	if req.Filter.Username != "" {
		query += ` and username ='` + req.Filter.Username + `'`
	}
	if req.Filter.Email != "" {
		query += ` and email ='` + req.Filter.Email + `'`
	}
	if req.Filter.Name != "" {
		query += ` and name ='` + req.Filter.Name + `'`
	}
	if req.Filter.IdentityNumber != "" {
		query += ` and identity_number ='` + req.Filter.IdentityNumber + `'`
	}
	if req.Order != "" {
		query += ` ORDER BY ` + req.Order
		if req.Sort != "" {
			query += ` ` + req.Sort
		} else {
			query += ` ASC`
		}
	}
	if req.Start != 0 {
		query += ` LIMIT ` + strconv.Itoa(int(req.Lenght))
		query += ` OFFSET ` + strconv.Itoa(int(req.Start))
	} else {
		query += ` LIMIT ` + strconv.Itoa(int(req.Lenght))
	}

	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	result, err = DataRowUserApp(rows)
	if err != nil {
		return result, err
	}
	if len(result) == 0 {
		return result, sql.ErrNoRows
	}
	return result, nil
}
func DataRowUserApp(rows *sql.Rows) (result []models.UserApp, err error) {
	for rows.Next() {
		var val models.UserApp
		err := rows.Scan(
			&val.ID,
			&val.Username,
			&val.Password,
			&val.Name,
			&val.IdentityType,
			&val.IdentityNumber,
			&val.Phone,
			&val.Email,
			&val.Gender,
			&val.Province,
			&val.City,
			&val.Address,
			&val.AccountID,
			&val.Status,
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
func (ctx hierarchy) GetUserApp(req models.ReqGetUserApp) (result models.UserApp, err error) {
	query := `select 
	a.id, 
	a.username,
	a.password,
	a.name,
	a.identity_type,
	a.identity_number,
	a.phone,
	a.email,
	a.gender,
	a.province,
	a.city,
	a.address,
	a.account_id,
	a.status, a.created_at, a.created_by, a.updated_at, a.updated_by,
	b.account_number, b.balance, b.saving_segment_id
	from user_apps as a
	join accounts as b on b.id=a.account_id where true `
	if req.Filter.ID != 0 {
		query += ` and a.id =` + strconv.Itoa(int(req.Filter.ID))
	}
	if req.Filter.Username != "" {
		query += ` and a.username ='` + req.Filter.Username + `'`
	}
	if req.Filter.Email != "" {
		query += ` and a.email ='` + req.Filter.Email + `'`
	}
	if req.Filter.Name != "" {
		query += ` and a.name ='` + req.Filter.Name + `'`
	}
	if req.Filter.IdentityNumber != "" {
		query += ` and a.identity_number ='` + req.Filter.IdentityNumber + `'`
	}
	fmt.Println("GetUserApp query: ", query)
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.Username,
		&result.Password,
		&result.Name,
		&result.IdentityType,
		&result.IdentityNumber,
		&result.Phone,
		&result.Email,
		&result.Gender,
		&result.Province,
		&result.City,
		&result.Address,
		&result.AccountID,
		&result.Status,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy,
		&result.AccountNumber,
		&result.Balance,
		&result.SavingSegmentID,
	)
	if err != nil {
		return result, err
	}
	return result, nil
}

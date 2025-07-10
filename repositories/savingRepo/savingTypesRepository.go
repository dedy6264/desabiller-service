package savingrepo

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/models"
	"desabiller/utils"
	"log"
	"strconv"
	"time"
)

func (ctx savingRepository) GetSavingTypeCount(req models.ReqGetSavingType) (result int, err error) {
	query := `select count(id) from saving_types where true `
	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func (ctx savingRepository) DropSavingType(id int, tx *sql.Tx) (err error) {
	query := `delete from saving_types where id = $1`
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
func (ctx savingRepository) UpdateSavingType(req models.ReqGetSavingType, tx *sql.Tx) (err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update saving_types set 
				saving_type_name=?,
				saving_type_desc=?,
				updated_at=?,
				updated_by=?
				where id = ?
				`
	query = utils.QuerySupport(query)
	if tx != nil {
		_, err = tx.Exec(query, req.Filter.SavingTypeName, req.Filter.SavingTypeDesc, dbTime, "sys", req.Filter.ID)
	} else {
		_, err = ctx.repo.Db.Exec(query, req.Filter.SavingTypeName, req.Filter.SavingTypeDesc, dbTime, "sys", req.Filter.ID)
	}
	if err != nil {
		return err
	}
	return nil
}
func (ctx savingRepository) AddSavingType(req models.ReqGetSavingType, tx *sql.Tx) (result models.SavingType, err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into saving_types (saving_type_name,saving_type_desc,created_at,updated_at, created_by,  updated_by) values (?,?,?,?,?,?) returning id,saving_type_name,saving_type_desc,created_at,updated_at, created_by,  updated_by`
	query = utils.QuerySupport(query)
	if tx != nil {
		err = tx.QueryRow(query, req.Filter.SavingTypeName, req.Filter.SavingTypeDesc,
			dbTime,
			dbTime, "sys", "sys").Scan(
			&result.ID,
			&result.SavingTypeName,
			&result.SavingTypeDesc,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	} else {
		err = ctx.repo.Db.QueryRow(query, req.Filter.SavingTypeName, req.Filter.SavingTypeDesc,
			dbTime,
			dbTime, "sys", "sys").Scan(
			&result.ID,
			&result.SavingTypeName,
			&result.SavingTypeDesc,
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
func (ctx savingRepository) GetSavingType(req models.ReqGetSavingType) (result models.SavingType, err error) {
	query := `select id,
saving_type_name,
saving_type_desc,
created_at, created_by, updated_at, updated_by from saving_types where true`
	if req.Filter.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.Filter.ID)
	}
	if req.Filter.SavingTypeName != "" {
		query += ` and saving_type_name ='` + req.Filter.SavingTypeName + `'`
	}

	err = ctx.repo.Db.QueryRow(query).Scan(&result.ID,
		&result.SavingTypeName,
		&result.SavingTypeDesc,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetSavingTypes(req models.ReqGetSavingType) (result []models.SavingType, err error) {
	query := `select id,
saving_type_name,
saving_type_desc,created_at, created_by, updated_at, updated_by from saving_types where true `
	if req.Filter.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.Filter.ID)
	}
	if req.Filter.SavingTypeName != "" {
		query += ` and saving_type_name ='` + req.Filter.SavingTypeName + `'`
	}
	if req.Lenght != 0 {
		query += ` limit  ` + strconv.Itoa(int(req.Lenght)) + `  offset  ` + strconv.Itoa(int(req.Start))
	} else {
		if req.Order != "" {
			query += `  order by '` + req.Order + `' asc`
		} else {
			query += `  order by saving_type_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	result, err = SavingTypeDataRow(rows)
	if err != nil {
		return result, err
	}
	if len(result) == 0 {
		return result, sql.ErrNoRows
	}
	return result, nil

}
func SavingTypeDataRow(rows *sql.Rows) (result []models.SavingType, err error) {
	for rows.Next() {
		var val models.SavingType
		err := rows.Scan(
			&val.ID,
			&val.SavingTypeName,
			&val.SavingTypeDesc,
			&val.CreatedAt,
			&val.CreatedBy,
			&val.UpdatedAt,
			&val.UpdatedBy,
		)
		if err != nil {
			log.Println("SavingTypeDataRow :: ", err.Error())
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}

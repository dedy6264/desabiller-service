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

func (ctx savingRepository) GetSavingTypeCount(req models.ReqGetSavingType) (result int, err error) {
	query := `select count(id) from saving_types where true `
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.SavingTypeName != "" {
		query += ` and saving_type_name ='` + req.SavingTypeName + `'`
	}

	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		log.Println("GetCount :: ", err.Error())
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
		log.Println("DropSavingType :: ", err.Error())
		return err
	}
	return nil
}
func (ctx savingRepository) UpdateSavingType(req models.ReqGetSavingType, tx *sql.Tx) (err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update saving_types set 
				saving_type_name=$1,
				saving_type_desc=$2,
				updated_at=$3,
				updated_by=$4
				where id = $5
				`
	if tx != nil {
		_, err = tx.Exec(query, req.SavingTypeName, req.SavingTypeDesc, dbTime, "sys", req.ID)
	} else {
		_, err = ctx.repo.Db.Exec(query, req.SavingTypeName, req.SavingTypeDesc, dbTime, "sys", req.ID)
	}
	if err != nil {
		log.Println("UpdateSavingType :: ", err.Error())
		return err
	}
	return nil
}
func (ctx savingRepository) AddSavingType(req models.ReqGetSavingType, tx *sql.Tx) (result models.RespGetSavingType, err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into saving_types (saving_type_name,saving_type_desc,created_at,updated_at, created_by,  updated_by) values ($1,$2,$3,$4,$5,$6) returning id,saving_type_name,saving_type_desc,created_at,updated_at, created_by,  updated_by`
	fmt.Println(query, dbTime)
	if tx != nil {
		err = tx.QueryRow(query, req.SavingTypeName, req.SavingTypeDesc,
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
		err = ctx.repo.Db.QueryRow(query, req.SavingTypeName, req.SavingTypeDesc,
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
		log.Println("AddSavingType :: ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetSavingType(req models.ReqGetSavingType) (result models.RespGetSavingType, err error) {
	query := `select id,
saving_type_name,
saving_type_desc,
created_at, created_by, updated_at, updated_by from saving_types where true`
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.SavingTypeName != "" {
		query += ` and saving_type_name ='` + req.SavingTypeName + `'`
	}

	err = ctx.repo.Db.QueryRow(query).Scan(&result.ID,
		&result.SavingTypeName,
		&result.SavingTypeDesc,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy)
	if err != nil {
		log.Println("GetSavingType :: ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetSavingTypes(req models.ReqGetSavingType) (result []models.RespGetSavingType, err error) {
	query := `select id,
saving_type_name,
saving_type_desc,created_at, created_by, updated_at, updated_by from saving_types where true `
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.SavingTypeName != "" {
		query += ` and saving_type_name ='` + req.SavingTypeName + `'`
	}

	if req.Filter.Length != 0 {
		query += ` limit  ` + strconv.Itoa(req.Filter.Length) + `  offset  ` + strconv.Itoa(req.Filter.Start)
	} else {
		if req.Filter.OrderBy != "" {
			query += `  order by '` + req.Filter.OrderBy + `' asc`
		} else {
			query += `  order by saving_type_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("GetSavingTypes :: ", err.Error())
		return result, err
	}
	defer rows.Close()
	result, err = SavingTypeDataRow(rows)
	if err != nil {
		log.Println("GetSavingTypes :: ", err.Error())
		return result, err
	}
	return result, nil

}
func SavingTypeDataRow(rows *sql.Rows) (result []models.RespGetSavingType, err error) {
	for rows.Next() {
		var val models.RespGetSavingType
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

package savingrepo

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/models"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func (ctx savingRepository) GetCifCount(req models.ReqGetCif) (result int, err error) {
	query := `select count(id) from cifs where true `
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.CifNik != "" {
		query += ` and cif_nik ='` + req.CifNik + `'`
	}
	if req.CifName != "" {
		query += ` and cif_name like '%` + req.CifName + `%'`
	}
	if req.CifPhone != "" {
		query += ` and cif_phone='` + req.CifPhone + `'`
	}

	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		log.Println("GetCount :: ", err.Error())
		return 0, err
	}
	return result, nil
}
func (ctx savingRepository) DropCif(id int, tx *sql.Tx) (err error) {
	query := `delete from cifs where id = $1`
	if tx != nil {
		_, err = tx.Exec(query, id)
	} else {
		_, err = ctx.repo.Db.Exec(query, id)
	}
	if err != nil {
		log.Println("DropCif :: ", err.Error())
		return err
	}
	return nil
}
func (ctx savingRepository) UpdateCif(req models.ReqGetCif, tx *sql.Tx) (err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update cifs set 
				cif_name=$1,
				cif_nik=$2,
				cif_phone=$3,
				cif_email=$4,
				cif_address=$5,
				updated_at = $6,
				updated_by =$7
				where id = $8
				`
	if tx != nil {
		_, err = tx.Exec(query, req.CifName, req.CifNik, req.CifPhone,
			req.CifEmail,
			req.CifAddress, dbTime, "sys", req.ID)
	} else {
		_, err = ctx.repo.Db.Exec(query, req.CifName, req.CifNik, req.CifPhone,
			req.CifEmail,
			req.CifAddress, dbTime, "sys", req.ID)
	}
	if err != nil {
		log.Println("UpdateCif :: ", err.Error())
		return err
	}
	return nil
}
func (ctx savingRepository) AddCif(req models.ReqGetCif, tx *sql.Tx) (result models.RespGetCif, err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into cifs (cif_name,cif_nik,cif_phone,cif_email,cif_address,created_at,updated_at, created_by,  updated_by) values ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id,cif_name,cif_nik,cif_phone,cif_email,cif_address,created_at,updated_at, created_by,  updated_by`
	fmt.Println(query, dbTime)
	if tx != nil {
		err = tx.QueryRow(query, req.CifName, req.CifNik, req.CifPhone,
			req.CifEmail,
			req.CifAddress,
			dbTime,
			dbTime, "sys", "sys").Scan(
			&result.ID,
			&result.CifName,
			&result.CifNik,
			&result.CifPhone,
			&result.CifEmail,
			&result.CifAddress,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	} else {
		err = ctx.repo.Db.QueryRow(query, req.CifName, req.CifNik, req.CifPhone,
			req.CifEmail,
			req.CifAddress,
			dbTime,
			dbTime, "sys", "sys").Scan(
			&result.ID,
			&result.CifName,
			&result.CifNik,
			&result.CifPhone,
			&result.CifEmail,
			&result.CifAddress,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	}

	if err != nil {
		log.Println("AddCif :: ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetCif(req models.ReqGetCif) (result models.RespGetCif, err error) {
	query := `select id,
cif_name,
cif_nik,
cif_phone,
cif_email,
cif_address,created_at, created_by, updated_at, updated_by from cifs where true`
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.CifNik != "" {
		query += ` and cif_nik ='` + req.CifNik + `'`
	}
	if req.CifName != "" {
		query += ` and cif_name like '%` + req.CifName + `%'`
	}
	if req.CifPhone != "" {
		query += ` and cif_phone='` + req.CifPhone + `'`
	}
	err = ctx.repo.Db.QueryRow(query).Scan(&result.ID,
		&result.CifName,
		&result.CifNik,
		&result.CifPhone,
		&result.CifEmail,
		&result.CifAddress,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy)
	if err != nil {
		log.Println("GetCif :: ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetCifs(req models.ReqGetCif) (result []models.RespGetCif, err error) {
	query := `select id,
cif_name,
cif_nik,
cif_phone,
cif_email,
cif_address,created_at, created_by, updated_at, updated_by from cifs where true `
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.CifNik != "" {
		query += ` and cif_nik ='` + req.CifNik + `'`
	}
	if req.CifName != "" {
		query += ` and cif_name like '%` + strings.ToUpper(req.CifName) + `%'`
	}
	if req.CifPhone != "" {
		query += ` and cif_phone='` + req.CifPhone + `'`
	}
	if req.Filter.Length != 0 {
		query += ` limit  ` + strconv.Itoa(req.Filter.Length) + `  offset  ` + strconv.Itoa(req.Filter.Start)
	} else {
		if req.Filter.OrderBy != "" {
			query += `  order by '` + req.Filter.OrderBy + `' asc`
		} else {
			query += `  order by cif_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("GetCifs :: ", err.Error())
		return result, err
	}
	defer rows.Close()
	result, err = CifDataRow(rows)
	if err != nil {
		log.Println("GetCifs :: ", err.Error())
		return result, err
	}
	return result, nil

}
func CifDataRow(rows *sql.Rows) (result []models.RespGetCif, err error) {
	for rows.Next() {
		var val models.RespGetCif
		err := rows.Scan(
			&val.ID,
			&val.CifName,
			&val.CifNik,
			&val.CifPhone,
			&val.CifEmail,
			&val.CifAddress,
			&val.CreatedAt,
			&val.CreatedBy,
			&val.UpdatedAt,
			&val.UpdatedBy,
		)
		if err != nil {
			log.Println("CifDataRow :: ", err.Error())
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}

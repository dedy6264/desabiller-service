package savingrepo

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/models"
	"desabiller/utils"
	"strconv"
	"time"
)

func (ctx savingRepository) GetCifCount(req models.ReqGetCIF) (result int, err error) {
	query := `select count(id) from cifs where true `
	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
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
		return err
	}
	return nil
}
func (ctx savingRepository) UpdateCif(req models.ReqGetCIF, tx *sql.Tx) (err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update cifs set 
				cif_name=?,
				cif_no_id=?,
				cif_type_id=?,
				cif_id_index=?,
				cif_phone=?,
				cif_email=?,
				cif_address=?,
				updated_at = ?,
				updated_by =?
				where id = ?
				`
	query = utils.QuerySupport(query)
	if tx != nil {
		_, err = tx.Exec(query, req.Filter.CifName,
			req.Filter.CifNoID,
			req.Filter.CifTypeID,
			req.Filter.CifIDIndex,
			req.Filter.CifPhone,
			req.Filter.CifEmail,
			req.Filter.CifAddress, dbTime, "sys", req.Filter.ID)
	} else {
		_, err = ctx.repo.Db.Exec(query, req.Filter.CifName,
			req.Filter.CifNoID,
			req.Filter.CifTypeID,
			req.Filter.CifIDIndex,
			req.Filter.CifPhone,
			req.Filter.CifEmail,
			req.Filter.CifAddress, dbTime, "sys", req.Filter.ID)
	}
	if err != nil {
		return err
	}
	return nil
}
func (ctx savingRepository) AddCif(req models.ReqGetCIF, tx *sql.Tx) (result models.CIF, err error) {

	query := `insert into cifs (
	cif_name,cif_no_id,
	cif_type_id,
	cif_id_index,cif_phone,cif_email,cif_address,created_at,updated_at, created_by,  updated_by) values (?,?,?,?,?,?,?,?,?,?,?) returning id,cif_name,cif_no_id,
	cif_type_id,cif_phone,cif_email,cif_address,created_at,updated_at, created_by,  updated_by`
	query = utils.QuerySupport(query)
	// fmt.Println(query, dbTime)
	if tx != nil {
		err = tx.QueryRow(query, req.Filter.CifName,
			req.Filter.CifNoID,
			req.Filter.CifTypeID,
			req.Filter.CifIDIndex,
			req.Filter.CifPhone,
			req.Filter.CifEmail,
			req.Filter.CifAddress,
			req.Filter.CreatedAt,
			req.Filter.UpdatedAt, req.Filter.CreatedBy, req.Filter.UpdatedBy).Scan(
			&result.ID,
			&result.CifName,
			&result.CifNoID,
			&result.CifTypeID,
			&result.CifPhone,
			&result.CifEmail,
			&result.CifAddress,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	} else {
		err = ctx.repo.Db.QueryRow(query, req.Filter.CifName,
			req.Filter.CifNoID,
			req.Filter.CifTypeID,
			req.Filter.CifIDIndex,
			req.Filter.CifPhone,
			req.Filter.CifEmail,
			req.Filter.CifAddress,
			req.Filter.CreatedAt,
			req.Filter.UpdatedAt, req.Filter.CreatedBy, req.Filter.UpdatedBy).Scan(
			&result.ID,
			&result.CifName,
			&result.CifNoID,
			&result.CifTypeID,
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
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetCif(req models.ReqGetCIF) (result models.CIF, err error) {
	query := `select id,
cif_name,
cif_no_id,
cif_type_id,
cif_phone,
cif_email,
cif_address,created_at, created_by, updated_at, updated_by from cifs where true`
	if req.Filter.ID != 0 {
		query += ` and id = ` + strconv.Itoa(int(req.Filter.ID))
	}
	if req.Filter.CifNoID != "" {
		query += ` and cif_no_id ='` + req.Filter.CifNoID + `'`
	}
	if req.Filter.CifName != "" {
		query += ` and cif_name like '%` + req.Filter.CifName + `%'`
	}
	if req.Filter.CifPhone != "" {
		query += ` and cif_phone='` + req.Filter.CifPhone + `'`
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.CifName,
		&result.CifNoID,
		&result.CifTypeID,
		&result.CifPhone,
		&result.CifEmail,
		&result.CifAddress,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetCifs(req models.ReqGetCIF) (result []models.CIF, err error) {
	query := `select id,
cif_name,
cif_no_id,
cif_type_id,
cif_phone,
cif_email,
cif_address,created_at, created_by, updated_at, updated_by from cifs where true `
	if req.Filter.ID != 0 {
		query += ` and id = ` + strconv.Itoa(int(req.Filter.ID))
	}
	if req.Filter.CifNoID != "" {
		query += ` and cif_no_id ='` + req.Filter.CifNoID + `'`
	}
	if req.Filter.CifName != "" {
		query += ` and cif_name like '%` + req.Filter.CifName + `%'`
	}
	if req.Filter.CifPhone != "" {
		query += ` and cif_phone='` + req.Filter.CifPhone + `'`
	}
	if req.Lenght != 0 {
		query += ` limit  ` + strconv.Itoa(int(req.Lenght)) + `  offset  ` + strconv.Itoa(int(req.Start))
	} else {
		if req.Order != "" {
			query += `  order by '` + req.Order + `' asc`
		} else {
			query += `  order by cif_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	result, err = CifDataRow(rows)
	if err != nil {
		return result, err
	}
	if len(result) == 0 {
		return result, sql.ErrNoRows
	}
	return result, nil

}
func CifDataRow(rows *sql.Rows) (result []models.CIF, err error) {
	for rows.Next() {
		var val models.CIF
		err := rows.Scan(
			&val.ID,
			&val.CifName,
			&val.CifNoID,
			&val.CifTypeID,
			&val.CifPhone,
			&val.CifEmail,
			&val.CifAddress,
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

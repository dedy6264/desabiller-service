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

func (ctx savingRepository) GetSavingSegmentCount(req models.ReqGetSavingSegment) (result int, err error) {
	query := `select count(a.id) from saving_segments as a 
	join saving_types as b on a.saving_type_id=b.id 
	where true `
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.SavingSegmentName != "" {
		query += ` and a.saving_segment_name ='` + req.SavingSegmentName + `'`
	}
	if req.SavingTypeID != 0 {
		query += ` and a.saving_type_id =` + strconv.Itoa(req.SavingTypeID)
	}

	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		log.Println("GetCount :: ", err.Error())
		return 0, err
	}
	return result, nil
}
func (ctx savingRepository) DropSavingSegment(id int, tx *sql.Tx) (err error) {
	query := `delete from saving_segments where id = $1`
	if tx != nil {
		_, err = tx.Exec(query, id)
	} else {
		_, err = ctx.repo.Db.Exec(query, id)
	}
	if err != nil {
		log.Println("DropSavingSegment :: ", err.Error())
		return err
	}
	return nil
}
func (ctx savingRepository) UpdateSavingSegment(req models.ReqGetSavingSegment, tx *sql.Tx) (err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update saving_segments set 
				saving_segment_name=$1,
				limit_amount=$2,
				saving_type_id=$3,
				updated_at=$4,
				updated_by=$5
				where id = $6
				`
	if tx != nil {
		_, err = tx.Exec(query, req.SavingSegmentName, req.LimitAmount, req.SavingTypeID, dbTime, "sys", req.ID)
	} else {
		_, err = ctx.repo.Db.Exec(query, req.SavingSegmentName, req.LimitAmount, req.SavingTypeID, dbTime, "sys", req.ID)
	}
	if err != nil {
		log.Println("UpdateSavingSegment :: ", err.Error())
		return err
	}
	return nil
}
func (ctx savingRepository) AddSavingSegment(req models.ReqGetSavingSegment, tx *sql.Tx) (result models.RespGetSavingSegment, err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into saving_segments (saving_segment_name,limit_amount,saving_type_id,created_at,updated_at, created_by,  updated_by) values ($1,$2,$3,$4,$5,$6,$7) returning id,saving_segment_name,limit_amount,saving_type_id,created_at,updated_at, created_by,  updated_by`
	fmt.Println(query, dbTime)
	if tx != nil {
		err = tx.QueryRow(query, req.SavingSegmentName, req.LimitAmount, req.SavingTypeID,
			dbTime,
			dbTime, "sys", "sys").Scan(
			&result.ID,
			&result.SavingSegmentName,
			&result.LimitAmount,
			&result.SavingTypeID,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	} else {
		err = ctx.repo.Db.QueryRow(query, req.SavingSegmentName, req.LimitAmount, req.SavingTypeID,
			dbTime,
			dbTime, "sys", "sys").Scan(
			&result.ID,
			&result.SavingSegmentName,
			&result.LimitAmount,
			&result.SavingTypeID,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	}

	if err != nil {
		log.Println("AddSavingSegment :: ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetSavingSegment(req models.ReqGetSavingSegment) (result models.RespGetSavingSegment, err error) {
	query := `select a.id,
a.saving_segment_name,a.limit_amount,a.saving_type_id,b.saving_type_name,
a.created_at, a.created_by, a.updated_at, a.updated_by from saving_segments as a
join saving_types as b on a.saving_type_id=b.id where true`
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.SavingSegmentName != "" {
		query += ` and a.saving_segment_name ='` + req.SavingSegmentName + `'`
	}
	if req.SavingTypeID != 0 {
		query += ` and a.saving_type_id =` + strconv.Itoa(req.SavingTypeID)
	}
	err = ctx.repo.Db.QueryRow(query).Scan(&result.ID,
		&result.SavingSegmentName,
		&result.LimitAmount,
		&result.SavingTypeID,
		&result.SavingTypeName,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy)
	if err != nil {
		log.Println("GetSavingSegment :: ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetSavingSegments(req models.ReqGetSavingSegment) (result []models.RespGetSavingSegment, err error) {
	query := `select a.id,
a.saving_segment_name,
a.limit_amount,
a.saving_type_id,
b.saving_type_name,
a.created_at, a.created_by, a.updated_at, a.updated_by from saving_segments as a
join saving_types as b on a.saving_type_id=b.id where true `
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.SavingSegmentName != "" {
		query += ` and a.saving_segment_name ='` + req.SavingSegmentName + `'`
	}
	if req.SavingTypeID != 0 {
		query += ` and a.saving_type_id =` + strconv.Itoa(req.SavingTypeID)
	}

	if req.Filter.Length != 0 {
		query += ` limit  ` + strconv.Itoa(req.Filter.Length) + `  offset  ` + strconv.Itoa(req.Filter.Start)
	} else {
		// if req.Filter.OrderBy != "" {
		// 	query += `  order by '` + req.Filter.OrderBy + `' asc`
		// } else {
		query += `  order by a.saving_segment_name asc`
		// }
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("GetSavingSegments :: ", err.Error())
		return result, err
	}
	defer rows.Close()
	result, err = SavingSegmentDataRow(rows)
	if err != nil {
		log.Println("GetSavingSegments :: ", err.Error())
		return result, err
	}
	return result, nil

}
func SavingSegmentDataRow(rows *sql.Rows) (result []models.RespGetSavingSegment, err error) {
	for rows.Next() {
		var val models.RespGetSavingSegment
		err := rows.Scan(
			&val.ID,
			&val.SavingSegmentName,
			&val.LimitAmount,
			&val.SavingTypeID,
			&val.SavingTypeName,
			&val.CreatedAt,
			&val.CreatedBy,
			&val.UpdatedAt,
			&val.UpdatedBy,
		)
		if err != nil {
			log.Println("SavingSegmentDataRow :: ", err.Error())
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}

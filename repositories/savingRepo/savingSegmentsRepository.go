package savingrepo

import (
	"database/sql"
	"desabiller/models"
	"desabiller/utils"
	"strconv"
)

func (ctx savingRepository) GetSavingSegmentCount(req models.ReqGetSavingSegment) (result int, err error) {
	query := `select count(a.id) from saving_segments as a 
	join saving_types as b on a.saving_type_id=b.id 
	where true `
	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
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
		return err
	}
	return nil
}
func (ctx savingRepository) UpdateSavingSegment(req models.ReqGetSavingSegment, tx *sql.Tx) (err error) {
	query := `update saving_segments set 
				saving_segment_name=?,
				limit_amount=?,
				saving_type_id=?,
				updated_at=?,
				updated_by=?
				where id = ?
				`
	query = utils.QuerySupport(query)
	if tx != nil {
		_, err = tx.Exec(query, req.Filter.SavingSegmentName, req.Filter.LimitAmount, req.Filter.SavingTypeID, req.Filter.UpdatedAt, req.Filter.UpdatedBy, req.Filter.ID)
	} else {
		_, err = ctx.repo.Db.Exec(query, req.Filter.SavingSegmentName, req.Filter.LimitAmount, req.Filter.SavingTypeID, req.Filter.UpdatedAt, req.Filter.UpdatedBy, req.Filter.ID)
	}
	if err != nil {
		return err
	}
	return nil
}
func (ctx savingRepository) AddSavingSegment(req models.ReqGetSavingSegment, tx *sql.Tx) (result models.SavingSegment, err error) {
	query := `insert into saving_segments (saving_segment_name,limit_amount,saving_type_id,created_at,updated_at, created_by,  updated_by) values (?,?,?,?,?,?,?) returning id,saving_segment_name,limit_amount,saving_type_id,created_at,updated_at, created_by,  updated_by`
	query = utils.QuerySupport(query)
	if tx != nil {
		err = tx.QueryRow(query, req.Filter.SavingSegmentName, req.Filter.LimitAmount, req.Filter.SavingTypeID,
			req.Filter.CreatedAt,
			req.Filter.UpdatedAt, req.Filter.CreatedBy, req.Filter.UpdatedBy).Scan(
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
		err = ctx.repo.Db.QueryRow(query, req.Filter.SavingSegmentName, req.Filter.LimitAmount, req.Filter.SavingTypeID,
			req.Filter.CreatedAt,
			req.Filter.UpdatedAt, req.Filter.CreatedBy, req.Filter.UpdatedBy).Scan(
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
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetSavingSegment(req models.ReqGetSavingSegment) (result models.SavingSegment, err error) {
	query := `select a.id,
a.saving_segment_name,a.limit_amount,a.saving_type_id,b.saving_type_name,
a.created_at, a.created_by, a.updated_at, a.updated_by from saving_segments as a
join saving_types as b on a.saving_type_id=b.id where true`
	if req.Filter.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.Filter.ID)
	}
	if req.Filter.SavingSegmentName != "" {
		query += ` and a.saving_segment_name ='` + req.Filter.SavingSegmentName + `'`
	}
	if req.Filter.SavingTypeID != 0 {
		query += ` and a.saving_type_id =` + strconv.Itoa(req.Filter.SavingTypeID)
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
		return result, err
	}
	return result, nil
}
func (ctx savingRepository) GetSavingSegments(req models.ReqGetSavingSegment) (result []models.SavingSegment, err error) {
	query := `select a.id,
a.saving_segment_name,
a.limit_amount,
a.saving_type_id,
b.saving_type_name,
a.created_at, a.created_by, a.updated_at, a.updated_by from saving_segments as a
join saving_types as b on a.saving_type_id=b.id where true `
	if req.Filter.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.Filter.ID)
	}
	if req.Filter.SavingSegmentName != "" {
		query += ` and a.saving_segment_name ='` + req.Filter.SavingSegmentName + `'`
	}
	if req.Filter.SavingTypeID != 0 {
		query += ` and a.saving_type_id =` + strconv.Itoa(req.Filter.SavingTypeID)
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
	result, err = SavingSegmentDataRow(rows)
	if err != nil {
		return result, err
	}
	if len(result) == 0 {
		return result, sql.ErrNoRows
	}
	return result, nil

}
func SavingSegmentDataRow(rows *sql.Rows) (result []models.SavingSegment, err error) {
	for rows.Next() {
		var val models.SavingSegment
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
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}

package productrepo

import (
	"desabiller/configs"
	"desabiller/models"
	"log"
	"strconv"
	"time"
)

func (ctx product) AddSegment(req models.ReqListSegment) (result models.ResListSegment, status bool) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` insert into segments (
		segment_name,
		created_at,
		updated_at,
		created_by,
		updated_by)
		values(
			$1,$2,$3,$4,$5
		) returning id `
	err := ctx.repo.Db.QueryRow(query,
		req.SegmentName,
		dbTime,
		dbTime,
		"sys",
		"sys").Scan(&result.ID)
	if err != nil {
		log.Println("Error failed ", err.Error())
		return result, false
	}
	return result, true
}
func (ctx product) GetListSegment(req models.ReqListSegment) (result []models.ResListSegment, status bool) {
	query := `select 
id,
segment_name,
created_at,
updated_at,
created_by,
updated_by
from segments where true 
`
	if req.SegmentName != "" {
		query += ` and segment_name= '` + req.SegmentName + `'`
	}

	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.StartDate != "" {
		query += ` and a.created_at between '` + req.StartDate + `' and '` + req.EndDate + `'`
	}
	if req.Limit != 0 {
		query += ` limit  ` + strconv.Itoa(req.Limit) + `  offset  ` + strconv.Itoa(req.Offset)
	} else {
		if req.OrderBy != "" {
			query += `  order by ` + req.OrderBy + ` asc`
		} else {
			query += `  order by segment_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println(" GetListSegment :: Failed : ", err.Error())
		return result, false
	}
	var val models.ResListSegment
	for rows.Next() {
		err := rows.Scan(
			&val.ID,
			&val.SegmentName,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy)
		if err != nil {
			return result, false
		}
		result = append(result, val)
	}
	if len(result) == 0 {
		return result, false
	}
	return result, false
}
func (ctx product) UpdateSegment(req models.ReqListSegment) (result models.ResListSegment, status bool) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` update segments set
	segment_name=$1,
	updated_at = $2,
	updated_by =$3
	where id = $4 returning id
	`
	err := ctx.repo.Db.QueryRow(query,
		req.SegmentName,
		dbTime,
		"sys",
		req.ID).Scan(&result.ID)
	if err != nil {
		log.Println(" UpdateSegment :: Failed : ", err.Error())
		return result, false
	}
	return result, true
}
func (ctx product) DropSegment(req models.ReqListSegment) (status bool) {
	query := ` delete from segments where id = $1`
	err := ctx.repo.Db.QueryRow(query, req.ID)
	if err.Err() != nil {
		log.Println("Drop product biller :: ", err.Err())
		return false
	}
	return true
	// t := time.Now()

	// dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)

	// query := `update product_biller_providers set
	// 				deleted_at = $1,
	// 				deleted_by =$2
	// 				where id = $3
	// 				`
	// err := ctx.repo.Db.QueryRow(query, dbTime, "sys", req.ID)
	// fmt.Println(":::", err.Err())
	// if err.Err() != nil {
	// 	log.Println("UpdateSegment :: ", err.Err())
	// 	return false
	// }
	// return true
}

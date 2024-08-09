package productrepo

import (
	"desabiller/configs"
	"desabiller/models"
	"log"
	"strconv"
	"time"
)

func (ctx product) AddProductBiller(req models.ReqGetListProductBiller) (result models.ResGetProductBiller, status bool) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` insert into product_billers (
		product_name,
		product_code,
		product_biller_provider_id,
		is_open,
		product_type_id,
		product_category_id,
		created_at,
		updated_at,
		created_by,
		updated_by)
		values(
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10
		) returning id `
	err := ctx.repo.Db.QueryRow(query,
		req.ProductName,
		req.ProductCode,
		req.ProductProviderId,
		true,
		req.ProductTypeId,
		req.ProductCategoryId,
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
func (ctx product) GetListProductBiller(req models.ReqGetListProductBiller) (result []models.ResGetProductBiller, status bool) {
	query := `select 
id,
product_name,
product_code,
product_biller_provider_id,
is_open,
product_type_id,
product_category_id,
created_at,
updated_at,
created_by,
updated_by
from product_billers where true 
`
	if req.ProductName != "" {
		query += ` and product_name= '` + req.ProductName + `'`
	}
	if req.ProductCode != "" {
		query += ` and product_code= '` + req.ProductCode + `'`
	}
	if req.ProductTypeId != 0 {
		query += ` and product_type_id = ` + strconv.Itoa(req.ProductTypeId)
	}
	if req.ProductCategoryId != 0 {
		query += ` and product_category_id = ` + strconv.Itoa(req.ProductCategoryId)
	}
	if req.ProductProviderId != 0 {
		query += ` and product_biller_provider_id = ` + strconv.Itoa(req.ProductProviderId)
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
			query += `  order by product_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println(" GetListProductBiller :: Failed : ", err.Error())
		return result, false
	}
	var val models.ResGetProductBiller
	for rows.Next() {
		err := rows.Scan(
			&val.ID,
			&val.ProductName,
			&val.ProductCode,
			&val.ProductProviderId,
			&val.IsOpen,
			&val.ProductTypeId,
			&val.ProductCategoryId,
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
func (ctx product) UpdateProductBiller(req models.ReqGetListProductBiller) (result models.ResGetProductBiller, status bool) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` update product_billers set
	product_name=$1,
	product_code=$2,
	product_biller_provider_id=$3,
	is_open=$4,
	product_type_id=$5,
	product_category_id=$6,
	updated_at = $7,
	updated_by =$8
	where id = $9 returning id
	`
	err := ctx.repo.Db.QueryRow(query,
		req.ProductName,
		req.ProductCode,
		req.ProductProviderId,
		req.IsOpen,
		req.ProductTypeId,
		req.ProductCategoryId,
		dbTime,
		"sys",
		req.ID).Scan(&result.ID)
	if err != nil {
		log.Println(" UpdateProductBiller :: Failed : ", err.Error())
		return result, false
	}
	return result, true
}
func (ctx product) DropProductBiller(req models.ReqGetListProductBiller) (status bool) {
	query := ` delete from product_billers where id = $1`
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
	// 	log.Println("UpdateProductBiller :: ", err.Err())
	// 	return false
	// }
	// return true
}

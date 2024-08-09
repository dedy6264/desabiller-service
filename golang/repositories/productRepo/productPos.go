package productrepo

import (
	"desabiller/configs"
	"desabiller/models"
	"log"
	"strconv"
	"time"
)

func (ctx product) GetListProductPosMany(productId []int, merchantId int) (result []models.ResGetProductPos, status bool) {
	query := `select 
	id,
	product_name,
	product_code,
	product_price_provider,
	is_open,
	product_type_id,
	product_category_id,
	merchant_id,
	merchant_name,
	product_price,
	created_at,
	updated_at,
	created_by,
	updated_by
	from product_pos where id in $1 and merchant_id = $2
	`
	rows, err := ctx.repo.Db.Query(query, productId)
	if err != nil {
		return result, false
	}
	defer rows.Close()
	for rows.Next() {
		var val models.ResGetProductPos
		err := rows.Scan(
			&val.ID,
			&val.ProductName,
			&val.ProductCode,
			&val.ProductPriceProvider,
			&val.IsOpen,
			&val.ProductTypeId,
			&val.ProductCategoryId,
			&val.MerchantId,
			&val.MerchantName,
			&val.ProductPrice,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy,
		)
		if err != nil {
			return result, false
		}
		result = append(result, val)
	}
	return result, true
}
func (ctx product) AddProductPos(req models.ReqGetListProductPos) (result models.ResGetProductPos, status bool) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` insert into product_pos (
		product_name,
		product_code,
		product_price_provider,
		is_open,
		product_type_id,
		product_category_id,
		merchant_id,
		merchant_name,
		product_price,
		created_at,
		updated_at,
		created_by,
		updated_by)
		values(
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13
		) returning id `
	err := ctx.repo.Db.QueryRow(query,
		req.ProductName,
		req.ProductCode,
		req.ProductPriceProvider,
		true,
		req.ProductTypeId,
		req.ProductCategoryId,
		req.MerchantId,
		req.MerchantName,
		req.ProductPrice,
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
func (ctx product) GetListProductPos(req models.ReqGetListProductPos) (result []models.ResGetProductPos, status bool) {
	query := `select 
id,
product_name,
product_code,
product_price_provider,
is_open,
product_type_id,
product_category_id,
merchant_id,
merchant_name,
product_price,
created_at,
updated_at,
created_by,
updated_by
from product_pos where true 
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
	if req.MerchantId != 0 {
		query += ` and merchant_id = ` + strconv.Itoa(req.MerchantId)
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
		log.Println(" GetListProductPos :: Failed : ", err.Error())
		return result, false
	}
	var val models.ResGetProductPos
	for rows.Next() {
		err := rows.Scan(
			&val.ID,
			&val.ProductName,
			&val.ProductCode,
			&val.ProductPriceProvider,
			&val.IsOpen,
			&val.ProductTypeId,
			&val.ProductCategoryId,
			&val.MerchantId,
			&val.MerchantName,
			&val.ProductPrice,
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
	return result, true
}
func (ctx product) UpdateProductPos(req models.ReqGetListProductPos) (result models.ResGetProductPos, status bool) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` update product_pos set
	product_name=$1,
	product_code=$2,
	product_price_provider=$3,
	is_open=$4,
	product_type_id=$5,
	product_category_id=$6,
	merchant_id=$7,
	merchant_name=$8,
	product_price=$9,
	updated_at = $10,
	updated_by =$11
	where id = $12 returning id
	`
	err := ctx.repo.Db.QueryRow(query,
		req.ProductName,
		req.ProductCode,
		req.ProductPriceProvider,
		req.IsOpen,
		req.ProductTypeId,
		req.ProductCategoryId,
		req.MerchantId,
		req.MerchantName,
		req.ProductPrice,
		dbTime,
		"sys",
		req.ID).Scan(&result.ID)
	if err != nil {
		log.Println(" UpdateProductPos :: Failed : ", err.Error())
		return result, false
	}
	return result, true
}
func (ctx product) DropProductPos(req models.ReqGetListProductPos) (status bool) {
	query := ` delete from product_pos where id = $1`
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
	// 	log.Println("UpdateProductPos :: ", err.Err())
	// 	return false
	// }
	// return true
}

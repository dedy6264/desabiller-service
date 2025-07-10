package productrepo

import (
	"desabiller/configs"
	"desabiller/models"
	"log"
	"strconv"
	"time"
)

func (ctx product) AddSegmentProduct(req models.ReqListSegmentProduct) (result models.ResListSegmentProduct, status bool) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` insert into segment_products (
		segment_product_prefix,
		segment_id,
		product_biller_id,
		product_biller_provider_id,
		product_price,
		product_admin_fee,
		product_merchant_fee,
		is_open,
		created_at,
		updated_at,
		created_by,
		updated_by)
		values(
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12
		) returning id `
	err := ctx.repo.Db.QueryRow(query,
		req.SegmentProductPrefix,
		req.SegmentId,
		req.ProductBillerId,
		req.ProductBillerProviderId,
		req.ProductPrice,
		req.ProductAdminFee,
		req.ProductMerchantFee,
		req.IsOpen,
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
func (ctx product) GetListSegmentProduct(req models.ReqListSegmentProduct) (result []models.ResListSegmentProduct, status bool) {
	query := `select 
id,
segment_product_prefix,
segment_id,
product_biller_id,
product_biller_provider_id,
product_price,
product_admin_fee,
product_merchant_fee,
is_open,
created_at,
updated_at,
created_by,
updated_by
from segment_products where true 
`
	if req.SegmentProductPrefix != "" {
		query += ` and segment_product_prefix= '` + req.SegmentProductPrefix + `'`
	}
	if req.SegmentId != 0 {
		query += ` and segment_id= ` + strconv.Itoa(req.SegmentId)
	}
	if req.ProductBillerId != 0 {
		query += ` and product_biller_id= ` + strconv.Itoa(req.ProductBillerId)
	}
	if req.ProductBillerProviderId != 0 {
		query += ` and product_biller_provider_id= ` + strconv.Itoa(req.ProductBillerProviderId)
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
			query += `  order by created_at asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println(" GetListSegmentProduct :: Failed : ", err.Error())
		return result, false
	}
	var val models.ResListSegmentProduct
	for rows.Next() {
		err := rows.Scan(
			&val.ID,
			&val.SegmentProductPrefix,
			&val.SegmentId,
			&val.ProductBillerId,
			&val.ProductBillerProviderId,
			&val.ProductPrice,
			&val.ProductAdminFee,
			&val.ProductMerchantFee,
			&val.IsOpen,
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
func (ctx product) UpdateSegmentProduct(req models.ReqListSegmentProduct) (result models.ResListSegmentProduct, status bool) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` update segment_products set
	segment_product_prefix=$1,
	segment_id=$2,
	product_biller_id=$3,
	product_biller_provider_id=$4,
	product_price=$5,
	product_admin_fee=$6,
	product_merchant_fee=$7,
	is_open=$8,
	updated_at = $9,
	updated_by =$10
	where id = $11 returning id
	`
	err := ctx.repo.Db.QueryRow(query,
		req.SegmentProductPrefix,
		req.SegmentId,
		req.ProductBillerId,
		req.ProductBillerProviderId,
		req.ProductPrice,
		req.ProductAdminFee,
		req.ProductMerchantFee,
		req.IsOpen,
		dbTime,
		"sys",
		req.ID).Scan(&result.ID)
	if err != nil {
		log.Println(" UpdateSegmentProduct :: Failed : ", err.Error())
		return result, false
	}
	return result, true
}
func (ctx product) DropSegmentProduct(req models.ReqListSegmentProduct) (status bool) {
	query := ` delete from segment_products where id = $1`
	err := ctx.repo.Db.QueryRow(query, req.ID)
	if err.Err() != nil {
		log.Println("Drop segment product :: ", err.Err())
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

package productrepo

import (
	"desabiller/configs"
	"desabiller/models"
	"log"
	"strconv"
	"time"
)

func (ctx product) AddProductCategory(req models.ReqGetListProductCategory) (result models.ResGetProductCategory, status bool) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` insert into product_categories (
		product_category_name,
		product_category_code,
		merchant_id,
		merchant_name,
		updateable,
		created_at,
		updated_at,
		created_by,
		updated_by)
		values(
			$1,$2,$3,$4,$5,$6,$7,$8,$9
		) returning id,
		product_category_name,
		product_category_code,
		merchant_id,
		merchant_name,
		updateable `
	err := ctx.repo.Db.QueryRow(query, req.ProductCategoryName, req.ProductCategoryCode, req.MerchantId, req.MerchantName,
		true, dbTime, dbTime, "sys", "sys").Scan(&result.ID, &result.ProductCategoryName, &result.ProductCategoryCode,
		&result.MerchantId, &result.MerchantName, &result.Updateable)
	if err != nil {
		log.Println("Error failed ", err.Error())
		return result, false
	}
	return result, true
}
func (ctx product) GetListProductCategory(req models.ReqGetListProductCategory) (result []models.ResGetProductCategory, status bool) {
	query := `select 
a.id,
a.product_category_name,
a.product_category_code,
b.id,
b.merchant_name,
a.updateable,
a.created_at,
a.updated_at,
a.created_by,
a.updated_by,
c.id,
c.client_name
from product_categories as a
join merchants as b on a.merchant_id = b.id
join clients as c on b.client_id=c.id

where true 
`
	if req.Updateable {
		query += ` and a.updateable= true`
	} else {
		query += ` and a.updateable= false`
	}
	if req.ProductCategoryName != "" {
		query += ` and a.product_category_name= '` + req.ProductCategoryName + `'`
	}
	if req.ProductCategoryCode != "" {
		query += ` and a.product_category_code= '` + req.ProductCategoryCode + `'`
	}
	if req.MerchantId != 0 {
		query += ` and b.id = ` + strconv.Itoa(req.MerchantId)
	}
	if req.ClientId != 0 {
		query += ` and c.id = ` + strconv.Itoa(req.ClientId)
	}
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
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
			query += `  order by a.product_category_name asc`
		}
		query += ` limit 100 offset 0`
	}
	log.Println(query)
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println(" GetListProductCategory :: Failed : ", err.Error())
		return result, false
	}
	var val models.ResGetProductCategory
	for rows.Next() {
		err := rows.Scan(&val.ID, &val.ProductCategoryName, &val.ProductCategoryCode, &val.MerchantId,
			&val.MerchantName, &val.Updateable, &val.CreatedAt, &val.UpdatedAt, &val.CreatedBy, &val.UpdatedBy, &val.ClientId, &val.ClientName)
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
func (ctx product) UpdateProductCategory(req models.ReqGetListProductCategory) (result models.ResGetProductCategory, status bool) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` update product_categories set
	product_category_name=$1,
	product_category_code=$2,
	merchant_id=$3,
	merchant_name=$4,
	updateable=$5,
	updated_at = $6,
	updated_by =$7
	where id = $8 returning id
	`
	err := ctx.repo.Db.QueryRow(query, req.ProductCategoryName, req.ProductCategoryCode, req.MerchantId, req.MerchantName,
		req.Updateable, dbTime, "sys", req.ID).Scan(&result.ID)
	if err != nil {
		log.Println(" UpdateProductCategory :: Failed : ", err.Error())
		return result, false
	}
	return result, true
}
func (ctx product) DropProductCategory(req models.ReqGetListProductCategory) (status bool) {
	// query := ` delete from merchants where id = $1`
	// err := ctx.repo.Db.QueryRow(query, req.ID)
	// if err != nil {
	// 	log.Println("UpdateUser :: ", err.Err())
	// 	return false
	// }
	// return true
	// t := time.Now()

	// dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)

	query := `delete from product_categories 
					where id = $1 returning id
					`
	err := ctx.repo.Db.QueryRow(query, req.ID)
	if err.Err() != nil {
		log.Println("UpdateProductCategory :: ", err.Err())
		return false
	}
	return true
}

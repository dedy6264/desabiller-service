package productrepo

import (
	"desabiller/configs"
	"desabiller/models"
	"log"
	"strconv"
	"time"
)

func (ctx product) AddProductCategory(req models.ReqGetProductCategory) (result models.RespGetProductCategory, err error) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` insert into product_categories (
		product_category_name,
		created_at,
		updated_at,
		created_by,
		updated_by)
		values(
			$1,$2,$3,$4,$5
		)  `
	_, err = ctx.repo.Db.Exec(query, req.ProductCategoryName, dbTime, dbTime, "sys", "sys")
	if err != nil {
		log.Println("Error failed ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx product) GetProductCategories(req models.ReqGetProductCategory) (result []models.RespGetProductCategory, err error) {
	query := `select
	id,
product_category_name,
created_at,
updated_at,
created_by,
updated_by
from product_categories 
where true
`
	if req.ProductCategoryName != "" {
		query += ` and product_category_name= '` + req.ProductCategoryName + `'`
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}

	if req.Filter.Limit != 0 {
		query += ` limit  ` + strconv.Itoa(req.Filter.Limit) + `  offset  ` + strconv.Itoa(req.Filter.Offset)
	} else {
		if req.Filter.OrderBy != "" {
			query += `  order by ` + req.Filter.OrderBy + ` asc`
		} else {
			query += `  order by product_category_name asc`
		}
		query += ` limit 100 offset 0`
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println(" GetProductCategory :: Failed : ", err.Error())
		return result, err
	}
	defer rows.Close()
	var val models.RespGetProductCategory
	for rows.Next() {
		err := rows.Scan(&val.ID, &val.ProductCategoryName, &val.CreatedAt, &val.UpdatedAt, &val.CreatedBy, &val.UpdatedBy)
		if err != nil {
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}
func (ctx product) UpdateProductCategory(req models.ReqGetProductCategory) (result models.RespGetProductCategory, err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` update product_categories set
	product_category_name=$1,
	updated_at = $2,
	updated_by =$3
	where id = $4 returning id
	`
	_, err = ctx.repo.Db.Exec(query, req.ProductCategoryName, dbTime, "sys", req.ID)
	if err != nil {
		log.Println(" UpdateProductCategory :: Failed : ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx product) DropProductCategory(req models.ReqGetProductCategory) (err error) {
	query := `delete from product_categories
					where id = $1 returning id
					`
	_, err = ctx.repo.Db.Exec(query, req.ID)
	if err != nil {
		log.Println("UpdateProductCategory :: ", err)
		return err
	}
	return nil
}
func (ctx product) GetProductCategory(req models.ReqGetProductCategory) (result models.RespGetProductCategory, err error) {
	query := `select
	id,
product_category_name,
created_at,
updated_at,
created_by,
updated_by
from product_categories 
where true
`
	if req.ProductCategoryName != "" {
		query += ` and product_category_name= '` + req.ProductCategoryName + `'`
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.ProductCategoryName,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.CreatedBy,
		&result.UpdatedBy,
	)
	if err != nil {
		log.Println(" GetProductCategory :: Failed : ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx product) GetProductCategoryCount(req models.ReqGetProductCategory) (result int, err error) {
	query := `select count(id)
from product_categories 
where true
`
	if req.ProductCategoryName != "" {
		query += ` and product_category_name= '` + req.ProductCategoryName + `'`
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result,
	)
	if err != nil {
		log.Println(" GetProductCategory :: Failed : ", err.Error())
		return result, err
	}
	return result, nil
}

package productrepo

import (
	"database/sql"
	"desabiller/models"
	"desabiller/utils"
	"strconv"
)

func (ctx product) AddProductCategory(req models.ReqGetProductCategory) (result models.ProductCategory, err error) {
	query := ` insert into product_categories (
		product_category_name,
		created_at,
		updated_at,
		created_by,
		updated_by)
		values(
			?,?,?,?,?
		)  `
	query = utils.QuerySupport(query)
	_, err = ctx.repo.Db.Exec(query, req.Filter.ProductCategoryName, req.Filter.CreatedAt, req.Filter.UpdatedAt, req.Filter.CreatedBy, req.Filter.UpdatedBy)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx product) GetProductCategories(req models.ReqGetProductCategory) (result []models.ProductCategory, err error) {

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
	if req.Filter.ProductCategoryName != "" {
		query += ` and product_category_name= '` + req.Filter.ProductCategoryName + `'`
	}
	if req.Filter.ID != 0 {
		query += ` and id = ` + strconv.Itoa(int(req.Filter.ID))
	}
	if req.Lenght != 0 {
		query += ` limit  ` + strconv.Itoa(int(req.Lenght)) + `  offset  ` + strconv.Itoa(int(req.Start))
	} else {
		if req.Order != "" {
			query += `  order by '` + req.Order + `' asc`
		} else {
			query += `  order by product_category_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	var val models.ProductCategory
	for rows.Next() {
		err := rows.Scan(&val.ID, &val.ProductCategoryName, &val.CreatedAt, &val.UpdatedAt, &val.CreatedBy, &val.UpdatedBy)
		if err != nil {
			return result, err
		}
		result = append(result, val)
	}
	if len(result) == 0 {
		return result, sql.ErrNoRows
	}
	return result, nil
}
func (ctx product) UpdateProductCategory(req models.ReqGetProductCategory) (result models.ProductCategory, err error) {

	query := ` update product_categories set
	product_category_name=?,
	updated_at = ?,
	updated_by =?
	where id = ? returning id
	`
	query = utils.QuerySupport(query)
	_, err = ctx.repo.Db.Exec(query, req.Filter.ProductCategoryName, req.Filter.UpdatedAt, req.Filter.UpdatedBy, req.Filter.ID)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx product) DropProductCategory(req models.ReqGetProductCategory) (err error) {
	query := `delete from product_categories
					where id = $1 returning id
					`
	_, err = ctx.repo.Db.Exec(query, req.Filter.ID)
	if err != nil {
		return err
	}
	return nil
}
func (ctx product) GetProductCategory(req models.ReqGetProductCategory) (result models.ProductCategory, err error) {
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
	if req.Filter.ProductCategoryName != "" {
		query += ` and product_category_name= '` + req.Filter.ProductCategoryName + `'`
	}
	if req.Filter.ID != 0 {
		query += ` and id = ` + strconv.Itoa(int(req.Filter.ID))
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
		return result, err
	}
	return result, nil
}
func (ctx product) GetProductCategoryCount(req models.ReqGetProductCategory) (result int, err error) {
	query := `select count(id)
from product_categories 
where true
`
	if req.Filter.ProductCategoryName != "" {
		query += ` and product_category_name= '` + req.Filter.ProductCategoryName + `'`
	}
	if req.Filter.ID != 0 {
		query += ` and id = ` + strconv.Itoa(int(req.Filter.ID))
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result,
	)
	if err != nil {
		return result, err
	}
	return result, nil
}

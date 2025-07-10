package productrepo

import (
	"database/sql"
	"desabiller/models"
	"desabiller/utils"
	"log"
	"strconv"
)

func (ctx product) AddProductType(req models.ReqGetProductType) (result models.ProductType, err error) {
	query := ` insert into product_types (
		product_type_name,
		created_at,
		updated_at,
		created_by,
		updated_by)
		values(
			?,?,?,?,?
		)  `
	query = utils.QuerySupport(query)
	_, err = ctx.repo.Db.Exec(query, req.Filter.ProductTypeName, req.Filter.CreatedAt, req.Filter.UpdatedAt, req.Filter.CreatedBy, req.Filter.UpdatedBy)
	if err != nil {
		log.Println("Error failed ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx product) GetProductTypes(req models.ReqGetProductType) (result []models.ProductType, err error) {
	query := `select
	id,
	product_type_name,
	created_at,
	updated_at,
	created_by,
	updated_by
	from product_types 
	where true
	`
	if req.Filter.ProductTypeName != "" {
		query += ` and product_type_name= '` + req.Filter.ProductTypeName + `'`
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
			query += `  order by product_type_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	var val models.ProductType
	for rows.Next() {
		err := rows.Scan(&val.ID, &val.ProductTypeName, &val.CreatedAt, &val.UpdatedAt, &val.CreatedBy, &val.UpdatedBy)
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
func (ctx product) UpdateProductType(req models.ReqGetProductType) (result models.ProductType, err error) {
	query := ` update product_types set
	product_type_name=?,
	updated_at = ?,
	updated_by =?
	where id = ? returning id
	`
	query = utils.QuerySupport(query)
	_, err = ctx.repo.Db.Exec(query, req.Filter.ProductTypeName, req.Filter.UpdatedAt, req.Filter.UpdatedBy, int(req.Filter.ID))
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx product) DropProductType(req models.ReqGetProductType) (err error) {
	query := `delete from product_types
					where id = $1 returning id
					`
	_, err = ctx.repo.Db.Exec(query, int(req.Filter.ID))
	if err != nil {
		return err
	}
	return nil
}
func (ctx product) GetProductType(req models.ReqGetProductType) (result models.ProductType, err error) {
	query := `select
	id,
product_type_name,
created_at,
updated_at,
created_by,
updated_by
from product_types 
where true
`
	if req.Filter.ProductTypeName != "" {
		query += ` and product_type_name= '` + req.Filter.ProductTypeName + `'`
	}
	if req.Filter.ID != 0 {
		query += ` and id = ` + strconv.Itoa(int(req.Filter.ID))
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.ProductTypeName,
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
func (ctx product) GetProductTypeCount(req models.ReqGetProductType) (result int, err error) {
	query := `select count(*)
	from product_types 
	where true
	`
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result,
	)
	if err != nil {
		return result, err
	}
	return result, nil
}

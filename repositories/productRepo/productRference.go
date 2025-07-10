package productrepo

import (
	"database/sql"
	"desabiller/models"
	"desabiller/utils"
	"strconv"
)

func (ctx product) AddProductReference(req models.ReqGetProductReference) (result models.ProductReference, err error) {
	query := ` insert into product_references (
		product_reference_name,
		product_reference_code,
		created_at,
		updated_at,
		created_by,
		updated_by)
		values(
			?,?,?,?,?,?,?,?,?,?,?,?,?,?
		)  `
	query = utils.QuerySupport(query)
	_, err = ctx.repo.Db.Exec(query,
		req.Filter.ProductReferenceName,
		req.Filter.ProductReferenceCode,
		req.Filter.CreatedAt,
		req.Filter.UpdatedAt,
		req.Filter.CreatedBy,
		req.Filter.UpdatedBy)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx product) GetProductReferences(req models.ReqGetProductReference) (result []models.ProductReference, err error) {

	query := `select
	id,
	product_reference_name,
	product_reference_code,
	created_at,
	updated_at,
	created_by,
	updated_by
	from product_references 
	where true
	`
	if req.Filter.ID != 0 {
		query += ` and id = ` + strconv.Itoa(int(req.Filter.ID))
	}
	if req.Filter.ProductReferenceCode != "" {
		query += ` and product_reference_code = '` + req.Filter.ProductReferenceCode + `'`
	}
	if req.Filter.ProductReferenceName != "" {
		query += ` and product_reference_name = '` + req.Filter.ProductReferenceName + `'`
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
	for rows.Next() {
		var val models.ProductReference
		err := rows.Scan(
			&val.ID,
			&val.ProductReferenceCode,
			&val.ProductReferenceName,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy,
		)
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
func (ctx product) GetProductReferenceCount(req models.ReqGetProductReference) (result int, err error) {
	query := `select count(id)
	from product_references
	`

	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx product) GetProduct(req models.ReqGetProductReference) (result models.ProductReference, err error) {
	query := `select
	id,
	product_reference_name,
	product_reference_code,
	created_at,
	updated_at,
	created_by,
	updated_by
	from product_references 
	where true
	`
	if req.Filter.ID != 0 {
		query += ` and id = ` + strconv.Itoa(int(req.Filter.ID))
	}
	if req.Filter.ProductReferenceCode != "" {
		query += ` and product_reference_code = '` + req.Filter.ProductReferenceCode + `'`
	}
	if req.Filter.ProductReferenceName != "" {
		query += ` and product_reference_name = '` + req.Filter.ProductReferenceName + `'`
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
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.ProductReferenceCode,
		&result.ProductReferenceName,
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
func (ctx product) UpdateProductReference(req models.ReqGetProductReference) (result models.ProductReference, err error) {

	query := ` update product_references set
	product_reference_name=?,
	product_reference_code=?,
	updated_at=?,
	updated_by=?
	where id =?
	`
	query = utils.QuerySupport(query)
	_, err = ctx.repo.Db.Exec(query,
		req.Filter.ProductReferenceName,
		req.Filter.ProductReferenceCode,
		req.Filter.UpdatedAt,
		req.Filter.UpdatedBy)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx product) DropProductReference(req models.ReqGetProductReference) (err error) {
	query := ` delete from product_references where id = $1`
	_, err = ctx.repo.Db.Exec(query, req.Filter.ID)
	if err != nil {
		return err
	}
	return nil
}

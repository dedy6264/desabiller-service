package productrepo

import (
	"database/sql"
	"desabiller/models"
	"desabiller/utils"
	"log"
	"strconv"
)

const fieldInsert = `
		product_name,
		product_code,
		product_price,
		product_admin_fee,
		product_merchant_fee,
		product_provider_name,
		product_provider_code,
		product_provider_price,
		product_provider_admin_fee,
		product_provider_merchant_fee,
		product_category_id,
		product_type_id,
		product_type_name,
		product_reference_id,
		product_reference_code,
		product_denom,
		created_by,
		created_at,
		updated_by,
		updated_at`

func (ctx product) AddProduct(req models.ReqGetProduct) (result models.Product, err error) {
	query := ` insert into products (
		` + fieldInsert + `
		)
		values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)  `
	query = utils.QuerySupport(query)
	_, err = ctx.repo.Db.Exec(query,
		req.Filter.ProductName,
		req.Filter.ProductCode,
		req.Filter.ProductPrice,
		req.Filter.ProductAdminFee,
		req.Filter.ProductMerchantFee,
		req.Filter.ProductProviderName,
		req.Filter.ProductProviderCode,
		req.Filter.ProductProviderPrice,
		req.Filter.ProductProviderAdminFee,
		req.Filter.ProductProviderMerchantFee,
		req.Filter.ProductCategoryID,
		req.Filter.ProductTypeID,
		req.Filter.ProductTypeName,
		req.Filter.ProductReferenceID,
		req.Filter.ProductReferenceCode,
		req.Filter.ProductDenom,
		req.Filter.CreatedBy,
		req.Filter.CreatedAt,
		req.Filter.UpdatedBy,
		req.Filter.UpdatedAt,
	)
	if err != nil {
		log.Println("Error failed ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx product) GetProducts(req models.ReqGetProduct) (result []models.Product, err error) {

	query := `select
	` + fieldInsert + `
	from products 
	where true
	`
	if req.Filter.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.Filter.ID)
	}
	if req.Filter.ProductName != "" {
		query += ` and product_name = '` + req.Filter.ProductName + `'`
	}
	if req.Filter.ProductCategoryID != 0 {
		query += ` and product_category_id = ` + strconv.Itoa(req.Filter.ProductCategoryID)
	}
	if req.Filter.ProductTypeID != 0 {
		query += ` and product_type_id = ` + strconv.Itoa(req.Filter.ProductTypeID)
	}
	if req.Filter.ProviderID != 0 {
		query += ` and provider_id = ` + strconv.Itoa(req.Filter.ProviderID)
	}
	if req.Filter.ProductReferenceID != 0 {
		query += ` and product_reference_id = ` + strconv.Itoa(req.Filter.ProductReferenceID)
	}
	if req.Filter.ProductReferenceCode != "" {
		query += ` and product_reference_code = '` + req.Filter.ProductReferenceCode + `'`
	}
	if req.Filter.ProductProviderID != 0 {
		query += ` and product_provider_id = ` + strconv.Itoa(req.Filter.ProductProviderID)
	}
	if req.Lenght != 0 {
		query += ` limit  ` + strconv.Itoa(int(req.Lenght)) + `  offset  ` + strconv.Itoa(int(req.Start))
	} else {
		if req.Order != "" {
			query += `  order by '` + req.Order + `' asc`
		} else {
			query += `  order by product_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var val models.Product
		err := rows.Scan(
			&val.ProductName,
			&val.ProductCode,
			&val.ProductPrice,
			&val.ProductAdminFee,
			&val.ProductMerchantFee,
			&val.ProductProviderName,
			&val.ProductProviderCode,
			&val.ProductProviderPrice,
			&val.ProductProviderAdminFee,
			&val.ProductProviderMerchantFee,
			&val.ProductCategoryID,
			&val.ProductTypeID,
			&val.ProductTypeName,
			&val.ProductReferenceID,
			&val.ProductReferenceCode,
			&val.ProductDenom,
			&val.CreatedBy,
			&val.CreatedAt,
			&val.UpdatedBy,
			&val.UpdatedAt,
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
func (ctx product) GetProductCount(req models.ReqGetProduct) (result int, err error) {
	query := `select count(*)
	from products `

	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx product) GetProduct(req models.ReqGetProduct) (result models.Product, err error) {
	query := `select` + fieldInsert + `
	where true
	`
	if req.Filter.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.Filter.ID)
	}
	if req.Filter.ProductName != "" {
		query += ` and product_name = '` + req.Filter.ProductName + `'`
	}
	if req.Filter.ProductCategoryID != 0 {
		query += ` and product_category_id = ` + strconv.Itoa(req.Filter.ProductCategoryID)
	}
	if req.Filter.ProductTypeID != 0 {
		query += ` and product_type_id = ` + strconv.Itoa(req.Filter.ProductTypeID)
	}
	if req.Filter.ProviderID != 0 {
		query += ` and provider_id = ` + strconv.Itoa(req.Filter.ProviderID)
	}
	if req.Filter.ProductReferenceID != 0 {
		query += ` and product_reference_id = ` + strconv.Itoa(req.Filter.ProductReferenceID)
	}
	if req.Filter.ProductReferenceCode != "" {
		query += ` and product_reference_code = '` + req.Filter.ProductReferenceCode + `'`
	}
	if req.Filter.ProductProviderID != 0 {
		query += ` and product_provider_id = ` + strconv.Itoa(req.Filter.ProductProviderID)
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ProductName,
		&result.ProductCode,
		&result.ProductPrice,
		&result.ProductAdminFee,
		&result.ProductMerchantFee,
		&result.ProductProviderName,
		&result.ProductProviderCode,
		&result.ProductProviderPrice,
		&result.ProductProviderAdminFee,
		&result.ProductProviderMerchantFee,
		&result.ProductCategoryID,
		&result.ProductTypeID,
		&result.ProductTypeName,
		&result.ProductReferenceID,
		&result.ProductReferenceCode,
		&result.ProductDenom,
		&result.CreatedBy,
		&result.CreatedAt,
		&result.UpdatedBy,
		&result.UpdatedAt,
	)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx product) UpdateProduct(req models.ReqGetProduct) (result models.Product, err error) {
	query := ` update products set
	product_name=?,
	product_code=?,
	product_price=?,
	product_admin_fee=?,
	product_merchant_fee=?,
	product_provider_name=?,
	product_provider_code=?,
	product_provider_price=?,
	product_provider_admin_fee=?,
	product_provider_merchant_fee=?,
	product_category_id=?,
	product_type_id=?,
	product_type_name=?,
	product_reference_id=?,
	product_reference_code=?,
	product_denom=?,
	updated_by=?,
	updated_at=?
	where id =?
	`
	query = utils.QuerySupport(query)
	_, err = ctx.repo.Db.Exec(query,
		req.Filter.ProductName,
		req.Filter.ProductCode,
		req.Filter.ProductPrice,
		req.Filter.ProductAdminFee,
		req.Filter.ProductMerchantFee,
		req.Filter.ProductProviderName,
		req.Filter.ProductProviderCode,
		req.Filter.ProductProviderPrice,
		req.Filter.ProductProviderAdminFee,
		req.Filter.ProductProviderMerchantFee,
		req.Filter.ProductCategoryID,
		req.Filter.ProductTypeID,
		req.Filter.ProductTypeName,
		req.Filter.ProductReferenceID,
		req.Filter.ProductReferenceCode,
		req.Filter.ProductDenom,
		req.Filter.UpdatedBy,
		req.Filter.UpdatedAt,
		req.Filter.ID,
	)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx product) DropProduct(req models.ReqGetProduct) (err error) {
	query := ` delete from products where id = $1`
	_, err = ctx.repo.Db.Exec(query, req.Filter.ID)
	if err != nil {
		return err
	}
	return nil
}

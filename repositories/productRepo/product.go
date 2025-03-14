package productrepo

import (
	"desabiller/configs"
	"desabiller/models"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (ctx product) AddProduct(req models.ReqGetProduct) (result models.RespGetProduct, err error) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` insert into products (
		product_provider_id,
		product_category_id,
		product_type_id,
		product_code,
		product_name,
		product_price,
		product_admin_fee,
		product_merchant_fee,
		created_at,
		updated_at,
		created_by,
		updated_by,
		product_reference_id,
		product_reference_code)
		values(
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14
		)  `
	_, err = ctx.repo.Db.Exec(query,
		req.ProductProviderId,
		req.ProductCategoryId,
		req.ProductTypeId,
		req.ProductCode,
		req.ProductName,
		req.ProductPrice,
		req.ProductAdminFee,
		req.ProductMerchantFee,
		dbTime,
		dbTime,
		"sys",
		"sys",
		req.ProductReferenceId,
		req.ProductReferenceCode)
	if err != nil {
		log.Println("Error failed ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx product) GetProducts(req models.ReqGetProduct) (result []models.RespGetProduct, err error) {
	var (
		limit, offset int
	)
	query := `select
	
	e.id,
	e.provider_name,
	c.id,
	c.product_category_name,
	b.id,
	b.product_type_name,
	a.id,
	a.product_name,
	a.product_code,
	a.product_price,
	a.product_admin_fee,
	a.product_merchant_fee,
	a.created_at,
	a.updated_at,
	a.created_by,
	a.updated_by,
	f.id,
	f.product_provider_name,
	f.product_provider_code,
	f.product_provider_price,
	f.product_provider_admin_fee,
	f.product_provider_merchant_fee,
	a.product_reference_id,
	a.product_reference_code
	from products as a
	join product_types as b on a.product_type_id=b.id
	join product_categories as c on a.product_category_id=c.id
	join product_providers as f on a.product_provider_id=f.id
	join providers as e on f.provider_id=e.id
	where true
	`
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.ProductName != "" {
		query += ` and a.product_name = '` + req.ProductName + `'`
	}
	if req.ProductCategoryId != 0 {
		query += ` and c.id = ` + strconv.Itoa(req.ProductCategoryId)
	}
	if req.ProductTypeId != 0 {
		query += ` and b.id = ` + strconv.Itoa(req.ProductTypeId)
	}
	if req.ProviderId != 0 {
		query += ` and e.id = ` + strconv.Itoa(req.ProviderId)
	}
	if req.ProductReferenceId != 0 {
		query += ` and a.product_reference_id = ` + strconv.Itoa(req.ProductReferenceId)
	}
	if req.ProductReferenceCode != "" {
		query += ` and a.product_reference_code = '` + req.ProductReferenceCode + `'`
	}
	if req.ProductProviderId != 0 {
		query += ` and f.id = ` + strconv.Itoa(req.ProductProviderId)
	}
	if req.Filter.Length != 0 {
		offset = req.Filter.Start * req.Filter.Length
		limit = req.Filter.Length
	} else {
		limit = 10
	}
	if req.Filter.Length != 0 {
		query += ` limit  ` + strconv.Itoa(limit) + `  offset  ` + strconv.Itoa(offset)
	} else {
		if req.Filter.OrderBy != "" {
			query += `  order by ` + req.Filter.OrderBy + ` asc`
		} else {
			query += `  order by a.product_name asc`
		}
		query += ` limit  ` + strconv.Itoa(limit) + ` offset 0`
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var val models.RespGetProduct
		err := rows.Scan(
			&val.ProviderId,
			&val.ProviderName,
			&val.ProductCategoryId,
			&val.ProductCategoryName,
			&val.ProductTypeId,
			&val.ProductTypeName,
			&val.ID,
			&val.ProductName,
			&val.ProductCode,
			&val.ProductPrice,
			&val.ProductAdminFee,
			&val.ProductMerchantFee,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy,
			&val.ProductProviderId,
			&val.ProductProviderName,
			&val.ProductProviderCode,
			&val.ProductProviderPrice,
			&val.ProductProviderAdminFee,
			&val.ProductProviderMerchantFee,
			&val.ProductReferenceId,
			&val.ProductReferenceCode,
		)
		if err != nil {
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}
func (ctx product) GetProductCount(req models.ReqGetProduct) (result int, err error) {
	query := `select count(a.id)
	from products as a
	join product_types as b on a.product_type_id=b.id
	join product_categories as c on a.product_category_id=c.id

	join product_providers as f on a.product_provider_id=f.id
	join providers as e on f.provider_id=e.id
	where true
	`
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.ProductName != "" {
		query += ` and a.product_name = '` + req.ProductName + `'`
	}
	if req.ProductCategoryId != 0 {
		query += ` and c.id = ` + strconv.Itoa(req.ProductCategoryId)
	}

	if req.ProductTypeId != 0 {
		query += ` and b.id = ` + strconv.Itoa(req.ProductTypeId)
	}
	if req.ProviderId != 0 {
		query += ` and e.id = ` + strconv.Itoa(req.ProviderId)
	}

	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx product) GetProduct(req models.ReqGetProduct) (result models.RespGetProduct, err error) {
	query := `select
	e.id,
	e.provider_name,
	c.id,
	c.product_category_name,
	b.id,
	b.product_type_name,
	a.id,
	a.product_name,
	a.product_code,
	a.product_price,
	a.product_admin_fee,
	a.product_merchant_fee,
	a.created_at,
	a.updated_at,
	a.created_by,
	a.updated_by,
	f.id,
	f.product_provider_name,
	f.product_provider_code,
	f.product_provider_price,
	f.product_provider_admin_fee,
	f.product_provider_merchant_fee,
	a.product_reference_id,
	a.product_reference_code
	from products as a
	join product_types as b on a.product_type_id=b.id
	join product_categories as c on a.product_category_id=c.id

	join product_providers as f on a.product_provider_id=f.id
	join providers as e on f.provider_id=e.id
	where true
	`
	// if req.ID != 0 {
	// 	query += ` and a.id = ` + strconv.Itoa(req.ID)
	// }
	if req.ProductName != "" {
		query += ` and a.product_name = '` + req.ProductName + `'`
	}
	if req.ProductCode != "" {
		query += ` and a.product_code = '` + req.ProductCode + `'`
	}
	if req.ProductCategoryId != 0 {
		query += ` and c.id = ` + strconv.Itoa(req.ProductCategoryId)
	}
	if req.ProductReferenceId != 0 {
		query += ` and a.product_reference_id = ` + strconv.Itoa(req.ProductReferenceId)
	}
	if req.ProductReferenceCode != "" {
		query += ` and a.product_reference_code = '` + req.ProductReferenceCode + `'`
	}
	if req.ProductTypeId != 0 {
		query += ` and b.id = ` + strconv.Itoa(req.ProductTypeId)
	}
	if req.ProviderId != 0 {
		query += ` and e.id = ` + strconv.Itoa(req.ProviderId)
	}
	if req.ProductProviderId != 0 {
		query += ` and f.id = ` + strconv.Itoa(req.ProductProviderId)
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ProviderId,
		&result.ProviderName,

		&result.ProductCategoryId,
		&result.ProductCategoryName,
		&result.ProductTypeId,
		&result.ProductTypeName,
		&result.ID,
		&result.ProductName,
		&result.ProductCode,
		&result.ProductPrice,
		&result.ProductAdminFee,
		&result.ProductMerchantFee,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.CreatedBy,
		&result.UpdatedBy,
		&result.ProductProviderId,
		&result.ProductProviderName,
		&result.ProductProviderCode,
		&result.ProductProviderPrice,
		&result.ProductProviderAdminFee,
		&result.ProductProviderMerchantFee,
		&result.ProductReferenceId,
		&result.ProductReferenceCode,
	)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx product) UpdateProduct(req models.ReqGetProduct) (result models.RespGetProduct, err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` update products set
	product_provider_id=$1,
	product_category_id=$2,
	product_type_id=$3,
	product_name=$4,
	product_price=$5,
	product_admin_fee=$6,
	product_merchant_fee=$7,
	updated_at=$8,
	updated_by=$9,
	product_code=$10,
	product_reference_id=$11,
	product_reference_code=$2
	where id =$13
	`
	aa, _ := json.Marshal(req)

	fmt.Println("sssss ", string(aa))
	fmt.Println("query ", query)
	_, err = ctx.repo.Db.Exec(query,
		req.ProductProviderId,
		req.ProductCategoryId,
		req.ProductTypeId,
		req.ProductName,
		req.ProductPrice,
		req.ProductAdminFee,
		req.ProductMerchantFee,
		dbTime,
		"sys",
		req.ProductCode,
		req.ID,
		req.ProductReferenceId,
		req.ProductReferenceCode)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx product) DropProduct(req models.ReqGetProduct) (err error) {
	query := ` delete from products where id = $1`
	_, err = ctx.repo.Db.Exec(query, req.ID)
	if err != nil {
		log.Println("Drop product biller :: ", err)
		return err
	}
	return nil
}

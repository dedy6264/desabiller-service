package productrepo

import (
	"desabiller/configs"
	"desabiller/models"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (ctx product) AddProduct(req models.ReqGetProduct) (result models.RespGetProduct, err error) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` insert into products (
		provider_id,
		product_clan_id,
		product_category_id,
		product_type_id,
		product_name,
		product_price,
		product_admin_fee,
		product_merchant_fee,
		created_at,
		updated_at,
		created_by,
		updated_by)
		values(
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12
		)  `
	_, err = ctx.repo.Db.Exec(query,
		req.ProviderId,
		req.ProductClanId,
		req.ProductCategoryId,
		req.ProductTypeId,
		req.ProductName,
		req.ProductAdminFee,
		req.ProductPrice,
		req.ProductMerchantFee,
		dbTime,
		dbTime,
		"sys",
		"sys")
	if err != nil {
		log.Println("Error failed ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx product) GetProducts(req models.ReqGetProduct) (result []models.RespGetProduct, err error) {
	query := `select
	e.id,
	e.provider_name,
	d.id,
	d.product_clan_name,
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
	f.product_provider_code,
	f.product_provider_price,
	f.product_provider_admin_fee,
	f.product_provider_merchant_fee
	from products as a
	join product_types as b on a.product_type_id=b.id
	join product_categories as c on a.product_category_id=c.id
	join product_clans as d on a.product_clan_id=d.id
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
	if req.ProductClanId != 0 {
		query += ` and d.id = ` + strconv.Itoa(req.ProductClanId)
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
	if req.Filter.Limit != 0 {
		query += ` limit  ` + strconv.Itoa(req.Filter.Limit) + `  offset  ` + strconv.Itoa(req.Filter.Offset)
	} else {
		if req.Filter.OrderBy != "" {
			query += `  order by ` + req.Filter.OrderBy + ` asc`
		} else {
			query += `  order by a.product_name asc`
		}
		query += ` limit 100 offset 0`
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
			&val.ProductClanId,
			&val.ProductClanName,
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
			&val.ProductProviderCode,
			&val.ProductProviderPrice,
			&val.ProductProviderAdminFee,
			&val.ProductProviderMerchantFee,
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
	join product_clans as d on a.product_clan_id=d.id
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
	if req.ProductClanId != 0 {
		query += ` and d.id = ` + strconv.Itoa(req.ProductClanId)
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
	d.id,
	d.product_clan_name,
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
	f.product_provider_code,
	f.product_provider_price,
	f.product_provider_admin_fee,
	f.product_provider_merchant_fee
	from products as a
	join product_types as b on a.product_type_id=b.id
	join product_categories as c on a.product_category_id=c.id
	join product_clans as d on a.product_clan_id=d.id
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
	if req.ProductClanId != 0 {
		query += ` and d.id = ` + strconv.Itoa(req.ProductClanId)
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
	fmt.Println("==", query)
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ProviderId,
		&result.ProviderName,
		&result.ProductClanId,
		&result.ProductClanName,
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
		&result.ProductProviderCode,
		&result.ProductProviderPrice,
		&result.ProductProviderAdminFee,
		&result.ProductProviderMerchantFee,
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
	provider_id=$1,
	product_clan_id=$2,
	product_category_id=$3,
	product_type_id=$4,
	product_name=$5,
	product_price=$6,
	product_admin_fee=$7,
	product_merchant_fee=$8,
	updated_at=$9,
	updated_by=$10
	where id = $11 
	`
	_, err = ctx.repo.Db.Exec(query,
		req.ProviderId,
		req.ProductClanId,
		req.ProductCategoryId,
		req.ProductTypeId,
		req.ProductName,
		req.ProductAdminFee,
		req.ProductPrice,
		req.ProductMerchantFee,
		dbTime,
		"sys",
		req.ID)
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

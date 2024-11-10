package productrepo

import (
	"desabiller/configs"
	"desabiller/models"
	"log"
	"strconv"
	"time"
)

func (ctx product) AddProductProvider(req models.ReqGetProductProvider) (result models.RespGetProductProvider, err error) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` insert into product_providers (
		provider_id,
		product_provider_name,
		product_provider_code,
		product_provider_price,
		product_provider_admin_fee,
		product_provider_merchant_fee,
		created_at,
		updated_at,
		created_by,
		updated_by)
		values(
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10
		)  `
	_, err = ctx.repo.Db.Exec(query,
		req.ProviderId,
		req.ProductProviderName,
		req.ProductProviderCode,
		req.ProductProviderPrice,
		req.ProductProviderAdminFee,
		req.ProductProviderMerchantFee,
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
func (ctx product) GetProductProviders(req models.ReqGetProductProvider) (result []models.RespGetProductProvider, err error) {
	query := `select
a.id,
b.id,
b.provider_name,
a.product_provider_name,
a.product_provider_code,
a.product_provider_price,
a.product_provider_admin_fee,
a.product_provider_merchant_fee,
a.created_at,
a.updated_at,
a.created_by,
a.updated_by
from product_providers as a
join providers as b on a.provider_id=b.id
where true
`
	if req.ProductProviderName != "" {
		query += ` and a.product_provider_name= '` + req.ProductProviderName + `'`
	}
	if req.ProductProviderCode != "" {
		query += ` and a.product_provider_code= '` + req.ProductProviderCode + `'`
	}
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.ProviderId != 0 {
		query += ` and b.id = ` + strconv.Itoa(req.ProviderId)
	}
	if req.Filter.Limit != 0 {
		query += ` limit  ` + strconv.Itoa(req.Filter.Limit) + `  offset  ` + strconv.Itoa(req.Filter.Offset)
	} else {
		if req.Filter.OrderBy != "" {
			query += `  order by ` + req.Filter.OrderBy + ` asc`
		} else {
			query += `  order by product_provider_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	var val models.RespGetProductProvider
	for rows.Next() {
		err := rows.Scan(
			&val.ID,
			&val.ProviderId,
			&val.ProviderName,
			&val.ProductProviderName,
			&val.ProductProviderCode,
			&val.ProductProviderPrice,
			&val.ProductProviderAdminFee,
			&val.ProductProviderMerchantFee,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy)
		if err != nil {
			return result, err
		}
		result = append(result, val)
	}

	return result, nil
}
func (ctx product) GetProductProviderCount(req models.ReqGetProductProvider) (result int, err error) {
	query := `select count(a.id)
from product_providers as a
join providers as b on a.provider_id=b.id
where true
`
	if req.ProductProviderName != "" {
		query += ` and a.product_provider_name= '` + req.ProductProviderName + `'`
	}
	if req.ProductProviderCode != "" {
		query += ` and a.product_provider_code= '` + req.ProductProviderCode + `'`
	}
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.ProviderId != 0 {
		query += ` and b.id = ` + strconv.Itoa(req.ProviderId)
	}

	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx product) UpdateProductProvider(req models.ReqGetProductProvider) (result models.RespGetProductProvider, err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` update product_providers set
		id=$1,
		provider_id=$2,
		product_provider_name=$3,
		product_provider_code=$4,
		product_provider_price=$5,
		product_provider_admin_fee=$6,
		product_provider_merchant_fee=$7,
		updated_at = $8,
		updated_by =$9
		where id = $10 
		`
	_, err = ctx.repo.Db.Exec(query,
		req.ID,
		req.ProviderId,
		req.ProductProviderName,
		req.ProductProviderCode,
		req.ProductProviderPrice,
		req.ProductProviderAdminFee,
		req.ProductProviderMerchantFee,
		dbTime,
		"sys",
		req.ID)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx product) DropProductProvider(req models.ReqGetProductProvider) (err error) {
	query := ` delete from product_providers where id = $1`
	_, err = ctx.repo.Db.Exec(query, req.ID)
	if err != nil {
		log.Println("UpdateUserOutlet :: ", err)
		return err
	}
	return nil
}
func (ctx product) GetProductProvider(req models.ReqGetProductProvider) (result models.RespGetProductProvider, err error) {
	query := `select
a.id,
b.provider_name,
a.product_provider_name,
a.product_provider_code,
a.product_provider_price,
a.product_provider_admin_fee,
a.product_provider_merchant_fee,
a.created_at,
a.updated_at,
a.created_by,
a.updated_by
from product_providers as a
join providers as b on a.provider_id=b.id
where true
`
	if req.ProductProviderName != "" {
		query += ` and a.product_provider_name= '` + req.ProductProviderName + `'`
	}
	if req.ProductProviderCode != "" {
		query += ` and a.product_provider_code= '` + req.ProductProviderCode + `'`
	}
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.ProviderId != 0 {
		query += ` and b.id = ` + strconv.Itoa(req.ProviderId)
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.ProviderName,
		&result.ProductProviderName,
		&result.ProductProviderCode,
		&result.ProductProviderPrice,
		&result.ProductProviderAdminFee,
		&result.ProductProviderMerchantFee,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.CreatedBy,
		&result.UpdatedBy)
	if err != nil {
		return result, err
	}
	return result, nil
}

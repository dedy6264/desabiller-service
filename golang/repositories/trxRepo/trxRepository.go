package trxrepo

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/models"
	"strconv"
	"time"
)

const insertQueryPos = `
product_clan_id,
product_clan_name,
product_category_id,
product_category_name,
product_type_id,
product_type_name,
product_id,
product_name,
product_code,
product_price,
product_admin_fee,
product_merchant_fee,
product_provider_id,
product_provider_name,
product_provider_code,
product_provider_price,
product_provider_admin_fee,
product_provider_merchant_fee,
status_code,
status_message,
status_desc,
reference_number,
provider_status_code,
provider_status_message,
provider_status_desc,
provider_reference_number,
client_id,
client_name,
group_id,
group_name,
merchant_id,
merchant_name,
merchant_outlet_id,
merchant_outlet_name,
merchant_outlet_username,
customer_id,
other_msg,
created_at,
updated_at,
created_by,
updated_by
`
const getQueryPos = `
id,
product_clan_id,
product_clan_name,
product_category_id,
product_category_name,
product_type_id,
product_type_name,
product_id,
product_name,
product_code,
product_price,
product_admin_fee,
product_merchant_fee,
product_provider_id,
product_provider_name,
product_provider_code,
product_provider_price,
product_provider_admin_fee,
product_provider_merchant_fee,
status_code,
status_message,
status_desc,
reference_number,
provider_status_code,
provider_status_message,
provider_status_desc,
provider_reference_number,
client_id,
client_name,
group_id,
group_name,
merchant_id,
merchant_name,
merchant_outlet_id,
merchant_outlet_name,
merchant_outlet_username,
customer_id,
other_msg,
created_at,
updated_at,
created_by,
updated_by`

func (ctx trxRepository) GetTrx(req models.ReqGetTrx) (result models.RespGetTrx, err error) {
	query := ` select ` + getQueryPos +
		` from biller_trxs where true `
	if req.Id != 0 {
		query += ` and id= ` + strconv.Itoa(req.Id)
	}
	if req.ProductCategoryId != 0 {
		query += ` and product_category_id= ` + strconv.Itoa(req.ProductCategoryId)
	}
	if req.ProductName != "" {
		query += ` and product_name= '` + req.ProductName + `'`
	}
	if req.StatusCode != "" {
		query += ` and status_code= '` + req.StatusCode + `'`
	}
	if req.ReferenceNumber != "" {
		query += ` and reference_number= '` + req.ReferenceNumber + `'`
	}
	if req.ClientId != 0 {
		query += ` and client_id= ` + strconv.Itoa(req.ClientId)
	}
	if req.GroupId != 0 {
		query += ` and group_id= ` + strconv.Itoa(req.GroupId)
	}
	if req.MerchantId != 0 {
		query += ` and merchant_id= ` + strconv.Itoa(req.MerchantId)
	}
	if req.MerchantOutletId != 0 {
		query += ` and merchantOutlet_id= ` + strconv.Itoa(req.MerchantOutletId)
	}
	if req.CustomerId != "" {
		query += ` and customer_id= '` + req.CustomerId + `'`
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.Id,
		&result.ProductClanId,
		&result.ProductClanName,
		&result.ProductCategoryId,
		&result.ProductCategoryName,
		&result.ProductTypeId,
		&result.ProductTypeName,
		&result.ProductId,
		&result.ProductName,
		&result.ProductCode,
		&result.ProductPrice,
		&result.ProductAdminFee,
		&result.ProductMerchantFee,
		&result.ProductProviderId,
		&result.ProductProviderName,
		&result.ProductProviderCode,
		&result.ProductProviderPrice,
		&result.ProductProviderAdminFee,
		&result.ProductProviderMerchantFee,
		&result.StatusCode,
		&result.StatusMessage,
		&result.StatusDesc,
		&result.ReferenceNumber,
		&result.ProviderStatusCode,
		&result.ProviderStatusMessage,
		&result.ProviderStatusDesc,
		&result.ProviderReferenceNumber,
		&result.ClientId,
		&result.ClientName,
		&result.GroupId,
		&result.GroupName,
		&result.MerchantId,
		&result.MerchantName,
		&result.MerchantOutletId,
		&result.MerchantOutletName,
		&result.MerchantOutletUsername,
		&result.CustomerId,
		&result.OtherMsg,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy,
	)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx trxRepository) GetTrxCount(req models.ReqGetTrx) (result int, err error) {
	query := ` select count(id) from biller_trxs where true `
	if req.Id != 0 {
		query += ` and id= ` + strconv.Itoa(req.Id)
	}
	if req.ProductCategoryId != 0 {
		query += ` and productCategory_id= ` + strconv.Itoa(req.ProductCategoryId)
	}
	if req.ProductName != "" {
		query += ` and productName= '` + req.ProductName + `'`
	}
	if req.StatusCode != "" {
		query += ` and statusCode= '` + req.StatusCode + `'`
	}
	if req.ReferenceNumber != "" {
		query += ` and referenceNumber= '` + req.ReferenceNumber + `'`
	}
	if req.ClientId != 0 {
		query += ` and client_id= ` + strconv.Itoa(req.ClientId)
	}
	if req.GroupId != 0 {
		query += ` and group_id= ` + strconv.Itoa(req.GroupId)
	}
	if req.MerchantId != 0 {
		query += ` and merchant_id= ` + strconv.Itoa(req.MerchantId)
	}
	if req.MerchantOutletId != 0 {
		query += ` and merchantOutlet_id= ` + strconv.Itoa(req.MerchantOutletId)
	}
	if req.CustomerId != "" {
		query += ` and customer_id= '` + req.CustomerId + `'`
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result,
	)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx trxRepository) GetTrxs(req models.ReqGetTrx) (result []models.RespGetTrx, err error) {
	query := ` select ` + getQueryPos +
		` from biller_trxs where true `
	if req.Id != 0 {
		query += ` and id= ` + strconv.Itoa(req.Id)
	}
	if req.ProductCategoryId != 0 {
		query += ` and product_category_id= ` + strconv.Itoa(req.ProductCategoryId)
	}
	if req.ProductName != "" {
		query += ` and product_name= '` + req.ProductName + `'`
	}
	if req.StatusCode != "" {
		query += ` and status_code= '` + req.StatusCode + `'`
	}
	if req.ReferenceNumber != "" {
		query += ` and reference_number= '` + req.ReferenceNumber + `'`
	}
	if req.ClientId != 0 {
		query += ` and client_id= ` + strconv.Itoa(req.ClientId)
	}
	if req.GroupId != 0 {
		query += ` and group_id= ` + strconv.Itoa(req.GroupId)
	}
	if req.MerchantId != 0 {
		query += ` and merchant_id= ` + strconv.Itoa(req.MerchantId)
	}
	if req.MerchantOutletId != 0 {
		query += ` and merchantOutlet_id= ` + strconv.Itoa(req.MerchantOutletId)
	}
	if req.CustomerId != "" {
		query += ` and customer_id= '` + req.CustomerId + `'`
	}

	// if req.OrderBy != "" {
	// 	query += ` order by ` + req.OrderBy + ` ` + req.SortBy
	// } else {
	// 	query += ` order by updated_at desc`
	// }
	if req.Filter.Limit != 0 {
		query += ` limit  ` + strconv.Itoa(req.Filter.Limit) + ` offset ` + strconv.Itoa(req.Filter.Offset)
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	n := 1
	for rows.Next() {
		var val models.RespGetTrx
		err = rows.Scan(
			&val.Id,
			&val.ProductClanId,
			&val.ProductClanName,
			&val.ProductCategoryId,
			&val.ProductCategoryName,
			&val.ProductTypeId,
			&val.ProductTypeName,
			&val.ProductId,
			&val.ProductName,
			&val.ProductCode,
			&val.ProductPrice,
			&val.ProductAdminFee,
			&val.ProductMerchantFee,
			&val.ProductProviderId,
			&val.ProductProviderName,
			&val.ProductProviderCode,
			&val.ProductProviderPrice,
			&val.ProductProviderAdminFee,
			&val.ProductProviderMerchantFee,
			&val.StatusCode,
			&val.StatusMessage,
			&val.StatusDesc,
			&val.ReferenceNumber,
			&val.ProviderStatusCode,
			&val.ProviderStatusMessage,
			&val.ProviderStatusDesc,
			&val.ProviderReferenceNumber,
			&val.ClientId,
			&val.ClientName,
			&val.GroupId,
			&val.GroupName,
			&val.MerchantId,
			&val.MerchantName,
			&val.MerchantOutletId,
			&val.MerchantOutletName,
			&val.MerchantOutletUsername,
			&val.CustomerId,
			&val.OtherMsg,
			&val.CreatedAt,
			&val.CreatedBy,
			&val.UpdatedAt,
			&val.UpdatedBy,
		)
		val.Index = n
		n++
		if err != nil {
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}
func (ctx trxRepository) InsertTrx(req models.ReqGetTrx, tx *sql.Tx) (err error) {
	query := ` insert into biller_trxs (` + insertQueryPos + `) values (
		$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,
		$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,
		$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,
		$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,
		$41
) returning id`
	if tx != nil {
		_, err = tx.Exec(query,
			req.ProductClanId,
			req.ProductClanName,
			req.ProductCategoryId,
			req.ProductCategoryName,
			req.ProductTypeId,
			req.ProductTypeName,
			req.ProductId,
			req.ProductName,
			req.ProductCode,
			req.ProductPrice,
			req.ProductAdminFee,
			req.ProductMerchantFee,
			req.ProductProviderId,
			req.ProductProviderName,
			req.ProductProviderCode,
			req.ProductProviderPrice,
			req.ProductProviderAdminFee,
			req.ProductProviderMerchantFee,
			req.StatusCode,
			req.StatusMessage,
			req.StatusDesc,
			req.ReferenceNumber,
			req.ProviderStatusCode,
			req.ProviderStatusMessage,
			req.ProviderStatusDesc,
			req.ProviderReferenceNumber,
			req.ClientId,
			req.ClientName,
			req.GroupId,
			req.GroupName,
			req.MerchantId,
			req.MerchantName,
			req.MerchantOutletId,
			req.MerchantOutletName,
			req.MerchantOutletUsername,
			req.CustomerId,
			req.OtherMsg,
			req.Filter.CreatedAt,
			req.Filter.CreatedAt,
			req.MerchantOutletUsername,
			req.MerchantOutletUsername,
		)
		if err != nil {
			return err
		}
	} else {
		_, err = ctx.repo.Db.Exec(query,
			req.ProductClanId,
			req.ProductClanName,
			req.ProductCategoryId,
			req.ProductCategoryName,
			req.ProductTypeId,
			req.ProductTypeName,
			req.ProductId,
			req.ProductName,
			req.ProductCode,
			req.ProductPrice,
			req.ProductAdminFee,
			req.ProductMerchantFee,
			req.ProductProviderId,
			req.ProductProviderName,
			req.ProductProviderCode,
			req.ProductProviderPrice,
			req.ProductProviderAdminFee,
			req.ProductProviderMerchantFee,
			req.StatusCode,
			req.StatusMessage,
			req.StatusDesc,
			req.ReferenceNumber,
			req.ProviderStatusCode,
			req.ProviderStatusMessage,
			req.ProviderStatusDesc,
			req.ProviderReferenceNumber,
			req.ClientId,
			req.ClientName,
			req.GroupId,
			req.GroupName,
			req.MerchantId,
			req.MerchantName,
			req.MerchantOutletId,
			req.MerchantOutletName,
			req.MerchantOutletUsername,
			req.CustomerId,
			req.OtherMsg,
			req.Filter.CreatedAt,
			req.Filter.CreatedAt,
			req.MerchantOutletUsername,
			req.MerchantOutletUsername,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
func (ctx trxRepository) UpdateTrx(req models.ReqGetTrx, tx *sql.Tx) (err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` update biller_trxs set
			product_clan_id=$1,
			product_clan_name=$2,
			product_category_id=$3,
			product_category_name=$4,
			product_type_id=$5,
			product_type_name=$6,
			product_id=$7,
			product_name=$8,
			product_code=$9,
			product_price=$10,
			product_admin_fee=$11,
			product_merchant_fee=$12,
			product_provider_id=$13,
			product_provider_name=$14,
			product_provider_code=$15,
			product_provider_price=$16,
			product_provider_admin_fee=$17,
			product_provider_merchant_fee=$18,
			status_code=$19,
			status_message=$20,
			status_desc=$21,
			reference_number=$22,
			provider_status_code=$23,
			provider_status_message=$24,
			provider_status_desc=$25,
			provider_reference_number=$26,
			client_id=$27,
			client_name=$28,
			group_id=$29,
			group_name=$30,
			merchant_id=$31,
			merchant_name=$32,
			merchant_outlet_id=$33,
			merchant_outlet_name=$34,
			merchant_outlet_username=$35,
			customer_id=$36,
			other_msg=$37,
			updated_at=$38,
			updated_by=$39
			where reference_number=$40
	`
	if tx != nil {
		_, err = tx.Exec(query,
			req.ProductClanId,
			req.ProductClanName,
			req.ProductCategoryId,
			req.ProductCategoryName,
			req.ProductTypeId,
			req.ProductTypeName,
			req.ProductId,
			req.ProductName,
			req.ProductCode,
			req.ProductPrice,
			req.ProductAdminFee,
			req.ProductMerchantFee,
			req.ProductProviderId,
			req.ProductProviderName,
			req.ProductProviderCode,
			req.ProductProviderPrice,
			req.ProductProviderAdminFee,
			req.ProductProviderMerchantFee,
			req.StatusCode,
			req.StatusMessage,
			req.StatusDesc,
			req.ReferenceNumber,
			req.ProviderStatusCode,
			req.ProviderStatusMessage,
			req.ProviderStatusDesc,
			req.ProviderReferenceNumber,
			req.ClientId,
			req.ClientName,
			req.GroupId,
			req.GroupName,
			req.MerchantId,
			req.MerchantName,
			req.MerchantOutletId,
			req.MerchantOutletName,
			req.MerchantOutletUsername,
			req.CustomerId,
			req.OtherMsg,
			dbTime,
			req.MerchantOutletUsername,
			req.ReferenceNumber,
		)
	} else {
		_, err = ctx.repo.Db.Exec(query,
			req.ProductClanId,
			req.ProductClanName,
			req.ProductCategoryId,
			req.ProductCategoryName,
			req.ProductTypeId,
			req.ProductTypeName,
			req.ProductId,
			req.ProductName,
			req.ProductCode,
			req.ProductPrice,
			req.ProductAdminFee,
			req.ProductMerchantFee,
			req.ProductProviderId,
			req.ProductProviderName,
			req.ProductProviderCode,
			req.ProductProviderPrice,
			req.ProductProviderAdminFee,
			req.ProductProviderMerchantFee,
			req.StatusCode,
			req.StatusMessage,
			req.StatusDesc,
			req.ReferenceNumber,
			req.ProviderStatusCode,
			req.ProviderStatusMessage,
			req.ProviderStatusDesc,
			req.ProviderReferenceNumber,
			req.ClientId,
			req.ClientName,
			req.GroupId,
			req.GroupName,
			req.MerchantId,
			req.MerchantName,
			req.MerchantOutletId,
			req.MerchantOutletName,
			req.MerchantOutletUsername,
			req.CustomerId,
			req.OtherMsg,
			dbTime,
			req.MerchantOutletUsername,
			req.ReferenceNumber,
		)
	}

	if err != nil {
		return err
	}
	return nil
}

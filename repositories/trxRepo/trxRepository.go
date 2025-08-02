package trxrepo

import (
	"database/sql"
	"desabiller/models"
	"desabiller/utils"
	"strconv"
)

// baru sampe sini mau update product reference code, perlu nambah field ti tb trx
const insertQueryPos = `
product_provider_name,
product_provider_code,
product_provider_price,
product_provider_admin_fee,
product_provider_merchant_fee,
product_id,
product_name,
product_code,
product_price,
product_admin_fee,
product_merchant_fee,
product_category_id,
product_category_name,
product_type_id,
product_type_name,
reference_number,
provider_reference_number,
status_code,
status_message,
status_desc,
status_code_detail,
status_message_detail,
status_desc_detail,
product_reference_id,
product_reference_code,
customer_id,
other_reff,
other_customer_info,
saving_account_name,
saving_account_id,
saving_account_number,
transaction_total_amount,
user_app_id,
username,
created_at,
updated_at,
created_by,
updated_by
`
const getQueryPos = `
id,
product_provider_name,
product_provider_code,
product_provider_price,
product_provider_admin_fee,
product_provider_merchant_fee,
product_id,
product_name,
product_code,
product_price,
product_admin_fee,
product_merchant_fee,
product_category_id,
product_category_name,
product_type_id,
product_type_name,
reference_number,
provider_reference_number,
status_code,
status_message,
status_desc,
status_code_detail,
status_message_detail,
status_desc_detail,
COALESCE(product_reference_id,0) as product_reference_id,
COALESCE(product_reference_code,'') as product_reference_code,
customer_id,
other_reff,
other_customer_info,
saving_account_name,
saving_account_id,
saving_account_number,
transaction_total_amount,
user_app_id,
username,
created_at,
updated_at,
created_by,
updated_by`

func (ctx trxRepository) GetTrx(req models.ReqGetTransaction) (result models.RespGetTrx, err error) {
	query := ` select ` + getQueryPos +
		` from transactions where true `
	if req.Filter.ID != 0 {
		query += ` and id= ` + strconv.Itoa(int(req.Filter.ID))
	}
	if req.Filter.ProductCategoryID != 0 {
		query += ` and product_category_id= ` + strconv.Itoa(int(req.Filter.ProductCategoryID))
	}
	if req.Filter.ProductName != "" {
		query += ` and product_name= '` + req.Filter.ProductName + `'`
	}
	if req.Filter.StatusCode != "" {
		query += ` and status_code= '` + req.Filter.StatusCode + `'`
	}
	if req.Filter.ReferenceNumber != "" {
		query += ` and reference_number= '` + req.Filter.ReferenceNumber + `'`
	}
	if req.Filter.CustomerID != "" {
		query += ` and customer_id= '` + req.Filter.CustomerID + `'`
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.Id,
		&result.ProductProviderName,
		&result.ProductProviderCode,
		&result.ProductProviderPrice,
		&result.ProductProviderAdminFee,
		&result.ProductProviderMerchantFee,
		&result.ProductID,
		&result.ProductName,
		&result.ProductCode,
		&result.ProductPrice,
		&result.ProductAdminFee,
		&result.ProductMerchantFee,
		&result.ProductCategoryID,
		&result.ProductCategoryName,
		&result.ProductTypeID,
		&result.ProductTypeName,
		&result.ReferenceNumber,
		&result.ProviderReferenceNumber,
		&result.StatusCode,
		&result.StatusMessage,
		&result.StatusDesc,
		&result.StatusCodeDetail,
		&result.StatusMessageDetail,
		&result.StatusDescDetail,
		&result.ProductReferenceID,
		&result.ProductReferenceCode,
		&result.CustomerID,
		&result.OtherReff,
		&result.OtherCustomerInfo,
		&result.SavingAccountName,
		&result.SavingAccountID,
		&result.SavingAccountNumber,
		&result.TransactionTotalAmount,
		&result.UserAppID,
		&result.Username,
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
func (ctx trxRepository) GetTrxCount(req models.ReqGetTransaction) (result int, err error) {
	query := ` select count(*) from transactions where true `
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result,
	)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (ctx trxRepository) GetTrxs(req models.ReqGetTransaction) (result []models.RespGetTrx, err error) {
	// var (
	// 	limit, offset int
	// )
	query := ` select ` + getQueryPos +
		` from transactions where true `
	if req.Filter.ID != 0 {
		query += ` and id= ` + strconv.Itoa(int(req.Filter.ID))
	}
	if req.Filter.ProductCategoryID != 0 {
		query += ` and product_category_id= ` + strconv.Itoa(int(req.Filter.ProductCategoryID))
	}
	if req.Filter.ProductName != "" {
		query += ` and product_name= '` + req.Filter.ProductName + `'`
	}
	if req.Filter.StatusCode != "" {
		query += ` and status_code= '` + req.Filter.StatusCode + `'`
	}
	if req.Filter.ReferenceNumber != "" {
		query += ` and reference_number= '` + req.Filter.ReferenceNumber + `'`
	}
	if req.Filter.CustomerID != "" {
		query += ` and customer_id= '` + req.Filter.CustomerID + `'`
	}
	if req.Lenght != 0 {
		query += ` order by updated_at desc limit  ` + strconv.Itoa(int(req.Lenght)) + `  offset  ` + strconv.Itoa(int(req.Start))
	} else {
		if req.Order != "" {
			query += `  order by '` + req.Order + `' asc`
		} else {
			query += `  order by id asc`
		}
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
			&val.ProductProviderName,
			&val.ProductProviderCode,
			&val.ProductProviderPrice,
			&val.ProductProviderAdminFee,
			&val.ProductProviderMerchantFee,
			&val.ProductID,
			&val.ProductName,
			&val.ProductCode,
			&val.ProductPrice,
			&val.ProductAdminFee,
			&val.ProductMerchantFee,
			&val.ProductCategoryID,
			&val.ProductCategoryName,
			&val.ProductTypeID,
			&val.ProductTypeName,
			&val.ReferenceNumber,
			&val.ProviderReferenceNumber,
			&val.StatusCode,
			&val.StatusMessage,
			&val.StatusDesc,
			&val.StatusCodeDetail,
			&val.StatusMessageDetail,
			&val.StatusDescDetail,
			&val.ProductReferenceID,
			&val.ProductReferenceCode,
			&val.CustomerID,
			&val.OtherReff,
			&val.OtherCustomerInfo,
			&val.SavingAccountName,
			&val.SavingAccountID,
			&val.SavingAccountNumber,
			&val.TransactionTotalAmount,
			&val.UserAppID,
			&val.Username,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy,
		)
		val.Index = n
		n++
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
func (ctx trxRepository) InsertTrx(req models.ReqGetTransaction, tx *sql.Tx) (err error) {
	query := ` insert into transactions (` + insertQueryPos + `) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) returning id`
	query = utils.QuerySupport(query)
	if tx != nil {
		_, err = tx.Exec(query,
			req.Filter.ProductProviderName,
			req.Filter.ProductProviderCode,
			req.Filter.ProductProviderPrice,
			req.Filter.ProductProviderAdminFee,
			req.Filter.ProductProviderMerchantFee,
			req.Filter.ProductID,
			req.Filter.ProductName,
			req.Filter.ProductCode,
			req.Filter.ProductPrice,
			req.Filter.ProductAdminFee,
			req.Filter.ProductMerchantFee,
			req.Filter.ProductCategoryID,
			req.Filter.ProductCategoryName,
			req.Filter.ProductTypeID,
			req.Filter.ProductTypeName,
			req.Filter.ReferenceNumber,
			req.Filter.ProviderReferenceNumber,
			req.Filter.StatusCode,
			req.Filter.StatusMessage,
			req.Filter.StatusDesc,
			req.Filter.StatusCodeDetail,
			req.Filter.StatusMessageDetail,
			req.Filter.StatusDescDetail,
			req.Filter.ProductReferenceID,
			req.Filter.ProductReferenceCode,
			req.Filter.CustomerID,
			req.Filter.OtherReff,
			req.Filter.OtherCustomerInfo,
			req.Filter.SavingAccountName,
			req.Filter.SavingAccountID,
			req.Filter.SavingAccountNumber,
			req.Filter.TransactionTotalAmount,
			req.Filter.UserAppID,
			req.Filter.Username,
			req.Filter.CreatedAt,
			req.Filter.UpdatedAt,
			req.Filter.CreatedBy,
			req.Filter.UpdatedBy,
		)
		if err != nil {
			return err
		}
	} else {
		_, err = ctx.repo.Db.Exec(query,
			req.Filter.ProductProviderName,
			req.Filter.ProductProviderCode,
			req.Filter.ProductProviderPrice,
			req.Filter.ProductProviderAdminFee,
			req.Filter.ProductProviderMerchantFee,
			req.Filter.ProductID,
			req.Filter.ProductName,
			req.Filter.ProductCode,
			req.Filter.ProductPrice,
			req.Filter.ProductAdminFee,
			req.Filter.ProductMerchantFee,
			req.Filter.ProductCategoryID,
			req.Filter.ProductCategoryName,
			req.Filter.ProductTypeID,
			req.Filter.ProductTypeName,
			req.Filter.ReferenceNumber,
			req.Filter.ProviderReferenceNumber,
			req.Filter.StatusCode,
			req.Filter.StatusMessage,
			req.Filter.StatusDesc,
			req.Filter.StatusCodeDetail,
			req.Filter.StatusMessageDetail,
			req.Filter.StatusDescDetail,
			req.Filter.ProductReferenceID,
			req.Filter.ProductReferenceCode,
			req.Filter.CustomerID,
			req.Filter.OtherReff,
			req.Filter.OtherCustomerInfo,
			req.Filter.SavingAccountName,
			req.Filter.SavingAccountID,
			req.Filter.SavingAccountNumber,
			req.Filter.TransactionTotalAmount,
			req.Filter.UserAppID,
			req.Filter.Username,
			req.Filter.CreatedAt,
			req.Filter.UpdatedAt,
			req.Filter.CreatedBy,
			req.Filter.UpdatedBy,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
func (ctx trxRepository) UpdateTrx(req models.ReqGetTransaction, tx *sql.Tx) (err error) {

	query := ` update transactions set
			product_provider_name=?,
product_provider_code=?,
product_provider_price=?,
product_provider_admin_fee=?,
product_provider_merchant_fee=?,
product_id=?,
product_name=?,
product_code=?,
product_price=?,
product_admin_fee=?,
product_merchant_fee=?,
product_category_id=?,
product_category_name=?,
product_type_id=?,
product_type_name=?,
reference_number=?,
provider_reference_number=?,
status_code=?,
status_message=?,
status_desc=?,
status_code_detail=?,
status_message_detail=?,
status_desc_detail=?,
product_reference_id=?,
 product_reference_code=?,
customer_id=?,
other_reff=?,
other_customer_info=?,
saving_account_name=?,
saving_account_id=?,
saving_account_number=?,
transaction_total_amount=?,
user_app_id=?,
username=?,
updated_at=?,
updated_by=?
			where reference_number=?
	`
	query = utils.QuerySupport(query)
	if tx != nil {
		_, err = tx.Exec(query,
			req.Filter.ProductProviderName,
			req.Filter.ProductProviderCode,
			req.Filter.ProductProviderPrice,
			req.Filter.ProductProviderAdminFee,
			req.Filter.ProductProviderMerchantFee,
			req.Filter.ProductID,
			req.Filter.ProductName,
			req.Filter.ProductCode,
			req.Filter.ProductPrice,
			req.Filter.ProductAdminFee,
			req.Filter.ProductMerchantFee,
			req.Filter.ProductCategoryID,
			req.Filter.ProductCategoryName,
			req.Filter.ProductTypeID,
			req.Filter.ProductTypeName,
			req.Filter.ReferenceNumber,
			req.Filter.ProviderReferenceNumber,
			req.Filter.StatusCode,
			req.Filter.StatusMessage,
			req.Filter.StatusDesc,
			req.Filter.StatusCodeDetail,
			req.Filter.StatusMessageDetail,
			req.Filter.StatusDescDetail,
			req.Filter.ProductReferenceID,
			req.Filter.ProductReferenceCode,
			req.Filter.CustomerID,
			req.Filter.OtherReff,
			req.Filter.OtherCustomerInfo,
			req.Filter.SavingAccountName,
			req.Filter.SavingAccountID,
			req.Filter.SavingAccountNumber,
			req.Filter.TransactionTotalAmount,
			req.Filter.UserAppID,
			req.Filter.Username,
			req.Filter.UpdatedAt,
			req.Filter.UpdatedBy,
			req.Filter.ReferenceNumber,
		)
	} else {
		_, err = ctx.repo.Db.Exec(query,
			req.Filter.ProductProviderName,
			req.Filter.ProductProviderCode,
			req.Filter.ProductProviderPrice,
			req.Filter.ProductProviderAdminFee,
			req.Filter.ProductProviderMerchantFee,
			req.Filter.ProductID,
			req.Filter.ProductName,
			req.Filter.ProductCode,
			req.Filter.ProductPrice,
			req.Filter.ProductAdminFee,
			req.Filter.ProductMerchantFee,
			req.Filter.ProductCategoryID,
			req.Filter.ProductCategoryName,
			req.Filter.ProductTypeID,
			req.Filter.ProductTypeName,
			req.Filter.ReferenceNumber,
			req.Filter.ProviderReferenceNumber,
			req.Filter.StatusCode,
			req.Filter.StatusMessage,
			req.Filter.StatusDesc,
			req.Filter.StatusCodeDetail,
			req.Filter.StatusMessageDetail,
			req.Filter.StatusDescDetail,
			req.Filter.ProductReferenceID,
			req.Filter.ProductReferenceCode,
			req.Filter.CustomerID,
			req.Filter.OtherReff,
			req.Filter.OtherCustomerInfo,
			req.Filter.SavingAccountName,
			req.Filter.SavingAccountID,
			req.Filter.SavingAccountNumber,
			req.Filter.TransactionTotalAmount,
			req.Filter.UserAppID,
			req.Filter.Username,
			req.Filter.UpdatedAt,
			req.Filter.UpdatedBy,
			req.Filter.ReferenceNumber,
		)
	}

	if err != nil {
		return err
	}
	return nil
}

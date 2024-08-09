package trxrepo

import (
	"database/sql"
	"desabiller/models"
	"fmt"
	"log"
	"strconv"
)

const insertQueryPos = `
trx_number,
trx_number_partner,
payment_at,
created_at,
updated_at,
created_by,
updated_by,
status_code,
status_message,
status_desc,
status_code_partner,
status_message_partner,
status_desc_partner,
segment_id,
product_admin_fee,
product_merchant_fee,
sub_total,
grand_total,
client_id,
client_name,
merchant_id,
merchant_name,
merchant_outlet_id,
merchant_outlet_name,
user_outlet_id,
user_outlet_name,
outlet_device_id,
outlet_device_type,
outlet_device_sn,
payment_method_id,
payment_method_name,
bill_nfo
`
const getQueryPos = `
id,
trx_number,
trx_number_partner,
payment_at,
created_at,
updated_at,
created_by,
updated_by,
status_code,
status_message,
status_desc,
status_code_partner,
status_message_partner,
status_desc_partner,
segment_id,
product_admin_fee,
product_merchant_fee,
sub_total,
grand_total,
client_id,
client_name,
merchant_id,
merchant_name,
merchant_outlet_id,
merchant_outlet_name,
user_outlet_id,
user_outlet_name,
outlet_device_id,
outlet_device_type,
outlet_device_sn,
payment_method_id,
payment_method_name,
bill_nfo`

func (ctx trxRepository) GetTrxListPos(req models.ReqTrx, table string) (result []models.RespTrxList, status bool) {
	var (
		repoName = "GetTrxList"
	)
	table = "trx_poses"
	query := ` select ` + getQueryPos +
		` from ` + table + ` where true `
	if req.Id != 0 {
		query += ` and id= ` + strconv.Itoa(req.Id)
	}
	if req.TrxNumber != "" {
		query += ` and trx_number= '` + req.TrxNumber + `'`
	}
	if req.TrxNumberPartner != "" {
		query += ` and trx_number= '` + req.TrxNumberPartner + `'`
	}
	if req.PaymentAt != "" {
		query += ` and payment_at= '` + req.PaymentAt + `'`
	}
	if req.UpdatedAt != "" {
		query += ` and updated_at= '` + req.UpdatedAt + `'`
	}
	if req.StatusCode != "" {
		query += ` and status_code= '` + req.StatusCode + `'`
	}
	if req.ClientId != 0 {
		query += ` and client_id= ` + strconv.Itoa(req.ClientId)
	}
	if req.MerchantId != 0 {
		query += ` and merchant_id= ` + strconv.Itoa(req.MerchantId)
	}
	if req.MerchantOutletId != 0 {
		query += `and  merchant_outlet_id= ` + strconv.Itoa(req.MerchantOutletId)
	}
	if req.UserOutletId != 0 {
		query += ` and user_outlet_id= ` + strconv.Itoa(req.UserOutletId)
	}
	if req.OutletDeviceId != 0 {
		query += ` and =outlet_device_id ` + strconv.Itoa(req.OutletDeviceId)
	}
	if req.OutletDeviceSn != "" {
		query += ` and outlet_device_sn= '` + req.OutletDeviceSn + `'`
	}
	if req.PaymentMethodId != 0 {
		query += ` and payment_method_id= ` + strconv.Itoa(req.PaymentMethodId)
	}
	if req.CustomerId != "" {
		query += ` and customer_id= ` + req.CustomerId + `'`
	}
	if req.OrderBy != "" {
		query += ` order by ` + req.OrderBy + ` ` + req.SortBy
	} else {
		query += ` order by updated_at desc`
	}
	if req.Limit != 0 {
		query += ` limit  ` + strconv.Itoa(req.Limit) + ` offset ` + strconv.Itoa(req.Offset)
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println(repoName+" ERR ::", err)
		return result, false
	}
	defer rows.Close()
	n := 1
	for rows.Next() {
		var val models.RespTrxList
		err = rows.Scan(
			&val.Id,
			&val.TrxNumber,
			&val.TrxNumberPartner,
			&val.PaymentAt,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy,
			&val.StatusCode,
			&val.StatusMessage,
			&val.StatusDesc,
			&val.StatusCodePartner,
			&val.StatusMessagePartner,
			&val.StatusDescPartner,
			&val.SegmentId,
			&val.ProductAdminFee,
			&val.ProductMerchantFee,
			&val.SubTotal,
			&val.GrandTotal,
			&val.ClientId,
			&val.ClientName,
			&val.MerchantId,
			&val.MerchantName,
			&val.MerchantOutletId,
			&val.MerchantOutletName,
			&val.UserOutletId,
			&val.UserOutletName,
			&val.OutletDeviceId,
			&val.OutletDeviceType,
			&val.OutletDeviceSn,
			&val.PaymentMethodId,
			&val.PaymentMethodName,
			&val.BillInfo,
		)
		val.Index = n
		n++
		if err != nil {
			log.Println(repoName+" ERR ::", err)
			return result, false
		}
		result = append(result, val)
	}
	fmt.Println("RESSS", result)
	if len(result) == 0 {
		log.Println(repoName + " DATA NOT FOUND")
		return result, true
	}
	return result, true
}
func (ctx trxRepository) GetTrxPos(req models.ReqTrx, table string) (result models.RespTrxList, status bool) {
	var (
		repoName = "GetTrx"
	)
	query := ` select ` + getQueryPos +
		` from ` + table + ` where true `
	if req.Id != 0 {
		query += ` and id= ` + strconv.Itoa(req.Id)
	}
	if req.TrxNumber != "" {
		query += ` and trx_number= '` + req.TrxNumber + `'`
	}
	if req.TrxNumberPartner != "" {
		query += ` and trx_number= '` + req.TrxNumberPartner + `'`
	}
	if req.PaymentAt != "" {
		query += ` and payment_at= '` + req.PaymentAt + `'`
	}
	if req.UpdatedAt != "" {
		query += ` and updated_at= '` + req.UpdatedAt + `'`
	}
	if req.StatusCode != "" {
		query += ` and status_code= '` + req.StatusCode + `'`
	}
	if req.ClientId != 0 {
		query += ` and client_id= ` + strconv.Itoa(req.ClientId)
	}
	if req.MerchantId != 0 {
		query += ` and merchant_id= ` + strconv.Itoa(req.MerchantId)
	}
	if req.MerchantOutletId != 0 {
		query += `and  merchant_outlet_id= ` + strconv.Itoa(req.MerchantOutletId)
	}
	if req.UserOutletId != 0 {
		query += ` and user_outlet_id= ` + strconv.Itoa(req.UserOutletId)
	}
	if req.OutletDeviceId != 0 {
		query += ` and =outlet_device_id ` + strconv.Itoa(req.OutletDeviceId)
	}
	if req.OutletDeviceSn != "" {
		query += ` and outlet_device_sn= '` + req.OutletDeviceSn + `'`
	}
	if req.PaymentMethodId != 0 {
		query += ` and payment_method_id= ` + strconv.Itoa(req.PaymentMethodId)
	}
	if req.CustomerId != "" {
		query += ` and customer_id= ` + req.CustomerId + `'`
	}
	if req.OrderBy != "" {
		query += ` order by ` + req.OrderBy + ` ` + req.SortBy
	} else {
		query += ` order by updated_at desc`
	}
	err := ctx.repo.Db.QueryRow(query).Scan(
		&result.Id,
		&result.TrxNumber,
		&result.TrxNumberPartner,
		&result.PaymentAt,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.CreatedBy,
		&result.UpdatedBy,
		&result.StatusCode,
		&result.StatusMessage,
		&result.StatusDesc,
		&result.StatusCodePartner,
		&result.StatusMessagePartner,
		&result.StatusDescPartner,
		&result.SegmentId,
		&result.ProductAdminFee,
		&result.ProductMerchantFee,
		&result.SubTotal,
		&result.GrandTotal,
		&result.ClientId,
		&result.ClientName,
		&result.MerchantId,
		&result.MerchantName,
		&result.MerchantOutletId,
		&result.MerchantOutletName,
		&result.UserOutletId,
		&result.UserOutletName,
		&result.OutletDeviceId,
		&result.OutletDeviceType,
		&result.OutletDeviceSn,
		&result.PaymentMethodId,
		&result.PaymentMethodName,
		&result.BillInfo,
	)
	if err != nil {
		log.Println(repoName+" ERR ::", err.Error())
		return result, false
	}
	return result, true
}
func (ctx trxRepository) InsertTrxDetails(req []models.ReqInsertTrxDetails, tx *sql.Tx) (status bool) {

	var (
		values string
		value  string
		totReq = len(req)
	)
	fmt.Println("TOTAL", len(req))
	for _, data := range req {
		value = ` (` + strconv.Itoa(data.ID) + `,` + strconv.Itoa(data.ProductTypeId) + `,'` + data.ProductTypeName + `',` + strconv.Itoa(data.ProductCategoryId) + `,'` + data.ProductCategoryName + `',` + strconv.Itoa(data.ProductId) + `,'` + data.ProductCode + `','` + data.ProductName + `',` + strconv.Itoa(int(data.ProductPrice)) + `,'` + data.CustomerId + `','` + data.BillInfo + `',` + strconv.Itoa(data.Quantity) + `) `
		if totReq-1 != 0 {
			value += `,`
		}
		totReq = totReq - 1
		values += value
	}
	query := ` insert into trx_pos_details (
		pos_trx_id,product_type_id,product_type_name,product_category_id,product_category_name,product_id,product_code,product_name,product_price,customer_id,bill_info, qty
		) values ` + values
	if tx != nil {
		err := tx.QueryRow(query)
		if err.Err() != nil {
			log.Println("Err :: ", err.Err())
			return false
		}
	} else {
		err := ctx.repo.Db.QueryRow(query)
		if err.Err() != nil {
			log.Println("Err :: ", err.Err())
			return false
		}
	}
	return true
}
func (ctx trxRepository) InsertTrxPos(req models.ReqInsertTrx, table string, tx *sql.Tx) (id int, status bool) {
	var (
		repoName = "InsertTrx"
	)
	table = "trx_poses"
	query := ` insert into ` + table + ` (` + insertQueryPos + `) values (
		$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,
		$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,
		$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,
		$31,$32) returning id`
	if tx != nil {
		err := tx.QueryRow(query,
			req.TrxNumber,
			req.TrxNumberPartner,
			req.PaymentAt,
			req.CreatedAt,
			req.UpdatedAt,
			req.CreatedBy,
			req.UpdatedBy,
			req.StatusCode,
			req.StatusMessage,
			req.StatusDesc,
			req.StatusCodePartner,
			req.StatusMessagePartner,
			req.StatusDescPartner,
			req.SegmentId,
			req.ProductAdminFee,
			req.ProductMerchantFee,
			req.SubTotal,
			req.GrandTotal,
			req.ClientId,
			req.ClientName,
			req.MerchantId,
			req.MerchantName,
			req.MerchantOutletId,
			req.MerchantOutletName,
			req.UserOutletId,
			req.UserOutletName,
			req.OutletDeviceId,
			req.OutletDeviceType,
			req.OutletDeviceSn,
			req.PaymentMethodId,
			req.PaymentMethodName,
			req.BillInfo).Scan(&id)
		fmt.Println(err)
		if err != nil {
			log.Println(repoName+" ERR ::=", err)
			return 0, false
		}
	} else {
		err := ctx.repo.Db.QueryRow(query,
			req.TrxNumber,
			req.TrxNumberPartner,
			req.PaymentAt,
			req.CreatedAt,
			req.UpdatedAt,
			req.CreatedBy,
			req.UpdatedBy,
			req.StatusCode,
			req.StatusMessage,
			req.StatusDesc,
			req.StatusCodePartner,
			req.StatusMessagePartner,
			req.StatusDescPartner,
			req.SegmentId,
			req.ProductAdminFee,
			req.ProductMerchantFee,
			req.SubTotal,
			req.GrandTotal,
			req.ClientId,
			req.ClientName,
			req.MerchantId,
			req.MerchantName,
			req.MerchantOutletId,
			req.MerchantOutletName,
			req.UserOutletId,
			req.UserOutletName,
			req.OutletDeviceId,
			req.OutletDeviceType,
			req.OutletDeviceSn,
			req.PaymentMethodId,
			req.PaymentMethodName,
			req.BillInfo).Scan(&id)
		fmt.Println(err)
		if err != nil {
			log.Println(repoName+" ERR ::=", err)
			return 0, false
		}
	}
	return id, true
}
func (ctx trxRepository) UpdateTrxPos(req models.ReqUpdateTrx, table string) (status bool) {
	var (
		repoName = "UpdateTrx"
	)
	query := ` update ` + table + ` set 
			trx_number_partner=$1,
			payment_at=$2,
			updated_by=$3,
			updated_at=$4,
			status_code=$5,
			status_message=$6,
			status_desc=$7,
			status_code_partner=$8,
			status_message_partner=$9,
			status_desc_partner=$10,
			bill_nfo=$11,
			payment_method_id=$12,
			payment_method_name=$13
			where trx_number=$14
	`
	err := ctx.repo.Db.QueryRow(query,
		req.TrxNumberPartner,
		req.PaymentAt,
		req.UpdatedBy,
		req.UpdatedAt,
		req.StatusCode,
		req.StatusMessage,
		req.StatusDesc,
		req.StatusCodePartner,
		req.StatusMessagePartner,
		req.StatusDescPartner,
		req.BillInfo,
		req.PaymentMethodId,
		req.PaymentMethodName,
		req.TrxNumber,
	)
	if err.Err() != nil {
		log.Println(repoName+" ERR ::", err.Err())
		return false
	}
	return true
}

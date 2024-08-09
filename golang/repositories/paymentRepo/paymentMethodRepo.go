package paymentrepo

import (
	"database/sql"
	"desabiller/models"
	"log"
	"strconv"
)

func (ctx paymentRepo) GetListPaymentMethod(req models.ReqGetListPaymentMethod) (result []models.ResPaymentMethod, status bool) {
	var (
		repoName = "GetListPaymentMethod"
		query    = ` select id, payment_method_name, payment_method_category_id from payment_method where true `
	)
	if req.Id != 0 {
		query += ` and id = ` + strconv.Itoa(req.Id)
	}
	if req.PaymentMethodName != "" {
		query += ` and payment_method_name = '` + req.PaymentMethodName + `'`
	}
	if req.OrderBy != "" {
		query += ` order by ` + req.OrderBy + ` ` + req.AscDesc
	}
	if req.Limit != 0 {
		query += ` limit ` + strconv.Itoa(req.Limit) + ` offset ` + strconv.Itoa(req.Offset)
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("Err ", repoName, " ", err.Error())
		return result, false
	}
	result, err = DTOPayment(rows)
	if err != nil {
		log.Println("Err ", repoName, " ", err)
		return result, false
	}
	return result, true
}
func DTOPayment(row *sql.Rows) (result []models.ResPaymentMethod, err error) {
	for row.Next() {
		var val models.ResPaymentMethod
		err := row.Scan(
			&val.Id,
			&val.PaymentMethodName,
			&val.PaymentMethodCategoryId,
		)
		if err != nil {
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}
func (ctx paymentRepo) GetPaymentMethod(req models.ReqGetListPaymentMethod) (result models.ResPaymentMethod, status bool) {
	var (
		repoName = "GetPaymentMethod"
		query    = ` select id, payment_method_name, payment_method_category_id from payment_method where true `
	)
	if req.Id != 0 {
		query += ` and id = ` + strconv.Itoa(req.Id)
	}
	if req.PaymentMethodName != "" {
		query += ` and payment_method_name = '` + req.PaymentMethodName + `'`
	}
	if req.OrderBy != "" {
		query += ` order by ` + req.OrderBy + ` ` + req.AscDesc
	}
	if req.Limit != 0 {
		query += ` limit ` + strconv.Itoa(req.Limit) + ` offset ` + strconv.Itoa(req.Offset)
	}
	err := ctx.repo.Db.QueryRow(query).Scan(&result.Id, &result.PaymentMethodName, &result.PaymentMethodCategoryId)
	if err != nil {
		log.Println("Err ", repoName, " ", err.Error())
		return result, false
	}
	return result, true
}
func (ctx paymentRepo) AddPaymentMethod(req models.ReqGetListPaymentMethod) (result models.ResPaymentMethod, status bool) {
	var (
		repoName = "AddPaymentMethod"
		query    = ` insert into payment_method  ( payment_method_name, payment_method_category_id) values ($1,$2) returning id`
	)
	err := ctx.repo.Db.QueryRow(query, req.PaymentMethodName, req.PaymentMethodCategoryId).Scan(&result.Id)
	if err != nil {
		log.Println("Error ", repoName, " ", err)
		return result, false
	}
	return result, true
}
func (ctx paymentRepo) UpdatePaymentMethod(req models.ReqGetListPaymentMethod) (result models.ResPaymentMethod, status bool) {
	var (
		repoName = "UpdatePaymentMethod"
		query    = ` update payment_method set  payment_method_name=$1, payment_method_category_id=$2 where id=$3 `
	)
	err := ctx.repo.Db.QueryRow(query, req.PaymentMethodName, req.PaymentMethodCategoryId, req.Id)
	if err.Err() != nil {
		log.Println("ERR "+repoName, " "+err.Err().Error())
		return result, false
	}
	return result, true
}
func (ctx paymentRepo) DropPaymentMethod(id int) (status bool) {
	var (
		repoName = "DropPaymentMethod"
		query    = ` delete from payment_method where id=$1 `
	)
	err := ctx.repo.Db.QueryRow(query, id)
	if err.Err() != nil {
		log.Println(" ERR ", repoName)
		return false
	}
	return true
}

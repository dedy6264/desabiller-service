package paymentrepo

import (
	"database/sql"
	"desabiller/models"
	"log"
	"strconv"
)

func (ctx paymentRepo) GetPaymentMethodCategory(req models.ReqGetListPaymentMethodCategory) (result models.ResPaymentMethodCategory, status bool) {
	var (
		repoName = "GetPaymentMethodCategory"
		query    = ` select id, payment_method_category_name from payment_method_category where true `
	)
	if req.Id != 0 {
		query += ` and id = ` + strconv.Itoa(req.Id)
	}
	if req.PaymentMethodCategoryName != "" {
		query += ` and payment_method_category_name = '` + req.PaymentMethodCategoryName + `'`
	}
	if req.OrderBy != "" {
		query += ` order by ` + req.OrderBy + ` ` + req.AscDesc
	}
	if req.Limit != 0 {
		query += ` limit ` + strconv.Itoa(req.Limit) + ` offset ` + strconv.Itoa(req.Offset)
	}
	err := ctx.repo.Db.QueryRow(query).Scan(&result.Id, &result.PaymentMethodCategoryName)
	if err != nil {
		log.Println("Err ", repoName, " ", err.Error())
		return result, false
	}
	return result, true
}
func (ctx paymentRepo) AddPaymentMethodCategory(req models.ReqGetListPaymentMethodCategory) (result models.ResPaymentMethodCategory, status bool) {
	var (
		repoName = "AddPaymentMethodCategory"
		query    = ` insert into payment_method_category  (payment_method_category_name) values ($1) returning id`
	)
	err := ctx.repo.Db.QueryRow(query, req.PaymentMethodCategoryName).Scan(&result.Id)
	if err != nil {
		log.Println("Error ", repoName, " "+err.Error())
		return result, false
	}
	return result, true
}
func (ctx paymentRepo) UpdatePaymentMethodCategory(req models.ReqGetListPaymentMethodCategory) (result models.ResPaymentMethodCategory, status bool) {
	var (
		repoName = "UpdatePaymentMethodCategory"
		query    = ` update payment_method_category set payment_method_category_name=$1 where id=$2 `
	)
	err := ctx.repo.Db.QueryRow(query, req.PaymentMethodCategoryName, req.Id)
	if err.Err() != nil {
		log.Println("ERR "+repoName, " "+err.Err().Error())
		return result, false
	}
	return result, true
}
func (ctx paymentRepo) DropPaymentMethodCategory(id int) (status bool) {
	var (
		repoName = "DropPaymentMethodCategory"
		query    = ` delete from payment_method_category where id=$1 `
	)
	err := ctx.repo.Db.QueryRow(query, id)
	if err.Err() != nil {
		log.Println(" ERR ", repoName)
		return false
	}
	return true
}
func (ctx paymentRepo) GetListPaymentMethodCategory(req models.ReqGetListPaymentMethodCategory) (result []models.ResPaymentMethodCategory, status bool) {
	var (
		repoName = "GetListPaymentMethodCategory"
		query    = ` select id, payment_method_category_name from payment_method_category where true `
	)
	if req.Id != 0 {
		query += ` and id = ` + strconv.Itoa(req.Id)
	}
	if req.PaymentMethodCategoryName != "" {
		query += ` and payment_method_category_name = '` + req.PaymentMethodCategoryName + `'`
	}
	if req.OrderBy != "" {
		query += ` order by ` + req.OrderBy + ` ` + req.AscDesc
	}
	if req.Limit != 0 {
		query += ` limit ` + strconv.Itoa(req.Limit) + ` offset ` + strconv.Itoa(req.Offset)
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("Err " + repoName)
		return result, false
	}
	result, err = DTO(rows)
	if err != nil {
		log.Println("Err " + repoName + " " + err.Error())
		return result, false
	}
	return result, true

}
func DTO(row *sql.Rows) (result []models.ResPaymentMethodCategory, err error) {
	for row.Next() {
		var val models.ResPaymentMethodCategory
		err := row.Scan(
			&val.Id,
			&val.PaymentMethodCategoryName,
		)
		if err != nil {
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}

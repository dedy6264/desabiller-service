package hierarchyrepo

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/models"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (ctx hierarchy) DropMerchant(req models.ReqGetListMerchant) (status bool) {
	// query := ` delete from merchants where id = $1`
	// err := ctx.repo.Db.QueryRow(query, req.ID)
	// if err != nil {
	// 	log.Println("UpdateMerchant :: ", err.Err())
	// 	return false
	// }
	// return true
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	var id int
	query := `update merchants set 
				deleted_at = $1,
				deleted_by =$2
				where id = $3 returning id
				`
	err := ctx.repo.Db.QueryRow(query, dbTime, "sys", req.ID).Scan(&id)
	if err != nil {
		log.Println("UpdateMerchant :: ", err.Error())
		return false
	}
	return true
}
func (ctx hierarchy) UpdateMerchant(req models.ReqGetListMerchant) (result models.ResGetMerchant, status bool) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update merchants set 
				merchant_name = $1,
				updated_at = $2,
				updated_by =$3
				where id = $4 returning id, merchant_name
				`
	err := ctx.repo.Db.QueryRow(query, req.MerchantName, dbTime, req.Username, req.ID).Scan(&result.ID, &result.MerchantName)
	if err != nil {
		log.Println("UpdateMerchant :: ", err.Error())
		return result, false
	}
	return result, true
}
func (ctx hierarchy) GetListMerchantCount(req models.ReqGetListMerchant) (result int, status bool) {
	query := `select count(a.id)
from merchants as a
join clients as b on a.client_id=b.id
where a.deleted_by=''`
	if req.ClientId != 0 {
		query += ` and a.client_id = ` + strconv.Itoa(req.ClientId)
	}
	if req.MerchantName != "" {
		query += ` and a.merchant_name = '` + req.MerchantName + `' `
	}
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.StartDate != "" {
		query += ` and a.created_at between '` + req.StartDate + `' and '` + req.EndDate + `'`
	}
	err := ctx.repo.Db.QueryRow(query).Scan(&result)
	fmt.Println(":::", query)
	if err != nil {
		log.Println("GetListMerchantCount :: ", err.Error())
		return 0, false
	}
	return result, true
}
func (ctx hierarchy) GetListMerchant(req models.ReqGetListMerchant) (result []models.ResGetMerchant, status bool) {
	query := `select 
	a.id,
	a.merchant_name,
	a.created_at,
	a.created_by,
	a.updated_at,
	a.updated_by,
	a.client_id,
	b.client_name from merchants as a
	join clients as b on a.client_id=b.id
	where a.deleted_by='' `
	if req.ClientId != 0 {
		query += ` and a.client_id = ` + strconv.Itoa(req.ClientId)
	}
	if req.MerchantName != "" {
		query += ` and a.merchant_name = '` + req.MerchantName + `' `
	}
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.StartDate != "" {
		query += ` and a.created_at between '` + req.StartDate + `' and '` + req.EndDate + `'`
	}
	if req.Limit != 0 {
		query += ` limit  ` + strconv.Itoa(req.Limit) + `  offset  ` + strconv.Itoa(req.Offset)
	} else {
		if req.OrderBy != "" {
			query += `  order by a.` + req.OrderBy + ` asc`
		} else {
			query += `  order by a.merchant_name asc`
		}
	}
	fmt.Println("::", query)
	rows, err := ctx.repo.Db.Query(query)
	fmt.Println(":::", query)
	if err != nil {
		log.Println("GetListMerchant :: ", err.Error())
		return result, false
	}
	result, status = DataRowMerchant(rows)
	if !status {
		return result, status
	}
	if len(result) == 0 {
		log.Println("Data not found")
		return result, false
	}
	return result, true
}
func DataRowMerchant(rows *sql.Rows) (result []models.ResGetMerchant, status bool) {
	for rows.Next() {
		var val models.ResGetMerchant
		err := rows.Scan(
			&val.ID,
			&val.MerchantName,
			&val.CreatedAt,
			&val.CreatedBy,
			&val.UpdatedAt,
			&val.UpdatedBy,
			&val.ClientId,
			&val.ClientName,
		)
		if err != nil {
			log.Println("DataRow :: ", err.Error())
			return result, false
		}
		result = append(result, val)
	}
	return result, true
}
func (ctx hierarchy) AddMerchant(req models.ReqGetListMerchant) (status bool) {
	t := time.Now()
	id := 0
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into merchants (
merchant_name,
client_id,
created_at,
created_by,
updated_at,
updated_by
) values ($1,$2,$3,$4,$5,$6) returning id
`
	err := ctx.repo.Db.QueryRow(query, req.MerchantName, req.ClientId, dbTime, "sys", dbTime, "sys").Scan(&id)
	// fmt.Println("::::", query)
	if err != nil {
		log.Println("Err AddMerchant :: ", err)
		return false
	}
	fmt.Println(id)
	return true
}

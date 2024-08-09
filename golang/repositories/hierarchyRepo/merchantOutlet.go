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

func (ctx hierarchy) DropMerchantOutlet(req models.ReqGetListMerchantOutlet) (status bool) {
	// query := ` delete from merchants where id = $1`
	// err := ctx.repo.Db.QueryRow(query, req.ID)
	// if err != nil {
	// 	log.Println("UpdateMerchantOutlet :: ", err.Err())
	// 	return false
	// }
	// return true
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	var id int
	query := `update merchant_outlets set 
				deleted_at = $1,
				deleted_by =$2
				where id = $3 returning id
				`
	err := ctx.repo.Db.QueryRow(query, dbTime, "sys", req.ID).Scan(&id)
	if err != nil {
		log.Println("UpdateMerchantOutlet :: ", err.Error())
		return false
	}
	return true
}
func (ctx hierarchy) UpdateMerchantOutlet(req models.ReqGetListMerchantOutlet) (result models.ResGetMerchantOutlet, status bool) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update merchant_outlets set 
				merchant_outlet_name = $1,
				updated_at = $2,
				updated_by =$3
				where id = $4 returning id, merchant_outlet_name
				`
	err := ctx.repo.Db.QueryRow(query, req.MerchantOutletName, dbTime, req.Username, req.ID).Scan(&result.ID, &result.MerchantOutletName)
	if err != nil {
		log.Println("UpdateMerchantOutlet :: ", err.Error())
		return result, false
	}
	return result, true
}
func (ctx hierarchy) GetListMerchantOutletCount(req models.ReqGetListMerchantOutlet) (result int, status bool) {
	query := `select count(a.id)
from merchant_outlets as a
join merchants as b on a.merchant_id=b.id
join clients as c on b.client_id=c.id
where a.deleted_by=''`
	if req.ClientId != 0 {
		query += ` and b.client_id = ` + strconv.Itoa(req.ClientId)
	}
	if req.MerchantId != 0 {
		query += ` and a.merchant_id = ` + strconv.Itoa(req.MerchantId)
	}
	if req.MerchantOutletName != "" {
		query += ` and a.merchant_outlet_name = '` + req.MerchantOutletName + `' `
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
		log.Println("GetListMerchantOutletCount :: ", err.Error())
		return 0, false
	}
	return result, true
}
func (ctx hierarchy) GetListMerchantOutlet(req models.ReqGetListMerchantOutlet) (result []models.ResGetMerchantOutlet, status bool) {
	query := `select 
	a.id,
	a.merchant_outlet_name,
	a.created_at,
	a.created_by,
	a.updated_at,
	a.updated_by,
	a.merchant_id,
	b.merchant_name,
	b.client_id,
	c.client_name from merchant_outlets as a
	join merchants as b on a.merchant_id=b.id
	join clients as c on b.client_id=c.id
	where a.deleted_by='' `
	if req.ClientId != 0 {
		query += ` and b.client_id = ` + strconv.Itoa(req.ClientId)
	}
	if req.MerchantId != 0 {
		query += ` and a.merchant_id = ` + strconv.Itoa(req.MerchantId)
	}
	if req.MerchantOutletName != "" {
		query += ` and a.merchant_outlet_name = '` + req.MerchantOutletName + `' `
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
			query += `  order by a.merchant_outlet_name asc`
		}
	}

	rows, err := ctx.repo.Db.Query(query)
	fmt.Println(":::", query)
	if err != nil {
		log.Println("GetListMerchantOutlet :: ", err.Error())
		return result, false
	}
	result, status = DataRowMerchantOutlet(rows)
	if !status {
		return result, status
	}
	if len(result) == 0 {
		log.Println("Data not found")
		return result, false
	}
	return result, true
}
func DataRowMerchantOutlet(rows *sql.Rows) (result []models.ResGetMerchantOutlet, status bool) {
	for rows.Next() {
		var val models.ResGetMerchantOutlet
		err := rows.Scan(
			&val.ID,
			&val.MerchantOutletName,
			&val.CreatedAt,
			&val.CreatedBy,
			&val.UpdatedAt,
			&val.UpdatedBy,
			&val.MerchantId,
			&val.MerchantName,
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
func (ctx hierarchy) AddMerchantOutlet(req models.ReqGetListMerchantOutlet) (status bool) {
	t := time.Now()
	id := 0
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into merchant_outlets (
merchant_outlet_name,
merchant_id,
created_at,
created_by,
updated_at,
updated_by
) values ($1,$2,$3,$4,$5,$6) returning id
`
	err := ctx.repo.Db.QueryRow(query, req.MerchantOutletName, req.MerchantId, dbTime, "sys", dbTime, "sys").Scan(&id)
	// fmt.Println("::::", query)
	if err != nil {
		log.Println("Err AddMerchantOutlet :: ", err)
		return false
	}
	fmt.Println(id)
	return true
}

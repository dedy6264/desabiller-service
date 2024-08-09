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

func (ctx hierarchy) DropUserOutlet(req models.ReqGetListUserOutlet) (status bool) {
	// query := ` delete from merchants where id = $1`
	// err := ctx.repo.Db.QueryRow(query, req.ID)
	// if err != nil {
	// 	log.Println("UpdateUserOutlet :: ", err.Err())
	// 	return false
	// }
	// return true
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	var id int
	query := `update user_outlets set
					deleted_at = $1,
					deleted_by =$2
					where id = $3 returning id
					`
	err := ctx.repo.Db.QueryRow(query, dbTime, "sys", req.ID).Scan(&id)
	if err != nil {
		log.Println("UpdateUserOutlet :: ", err.Error())
		return false
	}
	return true
}
func (ctx hierarchy) UpdateUserOutlet(req models.ReqGetListUserOutlet) (result models.ResGetUserOutlet, status bool) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update user_outlets set
					nickname=$1,
					outlet_username=$2,
					outlet_password=$3,
					updated_at = $4,
					updated_by =$5
					where id = $6 returning id
					`
	err := ctx.repo.Db.QueryRow(query, req.Nickname, req.OutletUsername, req.OutletPassword, dbTime, req.Username, req.ID).Scan(&result.ID)
	if err != nil {
		log.Println("UpdateUserOutlet :: ", err.Error())
		return result, false
	}
	return result, true
}
func (ctx hierarchy) GetListUserOutletCount(req models.ReqGetListUserOutlet) (result int, status bool) {
	query := `select count(a.id)
	from user_outlets as a
	join merchant_outlets as b on a.merchant_outlet_id=b.id
	join merchants as c on b.merchant_id=c.id
	join clients as d on c.client_id=d.id
	where a.deleted_by='' `

	if req.ClientId != 0 {
		query += ` and c.client_id = ` + strconv.Itoa(req.ClientId)
	}
	if req.MerchantId != 0 {
		query += ` and b.merchant_id = ` + strconv.Itoa(req.MerchantId)
	}
	if req.MerchantOutletId != 0 {
		query += ` and a.merchant_outlet_id = ` + strconv.Itoa(req.MerchantOutletId)
	}
	if req.MerchantOutletName != "" {
		query += ` and b.merchant_outlet_name = '` + req.MerchantOutletName + `' `
	}
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.StartDate != "" {
		query += ` and a.created_at between '` + req.StartDate + `' and '` + req.EndDate + `'`
	}
	err := ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		log.Println("GetListUserOutletCount :: ", err.Error())
		return 0, false
	}
	return result, true
}
func (ctx hierarchy) GetListUserOutlet(req models.ReqGetListUserOutlet) (result []models.ResGetUserOutlet, status bool) {
	query := `select
	a.id,
	a.nickname,
	a.outlet_username,
	a.outlet_password,
	a.merchant_outlet_id,
	b.merchant_outlet_name,
	b.merchant_id,
	c.merchant_name,
	c.client_id,
	d.client_name ,
	a.created_at,
	a.created_by,
	a.updated_at,
	a.updated_by
	from user_outlets as a
	join merchant_outlets as b on a.merchant_outlet_id=b.id
	join merchants as c on b.merchant_id=c.id
	join clients as d on c.client_id=d.id
	where a.deleted_by='' `
	if req.ClientId != 0 {
		query += ` and c.client_id = ` + strconv.Itoa(req.ClientId)
	}
	if req.OutletUsername != "" {
		query += ` and a.outlet_username = '` + req.OutletUsername + `'`
	}
	if req.MerchantId != 0 {
		query += ` and b.merchant_id = ` + strconv.Itoa(req.MerchantId)
	}
	if req.MerchantOutletId != 0 {
		query += ` and a.merchant_outlet_id = ` + strconv.Itoa(req.MerchantOutletId)
	}
	if req.MerchantOutletName != "" {
		query += ` and b.merchant_outlet_name = '` + req.MerchantOutletName + `' `
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
			query += `  order by a.nickname asc`
		}
	}
	fmt.Println(":MBUH:: ", query)
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("GetListUserOutlet :: ", err.Error())
		return result, false
	}
	result, status = DataRowUserOutlet(rows)
	if !status {
		return result, status
	}
	if len(result) == 0 {
		log.Println("Data not found")
		return result, false
	}
	return result, true
}

func DataRowUserOutlet(rows *sql.Rows) (result []models.ResGetUserOutlet, status bool) {
	for rows.Next() {
		var val models.ResGetUserOutlet
		err := rows.Scan(
			&val.ID,
			&val.Nickname,
			&val.OutletUsername,
			&val.OutletPassword,
			&val.MerchantOutletId,
			&val.MerchantOutletName,
			&val.MerchantId,
			&val.MerchantName,
			&val.ClientId,
			&val.ClientName,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy,
		)
		if err != nil {
			log.Println("DataRow :: ", err.Error())
			return result, false
		}
		result = append(result, val)
	}
	return result, true
}
func (ctx hierarchy) AddUserOutlet(req models.ReqGetListUserOutlet) (status bool) {
	t := time.Now()
	id := 0
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into user_outlets (
		nickname,
		outlet_username,
		outlet_password,
		merchant_outlet_id,
		created_at,
		created_by,
		updated_at,
		updated_by
		) values ($1,$2,$3,$4,$5,$6,$7,$8) returning id
		`
	err := ctx.repo.Db.QueryRow(query, req.Nickname, req.OutletUsername, req.OutletPassword, req.MerchantOutletId, dbTime, "sys", dbTime, "sys").Scan(&id)

	if err != nil {
		log.Println("Err AddUserOutlet :: ", err)
		return false
	}

	return true
}

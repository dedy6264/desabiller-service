package hierarchyrepo

import (
	"desabiller/models"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (ctx nHierarchy) NCreateUserOutlet(req models.ReqGetListNUserOutlet) (id int, err error) {
	dbTime := time.Now().Format(time.RFC3339)
	query := `insert into user_outlets (
		merchant_outlet_id,
		nickname,
		outlet_username,
		outlet_password,
		created_at,
		created_by,
		updated_at,
		updated_by
		) values ($1,$2,$3,$4,$5,$6,$7,$8) returning id
		`
	err = ctx.repo.Db.QueryRow(query, req.MerchantOutletId, req.Nickname, req.OutletUsername, req.OutletPassword, dbTime, "sys", dbTime, "sys").Scan(&id)
	// fmt.Println("::::", query)
	if err != nil {
		log.Println("Err NCreateUserOutlet :: ", err)
		return 0, err
	}
	fmt.Println(id)
	return id, err
}
func (ctx nHierarchy) NReadUserOutlet(req models.ReqGetListNUserOutlet) (result []models.ResGetNUserOutlet, err error) {
	query := `select
	d.id as user_outlet_id,
	a.id as merchant_outlet_id,
	b.id as merchant_id,
	c.id as client_id,
	
	c.client_name,
	b.merchant_name,
	a.merchant_outlet_name,
	d.nickname,
	d.outlet_username,
	a.created_at,
	a.updated_at,
	a.created_by,
	a.updated_by
	from user_outlets as d
	left join merchant_outlets as a on d.merchant_outlet_id=a.id
	left join merchants as b on a.merchant_id=b.id
	left join clients as c on b.client_id=c.id where true `
	if req.ClientId != 0 && req.ClientId != (-1) {
		query += ` and c.id =` + strconv.Itoa(req.ClientId)
	}
	if req.MerchantId != 0 {
		query += ` and b.id =` + strconv.Itoa(req.MerchantId)
	}
	if req.MerchantOutletId != 0 {
		query += ` and a.id =` + strconv.Itoa(req.MerchantOutletId)
	}
	if req.ID != 0 {
		query += ` and d.id =` + strconv.Itoa(req.ID)
	}
	if req.OutletUsername != "" {
		query += ` and d.outlet_username like '%` + req.OutletUsername + `%'`
	}
	if req.Nickname != "" {
		query += ` and a.nickname like '%` + req.Nickname + `%'`
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("Err Repo NReadUserOutlet " + err.Error())
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var val models.ResGetNUserOutlet
		err = rows.Scan(
			&val.ID,
			&val.MerchantOutletId,
			&val.MerchantId,
			&val.ClientId,
			&val.ClientName,
			&val.MerchantName,
			&val.MerchantOutletName,
			&val.Nickname,
			&val.OutletUsername,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy,
		)
		if err != nil {
			log.Println("Err Repo NReadUserOutlet" + err.Error())
			return result, err
		}
		result = append(result, val)
	}
	return result, err
}
func (ctx nHierarchy) NReadSingleUserOutlet(req models.ReqGetListNUserOutlet) (result models.ResGetNUserOutlet, err error) {
	query := `select
	d.id as user_outlet_id,
	a.id as merchant_outlet_id,
	b.id as merchant_id,
	c.id as client_id,
	
	c.client_name,
	b.merchant_name,
	a.merchant_outlet_name,
	d.nickname,
	d.outlet_username,
	d.outlet_password,
	a.created_at,
	a.updated_at,
	a.created_by,
	a.updated_by
	from user_outlets as d
	left join merchant_outlets as a on d.merchant_outlet_id=a.id
	left join merchants as b on a.merchant_id=b.id
	left join clients as c on b.client_id=c.id where true`
	if req.ClientId != 0 && req.ClientId != (-1) {
		query += ` and c.id =` + strconv.Itoa(req.ClientId)
	}
	if req.MerchantId != 0 {
		query += ` and b.id =` + strconv.Itoa(req.MerchantId)
	}
	if req.MerchantOutletId != 0 {
		query += ` and a.id =` + strconv.Itoa(req.MerchantOutletId)
	}
	if req.ID != 0 {
		query += ` and d.id =` + strconv.Itoa(req.ID)
	}
	if req.OutletUsername != "" {
		query += ` and d.outlet_username like '%` + req.OutletUsername + `%'`
	}
	if req.Nickname != "" {
		query += ` and a.nickname like '%` + req.Nickname + `%'`
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.MerchantOutletId,
		&result.MerchantId,
		&result.ClientId,
		&result.ClientName,
		&result.MerchantName,
		&result.MerchantOutletName,
		&result.Nickname,
		&result.OutletUsername,
		&result.OutletPassword,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.CreatedBy,
		&result.UpdatedBy,
	)
	if err != nil {
		log.Println("Err Repo NReadSingleUserOutlet " + err.Error())
		return result, err
	}
	return result, nil
}
func (ctx nHierarchy) NDropUserOutlet(id int) (status bool, err error) {
	query := `delete from user_outlets where id = $1`
	_, err = ctx.repo.Db.Exec(query, id)
	if err != nil {
		log.Println("Err Repo NDropUserOutlet " + err.Error())
		return false, err
	}
	return true, err
}
func (ctx nHierarchy) NUpdateUserOutlet(req models.ReqUpdateNUserOutlet) (result models.ResGetNUserOutlet, err error) {
	query := `update user_outlets set 
		merchant_outlet_id = $1,
		nickname= $2,
		outlet_username=$3,
		outlet_password=$4,
		updated_at = $5,
		updated_by =$6
		where id=$7 returning id, updated_at, updated_by
		`
	err = ctx.repo.Db.QueryRow(query, req.MerchantOutletId, req.Nickname, req.OutletUsername, req.OutletPassword, req.UpdatedAt, req.UpdatedBy, req.ID).Scan(&result.ID, &result.UpdatedAt, &result.UpdatedBy)
	if err != nil {
		log.Println("Err Repo NUpdateUserOutlet " + err.Error())
		return result, err
	}
	return result, err
}

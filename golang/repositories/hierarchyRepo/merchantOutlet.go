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

func (ctx hierarchy) DropMerchantOutlet(req models.ReqGetMerchantOutlet) (err error) {
	query := ` delete from merchant_outlets where id = $1`
	_, err = ctx.repo.Db.Exec(query, req.ID)
	if err != nil {
		log.Println("UpdateMerchantOutlet :: ", err)
		return err
	}
	return nil
}
func (ctx hierarchy) UpdateMerchantOutlet(req models.ReqGetMerchantOutlet) (result models.RespGetMerchantOutlet, err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update merchant_outlets set
				merchant_outlet_name = $1,
				merchant_outlet_password = $2,
				merchant_outlet_username = $3,
				merchant_id = $4,
				updated_at = $5,
				updated_by =$6
				where id = $7 
				`
	_, err = ctx.repo.Db.Exec(query, req.MerchantOutletName, req.MerchantOutletPassword, req.MerchantOutletUsername, req.MerchantId, dbTime, "sys", req.ID)
	if err != nil {
		log.Println("UpdateMerchantOutlet :: ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx hierarchy) GetMerchantOutletCount(req models.ReqGetMerchantOutlet) (result int, err error) {
	query := `select count(a.id)
from merchant_outlets as a
	join merchants as b on a.merchant_id=b.id
	join groups as c on b.group_id=c.id
	join clients as d on c.client_id=d.id
	where true`
	if req.ClientId != 0 {
		query += ` and c.client_id = ` + strconv.Itoa(req.ClientId)
	}
	if req.GroupId != 0 {
		query += ` and b.group_id = ` + strconv.Itoa(req.GroupId)
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
	// if req.StartDate != "" {
	// 	query += ` and a.created_at between '` + req.StartDate + `' and '` + req.EndDate + `'`
	// }
	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	fmt.Println(":::", query)
	if err != nil {
		log.Println("GetListMerchantOutletCount :: ", err.Error())
		return 0, err
	}
	return result, nil
}
func (ctx hierarchy) GetMerchantOutlets(req models.ReqGetMerchantOutlet) (result []models.RespGetMerchantOutlet, err error) {
	query := `select
	a.id,
	a.merchant_outlet_name,
	a.merchant_outlet_username,
	a.merchant_outlet_password,
	a.created_at,
	a.created_by,
	a.updated_at,
	a.updated_by,
	b.id,
	b.merchant_name,
	c.id,
	c.group_name,
	d.id,
	d.client_name from merchant_outlets as a
	join merchants as b on a.merchant_id=b.id
	join groups as c on b.group_id=c.id
	join clients as d on c.client_id=d.id
	where true `
	if req.ClientId != 0 {
		query += ` and c.client_id = ` + strconv.Itoa(req.ClientId)
	}
	if req.GroupId != 0 {
		query += ` and b.group_id = ` + strconv.Itoa(req.GroupId)
	}
	if req.MerchantId != 0 {
		query += ` and a.merchant_id = ` + strconv.Itoa(req.MerchantId)
	}
	if req.MerchantOutletName != "" {
		query += ` and a.merchant_outlet_name = '` + req.MerchantOutletName + `' `
	}
	if req.MerchantOutletUsername != "" {
		query += ` and a.merchant_outlet_username = '` + req.MerchantOutletUsername + `' `
	}
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	// if req.StartDate != "" {
	// 	query += ` and a.created_at between '` + req.StartDate + `' and '` + req.EndDate + `'`
	// }
	if req.Filter.Limit != 0 {
		query += ` limit  ` + strconv.Itoa(req.Filter.Limit) + `  offset  ` + strconv.Itoa(req.Filter.Offset)
	} else {
		if req.Filter.OrderBy != "" {
			query += `  order by a.` + req.Filter.OrderBy + ` asc`
		} else {
			query += `  order by a.merchant_outlet_name asc`
		}
	}

	rows, err := ctx.repo.Db.Query(query)
	fmt.Println(":::", query)
	if err != nil {
		log.Println("GetListMerchantOutlet :: ", err.Error())
		return result, err
	}
	result, err = DataRowMerchantOutlet(rows)
	if err != nil {
		return result, err
	}
	// if len(result) == 0 {
	// 	log.Println("Data not found")
	// 	return result, false
	// }
	return result, nil
}
func DataRowMerchantOutlet(rows *sql.Rows) (result []models.RespGetMerchantOutlet, err error) {
	for rows.Next() {
		var val models.RespGetMerchantOutlet
		err := rows.Scan(
			&val.ID,
			&val.MerchantOutletName,
			&val.MerchantOutletUsername,
			&val.MerchantOutletPassword,
			&val.CreatedAt,
			&val.CreatedBy,
			&val.UpdatedAt,
			&val.UpdatedBy,
			&val.MerchantId,
			&val.MerchantName,
			&val.GroupId,
			&val.GroupName,
			&val.ClientId,
			&val.ClientName,
		)
		if err != nil {
			log.Println("DataRow :: ", err.Error())
			return result, err
		}
		val.MerchantOutletPassword = ""
		result = append(result, val)
	}
	return result, nil
}
func (ctx hierarchy) AddMerchantOutlet(req models.ReqGetMerchantOutlet) (err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into merchant_outlets (
merchant_outlet_name,
merchant_outlet_username,
merchant_outlet_password,
merchant_id,
created_at,
created_by,
updated_at,
updated_by
) values ($1,$2,$3,$4,$5,$6,$7,$8) returning id
`
	_, err = ctx.repo.Db.Exec(query, req.MerchantOutletName, req.MerchantOutletUsername, req.MerchantOutletPassword, req.MerchantId, dbTime, "sys", dbTime, "sys")
	if err != nil {
		log.Println("Err AddMerchantOutlet :: ", err)
		return err
	}
	return nil
}
func (ctx hierarchy) GetMerchantOutlet(req models.ReqGetMerchantOutlet) (result models.RespGetMerchantOutlet, err error) {
	query := `select
	a.id,
	a.merchant_outlet_name,
	a.merchant_outlet_username,
	a.merchant_outlet_password,
	a.created_at,
	a.created_by,
	a.updated_at,
	a.updated_by,
	b.id,
	b.merchant_name,
	c.id,
	c.group_name,
	d.id,
	d.client_name from merchant_outlets as a
	join merchants as b on a.merchant_id=b.id
	join groups as c on b.group_id=c.id
	join clients as d on c.client_id=d.id
	where true `
	if req.ClientId != 0 {
		query += ` and c.client_id = ` + strconv.Itoa(req.ClientId)
	}
	if req.GroupId != 0 {
		query += ` and b.group_id = ` + strconv.Itoa(req.GroupId)
	}
	if req.MerchantId != 0 {
		query += ` and a.merchant_id = ` + strconv.Itoa(req.MerchantId)
	}
	if req.MerchantOutletName != "" {
		query += ` and a.merchant_outlet_name = '` + req.MerchantOutletName + `' `
	}
	if req.MerchantOutletUsername != "" {
		query += ` and a.merchant_outlet_username = '` + req.MerchantOutletUsername + `' `
	}
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	// if req.StartDate != "" {
	// 	query += ` and a.created_at between '` + req.StartDate + `' and '` + req.EndDate + `'`
	// }

	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.MerchantOutletName,
		&result.MerchantOutletUsername,
		&result.MerchantOutletPassword,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy,
		&result.MerchantId,
		&result.MerchantName,
		&result.GroupId,
		&result.GroupName,
		&result.ClientId,
		&result.ClientName,
	)
	// fmt.Println(":::", query)
	if err != nil {
		log.Println("GetListMerchantOutlet :: ", err.Error())
		return result, err
	}
	return result, nil
}

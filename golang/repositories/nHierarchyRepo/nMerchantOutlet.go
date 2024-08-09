package hierarchyrepo

import (
	"desabiller/models"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (ctx nHierarchy) NCreateMerchantOutlet(req models.ReqGetListNMerchantOutlet) (id int, err error) {
	dbTime := time.Now().Format(time.RFC3339)
	query := `insert into merchant_outlets (
		merchant_id,
		merchant_outlet_name,
		created_at,
		created_by,
		updated_at,
		updated_by
		) values ($1,$2,$3,$4,$5,$6) returning id
		`
	err = ctx.repo.Db.QueryRow(query, req.MerchantId, req.MerchantOutletName, dbTime, "sys", dbTime, "sys").Scan(&id)
	// fmt.Println("::::", query)
	if err != nil {
		log.Println("Err NCreateMerchantOutlet :: ", err)
		return 0, err
	}
	fmt.Println(id)
	return id, err
}
func (ctx nHierarchy) NReadMerchantOutlet(req models.ReqGetListNMerchantOutlet) (result []models.ResGetNMerchantOutlet, err error) {
	query := `select
	a.id as merchant_outlet_id,
	b.id as merchant_id,
	c.id as client_id,
	c.client_name,
	b.merchant_name,
	a.merchant_outlet_name,
	a.created_at,
	a.updated_at,
	a.created_by,
	a.updated_by
	from merchant_outlets as a
	left join merchants as b on a.merchant_id=b.id
	left join clients as c on b.client_id=c.id where true `
	if req.ClientId != 0 && req.ClientId != (-1) {
		query += ` and c.id =` + strconv.Itoa(req.ClientId)
	}
	if req.MerchantId != 0 {
		query += ` and b.id =` + strconv.Itoa(req.MerchantId)
	}
	if req.ID != 0 {
		query += ` and a.id =` + strconv.Itoa(req.ID)
	}
	if req.MerchantOutletName != "" {
		query += ` and a.merchant_outlet_name like '%` + req.MerchantOutletName + `%'`
	}
	if req.MerchantName != "" {
		query += ` and a.merchant_name like '%` + req.MerchantName + `%'`
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("Err Repo NReadMerchantOutlet " + err.Error())
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var val models.ResGetNMerchantOutlet
		err = rows.Scan(
			&val.ID,
			&val.MerchantId,
			&val.ClientId,
			&val.ClientName,
			&val.MerchantName,
			&val.MerchantOutletName,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy,
		)
		if err != nil {
			log.Println("Err Repo NReadMerchantOutlet" + err.Error())
			return result, err
		}
		result = append(result, val)
	}
	return result, err
}
func (ctx nHierarchy) NReadSingleMerchantOutlet(req models.ReqGetListNMerchantOutlet) (result models.ResGetNMerchantOutlet, err error) {
	query := `select
	a.id as merchant_outlet_id,
	b.id as merchant_id,
	c.id as client_id,
	c.client_name,
	b.merchant_name,
	a.merchant_outlet_name,
	a.created_at,
	a.updated_at,
	a.created_by,
	a.updated_by
	from merchant_outlets as a
	left join merchants as b on a.merchant_id=b.id
	left join clients as c on b.client_id=c.id where true`
	if req.ClientId != 0 && req.ClientId != (-1) {
		query += ` and c.id =` + strconv.Itoa(req.ClientId)
	}
	if req.MerchantId != 0 {
		query += ` and b.id =` + strconv.Itoa(req.MerchantId)
	}
	if req.ID != 0 {
		query += ` and a.id =` + strconv.Itoa(req.ID)
	}
	if req.MerchantOutletName != "" {
		query += ` and a.merchant_outlet_name like '%` + req.MerchantOutletName + `%'`
	}
	if req.MerchantName != "" {
		query += ` and a.merchant_name like '%` + req.MerchantName + `%'`
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.MerchantId,
		&result.ClientId,
		&result.ClientName,
		&result.MerchantName,
		&result.MerchantOutletName,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.CreatedBy,
		&result.UpdatedBy,
	)
	if err != nil {
		log.Println("Err Repo NReadSingleMerchantOutlet " + err.Error())
		return result, err
	}
	return result, nil
}
func (ctx nHierarchy) NDropMerchantOutlet(id int) (status bool, err error) {
	query := `delete from merchant_outlets where id = $1`
	_, err = ctx.repo.Db.Exec(query, id)
	if err != nil {
		log.Println("Err Repo NDropMerchantOutlet " + err.Error())
		return false, err
	}
	return true, err
}
func (ctx nHierarchy) NUpdateMerchantOutlet(req models.ReqUpdateNMerchantOutlet) (result models.ResGetNMerchantOutlet, err error) {
	query := `update merchant_outlets set 
		merchant_id = $1,
		merchant_outlet_name = $2,
		updated_at = $3,
		updated_by =$4
		where id=$5 returning id, updated_at, updated_by
		`
	err = ctx.repo.Db.QueryRow(query, req.MerchantId, req.MerchantOutletName, req.UpdatedAt, req.UpdatedBy, req.ID).Scan(&result.ID, &result.UpdatedAt, &result.UpdatedBy)
	if err != nil {
		log.Println("Err Repo NUpdateMerchantOutlet " + err.Error())
		return result, err
	}
	return result, err
}

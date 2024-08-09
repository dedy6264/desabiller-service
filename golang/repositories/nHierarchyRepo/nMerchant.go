package hierarchyrepo

import (
	"desabiller/models"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (ctx nHierarchy) NCreateMerchant(req models.ReqGetListNMerchant) (id int, err error) {
	dbTime := time.Now().Format(time.RFC3339)
	query := `insert into merchants (
		client_id,
		merchant_name,
		created_at,
		created_by,
		updated_at,
		updated_by
		) values ($1,$2,$3,$4,$5,$6) returning id
		`
	err = ctx.repo.Db.QueryRow(query, req.ClientId, req.MerchantName, dbTime, "sys", dbTime, "sys").Scan(&id)
	// fmt.Println("::::", query)
	if err != nil {
		log.Println("Err NCreateMerchant :: ", err)
		return 0, err
	}
	fmt.Println(id)
	return id, err
}
func (ctx nHierarchy) NReadMerchant(req models.ReqGetListNMerchant) (result []models.ResGetNMerchant, err error) {
	query := `select
	a.id as merchant_id,
	b.id as client_id,
	b.client_name,
	a.merchant_name,
	a.created_at,
	a.updated_at,
	a.created_by,
	a.updated_by
	from merchants as a
	left join clients as b on a.client_id=b.id where true `
	if req.ClientId != 0 && req.ClientId != (-1) {
		query += ` and b.id =` + strconv.Itoa(req.ClientId)
	}
	if req.ID != 0 {
		query += ` and a.id =` + strconv.Itoa(req.ID)
	}
	if req.ClientName != "" {
		query += ` and b.client_name like '%` + req.ClientName + `%'`
	}
	if req.MerchantName != "" {
		query += ` and a.merchant_name like '%` + req.MerchantName + `%'`
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("Err Repo NReadMerchant " + err.Error())
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var val models.ResGetNMerchant
		err = rows.Scan(
			&val.ID,
			&val.ClientId,
			&val.ClientName,
			&val.MerchantName,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy,
		)
		if err != nil {
			log.Println("Err Repo NReadMerchant" + err.Error())
			return result, err
		}
		result = append(result, val)
	}
	return result, err
}
func (ctx nHierarchy) NReadSingleMerchant(req models.ReqGetListNMerchant) (result models.ResGetNMerchant, err error) {
	query := `select
	a.id as merchant_id,
	b.id as client_id,
	b.client_name,
	a.merchant_name,
	a.created_at,
	a.updated_at,
	a.created_by,
	a.updated_by
	from merchants as a
	left join clients as b on a.client_id=b.id where true `
	if req.ClientId != 0 {
		query += ` and b.id =` + strconv.Itoa(req.ClientId)
	}
	if req.ID != 0 {
		query += ` and a.id =` + strconv.Itoa(req.ID)
	}
	if req.ClientName != "" {
		query += ` and b.client_name like '%` + req.ClientName + `%'`
	}
	if req.MerchantName != "" {
		query += ` and a.merchant_name like '%` + req.MerchantName + `%'`
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.ClientId,
		&result.ClientName,
		&result.MerchantName,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.CreatedBy,
		&result.UpdatedBy,
	)
	if err != nil {
		log.Println("Err Repo NReadSingleMerchant " + err.Error())
		return result, err
	}
	return result, nil
}

func (ctx nHierarchy) NDropMerchant(id int) (status bool, err error) {
	query := `delete from merchants where id = $1`
	_, err = ctx.repo.Db.Exec(query, id)
	if err != nil {
		log.Println("Err Repo NDropMerchant " + err.Error())
		return false, err
	}
	return true, err
}
func (ctx nHierarchy) NUpdateMerchant(req models.ReqUpdateNMerchant) (result models.ResGetNMerchant, err error) {
	query := `update merchants set 
		client_id = $1,
		merchant_name = $2,
		updated_at = $3,
		updated_by =$4
		where id=$5 returning id, updated_at, updated_by
		`
	err = ctx.repo.Db.QueryRow(query, req.ClientId, req.MerchantName, req.UpdatedAt, req.UpdatedBy, req.ID).Scan(&result.ID, &result.UpdatedAt, &result.UpdatedBy)
	if err != nil {
		log.Println("Err Repo NUpdateMerchant " + err.Error())
		return result, err
	}
	return result, err
}

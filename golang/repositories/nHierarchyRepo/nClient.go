package hierarchyrepo

import (
	"desabiller/models"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (ctx nHierarchy) NCreateClient(req models.ReqGetListNClient) (id int, err error) {
	dbTime := time.Now().Format(time.RFC3339)
	query := `insert into clients (
		client_name,
		created_at,
		created_by,
		updated_at,
		updated_by
		) values ($1,$2,$3,$4,$5) returning id
		`
	err = ctx.repo.Db.QueryRow(query, req.ClientName, dbTime, "sys", dbTime, "sys").Scan(&id)
	// fmt.Println("::::", query)
	if err != nil {
		log.Println("Err NCreateClient :: ", err)
		return 0, err
	}
	fmt.Println(id)
	return id, err
}
func (ctx nHierarchy) NReadClient(req models.ReqGetListNClient) (result []models.ResGetNClient, err error) {
	query := `select 
	id,
	client_name,
	created_at,
	updated_at,
	created_by,
	updated_by
	from clients where true `
	if req.ClientName != "" {
		query += ` and client_name like '%` + req.ClientName + `%'`
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("Err Repo NReadClient " + err.Error())
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var val models.ResGetNClient
		err = rows.Scan(
			&val.ID,
			&val.ClientName,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy,
		)
		if err != nil {
			log.Println("Err Repo NReadClient" + err.Error())
			return result, err
		}
		result = append(result, val)
	}
	return result, err
}
func (ctx nHierarchy) NReadSingleClient(req models.ReqGetListNClient) (result models.ResGetNClient, err error) {
	query := `select
	id,
	client_name,
	created_at,
	updated_at,
	created_by,
	updated_by
	from clients where true `
	if req.ClientName != "" {
		query += ` and client_name like '%` + req.ClientName + `%'`
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.ClientName,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.CreatedBy,
		&result.UpdatedBy,
	)
	if err != nil {
		log.Println("Err Repo NReadSingleClient " + err.Error())
		return result, err
	}
	return result, nil
}
func (ctx nHierarchy) NDropClient(id int) (status bool, err error) {
	query := `delete from clients where id = $1`
	_, err = ctx.repo.Db.Exec(query, id)
	if err != nil {
		log.Println("Err Repo NDropClient " + err.Error())
		return false, err
	}
	return true, err
}
func (ctx nHierarchy) NUpdateClient(req models.ReqUpdateNClient) (result models.ResGetNClient, err error) {
	query := `update clients set 
		client_name = $1,
		updated_at = $2,
		updated_by =$3
		where id=$4 returning id, client_name, updated_at, updated_by
		`
	err = ctx.repo.Db.QueryRow(query, req.ClientName, req.UpdatedAt, req.UpdatedBy, req.ID).Scan(&result.ID, &result.ClientName, &result.UpdatedAt, &result.UpdatedBy)
	if err != nil {
		log.Println("Err Repo NDropClient " + err.Error())
		return result, err
	}
	return result, err
}

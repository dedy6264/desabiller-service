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

const field = `
id, client_name, created_at, created_by, updated_at, updated_by
`

func (ctx hierarchy) AddClient(req models.ReqGetListClient) (status bool) {
	t := time.Now()
	id := 0
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into clients (
client_name,
created_at,
created_by,
updated_at,
updated_by
) values ($1,$2,$3,$4,$5) returning id
`
	err := ctx.repo.Db.QueryRow(query, req.ClientName, dbTime, "sys", dbTime, "sys").Scan(&id)
	// fmt.Println("::::", query)
	if err != nil {
		log.Println("Err AddClient :: ", err)
		return false
	}
	fmt.Println(id)
	return true
}
func (ctx hierarchy) DropClient(req models.ReqGetListClient) (status bool) {
	// query := ` delete from clients where id = $1`
	// err := ctx.repo.Db.QueryRow(query, req.ID)
	// if err.Err() != nil {
	// 	log.Println("UpdateClient :: ", err.Err())
	// 	return false
	// }
	// return true
	var id int
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update clients set 
				deleted_at = $1,
				deleted_by =$2
				where id = $3 returning id
				`
	err := ctx.repo.Db.QueryRow(query, dbTime, "sys", req.ID).Scan(&id)
	if err != nil {
		log.Println("UpdateClient :: ", err.Error())
		return false
	}
	log.Println("ID : ", id)
	return true
}
func (ctx hierarchy) UpdateClient(req models.ReqGetListClient) (result models.ResGetClient, status bool) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update clients set 
				client_name = $1,
				updated_at = $2,
				updated_by =$3
				where id = $4 returning id, client_name
				`
	err := ctx.repo.Db.QueryRow(query, req.ClientName, dbTime, "sys", req.ID).Scan(&result.ID, &result.ClientName)
	if err != nil {
		log.Println("UpdateClient :: ", err.Error())
		return result, false
	}
	return result, true
}
func (ctx hierarchy) GetListClientCount(req models.ReqGetListClient) (result int, status bool) {
	if req.ClientName != "" && req.ID != 0 {
		req.ID = 0
	}
	query := `select count(id) from clients where deleted_by='' `
	if req.ClientName != "" {
		query += ` and client_name = '` + req.ClientName + `' `
	}
	if req.ID != 0 {
		query += ` and id = '` + strconv.Itoa(req.ID) + `' `
	}
	if req.StartDate != "" {
		query += `and created_at between '` + req.StartDate + `' and '` + req.EndDate + `'`
	}

	err := ctx.repo.Db.QueryRow(query).Scan(&result)
	fmt.Println(":::", query, err)

	if err != nil {
		log.Println("GetListClientCount :: ", err.Error())
		return 0, false
	}
	return result, true
}
func (ctx hierarchy) GetListClient(req models.ReqGetListClient) (result []models.ResGetClient, status bool) {
	if req.ClientName != "" && req.ID != 0 {
		req.ID = 0
	}
	query := `select ` + field + ` from clients where deleted_by='' `
	if req.ClientName != "" {
		query += ` and client_name = '` + req.ClientName + `' `
	}
	if req.ID != 0 {
		query += ` and id = '` + strconv.Itoa(req.ID) + `' `
	}
	if req.StartDate != "" {
		query += ` and created_at between '` + req.StartDate + `' and '` + req.EndDate + `'`
	}
	if req.Limit != 0 {
		query += ` limit  ` + strconv.Itoa(req.Limit) + `  offset  ` + strconv.Itoa(req.Offset)
	} else {
		if req.OrderBy != "" {
			query += `  order by '` + req.OrderBy + `' asc`
		} else {
			query += `  order by client_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	fmt.Println(query)
	if err != nil {
		log.Println("GetListClient :: ", err.Error())
		return result, false
	}
	result, status = DataRow(rows)
	if !status {
		return result, status
	}
	if len(result) == 0 {
		log.Println("Data not found")
		return result, false
	}
	return result, true
}
func DataRow(rows *sql.Rows) (result []models.ResGetClient, status bool) {
	for rows.Next() {
		var val models.ResGetClient
		err := rows.Scan(
			&val.ID,
			&val.ClientName,
			&val.CreatedAt,
			&val.CreatedBy,
			&val.UpdatedAt,
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

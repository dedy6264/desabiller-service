package hierarchyrepo

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/models"
	"log"
	"strconv"
	"time"
)

const field = `id, client_name, created_at, created_by, updated_at, updated_by`

func (ctx hierarchy) AddClient(req models.ReqGetClient, tx *sql.DB) (err error) {

	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into clients (
				client_name,
				created_at,
				created_by,
				updated_at,
				updated_by
				) values ($1,$2,$3,$4,$5) returning id
				`
	if tx != nil {
		_, err = tx.Exec(query, req.ClientName, dbTime, req.Filter.CreatedBy, dbTime, req.Filter.CreatedBy)
	} else {
		_, err = ctx.repo.Db.Exec(query, req.ClientName, dbTime, req.Filter.CreatedBy, dbTime, req.Filter.CreatedBy)
	}
	if err != nil {
		log.Println("Err AddClient :: ", err)
		return err
	}
	return nil
}
func (ctx hierarchy) DropClient(id int, tx *sql.DB) (err error) {
	query := `delete from clients where id = $1`
	if tx != nil {
		_, err = tx.Exec(query, id)
	} else {
		_, err = ctx.repo.Db.Exec(query, id)
	}
	if err != nil {
		log.Println("DropClient :: ", err.Error())
		return err
	}
	return nil
}
func (ctx hierarchy) UpdateClient(req models.ReqGetClient, tx *sql.DB) (err error) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update clients set 
				client_name = $1,
				updated_at = $2,
				updated_by =$3
				where id = $4 
				`
	if tx != nil {
		_, err = tx.Exec(query, req.ClientName, dbTime, "sys", req.ID)
	} else {
		_, err = ctx.repo.Db.Exec(query, req.ClientName, dbTime, "sys", req.ID)
	}
	if err != nil {
		log.Println("UpdateClient :: ", err.Error())
		return err
	}
	return nil
}
func (ctx hierarchy) GetCount(req models.ReqGetClient) (result int, err error) {
	// if req.ClientName != "" && req.ID != 0 {
	// 	req.ID = 0
	// }
	query := `select count(id) from clients where true `
	if req.ClientName != "" {
		query += ` and client_name = '` + req.ClientName + `' `
	}
	if req.ID != 0 {
		query += ` and id = '` + strconv.Itoa(req.ID) + `' `
	}

	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		log.Println("GetCount :: ", err.Error())
		return 0, err
	}
	return result, nil
}
func (ctx hierarchy) GetClients(req models.ReqGetClient) (result []models.RespGetClient, err error) {
	// if req.ClientName != "" && req.ID != 0 {
	// 	req.ID = 0
	// }
	query := `select ` + field + ` from clients where true `
	if req.ClientName != "" {
		query += ` and client_name = '` + req.ClientName + `' `
	}
	if req.ID != 0 {
		query += ` and id = '` + strconv.Itoa(req.ID) + `' `
	}

	if req.Filter.Limit != 0 {
		query += ` limit  ` + strconv.Itoa(req.Filter.Limit) + `  offset  ` + strconv.Itoa(req.Filter.Offset)
	} else {
		if req.Filter.OrderBy != "" {
			query += `  order by '` + req.Filter.OrderBy + `' asc`
		} else {
			query += `  order by client_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("GetClients :: ", err.Error())
		return result, err
	}
	result, err = DataRow(rows)
	if err != nil {
		log.Println("GetClients :: ", err.Error())
		return result, err
	}
	// if len(result) == 0 {
	// 	log.Println("Data not found")
	// 	return result, false
	// }
	return result, nil
}
func DataRow(rows *sql.Rows) (result []models.RespGetClient, err error) {
	for rows.Next() {
		var val models.RespGetClient
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
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}
func (ctx hierarchy) GetClient(req models.ReqGetClient) (result models.RespGetClient, err error) {
	query := `select ` + field + ` from clients where true `
	if req.ClientName != "" {
		query += ` and client_name = '` + req.ClientName + `' `
	}
	if req.ID != 0 {
		query += ` and id = '` + strconv.Itoa(req.ID) + `' `
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.ClientName,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy,
	)
	if err != nil {
		log.Println("GetClient :: ", err.Error())
		return result, err
	}
	return result, nil
}

package hierarchyrepo

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/models"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (ctx hierarchy) DropGroup(id int) (err error) {
	query := ` delete from groups where id = $1`
	_, err = ctx.repo.Db.Exec(query, id)
	if err != nil {
		log.Println("UpdateMerchant :: ", err)
		return err
	}
	return nil
}
func (ctx hierarchy) UpdateGroup(req models.ReqGetGroup) (result models.RespGetGroup, err error) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update groups set 
				group_name = $1,
				client_id = $2,
				updated_at = $3,
				updated_by =$4
				where id = $5
				`
	_, err = ctx.repo.Db.Exec(query, req.GroupName, req.ClientId, dbTime, req.Filter.UpdatedBy, req.ID)
	if err != nil {
		log.Println("UpdateMerchant :: ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx hierarchy) GetGroupCount(req models.ReqGetGroup) (result int, err error) {
	query := `select count(a.id) from groups as a
	join clients as b on a.client_id=b.id
	where true `
	if req.ClientId != 0 {
		query += ` and a.client_id = ` + strconv.Itoa(req.ClientId)
	}
	if req.GroupName != "" {
		query += ` and a.group_name = '` + req.GroupName + `' `
	}
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	fmt.Println(":::", query)
	if err != nil {
		log.Println("GetListMerchantCount :: ", err.Error())
		return 0, err
	}
	return result, nil
}
func (ctx hierarchy) GetGroup(req models.ReqGetGroup) (result models.RespGetGroup, err error) {
	query := `select 
	a.id,
	a.group_name,
	a.created_at,
	a.created_by,
	a.updated_at,
	a.updated_by,
	a.client_id,
	b.client_name from groups as a
	join clients as b on a.client_id=b.id
	where true `
	if req.ClientId != 0 {
		query += ` and a.client_id = ` + strconv.Itoa(req.ClientId)
	}
	if req.GroupName != "" {
		query += ` and a.group_name = '` + req.GroupName + `' `
	}
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.Filter.Limit != 0 {
		query += ` limit  ` + strconv.Itoa(req.Filter.Limit) + `  offset  ` + strconv.Itoa(req.Filter.Offset)
	} else {
		if req.Filter.OrderBy != "" {
			query += `  order by a.` + req.Filter.OrderBy + ` asc`
		} else {
			query += `  order by a.group_name asc`
		}
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.GroupName,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.CreatedBy,
		&result.UpdatedBy,
		&result.ClientId,
		&result.ClientName,
	)
	if err != nil {
		log.Println("GetListMerchant :: ", err.Error())
		return result, err
	}
	// if len(result) == 0 {
	// 	log.Println("Data not found")
	// 	return result, sql.ErrNoRows
	// }
	return result, nil
}
func (ctx hierarchy) GetGroups(req models.ReqGetGroup) (result []models.RespGetGroup, err error) {
	query := `select 
	a.id,
	a.group_name,
	a.created_at,
	a.created_by,
	a.updated_at,
	a.updated_by,
	a.client_id,
	b.client_name from groups as a
	join clients as b on a.client_id=b.id
	where true `
	if req.ClientId != 0 {
		query += ` and a.client_id = ` + strconv.Itoa(req.ClientId)
	}
	if req.GroupName != "" {
		query += ` and a.group_name = '` + req.GroupName + `' `
	}
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.Filter.Limit != 0 {
		query += ` limit  ` + strconv.Itoa(req.Filter.Limit) + `  offset  ` + strconv.Itoa(req.Filter.Offset)
	} else {
		if req.Filter.OrderBy != "" {
			query += `  order by a.` + req.Filter.OrderBy + ` asc`
		} else {
			query += `  order by a.group_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("GetListMerchant :: ", err.Error())
		return result, err
	}
	defer rows.Close()
	result, err = DataRowGroup(rows)
	if err != nil {
		return result, err
	}
	// if len(result) == 0 {
	// 	log.Println("Data not found")
	// 	return result, sql.ErrNoRows
	// }
	return result, nil
}
func DataRowGroup(rows *sql.Rows) (result []models.RespGetGroup, err error) {
	for rows.Next() {
		var val models.RespGetGroup
		err := rows.Scan(
			&val.ID,
			&val.GroupName,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy,
			&val.ClientId,
			&val.ClientName,
		)
		if err != nil {
			log.Println("DataRow :: ", err.Error())
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}
func (ctx hierarchy) AddGroup(req models.ReqGetGroup, tx *sql.DB) (err error) {
	t := time.Now()
	id := 0
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	ss, _ := json.Marshal(req)
	fmt.Println(string(ss))
	query := `insert into groups (
client_id,
group_name,
created_at,
created_by,
updated_at,
updated_by
) values ($1,$2,$3,$4,$5,$6) returning id
`
	err = ctx.repo.Db.QueryRow(query, req.ClientId, req.GroupName, dbTime, "sys", dbTime, "sys").Scan(&id)
	if err != nil {
		log.Println("Err AddMerchant :: ", err)
		return err
	}
	fmt.Println(id)
	return nil
}

package hierarchyrepo

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/models"
	"log"
	"strconv"
	"time"
)

func (ctx hierarchy) DropMerchant(req models.ReqGetMerchant) (err error) {
	query := ` delete from merchants where id = $1`
	_, err = ctx.repo.Db.Exec(query, req.ID)
	if err != nil {
		log.Println("UpdateMerchant :: ", err)
		return err
	}
	return nil
}
func (ctx hierarchy) UpdateMerchant(req models.ReqGetMerchant) (result models.RespGetMerchant, err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update merchants set 
				merchant_name = $1,
				group_id=$2,
				updated_at = $3,
				updated_by =$4
				where id = $5
				`
	_, err = ctx.repo.Db.Exec(query, req.MerchantName, req.GroupId, dbTime, req.Filter.UpdatedBy, req.ID)
	if err != nil {
		log.Println("UpdateMerchant :: ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx hierarchy) GetMerchantCount(req models.ReqGetMerchant) (result int, err error) {
	query := `select count(a.id)
	from merchants as a
	join groups as b on a.group_id=b.id
	join clients as c on b.client_id=c.id
	where true `
	if req.ClientId != 0 {
		query += ` and c.id = ` + strconv.Itoa(req.ClientId)
	}
	if req.GroupId != 0 {
		query += ` and b.id = ` + strconv.Itoa(req.GroupId)
	}
	if req.MerchantName != "" {
		query += ` and a.merchant_name = '` + req.MerchantName + `' `
	}
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		log.Println("GetListMerchantCount :: ", err.Error())
		return 0, err
	}
	return result, nil
}
func (ctx hierarchy) GetMerchants(req models.ReqGetMerchant) (result []models.RespGetMerchant, err error) {
	query := `select 
	a.id,
	a.merchant_name,
	b.id,
	b.group_name,
	c.id,
	c.client_name,
	a.created_at,
	a.created_by,
	a.updated_at,
	a.updated_by from merchants as a
	join groups as b on a.group_id=b.id
	join clients as c on b.client_id=c.id
	where true `
	if req.ClientId != 0 {
		query += ` and c.id = ` + strconv.Itoa(req.ClientId)
	}
	if req.GroupId != 0 {
		query += ` and b.id = ` + strconv.Itoa(req.GroupId)
	}
	if req.MerchantName != "" {
		query += ` and a.merchant_name = '` + req.MerchantName + `' `
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
			query += `  order by a.merchant_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("GetListMerchant :: ", err.Error())
		return result, err
	}
	result, err = DataRowMerchant(rows)
	if err != nil {
		return result, err
	}
	// if len(result) == 0 {
	// 	log.Println("Data not found")
	// 	return result,
	// }
	return result, nil
}
func DataRowMerchant(rows *sql.Rows) (result []models.RespGetMerchant, err error) {
	for rows.Next() {
		var val models.RespGetMerchant
		err := rows.Scan(
			&val.ID,
			&val.MerchantName,
			&val.GroupId,
			&val.GroupName,
			&val.ClientId,
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
func (ctx hierarchy) AddMerchant(req models.ReqGetMerchant) (err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into merchants (
				merchant_name,
				group_id,
				created_at,
				created_by,
				updated_at,
				updated_by
				) values ($1,$2,$3,$4,$5,$6) 
				`
	_, err = ctx.repo.Db.Exec(query, req.MerchantName, req.GroupId, dbTime, "sys", dbTime, "sys")
	if err != nil {
		log.Println("Err AddMerchant :: ", err)
		return err
	}
	return nil
}
func (ctx hierarchy) GetMerchant(req models.ReqGetMerchant) (result models.RespGetMerchant, err error) {
	query := `select 
	a.id,
	a.merchant_name,
	b.id,
	b.group_name,
	c.id,
	c.client_name,
	a.created_at,
	a.created_by,
	a.updated_at,
	a.updated_by from merchants as a
	join groups as b on a.group_id=b.id
	join clients as c on b.client_id=c.id
	where true `
	if req.ClientId != 0 {
		query += ` and c.id = ` + strconv.Itoa(req.ClientId)
	}
	if req.GroupId != 0 {
		query += ` and b.id = ` + strconv.Itoa(req.GroupId)
	}
	if req.MerchantName != "" {
		query += ` and a.merchant_name = '` + req.MerchantName + `' `
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
			query += `  order by a.merchant_name asc`
		}
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.MerchantName,
		&result.GroupId,
		&result.GroupName,
		&result.ClientId,
		&result.ClientName,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy,
	)
	if err != nil {
		log.Println("GetListMerchant :: ", err.Error())
		return result, err
	}
	return result, nil
}

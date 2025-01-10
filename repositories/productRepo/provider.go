package productrepo

import (
	"desabiller/configs"
	"desabiller/models"
	"log"
	"strconv"
	"time"
)

func (ctx product) AddProvider(req models.ReqGetProvider) (err error) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` insert into providers (
		provider_name,
		created_at,
		updated_at,
		created_by,
		updated_by)
		values(
			$1,$2,$3,$4,$5
		) `
	_, err = ctx.repo.Db.Exec(query,
		req.ProviderName,
		dbTime,
		dbTime,
		"sys",
		"sys")
	if err != nil {
		log.Println("Err AddProvider ", err.Error())
		return err
	}
	return nil
}
func (ctx product) GetProviders(req models.ReqGetProvider) (result []models.RespGetProvider, err error) {
	var (
		limit, offset int
	)
	query := `select 
id,
provider_name,
created_at,
updated_at,
created_by,
updated_by
from providers where true 
`
	if req.ProviderName != "" {
		query += ` and provider_name= '` + req.ProviderName + `'`
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.Filter.Length != 0 {
		offset = req.Filter.Start * req.Filter.Length
		limit = req.Filter.Length
	} else {
		limit = 10
	}
	if req.Filter.Length != 0 {
		query += ` limit  ` + strconv.Itoa(limit) + `  offset  ` + strconv.Itoa(offset)
	} else {
		if req.Filter.OrderBy != "" {
			query += `  order by ` + req.Filter.OrderBy + ` asc`
		} else {
			query += `  order by provider_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("Err GetProviders ", err.Error())
		return result, err
	}
	defer rows.Close()
	var val models.RespGetProvider
	for rows.Next() {
		err := rows.Scan(
			&val.ID,
			&val.ProviderName,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy)
		if err != nil {
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}
func (ctx product) UpdateProvider(req models.ReqGetProvider) (result models.RespGetProvider, err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` update providers set
	provider_name=$1,
	updated_at = $2,
	updated_by =$3
	where id = $4
	`
	_, err = ctx.repo.Db.Exec(query,
		req.ProviderName,
		dbTime,
		"sys",
		req.ID)
	if err != nil {
		log.Println("Err UpdateProvider ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx product) DropProvider(req models.ReqGetProvider) (err error) {
	query := ` delete from providers where id = $1`
	_, err = ctx.repo.Db.Exec(query, req.ID)
	if err != nil {
		log.Println("Err DropProvider ", err)
		return err
	}
	return nil
}
func (ctx product) GetProviderCount(req models.ReqGetProvider) (result int, err error) {
	query := `select count(id)
from providers where true 
`
	if req.ProviderName != "" {
		query += ` and provider_name= '` + req.ProviderName + `'`
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}

	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		log.Println("Err GetProvider ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx product) GetProvider(req models.ReqGetProvider) (result models.RespGetProvider, err error) {
	var (
		limit, offset int
	)
	query := `select 
id,
provider_name,
created_at,
updated_at,
created_by,
updated_by
from providers where true 
`
	if req.ProviderName != "" {
		query += ` and provider_name= '` + req.ProviderName + `'`
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.Filter.Length != 0 {
		offset = req.Filter.Start * req.Filter.Length
		limit = req.Filter.Length
	} else {
		limit = 10
	}
	if req.Filter.Length != 0 {
		query += ` limit  ` + strconv.Itoa(limit) + `  offset  ` + strconv.Itoa(offset)
	} else {
		if req.Filter.OrderBy != "" {
			query += `  order by ` + req.Filter.OrderBy + ` asc`
		} else {
			query += `  order by provider_name asc`
		}
	}
	err = ctx.repo.Db.QueryRow(query).Scan(&result.ID,
		&result.ProviderName,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.CreatedBy,
		&result.UpdatedBy)
	if err != nil {
		log.Println("Err GetProvider ", err.Error())
		return result, err
	}
	return result, nil
}

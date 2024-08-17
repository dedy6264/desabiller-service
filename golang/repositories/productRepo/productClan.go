package productrepo

import (
	"desabiller/configs"
	"desabiller/models"
	"log"
	"strconv"
	"time"
)

func (ctx product) AddClan(req models.ReqGetProductClanClan) (err error) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` insert into product_clans (
		product_clan_name,
		created_at,
		updated_at,
		created_by,
		updated_by)
		values(
			$1,$2,$3,$4,$5
		) `
	_, err = ctx.repo.Db.Exec(query,
		req.ProductClanName,
		dbTime,
		dbTime,
		"sys",
		"sys")
	if err != nil {
		log.Println("Err AddProductClan ", err.Error())
		return err
	}
	return nil
}
func (ctx product) GetProductClans(req models.ReqGetProductClan) (result []models.RespGetProductClan, err error) {
	query := `select 
		id,
		product_cln_name,
		created_at,
		updated_at,
		created_by,
		updated_by
		from product_clans where true 
		`
	if req.ProductClanName != "" {
		query += ` and product_clan_name= '` + req.ProductClanName + `'`
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.Filter.Limit != 0 {
		query += ` limit  ` + strconv.Itoa(req.Filter.Limit) + `  offset  ` + strconv.Itoa(req.Filter.Offset)
	} else {
		if req.Filter.OrderBy != "" {
			query += `  order by ` + req.Filter.OrderBy + ` asc`
		} else {
			query += `  order by product_clan_name asc`
		}
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("Err GetProductClans ", err.Error())
		return result, err
	}
	var val models.RespGetProductClan
	for rows.Next() {
		err := rows.Scan(
			&val.ID,
			&val.ProductClanName,
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
func (ctx product) UpdateProductClan(req models.ReqGetProductClan) (result models.RespGetProductClan, err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` update product_clans set
				product_clan_name=$1,
				updated_at = $2,
				updated_by =$3
				where id = $4`
	_, err = ctx.repo.Db.Exec(query,
		req.ProductClanName,
		dbTime,
		"sys",
		req.ID)
	if err != nil {
		log.Println("Err UpdateProductClan ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx product) DropProductClan(req models.ReqGetProductClan) (err error) {
	query := ` delete from product_clans where id = $1`
	_, err = ctx.repo.Db.Exec(query, req.ID)
	if err != nil {
		log.Println("Err DropProductClan ", err)
		return err
	}
	return nil
}
func (ctx product) GetProductClanCount(req models.ReqGetProductClan) (result int, err error) {
	query := `select count(id)
	from product_clans where true `
	if req.ProductClanName != "" {
		query += ` and product_clan_name= '` + req.ProductClanName + `'`
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}

	err = ctx.repo.Db.QueryRow(query).Scan(&result)
	if err != nil {
		log.Println("Err GetProductClan ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx product) GetProductClan(req models.ReqGetProductClan) (result models.RespGetProductClan, err error) {
	query := `select 
	id,
	product_clan_name,
	created_at,
	updated_at,
	created_by,
	updated_by
	from products where true 
	`
	if req.ProductClanName != "" {
		query += ` and product_clan_name= '` + req.ProductClanName + `'`
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	if req.Filter.Limit != 0 {
		query += ` limit  ` + strconv.Itoa(req.Filter.Limit) + `  offset  ` + strconv.Itoa(req.Filter.Offset)
	} else {
		if req.Filter.OrderBy != "" {
			query += `  order by ` + req.Filter.OrderBy + ` asc`
		} else {
			query += `  order by product_clan_name asc`
		}
	}
	err = ctx.repo.Db.QueryRow(query).Scan(&result.ID,
		&result.ProductClanName,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.CreatedBy,
		&result.UpdatedBy)
	if err != nil {
		log.Println("Err GetProductClan ", err.Error())
		return result, err
	}
	return result, nil
}

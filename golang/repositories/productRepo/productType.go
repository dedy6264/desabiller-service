package productrepo

import (
	"desabiller/configs"
	"desabiller/models"
	"log"
	"strconv"
	"time"
)

func (ctx product) AddProductType(req models.ReqGetProductType) (result models.RespGetProductType, err error) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` insert into product_types (
		product_type_name,
		created_at,
		updated_at,
		created_by,
		updated_by)
		values(
			$1,$2,$3,$4,$5
		)  `
	_, err = ctx.repo.Db.Exec(query, req.ProductTypeName, dbTime, dbTime, "sys", "sys")
	if err != nil {
		log.Println("Error failed ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx product) GetProductTypes(req models.ReqGetProductType) (result []models.RespGetProductType, err error) {
	query := `select
	id,
product_type_name,
created_at,
updated_at,
created_by,
updated_by
from product_types 
where true
`
	if req.ProductTypeName != "" {
		query += ` and product_type_name= '` + req.ProductTypeName + `'`
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
			query += `  order by product_type_name asc`
		}
		query += ` limit 100 offset 0`
	}
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println(" GetProductType :: Failed : ", err.Error())
		return result, err
	}
	defer rows.Close()
	var val models.RespGetProductType
	for rows.Next() {
		err := rows.Scan(&val.ID, &val.ProductTypeName, &val.CreatedAt, &val.UpdatedAt, &val.CreatedBy, &val.UpdatedBy)
		if err != nil {
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}
func (ctx product) UpdateProductType(req models.ReqGetProductType) (result models.RespGetProductType, err error) {
	t := time.Now()
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := ` update product_types set
	product_type_name=$1,
	updated_at = $2,
	updated_by =$3
	where id = $4 returning id
	`
	_, err = ctx.repo.Db.Exec(query, req.ProductTypeName, dbTime, "sys", req.ID)
	if err != nil {
		log.Println(" UpdateProductType :: Failed : ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx product) DropProductType(req models.ReqGetProductType) (err error) {
	query := `delete from product_types
					where id = $1 returning id
					`
	_, err = ctx.repo.Db.Exec(query, req.ID)
	if err != nil {
		log.Println("UpdateProductType :: ", err)
		return err
	}
	return nil
}
func (ctx product) GetProductType(req models.ReqGetProductType) (result models.RespGetProductType, err error) {
	query := `select
	id,
product_type_name,
created_at,
updated_at,
created_by,
updated_by
from product_types 
where true
`
	if req.ProductTypeName != "" {
		query += ` and product_type_name= '` + req.ProductTypeName + `'`
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.ProductTypeName,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.CreatedBy,
		&result.UpdatedBy,
	)
	if err != nil {
		log.Println(" GetProductType :: Failed : ", err.Error())
		return result, err
	}
	return result, nil
}
func (ctx product) GetProductTypeCount(req models.ReqGetProductType) (result int, err error) {
	query := `select count(id)
from product_types 
where true
`
	if req.ProductTypeName != "" {
		query += ` and product_type_name= '` + req.ProductTypeName + `'`
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	err = ctx.repo.Db.QueryRow(query).Scan(
		&result,
	)
	if err != nil {
		log.Println(" GetProductType :: Failed : ", err.Error())
		return result, err
	}
	return result, nil
}

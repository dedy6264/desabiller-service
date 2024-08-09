package nuserdashboardsrepo

import (
	"desabiller/models"
	"desabiller/repositories"
	"log"
	"strconv"
)

type nUserDasboardsRepo struct {
	Repo repositories.Repositories
}

func NewNUserDashboardsRepo(Repo repositories.Repositories) nUserDasboardsRepo {
	return nUserDasboardsRepo{
		Repo: Repo,
	}
}

func (ctx nUserDasboardsRepo) NCreateUserDashboard(req models.ReqCreateNUserDashboard) (id int, err error) {
	query := `insert into user_dashboards (
		username,
		email,
		password,
		role,
		client,
		merchant,
		merchant_outlet,
		created_at,
		updated_at,
		created_by,
		updated_by
	) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) returning id`
	err = ctx.Repo.Db.QueryRow(query, req.Username, req.Email, req.Password, req.Role, req.ClientId, req.MerchantId, req.MerchantId, req.CreatedAt, req.UpdatedAt, req.CreatedBy, req.UpdatedBy).Scan(&id)
	if err != nil {
		log.Println("ERROR REPO CreateUserDashboard ", err)
		return 0, err
	}
	return id, err
}
func (ctx nUserDasboardsRepo) NReadUserDashboard(req models.ReqGetListNUserDashboard) (result []models.RespGetListNUserDashboard, err error) {
	query := `select 
	id,
	username,
	email,
	role,
	client,
	merchant,
	merchant_outlet,
	created_at,
	updated_at,
	created_by,
	updated_by
	from user_dashboards where true `
	if req.Data.Email != "" {
		query += ` and email like '%` + req.Data.Email + `%'`
	}
	if req.Data.Username != "" {
		query += ` and username like '%` + req.Data.Username + `%'`
	}
	if req.Data.Password != "" {
		query += ` and password like '%` + req.Data.Password + `%'`
	}
	if req.Data.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.Data.ID)
	}
	rows, err := ctx.Repo.Db.Query(query)
	if err != nil {
		log.Println("Err Repo ReadUserDashboard " + err.Error())
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var val models.RespGetListNUserDashboard
		err = rows.Scan(
			&val.ID,
			&val.Username,
			&val.Email,
			&val.Role,
			&val.ClientId,
			&val.MerchantId,
			&val.MerchantOutletId,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy,
		)
		if err != nil {
			log.Println("Err Repo ReadUserDashboard" + err.Error())
			return result, err
		}
		result = append(result, val)
	}
	return result, err
}
func (ctx nUserDasboardsRepo) NReadSingleUserDashboard(req models.ReqCreateNUserDashboard) (result models.RespGetListNUserDashboard, err error) {
	query := `select 
	id,
	username,
	email,
	password,
	role,
	client,
	merchant,
	merchant_outlet,
	created_at,
	updated_at,
	created_by,
	updated_by
	from user_dashboards where true `
	if req.Email != "" {
		query += ` and email like '%` + req.Email + `%'`
	}
	if req.Username != "" {
		query += ` and username like '%` + req.Username + `%'`
	}
	if req.Password != "" {
		query += ` and password like '%` + req.Password + `%'`
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	err = ctx.Repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.Username,
		&result.Email,
		&result.Password,
		&result.Role,
		&result.ClientId,
		&result.MerchantId,
		&result.MerchantOutletId,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.CreatedBy,
		&result.UpdatedBy,
	)
	if err != nil {
		log.Println("Err Repo ReadSingleUserDashboard " + err.Error())
		return result, err
	}
	return result, nil
}
func (ctx nUserDasboardsRepo) NDropUserDashboard(id int) (status bool, err error) {
	query := `delete from user_dashboards where id = $1`
	_, err = ctx.Repo.Db.Exec(query, id)
	if err != nil {
		log.Println("Err Repo DropUserDashboard " + err.Error())
		return false, err
	}
	return true, err
}
func (ctx nUserDasboardsRepo) NUpdateUserDashboard(req models.ReqCreateNUserDashboard) (result models.RespGetListNUserDashboard, err error) {
	query := `update user_dashboards set 
		username=$1,
		email=$2,
		password=$3,
		role=$4,
		updated_at=$5,
		updated_by=$6,
		client=$7,
		merchant=$8,
		merchant_outlet=$9
		where id=$10 returning id, username,
		email,
		password,
		role,
		client,
		merchant,
		merchant_outlet,
		updated_at,
		updated_by
		`
	err = ctx.Repo.Db.QueryRow(query, req.Username, req.Email, req.Password, req.Role, req.UpdatedAt, req.UpdatedBy, req.ID, req.ClientId, req.MerchantId, req.MerchantOutletId).Scan(&result.ID, &result.Username, &result.Email, &result.Password, &result.Role, &result.ClientId, &result.MerchantId, &result.MerchantOutletId, &result.UpdatedAt, &result.UpdatedBy)
	if err != nil {
		log.Println("Err Repo NUpdateUserDashboard " + err.Error())
		return result, err
	}
	return result, err
}

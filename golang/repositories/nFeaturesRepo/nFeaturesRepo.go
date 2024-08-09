package nfeaturesrepo

import (
	"desabiller/models"
	"desabiller/repositories"
	"log"
	"strconv"
)

type nFeaturesRepo struct {
	Repo repositories.Repositories
}

func NewNFeaturesRepo(Repo repositories.Repositories) nFeaturesRepo {
	return nFeaturesRepo{
		Repo: Repo,
	}
}

func (ctx nFeaturesRepo) NCreateFeature(req models.ReqCreateNFeature) (id int, err error) {
	query := `insert into features (
		feature_name,
		created_at,
		updated_at,
		created_by,
		updated_by
	) values ($1,$2,$3,$4,$5) returning id`
	err = ctx.Repo.Db.QueryRow(query, req.FeatureName, req.CreatedAt, req.UpdatedAt, req.CreatedBy, req.UpdatedBy).Scan(&id)
	if err != nil {
		log.Println("ERROR REPO CreateFeature ", err)
		return 0, err
	}
	return id, err
}
func (ctx nFeaturesRepo) NReadFeature(req models.ReqGetListNFeature) (result []models.RespGetListNFeature, err error) {
	query := `select 
	id,
	feature_name,
	created_at,
	updated_at,
	created_by,
	updated_by
	from features where true `
	if req.Data.FeatureName != "" {
		query += ` and feature_name like '%` + req.Data.FeatureName + `%'`
	}
	if req.Data.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.Data.ID)
	}
	rows, err := ctx.Repo.Db.Query(query)
	if err != nil {
		log.Println("Err Repo ReadFeature " + err.Error())
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var val models.RespGetListNFeature
		err = rows.Scan(
			&val.ID,
			&val.FeatureName,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy,
		)
		if err != nil {
			log.Println("Err Repo ReadFeature" + err.Error())
			return result, err
		}
		result = append(result, val)
	}
	return result, err
}
func (ctx nFeaturesRepo) NReadSingleFeature(req models.ReqCreateNFeature) (result models.RespGetListNFeature, err error) {
	query := `select 
	id,
	feature_name,
	created_at,
	updated_at,
	created_by,
	updated_by
	from features where true `
	if req.FeatureName != "" {
		query += ` and feature_name like '%` + req.FeatureName + `%'`
	}
	if req.ID != 0 {
		query += ` and id = ` + strconv.Itoa(req.ID)
	}
	err = ctx.Repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.FeatureName,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.CreatedBy,
		&result.UpdatedBy,
	)
	if err != nil {
		log.Println("Err Repo ReadSingleFeature " + err.Error())
		return result, err
	}
	return result, nil
}
func (ctx nFeaturesRepo) NDropFeature(id int) (status bool, err error) {
	query := `delete from features where id = $1`
	_, err = ctx.Repo.Db.Exec(query, id)
	if err != nil {
		log.Println("Err Repo DropFeature " + err.Error())
		return false, err
	}
	return true, err
}
func (ctx nFeaturesRepo) NUpdateFeature(req models.ReqCreateNFeature) (result models.RespGetListNFeature, err error) {
	query := `update features set 
		feature_name=$1,
		updated_at=$2,
		updated_by=$3
		where id=$4 returning id, feature_name,
		updated_at,
		updated_by
		`
	err = ctx.Repo.Db.QueryRow(query, req.FeatureName, req.UpdatedAt, req.UpdatedBy, req.ID).Scan(&result.ID, &result.FeatureName, &result.UpdatedAt, &result.UpdatedBy)
	if err != nil {
		log.Println("Err Repo NUpdateFeature " + err.Error())
		return result, err
	}
	return result, err
}

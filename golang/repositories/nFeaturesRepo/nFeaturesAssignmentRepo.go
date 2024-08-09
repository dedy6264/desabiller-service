package nfeaturesrepo

import (
	"desabiller/models"
	"log"
	"strconv"
)

func (ctx nFeaturesRepo) NCreateFeatureAssignment(req models.ReqCreateNFeatureAssignment) (id int, err error) {
	query := `insert into feature_assignment (
		feature_id,
		merchant_id,
		created_at,
		updated_at,
		created_by,
		updated_by,
		index
	) values ($1,$2,$3,$4,$5,$6,$7) returning id`
	err = ctx.Repo.Db.QueryRow(query, req.FeatureId, req.MerchantId, req.CreatedAt, req.UpdatedAt, req.CreatedBy, req.UpdatedBy, req.Index).Scan(&id)
	if err != nil {
		log.Println("ERROR REPO CreateFeatureAssignment ", err)
		return 0, err
	}
	return id, err
}
func (ctx nFeaturesRepo) NReadFeatureAssignment(req models.ReqGetListNFeatureAssignment) (result []models.RespGetListNFeatureAssignment, err error) {
	query := `select 
	a.id,
	a.feature_id,
	a.merchant_id,
	a.created_at,
	a.updated_at,
	a.created_by,
	a.updated_by,
	b.merchant_name,
	c.feature_name
	from feature_assignment as a
	left join merchants as b on a.merchant_id=b.id
	left join features as c on a.feature_id=c.id
	where true `
	if req.Data.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.Data.ID)
	}
	if req.Data.FeatureId != 0 {
		query += ` and a.feature_id = ` + strconv.Itoa(req.Data.FeatureId)
	}
	if req.Data.MerchantId != 0 {
		query += ` and a.merchant_id = ` + strconv.Itoa(req.Data.MerchantId)
	}
	rows, err := ctx.Repo.Db.Query(query)
	if err != nil {
		log.Println("Err Repo ReadFeatureAssignment " + err.Error())
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var val models.RespGetListNFeatureAssignment
		err = rows.Scan(
			&val.ID,
			&val.FeatureId,
			&val.MerchantId,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy,
			&val.MerchantName,
			&val.FeatureName,
		)
		if err != nil {
			log.Println("Err Repo ReadFeature" + err.Error())
			return result, err
		}
		result = append(result, val)
	}
	return result, err
}
func (ctx nFeaturesRepo) NReadSingleFeatureAssignment(req models.ReqCreateNFeatureAssignment) (result models.RespGetListNFeatureAssignment, err error) {
	query := `select 
	a.id,
	a.feature_id,
	a.merchant_id,
	a.created_at,
	a.updated_at,
	a.created_by,
	a.updated_by,
	b.merchant_name,
	c.feature_name
	from feature_assignment as a
	left join merchants as b on a.merchant_id=b.id
	left join features as c on a.feature_id=c.id
	 where true `
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.FeatureId != 0 {
		query += ` and a.feature_id = ` + strconv.Itoa(req.FeatureId)
	}
	if req.MerchantId != 0 {
		query += ` and a.merchant_id = ` + strconv.Itoa(req.MerchantId)
	}
	err = ctx.Repo.Db.QueryRow(query).Scan(
		&result.ID,
		&result.FeatureId,
		&result.MerchantId,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.CreatedBy,
		&result.UpdatedBy,
		&result.MerchantName,
		&result.FeatureName,
	)
	if err != nil {
		log.Println("Err Repo ReadSingleFeatureAssignment " + err.Error())
		return result, err
	}
	return result, nil
}
func (ctx nFeaturesRepo) NDropFeatureAssignment(id int) (status bool, err error) {
	query := `delete from feature_assignment where id = $1`
	_, err = ctx.Repo.Db.Exec(query, id)
	if err != nil {
		log.Println("Err Repo DropFeatureAssignment " + err.Error())
		return false, err
	}
	return true, err
}
func (ctx nFeaturesRepo) NUpdateFeatureAssignment(req models.ReqCreateNFeatureAssignment) (result models.RespGetListNFeatureAssignment, err error) {
	query := `update feature_assignment set 
		feature_id=$1,
		merchant_id=$2,
		updated_at=$3,
		updated_by=$4
		where id=$5 returning id, feature_id,merchant_id,
		updated_at,
		updated_by
		`
	err = ctx.Repo.Db.QueryRow(query, req.FeatureId, req.MerchantId, req.UpdatedAt, req.UpdatedBy, req.ID).Scan(&result.ID, &result.FeatureId, &result.MerchantId, &result.UpdatedAt, &result.UpdatedBy)
	if err != nil {
		log.Println("Err Repo NUpdateFeatureAssignment " + err.Error())
		return result, err
	}
	return result, err
}

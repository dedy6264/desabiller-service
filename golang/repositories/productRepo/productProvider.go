package productrepo

// func (ctx product) AddProductBillerProvider(req models.ReqGetListProductBillerProvider) (result models.ResGetProductBillerProvider, status bool) {
// 	t := time.Now()

// 	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
// 	query := ` insert into product_biller_providers (
// 		product_provider_name,
// 		product_provider_code,
// 		product_provider_price,
// 		product_provider_admin_fee,
// 		product_provider_merchant_fee,
// 		is_open,
// 		product_type_id,
// 		product_category_id,
// 		created_at,
// 		updated_at,
// 		created_by,
// 		updated_by)
// 		values(
// 			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12
// 		) returning id `
// 	err := ctx.repo.Db.QueryRow(query,
// 		req.ProductProviderName,
// 		req.ProductProviderCode,
// 		req.ProductProviderPrice,
// 		req.ProductProviderAdminFee,
// 		req.ProductProviderMerchantFee,
// 		true,
// 		req.ProductTypeId,
// 		req.ProductCategoryId,
// 		dbTime,
// 		dbTime,
// 		"sys",
// 		"sys").Scan(&result.ID)
// 	if err != nil {
// 		log.Println("Error failed ", err.Error())
// 		return result, false
// 	}
// 	return result, true
// }
// func (ctx product) GetListProductBillerProvider(req models.ReqGetListProductBillerProvider) (result []models.ResGetProductBillerProvider, status bool) {
// 	query := `select
// id,
// product_provider_name,
// product_provider_code,
// product_provider_price,
// product_provider_admin_fee,
// product_provider_merchant_fee,
// is_open,
// product_type_id,
// product_category_id,
// created_at,
// updated_at,
// created_by,
// updated_by
// from product_biller_providers where true
// `
// 	if req.ProductProviderName != "" {
// 		query += ` and product_provider_name= '` + req.ProductProviderName + `'`
// 	}
// 	if req.ProductProviderCode != "" {
// 		query += ` and product_provider_code= '` + req.ProductProviderCode + `'`
// 	}
// 	if req.ProductTypeId != 0 {
// 		query += ` and product_type_id = ` + strconv.Itoa(req.ProductTypeId)
// 	}
// 	if req.ProductCategoryId != 0 {
// 		query += ` and product_category_id = ` + strconv.Itoa(req.ProductCategoryId)
// 	}
// 	if req.ID != 0 {
// 		query += ` and id = ` + strconv.Itoa(req.ID)
// 	}
// 	if req.StartDate != "" {
// 		query += ` and a.created_at between '` + req.StartDate + `' and '` + req.EndDate + `'`
// 	}
// 	if req.Limit != 0 {
// 		query += ` limit  ` + strconv.Itoa(req.Limit) + `  offset  ` + strconv.Itoa(req.Offset)
// 	} else {
// 		if req.OrderBy != "" {
// 			query += `  order by ` + req.OrderBy + ` asc`
// 		} else {
// 			query += `  order by product_provider_name asc`
// 		}
// 	}
// 	rows, err := ctx.repo.Db.Query(query)
// 	if err != nil {
// 		log.Println(" GetListProductBillerProvider :: Failed : ", err.Error())
// 		return result, false
// 	}
// 	var val models.ResGetProductBillerProvider
// 	for rows.Next() {
// 		err := rows.Scan(
// 			&val.ID,
// 			&val.ProductProviderName,
// 			&val.ProductProviderCode,
// 			&val.ProductProviderPrice,
// 			&val.ProductProviderAdminFee,
// 			&val.ProductProviderMerchantFee,
// 			&val.IsOpen,
// 			&val.ProductTypeId,
// 			&val.ProductCategoryId,
// 			&val.CreatedAt,
// 			&val.UpdatedAt,
// 			&val.CreatedBy,
// 			&val.UpdatedBy)
// 		if err != nil {
// 			return result, false
// 		}
// 		result = append(result, val)
// 	}
// 	if len(result) == 0 {
// 		return result, false
// 	}
// 	return result, false
// }
// func (ctx product) UpdateProductBillerProvider(req models.ReqGetListProductBillerProvider) (result models.ResGetProductBillerProvider, status bool) {
// 	t := time.Now()
// 	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
// 	query := ` update product_biller_providers set
// 	product_provider_name=$1,
// 	product_provider_code=$2,
// 	product_provider_price=$3,
// 	product_provider_admin_fee=$4,
// 	product_provider_merchant_fee=$5,
// 	is_open=$6,
// 	product_type_id=$7,
// 	product_category_id=$8,
// 	updated_at = $9,
// 	updated_by =$10
// 	where id = $11 returning id
// 	`
// 	err := ctx.repo.Db.QueryRow(query,
// 		req.ProductProviderName,
// 		req.ProductProviderCode,
// 		req.ProductProviderPrice,
// 		req.ProductProviderAdminFee,
// 		req.ProductProviderMerchantFee,
// 		req.IsOpen,
// 		req.ProductTypeId,
// 		req.ProductCategoryId,
// 		dbTime,
// 		"sys",
// 		req.ID).Scan(&result.ID)
// 	if err != nil {
// 		log.Println(" UpdateProductCategory :: Failed : ", err.Error())
// 		return result, false
// 	}
// 	return result, true
// }
// func (ctx product) DropProductBillerProvider(req models.ReqGetListProductBillerProvider) (status bool) {
// query := ` delete from product_biller_providers where id = $1`
// err := ctx.repo.Db.QueryRow(query, req.ID)
// if err.Err() != nil {
// 	log.Println("UpdateUserOutlet :: ", err.Err())
// 	return false
// }
// return true
// t := time.Now()

// dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)

// query := `update product_biller_providers set
// 				deleted_at = $1,
// 				deleted_by =$2
// 				where id = $3
// 				`
// err := ctx.repo.Db.QueryRow(query, dbTime, "sys", req.ID)
// fmt.Println(":::", err.Err())
// if err.Err() != nil {
// 	log.Println("UpdateProductBillerProvider :: ", err.Err())
// 	return false
// }
// return true
// }

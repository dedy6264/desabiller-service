package productrepo

// func (ctx product) GetListProductType() (result []models.ListProductType, err error) {
// 	query := `select
// 	id,
// 	product_type_name,
// 	product_type_code
// 	from product_types
// 	`
// 	rows, err := ctx.repo.Db.Query(query)
// 	if err != nil {
// 		log.Println(" GetListProductType :: Failed : ", err.Error())
// 		return result, err
// 	}
// 	var val models.ListProductType
// 	for rows.Next() {
// 		err := rows.Scan(&val.Id, &val.ProductTypeName, &val.ProductTypeCode)
// 		if err != nil {
// 			return result, err
// 		}
// 		result = append(result, val)
// 	}
// 	if len(result) == 0 {
// 		return result, nil
// 	}
// 	return result, nil
// }

package helperrepo

import (
	"database/sql"
	"desabiller/models"
	"log"
)

func (ctx helper) GetProductReferenceById(subscriberId string) (result models.RespGetPrefix, err error) {
	query := `select
a.id,	
a.product_reference_name,
a.product_reference_code
from product_references as a 
join product_helpers as b on b.product_reference_id=a.id
where b.product_prefix=$1
`
	err = ctx.repo.Db.QueryRow(query, subscriberId).Scan(
		&result.ProductReferenceId,
		&result.ProductReferenceName,
		&result.ProductReferenceCode,
	)
	if err != nil || err == sql.ErrNoRows {
		log.Println(err)
		return
	}
	return
}

package hierarchyrepo

import (
	"desabiller/models"
	"log"
)

func (ctx hierarchy) GetHierarchy(mID int) (result models.RespHierarchy, status bool) {
	query := ` select 
a.id,
a.merchant_name,
b.id,
b.client_name
from merchants as a
join clients as b on b.id=a.client_id
where a.id = $1
`
	err := ctx.repo.Db.QueryRow(query, mID).Scan(&result.MerchantId, &result.MerchantName, &result.ClientId, &result.ClientName)
	if err != nil {
		log.Println("GetHierarchy :: ", err.Error())
		return result, false
	}
	return result, true
}
func (ctx hierarchy) GetHierarchyByOutlet(oUID int) (result models.RespHierarchy, status bool) {
	query := ` select 
	a.id,
	a.merchant_name,
	b.id,
	b.client_name,
	c.id,
	c.merchant_outlet_name
	from merchants as a
	join clients as b on b.id=a.client_id
	join merchant_outlets as c on a.id=c.merchant_id
	where c.id = $1
	`
	err := ctx.repo.Db.QueryRow(query, oUID).Scan(&result.MerchantId, &result.MerchantName, &result.ClientId, &result.ClientName, &result.MerchantOutletId, &result.MerchantOutletName)
	if err != nil {
		log.Println("GetHierarchy :: ", err.Error())
		return result, false
	}
	return result, true
}

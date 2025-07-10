package hierarchyrepo

import (
	"desabiller/models"
	"fmt"
	"log"
)

func (ctx hierarchy) GetListUser(req models.ReqUserList) (resp []models.RespUserList, err error, status bool) {
	// t := time.Now()
	// dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `select 
	a.password,
	a.email,
	a.role_segment_id,
	COALESCE(a.hierarchy_type,''),
	COALESCE(a.hierarchy_id,0),
	b.role_segment_name,
	b.role_segment_code,
	COALESCE(b.role,'')
	from user_dashboard as a
	join role_segment_dashboards as b on a.role_segment_id=b.id
	where true 
	`
	if req.Email != "" {
		query += ` and a.email ='` + req.Email + `' `
	}
	if req.Password != "" {
		query += ` and a.password ='` + req.Password + `' `
	}

	if req.SortBy != "" {
		query += ` order by ` + req.SortBy + ` asc`
	} else {
		query += ` order by a.name asc`
	}
	fmt.Println("::::", query)
	rows, err := ctx.repo.Db.Query(query)
	if err != nil {
		log.Println("GetListUser :: ", err)
		return resp, err, false
	}
	defer rows.Close()
	for rows.Next() {
		var val models.RespUserList
		err := rows.Scan(
			&val.Password,
			&val.Email,
			&val.RoleSegmentId,
			&val.HierarchyType,
			&val.HierarchyId,
			&val.RoleSegmentName,
			&val.RoleSegmentCode,
			&val.Role,
		)
		if err != nil {
			log.Println("GetListUser :: ", err)
			return resp, err, false
		}
		resp = append(resp, val)
	}
	return resp, nil, true
}

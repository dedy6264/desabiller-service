package hierarchyrepo

import (
	"database/sql"
	"desabiller/configs"
	"desabiller/models"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (ctx hierarchy) DropOutletDevice(req models.ReqGetListOutletDevice) (status bool) {
	// query := ` delete from merchants where id = $1`
	// err := ctx.repo.Db.QueryRow(query, req.ID)
	// if err != nil {
	// 	log.Println("UpdateOutletDevice :: ", err.Err())
	// 	return false
	// }
	// return true
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	var id int
	query := `update outlet_devices set
					deleted_at = $1,
					deleted_by =$2
					where id = $3 returning id
					`
	err := ctx.repo.Db.QueryRow(query, dbTime, "sys", req.ID).Scan(&id)
	if err != nil {
		log.Println("UpdateOutletDevice :: ", err.Error())
		return false
	}
	return true
}
func (ctx hierarchy) UpdateOutletDevice(req models.ReqGetListOutletDevice) (result models.ResGetOutletDevice, status bool) {
	t := time.Now()

	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `update outlet_devices set
					device_type=$1,
					device_sn=$2,
					updated_at = $3,
					updated_by =$4
					where id = $5 returning id
					`
	err := ctx.repo.Db.QueryRow(query, req.DeviceType, req.DeviceSn, dbTime, req.Username, req.ID).Scan(&result.ID)
	if err != nil {
		log.Println("UpdateOutletDevice :: ", err.Error())
		return result, false
	}
	return result, true
}
func (ctx hierarchy) GetListOutletDeviceCount(req models.ReqGetListOutletDevice) (result int, status bool) {
	query := `select count(a.id)
	from outlet_devices as a
	join merchant_outlets as b on a.merchant_outlet_id=b.id
	join merchants as c on b.merchant_id=c.id
	join clients as d on c.client_id=d.id
	where a.deleted_by='' `

	if req.ClientId != 0 {
		query += ` and c.client_id = ` + strconv.Itoa(req.ClientId)
	}
	if req.MerchantId != 0 {
		query += ` and b.merchant_id = ` + strconv.Itoa(req.MerchantId)
	}
	if req.MerchantOutletId != 0 {
		query += ` and a.merchant_outlet_id = ` + strconv.Itoa(req.MerchantOutletId)
	}
	if req.MerchantOutletName != "" {
		query += ` and b.merchant_outlet_name = '` + req.MerchantOutletName + `' `
	}
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.StartDate != "" {
		query += ` and a.created_at between '` + req.StartDate + `' and '` + req.EndDate + `'`
	}
	err := ctx.repo.Db.QueryRow(query).Scan(&result)
	fmt.Println(":::", query)
	if err != nil {
		log.Println("GetListOutletDevice :: ", err.Error())
		return 0, false
	}
	return result, true
}
func (ctx hierarchy) GetListOutletDevice(req models.ReqGetListOutletDevice) (result []models.ResGetOutletDevice, status bool) {
	query := `select
	a.id,
	a.device_type,
	a.device_sn,
	a.merchant_outlet_id,
	b.merchant_outlet_name,
	b.merchant_id,
	c.merchant_name,
	c.client_id,
	d.client_name ,
	a.created_at,
	a.created_by,
	a.updated_at,
	a.updated_by
	from outlet_devices as a
	join merchant_outlets as b on a.merchant_outlet_id=b.id
	join merchants as c on b.merchant_id=c.id
	join clients as d on c.client_id=d.id
	where a.deleted_by='' `
	fmt.Println("===", req.DeviceSn)
	if req.DeviceSn != "" {
		query += ` and a.device_sn = '` + req.DeviceSn + `'`
	}
	if req.ClientId != 0 {
		query += ` and c.client_id = ` + strconv.Itoa(req.ClientId)
	}
	if req.MerchantId != 0 {
		query += ` and b.merchant_id = ` + strconv.Itoa(req.MerchantId)
	}
	if req.MerchantOutletId != 0 {
		query += ` and a.merchant_outlet_id = ` + strconv.Itoa(req.MerchantOutletId)
	}
	if req.MerchantOutletName != "" {
		query += ` and b.merchant_outlet_name = '` + req.MerchantOutletName + `' `
	}
	if req.ID != 0 {
		query += ` and a.id = ` + strconv.Itoa(req.ID)
	}
	if req.StartDate != "" {
		query += ` and a.created_at between '` + req.StartDate + `' and '` + req.EndDate + `'`
	}
	if req.Limit != 0 {
		query += ` limit  ` + strconv.Itoa(req.Limit) + `  offset  ` + strconv.Itoa(req.Offset)
	} else {
		if req.OrderBy != "" {
			query += `  order by a.` + req.OrderBy + ` asc`
		} else {
			query += `  order by a.id asc`
		}
	}

	rows, err := ctx.repo.Db.Query(query)
	fmt.Println(":::", query)
	if err != nil {
		log.Println("GetListOutletDevice :: ", err.Error())
		return result, false
	}
	result, status = DataRowOutletDevice(rows)
	if !status {
		return result, status
	}
	if len(result) == 0 {
		log.Println("Data not found")
		return result, false
	}
	return result, true
}

func DataRowOutletDevice(rows *sql.Rows) (result []models.ResGetOutletDevice, status bool) {
	for rows.Next() {
		var val models.ResGetOutletDevice
		err := rows.Scan(
			&val.ID,
			&val.DeviceType,
			&val.DeviceSn,
			&val.MerchantOutletId,
			&val.MerchantOutletName,
			&val.MerchantId,
			&val.MerchantName,
			&val.ClientId,
			&val.ClientName,
			&val.CreatedAt,
			&val.UpdatedAt,
			&val.CreatedBy,
			&val.UpdatedBy,
		)
		if err != nil {
			log.Println("DataRow :: ", err.Error())
			return result, false
		}
		result = append(result, val)
	}
	return result, true
}
func (ctx hierarchy) AddOutletDevice(req models.ReqGetListOutletDevice) (status bool) {
	fmt.Println("::::", req.MerchantOutletId)
	t := time.Now()
	id := 0
	dbTime := t.Local().Format(configs.LAYOUT_TIMESTAMP)
	query := `insert into outlet_devices (
		device_type,
		device_sn,
		merchant_outlet_id,
		created_at,
		created_by,
		updated_at,
		updated_by
		) values ($1,$2,$3,$4,$5,$6,$7) returning id
		`
	err := ctx.repo.Db.QueryRow(query, req.DeviceType, req.DeviceSn, req.MerchantOutletId, dbTime, "sys", dbTime, "sys").Scan(&id)
	fmt.Println("::::", query)
	if err != nil {
		log.Println("Err AddOutletDevice :: ", err)
		return false
	}
	fmt.Println(id)
	return true
}

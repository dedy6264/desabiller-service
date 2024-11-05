package trxrepo

import (
	"database/sql"
	"desabiller/models"
)

const insertQuery = `
reference_number,
provider_reference_number,
status_code,
status_message
`

func (ctx trxRepository) InsertTrxStatus(req models.ReqGetTrxStatus, tx *sql.Tx) (err error) {
	query := ` insert into trx_statuses (` + insertQuery + `) values (
		$1,$2,$3,$4) `
	if tx != nil {
		_, err = tx.Exec(query,
			req.ReferenceNumber,
			req.ProviderReferenceNumber,
			req.StatusCode,
			req.StatusMessage)
	} else {
		_, err = ctx.repo.Db.Exec(query,
			req.ReferenceNumber,
			req.ProviderReferenceNumber,
			req.StatusCode,
			req.StatusMessage)
	}
	if err != nil {
		return err
	}
	return nil
}

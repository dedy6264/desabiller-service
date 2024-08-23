package trxrepo

import (
	"desabiller/models"
)

const insertQuery = `
reference_number,
provider_reference_number,
status_code,
status_message
`

func (ctx trxRepository) InsertTrxStatus(req models.ReqGetTrxStatus) (err error) {
	query := ` insert into trx_statuses (` + insertQuery + `) values (
		$1,$2,$3,$4) `
	_, err = ctx.repo.Db.Exec(query,
		req.ReferenceNumber,
		req.ProviderReferenceNumber,
		req.StatusCode,
		req.StatusMessage)
	if err != nil {
		return err
	}
	return nil
}

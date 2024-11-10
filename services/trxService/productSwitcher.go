package trxservice

import (
	"desabiller/models"
	helperIakservice "desabiller/services/helperIakService"
	"errors"
	"strconv"
)

func (svc trxService) InqProviderSwitcher(req models.ProviderInqRequest) (respWorker models.ResponseWorkerInquiry, err error) {
	switch req.ProductClan {
	// case "BPJSKS":
	// 	respWorker, err = helperIakservice.IakPLNPostpaidWorkerInquiry(models.ReqInqIak{
	// 		ProductCode: req.ProductCode,
	// 		CustomerId:  req.SubscriberNumber,
	// 		RefId:       req.ReferenceNumber,
	// 		Url:         req.Url,
	// 		Month:       strconv.Itoa(req.Periode),
	// 	})
	case "PLN POST":
		respWorker, err = helperIakservice.IakPLNPostpaidWorkerInquiry(models.ReqInqIak{
			ProductCode: req.ProductCode,
			CustomerId:  req.SubscriberNumber,
			RefId:       req.ReferenceNumber,
			Url:         req.Url,
			Month:       strconv.Itoa(req.Periode),
		})
	default:
		err = errors.New("invalid product clan")
	}
	return
}

package trxservice

import (
	"desabiller/configs"
	"desabiller/models"
	iakworkerservice "desabiller/services/IAKWorkerService"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
)

func (svc trxService) InqProviderSwitcher(req models.ProviderInqRequest) (respWorker models.ResponseWorkerInquiry, err error) {
	switch req.ProviderName {
	case "IAK":
		switch req.ProductReferenceId {
		case 8:
			respWorker, err = iakworkerservice.IakBPJSWorkerInquiry(models.ReqInqIak{
				ProductCode: req.ProductCode,
				CustomerId:  req.SubscriberNumber,
				RefId:       req.ReferenceNumber,
				Url:         req.Url,
				Month:       strconv.Itoa(req.Periode),
			})
		case 9:
			respWorker, err = iakworkerservice.IakPLNPostpaidWorkerInquiry(models.ReqInqIak{
				ProductCode: req.ProductCode,
				CustomerId:  req.SubscriberNumber,
				RefId:       req.ReferenceNumber,
				Url:         req.Url,
				Month:       strconv.Itoa(req.Periode),
			})
		case 10:
			respWorker, err = iakworkerservice.IakPLNPrepaidWorkerInquiry(models.ReqInqIak{
				ProductCode: req.ProductCode,
				CustomerId:  req.SubscriberNumber,
				RefId:       req.ReferenceNumber,
				Url:         req.Url,
				Month:       strconv.Itoa(req.Periode),
			})
		default:
			err = errors.New("invalid product clan")
		}
	}
	return
}
func (svc trxService) PayProviderSwitcher(req models.ProviderPayRequest) (respWorker models.ResponseWorkerPayment, err error) {
	cc, _ := json.Marshal(req)
	fmt.Println("SINI 3", string(cc), req.ProductReferenceCode)
	switch req.ProviderName {
	case "IAK":
		switch req.ProductReferenceId {
		case 8:
			respWorker, err = iakworkerservice.IakBPJSWorkerPayment(models.ReqInqIak{
				// ProductCode: req.ProductCode,
				Commands: "pay-pasca",
				RefId:    req.ProviderReferenceNumber,
				Url:      req.Url,
			})
		case 9:
			respWorker, err = iakworkerservice.IakPLNPostpaidWorkerPayment(models.ReqInqIak{
				RefId:    req.ProviderReferenceNumber,
				Url:      req.Url,
				Commands: "pay-pasca",
			})
		default:
			err = errors.New("invalid product clan")
		}
	}
	return
}

func (svc trxService) CheckStatusProviderSwitcher(req models.ProviderInqRequest) (respWorker models.ResponseWorkerPayment, err error) {
	switch req.ProviderName {
	case "IAK":
		switch req.ProductReferenceId {
		case 8:
			respWorker, err = iakworkerservice.IakPostpaidWorkerCheckStatus(models.ReqInqIak{
				RefId:                req.ReferenceNumber,
				Commands:             "checkstatus",
				Url:                  req.Url,
				ProductReferenceCode: req.ProductReferenceCode,
			})
		case 10:
			respWorker, err = iakworkerservice.IakPrepaidWorkerCheckStatus(models.ReqInqIak{
				RefId: req.ReferenceNumber,
				Url:   req.Url,
			})
		default:
			respWorker = models.ResponseWorkerPayment{
				PaymentStatus:           configs.WORKER_PENDING_CODE,
				PaymentStatusDesc:       configs.PENDING_MSG,
				PaymentStatusDescDetail: "PENDING",
			}
			log.Println("CheckStatusProviderSwitcher :: invalid product clan")
		}
	}
	fmt.Println(req.ProductReferenceId)
	return
}

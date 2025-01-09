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
		switch req.ProductReferenceCode {
		case "BPJSKS":
			respWorker, err = iakworkerservice.IakBPJSWorkerInquiry(models.ReqInqIak{
				ProductCode: req.ProductCode,
				CustomerId:  req.SubscriberNumber,
				RefId:       req.ReferenceNumber,
				Url:         req.Url,
				Month:       strconv.Itoa(req.Periode),
			})
		case "PLN POST":
			respWorker, err = iakworkerservice.IakPLNPostpaidWorkerInquiry(models.ReqInqIak{
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
	fmt.Println("SINI 3", string(cc))
	switch req.ProviderName {
	case "IAK":
		switch req.ProductReferenceCode {
		case "BPJSKS":
			respWorker, err = iakworkerservice.IakBPJSWorkerPayment(models.ReqInqIak{
				// ProductCode: req.ProductCode,
				Commands: "pay-pasca",
				RefId:    req.ProviderReferenceNumber,
				Url:      req.Url,
			})
		case "PLN POST":
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
		switch req.ProductReferenceCode {
		case "BPJSKS":
			respWorker, err = iakworkerservice.IakPostpaidWorkerCheckStatus(models.ReqInqIak{
				RefId:                req.ReferenceNumber,
				Commands:             "checkstatus",
				Url:                  req.Url,
				ProductReferenceCode: req.ProductReferenceCode,
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
	return
}

package utils

import (
	"bytes"
	"crypto/tls"
	"desabiller/configs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func WorkerPostWithBearer(suffixUrl string, accessToken string, dataRequest interface{}, contentType string) ([]byte, int, error) {

	bodyRequest, err := json.Marshal(dataRequest)
	if err != nil {
		fmt.Println("Err Worker Post - json.Marshal : ", err.Error())
		return nil, http.StatusBadRequest, err
	}

	httpReq, err := http.NewRequest("POST", suffixUrl, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return nil, 0, err
	}
	defer httpReq.Body.Close()
	if contentType == "json" {
		httpReq.Header.Add("Content-Type", "application/json")
	}
	httpReq.Header.Set("Connection", "close")

	if accessToken != configs.EMPTY_VALUE {
		httpReq.Header.Add("Authorization", "Bearer "+accessToken)
	}

	httpReq.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, http.StatusBadGateway, err
	}

	resp.Header.Set("Connection", "close")
	defer resp.Body.Close()
	resp.Close = true

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	log.Println("Worker Request Url : ", suffixUrl)
	log.Println("Worker Request Data : ", string(bodyRequest))
	log.Println("Worker Response Data : ", string(bodyBytes))

	return bodyBytes, resp.StatusCode, nil
}

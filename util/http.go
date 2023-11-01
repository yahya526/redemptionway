package util

import (
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func DoHttp(req *http.Request) error {
	client := new(http.Client)
	client.Timeout = time.Second * 30

	reqBytes, err := httputil.DumpRequest(req, true)
	if err == nil {
		log.Println("请求：", string(reqBytes))
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	respBytes, err := httputil.DumpResponse(resp, true)
	if err == nil {
		log.Println("响应：", string(respBytes))
	}
	return nil
}

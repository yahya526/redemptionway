package util

import (
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func DoHttp(request *http.Request) error {
	client := new(http.Client)
	client.Timeout = time.Second * 30

	reqBytes, err := httputil.DumpRequest(request, true)
	if err == nil {
		log.Println("请求：", string(reqBytes))
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	respBytes, err := httputil.DumpResponse(resp, true)
	if err == nil {
		log.Println("响应：", string(respBytes))
	}
	return nil
}

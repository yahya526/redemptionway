package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func DoHttp(request *http.Request, result interface{}) error {
	client := new(http.Client)
	client.Timeout = time.Second * 30

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, result)
	if err != nil {
		return err
	}
	return nil
}

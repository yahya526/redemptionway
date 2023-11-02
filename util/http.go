package util

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
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

type HttpTemplateParser struct {
	Method  string
	Url     string
	Headers map[string]any
	Body    string
}

func (htp *HttpTemplateParser) ParseCurl(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(bytes), " \\")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, " '")
		if len(parts) == 0 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		if strings.Index(key, "curl") == 0 {
			key = strings.TrimSpace(key[4:])
		}
		key = TrimStartSwapLine(key)
		val = TrimEndSingleQuotationMark(val)
		switch key {
		case "-L", "--location":
			htp.Url = val
			break
		case "-X", "--request":
			htp.Method = val
			break
		case "-H", "--header":
			if htp.Headers == nil {
				htp.Headers = make(map[string]any)
			}
			hArr := strings.Split(val, ":")
			if len(hArr) < 2 {
				break
			}
			htp.Headers[strings.TrimSpace(hArr[0])] = strings.TrimSpace(hArr[1])
			break
		case "--data":
			htp.Body = val
			break
		default:
			log.Printf("order: %s not support\n", key)
		}
	}
	if len(htp.Method) == 0 {
		htp.Method = "GET"
		if len(htp.Body) > 0 {
			htp.Method = "POST"
		}
	}
	return nil
}

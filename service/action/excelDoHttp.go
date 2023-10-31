package action

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"redemptionway/constant"
	"redemptionway/entity"
	"redemptionway/util"
	"strings"
)

type ExcelDoHttpRedemption struct {
}

func (entity *ExcelDoHttpRedemption) Support(input string, action string) bool {
	return action == constant.ACTION_HTTP && input == constant.INPUT_EXCEL
}

func (entity *ExcelDoHttpRedemption) Redemption(config *entity.Config) {
	excel, err := excelize.OpenFile(config.Input.File)
	if err != nil {
		log.Println(fmt.Sprintf("读取Excel文件异常, 原因: %v", err))
		return
	}
	defer func(excel *excelize.File) {
		_ = excel.Close()
	}(excel)
	sheet := excel.GetSheetName(0)
	rows, err := excel.GetRows(sheet)
	if err != nil {
		log.Println(fmt.Sprintf("读取sheet %s 的行异常, 原因: %v", sheet, err))
		return
	}
	log.Println("读取到", len(rows), "行数据")
	if len(rows) <= 1 {
		return
	}
	head := rows[0]
	log.Println("表头", strings.Join(head, ","))
	ctx := (config.Action.Context).(map[string]interface{})
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		obj := make(map[string]interface{})
		for j := 0; j < len(head); j++ {
			obj[head[j]] = row[j]
		}
		objBytes, _ := json.Marshal(obj)

		req := new(http.Request)
		req.Method = method(ctx, objBytes)
		req.URL = url(ctx, objBytes)
		req.Header = header(ctx, objBytes)
		if http.MethodGet != req.Method {
			req.Body = body(ctx, objBytes)
		}
		err := util.DoHttp(req)
		if err != nil {
			log.Println("第", i, "行请求异常，原因：", fmt.Sprintf("%v", err))
		}
	}
}

func method(ctx map[string]interface{}, objBytes []byte) string {
	v, exist := ctx[constant.FILED_METHOD]
	if !exist {
		return http.MethodGet
	}
	return propertyResolve(v.(string), objBytes)
}

func url(ctx map[string]interface{}, objBytes []byte) *url2.URL {
	v, exist := ctx[constant.FILED_URL]
	if !exist {
		return nil
	}
	url, err := url2.Parse(propertyResolve(v.(string), objBytes))
	if err != nil {
		log.Println(fmt.Sprintf("解析url异常，原因：%v", err))
		return nil
	}
	return url
}

func header(ctx map[string]interface{}, objBytes []byte) http.Header {
	h := http.Header{}
	v, exist := ctx[constant.FILED_HEADER]
	if !exist {
		return h
	}
	ctxH := v.(map[string]interface{})
	if len(ctxH) == 0 {
		return h
	}
	for k, v := range ctxH {
		h[k] = []string{propertyResolve(fmt.Sprintf("%v", v), objBytes)}
	}
	return h
}

func body(ctx map[string]interface{}, objBytes []byte) io.ReadCloser {
	return nil
}

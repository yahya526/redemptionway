package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
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
	// 解析curl模板
	reqTemplate, err := parseCurlFile(config.Action.Param.(string))
	if err != nil {
		log.Printf("解析模板文件异常, 原因: %v\n", err)
		return
	}
	fmt.Println(reqTemplate.Url)

	excel, err := excelize.OpenFile(config.Input.File)
	if err != nil {
		log.Printf("读取Excel文件异常, 原因: %v\n", err)
		return
	}
	defer excel.Close()
	sheet := excel.GetSheetName(0)
	excel.SetActiveSheet(0)
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
	reqColumn, rspColumn := len(head)+1, len(head)+2
	_ = excel.SetCellStr(sheet, util.MustCell(reqColumn, 1), "请求报文")
	_ = excel.SetCellStr(sheet, util.MustCell(rspColumn, 1), "响应报文")
	ctx := (config.Action.Param).(map[string]interface{})
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		obj := make(map[string]interface{})
		for j := 0; j < len(head); j++ {
			obj[head[j]] = row[j]
		}
		objBytes, _ := json.Marshal(obj)
		reqCell := util.MustCell(reqColumn, i+1)
		repCell := util.MustCell(rspColumn, i+1)

		req, err := http.NewRequest(method(ctx, objBytes), url(ctx, objBytes), body(ctx, objBytes))
		if err != nil {
			_ = excel.SetCellStr(sheet, reqCell, fmt.Sprintf("%v", err))
			continue
		}
		req.Header = header(ctx, objBytes)
		reqBytes, _ := httputil.DumpRequest(req, true)
		_ = excel.SetCellStr(sheet, reqCell, string(reqBytes))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			_ = excel.SetCellStr(sheet, repCell, fmt.Sprintf("%v", err))
			continue
		}
		respBytes, _ := httputil.DumpResponse(resp, true)
		_ = excel.SetCellStr(sheet, repCell, string(respBytes))
	}
	err = excel.SaveAs(fmt.Sprintf("处理结果_%s.xlsx", util.NowStrSimple()))
	if err != nil {
		log.Printf("保存文件异常，原因：%v\n", err)
	}
}

func parseCurlFile(path string) (*UnresolvedHttpRequest, error) {
	buf := new(bytes.Buffer)
	callback := func(line string) {
		buf.WriteString(line)
	}
	err := util.Scan(path, callback)
	if err != nil {
		return nil, err
	}
	elements := strings.Split(buf.String(), "' \\")
	for _, line := range elements {

	}
	fmt.Println(len(elements))
	return nil, err
}

func method(ctx map[string]interface{}, objBytes []byte) string {
	v, exist := ctx[constant.HttpMethod]
	if !exist {
		return http.MethodGet
	}
	return propertyResolve(v.(string), objBytes)
}

func url(ctx map[string]interface{}, objBytes []byte) string {
	v, exist := ctx[constant.HttpURL]
	if !exist {
		return ""
	}
	return propertyResolve(v.(string), objBytes)
}

func header(ctx map[string]interface{}, objBytes []byte) http.Header {
	h := http.Header{}
	v, exist := ctx[constant.HttpHeaders]
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

func body(ctx map[string]interface{}, objBytes []byte) io.Reader {
	buf := new(bytes.Buffer)
	configBody, exist := ctx[constant.HttpBody]
	if !exist {
		return buf
	}
	bodyStr, ok := configBody.(string)
	if ok {
		return strings.NewReader(propertyResolve(bodyStr, objBytes))
	}
	bodyObj, ok := configBody.(map[string]interface{})
	if ok {
		for bodyK, bodyV := range bodyObj {
			bodyObj[bodyK] = propertyResolve(fmt.Sprintf("%v", bodyV), objBytes)
		}
	}
	buf.WriteString(util.MustMarshal(bodyObj))
	return buf
}

type UnresolvedHttpRequest struct {
	Url    string
	Header map[string]any
	Body   string
}

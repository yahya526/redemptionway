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
)

type ExcelDoHttpRedemption struct {
}

func (entity *ExcelDoHttpRedemption) Support(input string, action string) bool {
	return action == constant.ACTION_HTTP && input == constant.INPUT_EXCEL
}

func (entity *ExcelDoHttpRedemption) Redemption(config *entity.Config) {
	// 解析报文模板
	template := util.HttpTemplateParser{}
	err := template.ParseCurl(config.Action.Param.(string))
	if err != nil {
		log.Printf("解析请求模板文件异常, 原因: %v\n", err)
		return
	}
	// 读取excel文件
	excel, err := excelize.OpenFile(config.Input.File)
	if err != nil {
		log.Printf("读取Excel文件异常, 原因: %v\n", err)
		return
	}
	defer excel.Close()
	// 解析占位符并发送请求
	sheet := excel.GetSheetName(0)
	excel.SetActiveSheet(0)
	rows, err := excel.GetRows(sheet)
	if err != nil {
		log.Printf("读取sheet %s 的数据异常, 原因: %v\n", sheet, err)
		return
	}
	if len(rows) <= 1 {
		return
	}
	head := rows[0]
	reqColumn, rspColumn := len(head)+1, len(head)+2
	_ = excel.SetCellStr(sheet, util.MustCell(reqColumn, 1), "请求报文")
	_ = excel.SetCellStr(sheet, util.MustCell(rspColumn, 1), "响应报文")
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		obj := make(map[string]interface{})
		for j := 0; j < len(head); j++ {
			obj[head[j]] = row[j]
		}
		objBytes, _ := json.Marshal(obj)
		reqCell := util.MustCell(reqColumn, i+1)
		repCell := util.MustCell(rspColumn, i+1)

		req, err := http.NewRequest(method(&template, objBytes), url(&template, objBytes), body(&template, objBytes))
		if err != nil {
			_ = excel.SetCellStr(sheet, reqCell, fmt.Sprintf("%v", err))
			continue
		}
		req.Header = header(&template, objBytes)
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

func method(template *util.HttpTemplateParser, objBytes []byte) string {
	return propertyResolve(template.Method, objBytes)
}

func url(template *util.HttpTemplateParser, objBytes []byte) string {
	return propertyResolve(template.Url, objBytes)
}

func header(template *util.HttpTemplateParser, objBytes []byte) http.Header {
	h := http.Header{}
	if len(template.Headers) == 0 {
		return h
	}
	for k, v := range template.Headers {
		h[propertyResolve(k, objBytes)] = []string{propertyResolve(fmt.Sprintf("%v", v), objBytes)}
	}
	return h
}

func body(template *util.HttpTemplateParser, objBytes []byte) io.Reader {
	buf := new(bytes.Buffer)
	if len(template.Body) == 0 {
		return buf
	}
	buf.WriteString(propertyResolve(template.Body, objBytes))
	return buf
}

type UnresolvedHttpRequest struct {
	Url    string
	Header map[string]any
	Body   string
}

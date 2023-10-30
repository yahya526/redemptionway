package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"log"
	"redemptionway/constant"
	"redemptionway/entity"
	"redemptionway/util"
)

type JsonDoPrintRedemption struct {
}

func (jpr *JsonDoPrintRedemption) Support(input string, action string) bool {
	return action == constant.ACTION_PRINT && input == constant.INPUT_JSON
}

func (jpr *JsonDoPrintRedemption) Redemption(config *entity.Config) {
	arr := make([]map[string]interface{}, 0)
	err := util.ReadEntityFile(config.Input.File, &arr)
	if err != nil {
		log.Println(fmt.Sprintf("读取输入文件异常, 原因: %v", err))
		return
	}
	log.Println(fmt.Sprintf("读取到%d条数据", len(arr)))
	parser := entity.ParseContent(config.Action.Context)
	if len(parser.Subs) == 0 {
		log.Println(fmt.Sprintf("未解析出任何打印内容：%s", config.Action.Context))
		return
	}
	for _, obj := range arr {
		buf := new(bytes.Buffer)
		objBytes, _ := json.Marshal(obj)
		for _, sub := range parser.Subs {
			if !sub.Param {
				buf.WriteString(sub.Text)
				continue
			}
			buf.WriteString(fmt.Sprintf("%v", readJsonValue(objBytes, sub.Text)))
		}
		fmt.Println(buf.String())
	}
}

func readJsonValue(bytes []byte, key string) interface{} {
	res := gjson.GetBytes(bytes, key)
	return res.String()
}

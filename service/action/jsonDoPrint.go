package action

import (
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
	text := (config.Action.Context).(string)
	for _, obj := range arr {
		objBytes, _ := json.Marshal(obj)
		fmt.Println(propertyResolve(text, objBytes))
	}
}

func readJsonValue(bytes []byte, key string) interface{} {
	res := gjson.GetBytes(bytes, key)
	return res.Value()
}

package action

import (
	"redemptionway/constant"
	"redemptionway/entity"
)

type ExcelDoHttpRedemption struct {
}

func (entity *ExcelDoHttpRedemption) Support(input string, action string) bool {
	return action == constant.ACTION_HTTP && input == constant.INPUT_EXCEL
}

func (entity *ExcelDoHttpRedemption) Redemption(config *entity.Config) {

}

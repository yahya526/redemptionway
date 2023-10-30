package action

import (
	"redemptionway/constant"
	"redemptionway/entity"
)

type ExcelDoPrintRedemption struct {
}

func (entity *ExcelDoPrintRedemption) Support(input string, action string) bool {
	return action == constant.ACTION_PRINT && input == constant.INPUT_EXCEL
}

func (entity *ExcelDoPrintRedemption) Redemption(config *entity.Config) {
}

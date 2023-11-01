package util

import (
	"github.com/xuri/excelize/v2"
	"log"
)

func MustCell(column, row int) string {
	result, err := excelize.CoordinatesToCellName(column, row)
	if err != nil {
		log.Printf("构造excel单元格坐标失败，原因：%v", err)
	}
	return result
}

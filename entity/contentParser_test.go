package entity

import (
	"fmt"
	"redemptionway/util"
	"testing"
)

func TestParseContent(t *testing.T) {
	fmt.Println(util.MustMarshal(ParseContent("姓名：${name}，年龄：${age}，妻子：${wife.name}")))
	fmt.Println(util.MustMarshal(ParseContent("姓名：")))
	fmt.Println(util.MustMarshal(ParseContent("a${name}${wife.name}a")))
	fmt.Println(util.MustMarshal(ParseContent("${name}${wife.name}")))
}

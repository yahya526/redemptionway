package entity

import (
	"log"
	"redemptionway/constant"
	"strings"
)

type ContentParser struct {
	origin string
	Subs   []*SubContent
}

type SubContent struct {
	Text string

	Param bool
}

func ParseContent(context string) *ContentParser {
	result := new(ContentParser)
	result.origin = context
	if len(context) == 0 {
		return result
	}
	arrL := strings.Split(context, constant.PARAM_L)
	for _, v := range arrL {
		if len(v) == 0 {
			continue
		}
		arrR := strings.Split(v, constant.PARAM_R)
		result.Subs = append(result.Subs, &SubContent{
			Text:  arrR[0],
			Param: len(arrR) > 1,
		})
		if len(arrR) >= 2 {
			result.Subs = append(result.Subs, &SubContent{
				Text: arrR[1],
			})
		}
		if len(arrR) >= 3 {
			log.Printf("%s解析异常, 忽略", strings.Join(arrR[2:], ";"))
		}
	}
	return result
}

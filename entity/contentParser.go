package entity

import (
	"redemptionway/constant"
	"strings"
)

type ContentParser struct {
	origin string
	Subs   []*SubContent
}

type SubContent struct {
	Text string

	PlaceHolder bool
}

func ParseContent(context string) *ContentParser {
	result := ContentParser{
		origin: context,
	}
	for {
		if len(context) == 0 {
			break
		}
		idx := strings.Index(context, constant.PlaceHolderL)
		if idx < 0 {
			result.Subs = append(result.Subs, &SubContent{
				Text: context,
			})
			break
		}
		text := context[0:idx]
		if len(text) > 0 {
			result.Subs = append(result.Subs, &SubContent{
				Text: text,
			})
		}
		context = context[idx+len(constant.PlaceHolderL):]
		idx = strings.Index(context, constant.PlaceHolderR)
		if idx > 0 {
			text = context[0:idx]
			result.Subs = append(result.Subs, &SubContent{
				Text:        text,
				PlaceHolder: true,
			})
			context = context[idx+len(constant.PlaceHolderR):]
		} else if idx == 0 {
			context = context[idx+len(constant.PlaceHolderR):]
		} else {
			// next round
		}
	}
	return &result
}

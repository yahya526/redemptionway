package action

import (
	"bytes"
	"fmt"
	"redemptionway/entity"
)

func propertyResolve(text string, objBytes []byte) string {
	parser := entity.ParseContent(text)
	if len(parser.Subs) == 0 {
		return ""
	}
	buf := new(bytes.Buffer)
	for _, sub := range parser.Subs {
		if !sub.PlaceHolder {
			buf.WriteString(sub.Text)
			continue
		}
		buf.WriteString(fmt.Sprintf("%v", readJsonValue(objBytes, sub.Text)))
	}
	return buf.String()
}

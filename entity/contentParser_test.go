package entity

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestParseContent(t *testing.T) {
	jsonStr(ParseContent("姓名：{{name}}，年龄：{{age}}，妻子：{{wife.name}}"))
	jsonStr(ParseContent("姓名："))
	jsonStr(ParseContent("{{name}}{{wife}}.name}}"))
	jsonStr(ParseContent("{{name}} {{wife.name}}"))
}

func jsonStr(obj interface{}) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	fmt.Println(string(bytes))
}

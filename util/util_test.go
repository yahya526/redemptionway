package util

import (
	"log"
	"testing"
)

func TestHttpTemplateParser(t *testing.T) {
	parser := HttpTemplateParser{}
	err := parser.ParseCurl("E:\\workspace\\golang\\redemptionway\\conf\\http_template.txt")
	log.Printf("%v", err)
}

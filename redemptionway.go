package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"redemptionway/entity"
	"redemptionway/service"
	"redemptionway/service/action"
	"redemptionway/util"
)

var redemptionWays []service.RedemptionWay
var configFile string

func main() {
	flag.Parse()
	defer exit()

	config := new(entity.Config)
	err := util.ReadEntityFile(configFile, &config)
	if err != nil {
		log.Println(fmt.Sprintf("解析配置异常，原因：%v", err))
		return
	}
	for _, redemptionWay := range redemptionWays {
		if !redemptionWay.Support(config.Input.Type, config.Action.Type) {
			continue
		}
		redemptionWay.Redemption(config)
		return
	}
	log.Println(fmt.Sprintf("不支持的操作：%s-%s", config.Input.Type, config.Action.Type))
}

func init() {
	redemptionWays = make([]service.RedemptionWay, 0)
	redemptionWays = append(redemptionWays, new(action.ExcelDoHttpRedemption))
	redemptionWays = append(redemptionWays, new(action.JsonDoPrintRedemption))

	flag.StringVar(&configFile, "c", "config.json", "config file path")
}

func exit() {
	fmt.Println("请输入回车键退出")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	os.Exit(1)
}

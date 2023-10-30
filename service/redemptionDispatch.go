package service

import "redemptionway/entity"

type RedemptionWay interface {
	Support(input string, action string) bool

	Redemption(config *entity.Config)
}

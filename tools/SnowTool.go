package tools

import (
	"CloudRestaurant/common"
	"CloudRestaurant/config"
)

type SnowTool struct {
}

func GenerateNextId() (id int64) {
	worker, _ := common.NewWorker(config.Conf.WorkId)
	return worker.GetId()
}

package initialize

import (
	"fmt"
	"gin-starter/config"
	"gin-starter/global"
	"gin-starter/utils"
)

func Timer() {
	if global.CONFIG.Timer.Start {
		for _, detail := range global.CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				_, _ = global.Timer.AddTaskByFunc("ClearDB", global.CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(detail)
		}
	}
}

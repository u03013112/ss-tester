package tester

import (
	"fmt"
	"time"
)

// ScheduleInit :
func ScheduleInit() {
	go func() {
		for {
			check()
			time.Sleep(time.Second * 60 * 30 * 1)
		}
	}()
	return
}

func check() {
	// 读取数据库中需要检测配置
	scList := getSSList()
	if len(scList) == 0 {
		return
	}
	// 读取数据库中需要检测url
	urlList := getTestList()
	if len(urlList) == 0 {
		return
	}
	// 检测并更新结果
	for _, sc := range scList {
		ret, err := ssTest(&sc, 5, urlList)
		if err != nil {
			fmt.Printf("check err:%v\n", err)
		}
		// 简单做个处理，计算成功率
		success := 0
		for _, r := range ret.Result {
			if r.Code == 200 {
				success++
			}
		}
		rate := int64((success * 100) / len(urlList))
		fmt.Printf("updateSSConfig: ID[%d] -> IP[%s] rate[%d]\n", ret.ID, sc.IP, rate)
		updateSSConfig(ret.ID, sc.IP, rate)
	}
}

package tester

import (
	"github.com/u03013112/ss-tester/mod"
	"github.com/u03013112/ss-tester/sql"
)

// TestWebsite : 测试用网站，和测试次数，以及通过次数，用于发现有些url并不适合作为测试标准
type TestWebsite struct {
	sql.BaseModel
	URL     string
	Count   int64
	Success int64
}

// InitDB :
func InitDB() {
	sql.GetInstance().AutoMigrate(&mod.TestSSConfig{}, &TestWebsite{})
}

func getTestList() []string {
	var webList []TestWebsite
	sql.GetInstance().Find(&webList)
	ret := []string{}
	for _, web := range webList {
		ret = append(ret, web.URL)
	}
	return ret
}

func getSSList() []SSConfig {
	var ssList []mod.TestSSConfig
	sql.GetInstance().Find(&ssList)
	ret := []SSConfig{}
	for _, ss := range ssList {
		s := SSConfig{
			ID:     ss.ID,
			Domain: ss.Domain,
			IP:     ss.IP,
			Port:   ss.Port,
			Passwd: ss.Passwd,
			Method: ss.Method,
		}
		ret = append(ret, s)
	}
	return ret
}

// 更新指定ID配置中的IP和成功率
func updateSSConfig(ID uint, IP string, rate int64) {
	sql.GetInstance().Model(new(mod.TestSSConfig)).Where("id=?", ID).Select("ip", "rate").Updates(map[string]interface{}{"ip": IP, "rate": rate})
	return
}

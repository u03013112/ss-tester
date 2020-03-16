package tester

import (
	"github.com/u03013112/ss-tester/sql"
)

// TestWebsite : 测试用网站，和测试次数，以及通过次数，用于发现有些url并不适合作为测试标准
type TestWebsite struct {
	sql.BaseModel
	Url     string
	Count   int64
	Success int64
}

// TestSSConfig :有待测试的ss配置,rate 是一个综合评价，暂时就定为成功率，目前延时不作为判断标准
type TestSSConfig struct {
	sql.BaseModel
	Domain string `json:"domain,omitempty"`
	IP     string `json:"ip,omitempty"`
	Port   string `json:"port,omitempty"`
	Passwd string `json:"passwd,omitempty"`
	Method string `json:"method,omitempty"`
	Source string `json:"source,omitempty"`
	Backup string `json:"backup,omitempty"`
	Rate   int64  `json:"rate,omitempty"`
}

// InitDB :
func InitDB() {
	sql.GetInstance().AutoMigrate(&TestSSConfig{}, &TestWebsite{})
}

func getTestList() []string {
	var webList []TestWebsite
	sql.GetInstance().Find(&webList)
	ret := []string{}
	for _, web := range webList {
		ret = append(ret, web.Url)
	}
	return ret
}

func getSSList() []SSConfig {
	var ssList []TestSSConfig
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

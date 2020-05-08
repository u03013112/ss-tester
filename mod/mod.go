package mod

import "github.com/u03013112/ss-tester/sql"

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

// AddTestSSConfig :添加需要测试的SS配置，目前是用Domain做唯一键
func AddTestSSConfig(configList []TestSSConfig) {
	for _, config := range configList {
		var ret TestSSConfig
		db := sql.GetInstance().Where("domain=?", config.Domain).First(&ret)
		if db.Error != nil {
			sql.GetInstance().Create(&config)
		} else {
			sql.GetInstance().Model(&config).Updates(&config)
		}
	}
}

package spider

import "github.com/u03013112/ss-tester/sql"

// ProxyConfig : ss配置
type ProxyConfig struct {
	sql.BaseModel
	Domain string
	Port   string
	Passwd string
	Method string
}

// InitDB :
func InitDB() {
	sql.GetInstance().AutoMigrate(&ProxyConfig{})
}

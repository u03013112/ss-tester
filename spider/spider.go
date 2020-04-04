package spider

import (
	"encoding/base64"
	"strings"

	"github.com/u03013112/ss-tester/mod"
)

func start() {
	// 添加重试
	for i := 0; i < 10; i++ {
		if err := startShadowsocksRRShare(); err == nil {
			break
		}
	}
	for i := 0; i < 10; i++ {
		if err := startMyIshadowxBiz(); err == nil {
			break
		}
	}
}

// 输入url，可以支持多个，用换行分割
func urlParse(url string, source string) []mod.TestSSConfig {
	ret := []mod.TestSSConfig{}

	list := strings.Split(url, "\n")

	for _, line := range list {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "ss://") == false {
			continue
		}
		config := mod.TestSSConfig{
			Source: source,
		}
		line = strings.TrimPrefix(line, "ss://")
		list2 := strings.Split(line, "#")
		if len(list2) == 2 {
			config.Backup = list2[1]
		}
		// fmt.Println(list2[0])
		decoded, err := base64.StdEncoding.DecodeString(list2[0])
		if err != nil {
			continue
		}
		decodestr := string(decoded)
		list3 := strings.SplitN(decodestr, ":", 2)
		config.Method = list3[0]
		list4 := strings.SplitN(list3[1], "@", 2)
		config.Passwd = list4[0]
		list5 := strings.SplitN(list4[1], ":", 2)
		if len(list5) != 2 {
			continue
		}
		config.Domain = list5[0]
		config.Port = list5[1]

		if config.Method == "rc4" {
			continue
		}

		ret = append(ret, config)
	}
	return ret
}

package spider

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/u03013112/ss-tester/mod"
	"golang.org/x/net/proxy"
	"gopkg.in/xmlpath.v2"
)

func startMyIshadowxBiz() error {
	url0 := "https://my.ishadowx.biz/"
	stopSSLocal()
	_, err := startSSLocal()
	if err != nil {
		return err
	}
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:2080", nil, proxy.Direct)
	if err != nil {
		fmt.Println("can't connect to the proxy:", err)
		return err
	}
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	httpTransport.Dial = dialer.Dial
	resp, err := httpClient.Get(url0)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		n, err := xmlpath.ParseHTML(resp.Body)
		if err != nil {
			fmt.Println(err)
			return err
		}
		path, err := xmlpath.Compile("//div[@class='hover-text']")
		if err != nil {
			fmt.Println(err)
			return err
		}
		iter := path.Iter(n)
		configList := []mod.TestSSConfig{}
		for iter.Next() {
			config := mod.TestSSConfig{
				Source: "MyIshadowxBiz",
				Backup: "New",
			}
			h4List, err := xmlpath.Compile(".//h4")
			if err != nil {
				fmt.Println(err)
				return err
			}
			iter1 := h4List.Iter(iter.Node())
			for iter1.Next() {
				str := iter1.Node().String()
				str = strings.TrimSpace(str)
				if strings.HasPrefix(str, "IP Address:") {
					config.Domain = strings.TrimPrefix(str, "IP Address:")
				}
				if strings.HasPrefix(str, "Port:") {
					config.Port = strings.TrimPrefix(str, "Port:")
				}
				if strings.HasPrefix(str, "Password:") {
					config.Passwd = strings.TrimPrefix(str, "Password:")
				}
				if strings.HasPrefix(str, "Method:") {
					config.Method = strings.TrimPrefix(str, "Method:")
				}
			}
			if len(config.Domain) > 0 {
				configList = append(configList, config)
			}
		}
		mod.AddTestSSConfig(configList)
	}
	return nil
}

// Test :
func Test() error {
	fmt.Println("test")
	return nil
}

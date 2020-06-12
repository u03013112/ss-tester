package spider

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/u03013112/ss-tester/mod"
	"golang.org/x/net/proxy"
	"gopkg.in/xmlpath.v2"
)

func startShadowsocksRRShare() error {
	herf1 := shadowsocksRRShareS1()
	if herf1 != "" {
		herf2 := shadowsocksRRShareS2(herf1)
		if herf2 != "" {
			text := shadowsocksRRShareS3(herf2)
			if text != "" {
				// fmt.Println(text)
				configList := urlParse(text, "ShadowsocksRRShare")
				mod.AddTestSSConfig(configList)
				return nil
			}
		}
	}
	fmt.Println("startShadowsocksRRShare failed")
	return errors.New("failed")
}

func shadowsocksRRShareS3(href string) string {
	url0 := "https://www.github.com" + href
	stopSSLocal()
	_, err := startSSLocal()
	if err != nil {
		return ""
	}
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:2080", nil, proxy.Direct)
	if err != nil {
		fmt.Println("can't connect to the proxy:", err)
		return ""
	}
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	httpTransport.Dial = dialer.Dial
	resp, err := httpClient.Get(url0)
	if err != nil {
		fmt.Println("s3:", err)
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			return string(body)
		}
	}
	return ""
}
func shadowsocksRRShareS2(href string) string {
	url0 := "https://www.github.com" + href
	resp, err := http.Get(url0)
	if err != nil {
		fmt.Println("s2:", err)
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		n, err := xmlpath.ParseHTML(resp.Body)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		path, err := xmlpath.Compile("//a[@id='raw-url']//@href")
		iter := path.Iter(n)
		for iter.Next() {
			return iter.Node().String()
		}
	}
	return ""
}
func shadowsocksRRShareS1() string {
	url0 := "https://github.com/ruanfei/ShadowsocksRRShare/tree/master/ss"
	resp, err := http.Get(url0)
	if err != nil {
		fmt.Println("s1:", err)
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		n, err := xmlpath.ParseHTML(resp.Body)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		path, err := xmlpath.Compile("//tr[@class='js-navigation-item']")
		if err != nil {
			fmt.Println(err)
			return ""
		}
		iter := path.Iter(n)

		lastest := int64(0)
		var node *xmlpath.Node
		node = nil
		for iter.Next() {
			path1, err := xmlpath.Compile(".//@datetime")
			if err != nil {
				fmt.Println(err)
				return ""
			}
			iter1 := path1.Iter(iter.Node())

			for iter1.Next() {
				// s := iter1.Node().String()
				// fmt.Println(s)
				t, err := time.Parse(time.RFC3339, iter1.Node().String())
				if err == nil {
					d := t.Unix()
					if d > lastest {
						lastest = d
						node = iter.Node()
					}
				}
			}
		}
		if node != nil {
			path2, _ := xmlpath.Compile(".//a[@class='js-navigation-open ']//@href")
			iter2 := path2.Iter(node)
			for iter2.Next() {
				return iter2.Node().String()
			}
		}

	}
	return ""
}

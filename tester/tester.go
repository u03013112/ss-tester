package tester

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"os/exec"
	"time"
)

// Srv ：服务
type Srv struct{}

// SSConfig : ss配置
type SSConfig struct {
	ID     uint
	Domain string
	IP     string
	Port   string
	Passwd string
	Method string
}

// CurlResult : curl 检测结果
type CurlResult struct {
	URL   string
	Code  int64
	Delay int64
}

// SSTestResult : ss检测结果
type SSTestResult struct {
	ID     int32
	Result []CurlResult
}

func execShell(cmdName string, args []string) (string, error) {
	cmd := exec.Command(cmdName, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Printf("exec [%s] %v\n", cmdName, err)
		return "", err
	}
	return out.String(), nil
}

func startSSLocal(sc *SSConfig) (string, error) {
	// sslocal -s $SS_SERVER -p $SS_SERVER_PORT -k $SS_SERVER_PASSWD \
	// -b 0.0.0.0 -l $SOCKS_LOCAL_PORT -m $ENCRYPT_METHOD \
	// -d start
	args := []string{
		"-s", sc.IP,
		"-p", sc.Port,
		"-k", sc.Passwd,
		"-b", "127.0.0.1",
		"-l", "1080",
		"-m", sc.Method,
		"-d", "start",
	}
	return execShell("sslocal", args)
}

func stopSSLocal() {
	args := []string{
		"-9",
		"sslocal",
	}
	execShell("killall", args)
}

func testURL(URL string, timeout int64) *CurlResult {
	ret := new(CurlResult)
	ret.URL = URL
	fmt.Printf("test >%s<\n", URL)
	args := []string{
		"-c",
		fmt.Sprintf("curl --connect-timeout %d -m %d -I  \"%s\" --socks5 127.0.0.1:1080", timeout, timeout, URL),
	}
	t1 := time.Now().UnixNano()
	_, err := execShell("/bin/sh", args)
	if err != nil {
		fmt.Printf("curl error: %v\n", err)
		ret.Code = 500
		fmt.Printf("test >%s< failed\n", URL)
	} else {
		t2 := time.Now().UnixNano()
		ret.Delay = t2 - t1
		if err == nil {
			ret.Code = 200
		}
		fmt.Printf("test >%s< ok,uesd %d nano\n", URL, t2-t1)
	}
	return ret
}

// 输入：ss配置，超时：单位秒（连接超时和数据超时一致），需要检测url列表。返回结果
func ssTest(sc *SSConfig, timeout int64, URLList []string) (*SSTestResult, error) {
	if sc == nil {
		return nil, errors.New("sc is nil")
	}
	if len(URLList) <= 0 {
		return nil, errors.New("no URLLIst")
	}

	// 1、连接ss服务，填入域名的在这里临时dns一下，防止dns缓存，最终返回的也应该是IP方式的配置
	if sc.Domain != "" {
		if addrs, err := net.LookupHost(sc.Domain); err == nil {
			if len(addrs) > 0 {
				sc.IP = addrs[0]
			} else {
				return nil, errors.New("domain dns failed")
			}
		} else {
			return nil, errors.New("domain dns failed")
		}
	}
	// 先把之前的停掉
	stopSSLocal()
	time.Sleep(time.Second)
	str, err := startSSLocal(sc)
	if err == nil {
		fmt.Println(str)
	}
	fmt.Printf("connect to >>%s<<\n", sc.IP)
	// 2、等待一会，确保ss服务稳定
	time.Sleep(time.Second * 3)
	// 3、逐个url测试
	ret := new(SSTestResult)
	ret.ID = sc.ID
	for _, URL := range URLList {
		r := testURL(URL, timeout)
		if r != nil {
			ret.Result = append(ret.Result, *r)
		}
	}
	return ret, nil
}

func Test() {
	ssConfig := SSConfig{
		IP:     "107.182.186.33",
		Port:   "58700",
		Method: "aes-256-gcm",
		Passwd: "VVV5hw9PYb",
	}
	URLList := []string{
		"google.com",
		"https://www.youtube.com",
		"pornhub.com",
		"www.tumblr.com",
	}
	ret, err := ssTest(&ssConfig, 5, URLList)
	if err == nil {
		fmt.Printf("%#v\n", ret)
	}
}

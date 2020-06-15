package spider

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/u03013112/ss-tester/sql"
)

func startSSLocal() (string, error) {
	var sc ProxyConfig
	db := sql.GetInstance().First(&sc)
	if db.Error != nil {
		return "readDB Error", db.Error
	}
	args := []string{
		"-s", sc.Domain,
		"-p", sc.Port,
		"-k", sc.Passwd,
		"-b", "127.0.0.1",
		"-l", "2080",
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

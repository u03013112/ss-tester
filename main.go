package main

import (
	// pb "github.com/u03013112/ss-pb/tester"

	"time"

	"github.com/u03013112/ss-tester/spider"
	"github.com/u03013112/ss-tester/tester"
)

const (
	port = ":50004"
)

// for ci
func main() {
	tester.InitDB()
	tester.ScheduleInit()
	spider.ScheduleInit()
	for {
		time.Sleep(time.Second * 600)
	}
}

package spider

import (
	"time"
)

// ScheduleInit :
func ScheduleInit() {
	go func() {
		for {
			start()
			time.Sleep(time.Second * 60 * 60 * 1)
		}
	}()
	return
}

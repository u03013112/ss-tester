package spider

import (
	"time"

	"github.com/u03013112/ss-tester/mod"
)

// ScheduleInit :
func ScheduleInit() {
	go func() {
		for {
			for {
				if mod.Spiding == false {
					mod.Spiding = true
					start()
					mod.Spiding = false
					break
				} else {
					time.Sleep(time.Second * 1)
				}
			}
			time.Sleep(time.Second * 60 * 60 * 2)
		}
	}()
	return
}

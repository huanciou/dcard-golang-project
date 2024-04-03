package workers

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func CronJob() {
	c := cron.New()
	c.AddFunc("@daily", func() {
		fmt.Println("This is a cron job running every 5 seconds.")
	})
	c.Start()
	time.Sleep(time.Second * 3600)
}

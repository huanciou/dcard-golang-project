package workers

import (
	"dcard-golang-project/utils"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func CronJob() {
	c := cron.New()

	c.AddFunc("@daily", func() {
		utils.SetBitmaps()
		fmt.Println("Update redis bitmaps successfully")
	})

	c.Start()
	time.Sleep(time.Second * 3600)
}

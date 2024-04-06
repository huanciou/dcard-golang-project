package workers

import (
	"context"
	"dcard-golang-project/models"
	"dcard-golang-project/utils"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func CronJob() {
	c := cron.New()

	c.AddFunc("@daily", func() {

		ctx := context.Background()

		/* cron job 1: 將 redis 中的 post_queue 推送至 db 當中 */
		if data, isNotEmpty := utils.Dequeue(); isNotEmpty {
			models.DB.Create(&data)
		}

		/* cron job 2: 確保 數據於 db 儲存成功後，清洗 redis 快取 */
		models.Client.FlushAll(ctx)

		/* cron job 3: 將 mysql 中的 Ads 按照時序 拿出，並且放入 redis 建立 bitmap */
		utils.SetBitmaps()

		fmt.Println("Cron Job all Done!")
	})

	c.Start()
	time.Sleep(time.Second * 3600)
}

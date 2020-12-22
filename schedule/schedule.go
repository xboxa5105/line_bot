package schedule

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func ScheduleLine() {

	spec := "0 0 1,4,7 * * *"
	c := cron.New()
	c.AddFunc(spec, func() {
		rate_data, _ := HttpBank()
		// exchange_rate_data, _ := json.Marshal(rate_data)
		usd := fmt.Sprintf("美金 : %s\n日圓 : %s\n英鎊 : %s\n歐元 : %s", rate_data.USD, rate_data.JPY, rate_data.GBP, rate_data.EUR)
		LinePost(usd)
	})
	c.Start()
}

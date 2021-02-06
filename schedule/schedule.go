package schedule

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func ScheduleLine() {

	spec := "0 0 1,4,7 * * *"
	c := cron.New()
	c.AddFunc(spec, func() {
		ch := make(chan exchangeRate)
		go func() {
			data, _ := HTTPBank()
			ch <- *data
			close(ch)
		}()
		rateData, _ := HTTPBank()
		// exchange_rateData, _ := json.Marshal(rateData)
		usd := fmt.Sprintf("美金 : %s\n日圓 : %s\n英鎊 : %s\n歐元 : %s", rateData.USD, rateData.JPY, rateData.GBP, rateData.EUR)
		LinePost(usd)
	})
	c.Start()
}

package schedule

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type exchange_rate struct {
	USD string `json:"美金"`
	JPY string `json:"日圓"`
	GBP string `json:"英鎊"`
	EUR string `json:"歐元"`
}

func HttpBank() (*exchange_rate, error) {
	// var rate []string
	resp, err := http.Get("https://rate.bot.com.tw/xrt?Lang=zh-TW")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rate := doc.Find("[class=\"rate-content-sight text-right print_hide\"]").Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})
	rate_data := &exchange_rate{
		USD: strings.Replace(rate[1], " ", "", -1),
		JPY: strings.Replace(rate[15], " ", "", -1),
		GBP: strings.Replace(rate[5], " ", "", -1),
		EUR: strings.Replace(rate[29], " ", "", -1),
	}
	return rate_data, nil
}

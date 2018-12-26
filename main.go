package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/robfig/cron"
)

type exchange_rate struct {
	USD string `json:"美金"`
	JPY string `json:"日圓"`
	GBP string `json:"英鎊"`
	EUR string `json:"歐元"`
}

type body struct {
	To       string    `json:"to"`
	Messages []message `json:"messages"`
}
type message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func main() {
	rate_data, _ := http_bank()
	// exchange_rate_data, _ := json.Marshal(rate_data)
	usd := fmt.Sprintf("美金 : %s\n日圓 : %s\n英鎊 : %s\n歐元 : %s", rate_data.USD, rate_data.JPY, rate_data.GBP, rate_data.EUR)
	spec := "0 0 * * * *"
	c := cron.New()
	c.AddFunc(spec, func() {
		line_post(usd)
		fmt.Println(usd)
	})
	c.Start()
	// select {}

	fmt.Println("Server is on 8001")
	http.ListenAndServe(":8001", nil)
}

func http_bank() (*exchange_rate, error) {
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
		// For each item found, get the band and title
		// rate = append(rate, s.Text())
		// fmt.Println("here")
		fmt.Println(s.Text())

		// band := s.Find("a").Text()
		// title := s.Find("i").Text()
		// fmt.Printf("Review %d: %s - %s\n", i, band, title)

		return s.Text()
	})
	// fmt.Println(rate)
	rate_data := &exchange_rate{
		USD: rate[1],
		JPY: rate[15],
		GBP: rate[5],
		EUR: rate[29],
	}
	return rate_data, nil
}

func line_post(r string) string {
	CHANNEL_ACCESS_TOKEN := "X34ZiD0OaA+ZBUUVsGuXTRMwhb3S7G/XHv6oAW7lSbDyu+ZSeyuqoywj6D04fWLCtdvfiAiM9KZ+wCtwByMEFsp2DroqiOWfDbs5SJQgEGRNTidVJqGgag8gtKLVI5MyUE4VWq8v9uSTR0LBAWIyPwdB04t89/1O/w1cDnyilFU="
	line_url := "https://api.line.me/v2/bot/message/push"
	basictest := fmt.Sprintf("Bearer %s", CHANNEL_ACCESS_TOKEN)

	// body := []byte(`{
	// 		"to": "Ubfba3a940d4441bf51e81bfce0159bc1" ,
	// 		"messages": [{
	// 			"type":"text",
	// 			"text":"YYYYYYYYYES"
	// 		}]
	// }`)
	// json_b := []byte(body)
	fmt.Println(r)

	b := &body{
		To: "Ubfba3a940d4441bf51e81bfce0159bc1",
		Messages: []message{message{
			Type: "text",
			Text: r,
		}},
	}
	json_b, err := json.Marshal(b)

	// req, err := http.NewRequest("POST", line_url, bytes.NewBuffer(json_b))
	req, err := http.NewRequest("POST", line_url, bytes.NewReader(json_b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", basictest)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body1, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body1))
	return "success"
}

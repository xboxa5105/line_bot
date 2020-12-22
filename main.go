package main

import (
	"lineBot/schedule"
	"log"
	"net/http"
)

type exchange_rate struct {
	USD string `json:"美金"`
	JPY string `json:"日圓"`
	GBP string `json:"英鎊"`
	EUR string `json:"歐元"`
}

func main() {

	schedule.ScheduleLine()
	// rate_data, _ := schedule.HttpBank()
	// usd := fmt.Sprintf("美金 : %s\n日圓 : %s\n英鎊 : %s\n歐元 : %s", rate_data.USD, rate_data.JPY, rate_data.GBP, rate_data.EUR)
	// schedule.LinePost(usd)

	log.Println("Server is on 8001")
	http.ListenAndServe(":8001", nil)
}

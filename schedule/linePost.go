package schedule

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type body struct {
	To       string    `json:"to"`
	Messages []message `json:"messages"`
}
type message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func LinePost(r string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("CHANNEL_ACCESS_TOKEN"))
	basictest := fmt.Sprintf("Bearer %s", os.Getenv("CHANNEL_ACCESS_TOKEN"))

	b := &body{
		To: "Ubfba3a940d4441bf51e81bfce0159bc1",
		Messages: []message{message{
			Type: "text",
			Text: r,
		}},
	}
	json_b, err := json.Marshal(b)

	req, err := http.NewRequest("POST", os.Getenv("LINE_URL"), bytes.NewReader(json_b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", basictest)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body1, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body: " + string(body1))
	return "success"
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dghubble/oauth1"
)

type UserInfo struct {
	Name string `json:"name"`
}

type Tweet struct {
	User UserInfo `json:"user"`
	Text string   `json:"text"`
}

type tweets []Tweet

func main() {
	count := flag.Int("count", 50, "表示したいツイート数を入力して下さい")
	flag.Parse()

	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	path := fmt.Sprintf("https://api.twitter.com/1.1/statuses/home_timeline.json?count=%d", *count)
	resp, _ := httpClient.Get(path)
	defer resp.Body.Close()

	var tweets tweets
	reader := resp.Body
	if err := json.NewDecoder(reader).Decode(&tweets); err != nil {
		log.Fatal(err)
	}

	for _, tweet := range tweets {
		fmt.Printf("%s: %s\n\n", tweet.User.Name, tweet.Text)
	}
}
